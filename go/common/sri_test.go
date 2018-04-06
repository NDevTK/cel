// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"testing"
)

// Also tests ValidateIntegrity
func TestGetIntegrity_knownVectors(t *testing.T) {
	testCases := [...]struct {
		input    []byte
		expected string
	}{

		// Example in https://w3c.github.io/webappsec-subresource-integrity/#integrity-metadata-description
		{[]byte("alert('Hello, world.');"), "sha384-H8BRh8j48O9oYatfu5AZzq6A9RINhZO5H16dQZngK7T62em8MUt1FLm52t+eX6xO"},

		// Generated via: echo -n "" | openssl dgst -sha384 -binary | openssl base64 -A
		{[]byte{}, "sha384-OLBgp1GsljhM2TJ+sbHjaiH9txEUvgdDTAzHv2P24donTt6/529l+9Ua0vFImLlb"},
	}

	for _, tc := range testCases {
		result := GetIntegrity(tc.input)
		if result != tc.expected {
			t.Errorf("Got \"%s\". Want \"%s\"", result, tc.expected)
		}

		err := ValidateIntegrity(tc.input, tc.expected)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestValidateIntegrity_badDigest(t *testing.T) {
	err := ValidateIntegrity([]byte{}, "")
	if err != UnknownDigestError {
		t.Error(err)
	}

	err = ValidateIntegrity([]byte{}, "sha256-abcdef")
	if err != UnknownDigestError {
		t.Error(err)
	}
}

func TestValidateIntegrity_badEncoding(t *testing.T) {
	// the trailing '?' is not a valid base64 character.
	err := ValidateIntegrity([]byte{},
		"sha384-OLBgp1GsljhM2TJ+sbHjaiH9txEUvgdDTAzHv2P24donTt6/529l+9Ua0vFImLl?")
	if err == nil {
		t.Error("should've failed to decode, but succeeded")
	}
}
