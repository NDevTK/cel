// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

import (
	"context"
	"fmt"
	compute "google.golang.org/api/compute/v1"
	"net/http"
	"strings"
	"time"
)

const (
	// Commands time out after one minute by default.
	CommandRetryTimeout = time.Minute
)

func extractLastMatchingPasswordEntry(output string, windows_key *WindowsKeyMetadataEntry) *PasswordResetLogEntry {
	lines := strings.Split(output, "\n")
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		credentials, err := NewPasswordResetLogEntry([]byte(line))
		if err != nil {
			continue
		}

		if credentials.Succeeded() && credentials.MatchesWindowsKey(windows_key) {
			return credentials
		}
	}

	return nil
}

// ResetWindowsPassword resets the password of a Windows VM hosted on Google
// Compute Engine.
//
// It uses the GCE Windows Agent's password reset mechanism and hence only
// works on the public Windows VMs which are running the agent. The code for
// the agent can be found at
// https://github.com/GoogleCloudPlatform/compute-image-windows .
//
// At a high level, password resets work as follows:
//
// 1. Someone with edit privileges over the GCE project adds a metadata entry
//    for a GCE instance with key 'windows-key', and JSON object as a value.
//    The value encodes an RSA public key and a username to use for the
//    new/existing account.
//
// 2. The agent running in the VM observes the metadata change, picks up the
//    new 'windows-key' metdata entry. The agent differentiates the new
//    'windows-key' entry from a prior entry by comparing the public key.
//
// 3. The agent creates the user account if one doesn't exist, or initiates a
//    password reset for an existing account with the same username. The new
//    password is generated at random. Upon success, the new password is
//    encrypted using the public key in the 'windows-key' metadata and written
//    as a string encoded JSON object into the COM4 port. The JSON object also
//    includes the public key in order to identify the password reset request
//    associated with this response.
//
// 4. Meanwhile the calling code polls the COM4 output for the VM. If a line
//    shows up with the correct public key, the calling code can use the
//    private key associated with the password change request to decrypt the
//    encrypted password.
//
// The http.Client referenced by |client| will be used for interacting with the
// Google Cloud API and should already incorporate any required authentication.
//
// If the request is successful, the function returns the new password as a
// string. Otherwise there will be an error.
func ResetWindowsPassword(ctx context.Context, client *http.Client,
	project, zone, instance, username, email string) (string, error) {

	service, err := compute.New(client)
	if err != nil {
		return "", err
	}

	instance_data, err := service.Instances.Get(project, zone, instance).
		Context(ctx).Fields("status", "metadata").Do()
	if err != nil {
		return "", err
	}

	if instance_data.Status != "RUNNING" {
		return "", fmt.Errorf("expected status to be RUNNING, but found %s", instance_data.Status)
	}

	metadata := instance_data.Metadata
	windows_key, err := NewWindowsKey(username, email)
	if err != nil {
		return "", err
	}

	err = UpdateGceMetadataWithWindowsKey(metadata, windows_key)
	if err != nil {
		return "", err
	}

	_, err = service.Instances.SetMetadata(project, zone, instance, metadata).Context(ctx).Do()
	if err != nil {
		return "", err
	}

	// One minute timeout. Slow VMs may take a while to reset the password
	// of an existing user.
	timeout_at := time.Now().Add(CommandRetryTimeout)

	var password_entry *PasswordResetLogEntry

	for time.Now().Before(timeout_at) {
		output, err := service.Instances.GetSerialPortOutput(project, zone, instance).
			Port(4).Context(ctx).Do()
		if err != nil {
			return "", err
		}

		password_entry = extractLastMatchingPasswordEntry(output.Contents, windows_key)
		if password_entry != nil {
			break
		}
	}

	if password_entry == nil {
		return "", fmt.Errorf("timed out while waiting for password reset")
	}

	return password_entry.DecryptPassword(windows_key.PrivateKey)
}
