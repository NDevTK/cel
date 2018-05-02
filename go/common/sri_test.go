// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"testing"
)

// Also tests ValidateIntegrity
func TestIntegrity_knownVectors(t *testing.T) {
	testCases := [...]struct {
		input    []byte
		expected string
	}{
		// Generated via: printf 'hello world!' | openssl dgst -sha256 -binary | openssl base64 -A
		{[]byte("hello world!"), "sha256-dQnlvaDHYtK6x/kNdYtbImP6Acy8VCq1498WO+CObKk="},

		// Generated via: echo -n "" | openssl dgst -sha256 -binary | openssl base64 -A
		{[]byte{}, "sha256-47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU="},
	}

	for _, tc := range testCases {
		result := IntegrityToken(tc.input)
		if result != tc.expected {
			t.Errorf("Got \"%s\". Want \"%s\" for input [%s]", result, tc.expected, tc.input)
		}

		err := ValidateIntegrity(tc.input, tc.expected)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestIntegrityURLToken_sanity(t *testing.T) {
	input := []byte("hello world!")
	s := IntegrityLabel(input)
	if s != "lsha256-ek4ubfd0othd5em7v46nb2qr49hvk0ecnha2ldf3rsb3no4edikg" {
		t.Errorf("Got \"%s\"", s)
	}

	err := ValidateIntegrity(input, s)
	if err != nil {
		t.Error(err)
	}
}

// As long as the length limit for GCP metadata strings are at 64, this test
// should ensure that the token emitted by IntegrityURLToken is less than that.
func TestIntegrityURLToken_length(t *testing.T) {
	s := IntegrityLabel([]byte("hello world!"))
	if len(s) > 63 {
		t.Errorf("IntegrityURLToken returned a token of length %d. Want less than 64", len(s))
	}
}

func TestValidateIntegrity_badDigest(t *testing.T) {
	err := ValidateIntegrity([]byte{}, "")
	if err != UnknownDigestAlgorithmError {
		t.Error(err)
	}

	err = ValidateIntegrity([]byte{}, "sha512-abcdef")
	if err != UnknownDigestAlgorithmError {
		t.Error(err)
	}
}

func TestValidateIntegrity_badEncoding(t *testing.T) {
	// the trailing '?' is not a valid base64 character.
	err := ValidateIntegrity([]byte{},
		"sha256-OLBgp1GsljhM2TJ+sbHjaiH9txEUvgdDTAzHv2P24donTt6/529l+9Ua0vFImLl?")
	if err == nil {
		t.Error("should've failed to decode, but succeeded")
	}
}
