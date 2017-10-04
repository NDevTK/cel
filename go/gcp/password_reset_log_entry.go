// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// PasswordResetLogEntry describes the JSON object that's written out to COM4
// of a Windows VM after processing a password reset request via GCE instance
// metadata. The corresponding password reset request is described in
// WindowsKeyMetadataEntry.
type PasswordResetLogEntry struct {
	PasswordFound     bool   `json:"passwordFound"`
	Exponent          string `json:"exponent"`
	Modulus           string `json:"modulus"`
	UserName          string `json:"userName"`
	EncryptedPassword string `json:"encryptedPassword,omitempty"`
	ErrorMessage      string `json:"errorMessage,omitempty"`
}

// DecryptPassword uses |private_key| to decrypt the |EncryptedPassword|. This
// process should reverse the encryption performed in
// https://github.com/GoogleCloudPlatform/compute-image-windows
func (c *PasswordResetLogEntry) DecryptPassword(private_key *rsa.PrivateKey) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(c.EncryptedPassword)
	if err != nil {
		return "", err
	}

	plaintext, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, private_key, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("error decrypting password: %v", err)
	}
	return string(plaintext), nil
}

// Succeeded returns true if the password reset request succeeded.
func (c *PasswordResetLogEntry) Succeeded() bool {
	return c.PasswordFound
}

// MatchesWindowsKey return true if |c| is the response to |windows_key|.
func (c *PasswordResetLogEntry) MatchesWindowsKey(windows_key *WindowsKeyMetadataEntry) bool {
	return c.Modulus == windows_key.Modulus &&
		c.Exponent == windows_key.Exponent &&
		c.UserName == windows_key.UserName
}

// NewPasswordResetLogEntry creates a PasswordResetLogEntry based on a byte
// stream that was read from the COM4 port of  Windwos VM. After calling this
// function you can call MatchesWindowsKey() method to determine if the
// returned PasswordResetLogEntry matches a specific password reset request.
func NewPasswordResetLogEntry(json_encoded_bytes []byte) (*PasswordResetLogEntry, error) {
	var entry PasswordResetLogEntry
	err := json.Unmarshal(json_encoded_bytes, &entry)
	if err != nil {
		return nil, err
	}
	return &entry, nil
}
