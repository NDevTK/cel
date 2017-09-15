// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package lab

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	compute "google.golang.org/api/compute/v1"
	"math/big"
	"time"
)

const (
	// Name of metadata key.
	WindowsKeysKey = "windows-keys"

	// Lifetime of 'windows-keys' metadata key. If the VM sees the metadata
	// reqeust outside of the litemtime, then that entry will be ignored.
	WindowsKeysLifetime = time.Minute * 5
)

type WindowsKeyMetadataEntry struct {
	// AFAICT, this field is ignored.
	Email string `json:"email"`

	// Time after which this metadata entry should be ignored. Should be
	// encoded as per RFC 3339.
	ExpiresOn string `json:"expireOn"`

	// Base64 encoded BE integer.
	Exponent string `json:"exponent"`

	// Base64 encoded BE integer.
	Modulus string `json:"modulus"`

	// UserName to use. AFAICT can't be used to set domain credentials.
	UserName string `json:"userName"`

	PrivateKey *rsa.PrivateKey `json:"-"`
}

func NewWindowsKey(username, email string) (*WindowsKeyMetadataEntry, error) {
	k := WindowsKeyMetadataEntry{
		UserName:  username,
		Email:     email,
		ExpiresOn: time.Now().Add(WindowsKeysLifetime).Format(time.RFC3339)}
	private_key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	k.PrivateKey = private_key
	k.Modulus = base64.StdEncoding.EncodeToString(private_key.PublicKey.N.Bytes())
	k.Exponent = base64.StdEncoding.EncodeToString(big.NewInt(int64(private_key.PublicKey.E)).Bytes())

	return &k, nil
}

func UpdateGceMetadataWithWindowsKey(m *compute.Metadata, windows_key *WindowsKeyMetadataEntry) error {
	value_set := false
	for _, item := range m.Items {
		if item.Key == WindowsKeysKey {
			value_buffer, err := json.Marshal(windows_key)
			if err != nil {
				return err
			}
			item.Value = new(string)
			*item.Value = string(value_buffer)
			value_set = true
			break
		}
	}
	if !value_set {
		value_buffer, err := json.Marshal(windows_key)
		if err != nil {
			return err
		}
		value := new(string)
		*value = string(value_buffer)
		m.Items = append(m.Items, &compute.MetadataItems{
			Key:   WindowsKeysKey,
			Value: value})
	}
	return nil
}
