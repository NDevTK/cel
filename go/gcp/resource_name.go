// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

import (
	"fmt"
	"strings"
)

// LastPathComponent extracts the label out of a full URL. Many APIs require the
// short form label rather than the full URL.
func LastPathComponent(s string) string {
	i := strings.LastIndex(s, "/")
	if i == -1 {
		return s
	}
	return s[i+1:]
}

// ProjectResource returns the partial URL for a project.
func ProjectResource(project string) string {
	return fmt.Sprintf("projects/%s", project)
}

// ServiceAccountResource returns the partial URL for a service account.
func ServiceAccountResource(project, email string) string {
	return fmt.Sprintf("projects/%s/serviceAccounts/%s", project, email)
}

// ServiceAccountEmail returns the email address given a project ID and an account ID.
func ServiceAccountEmail(project, accountId string) string {
	// If the project ID has an org prefix (e.g.: "google.com:my-project"),
	// then the org domain comes after the project name in the domain portion.
	if strings.Contains(project, ":") {
		s := strings.Split(project, ":")
		return fmt.Sprintf("%s@%s.%s.iam.gserviceaccount.com", accountId, s[1], s[0])
	}
	return fmt.Sprintf("%s@%s.iam.gserviceaccount.com", accountId, project)
}

// ServiceAccountAclBinding returns the binding string to be used in a iam
// Policy object to refer to a service account by email.
func ServiceAccountAclBinding(saEmail string) string {
	return fmt.Sprintf("serviceAccount:%s", saEmail)
}

// MachineTypeResource returns the partial URL for a machine type resource.
func MachineTypeResource(project, zone, machtype string) string {
	return fmt.Sprintf("projects/%s/zones/%s/machineTypes/%s", project, zone, machtype)
}

// KmsLocationResource returns the partial URL for a KMS Location resource.
func KmsLocationResource(project, location string) string {
	return fmt.Sprintf("projects/%s/locations/%s", project, location)
}

// KeyringResource return the partial URL for a KMS Keyring resource.
func KeyringResource(project, location, keyring string) string {
	return fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", project, location, keyring)
}

// CryptoKeyResource returns the partial URL for a KMS crypto key resource.
func CryptoKeyResource(project, location, keyring, cryptokey string) string {
	return fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", project, location, keyring, cryptokey)
}
