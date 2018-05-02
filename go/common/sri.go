// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"strings"
)

// UnknownDigestAlgorithmError indicates that the subresource integrity
// metadata uses a digest algorithm that is not supported by this
// implementation.
var UnknownDigestAlgorithmError = errors.New(
	"unknown digest while validating integrity string. Only SHA-256 is supported.")

// DigestMismatchError indicates that the subresource integrity metadata does
// not match the data.
var DigestMismatchError = errors.New("integrity check failure")

// Prefix for SRI compliant integrity token.
const sriMetadataPrefix = "sha256-"

// IntegrityToken calculates the subresource integrity metadata for a given
// blob of data.
//
// Subresource Integrity is described in https://www.w3.org/TR/SRI/ with
// integrity metadata described in
// https://www.w3.org/TR/SRI/#integrity-metadata . The calculations made by
// this function is limited to SHA-256 and does not attempt to capture anything
// beyond a content hash.
//
// The generated value can later be validated by the ValidateIntegrity()
// function, or any other SRI compliant mechanism as long as that mechanism
// supports the digest algorithm used when calculating the metadata.
func IntegrityToken(data []byte) string {
	digest := sha256.Sum256(data)
	return sriMetadataPrefix + base64.StdEncoding.EncodeToString(digest[:])
}

// IntegrityURLToken is similar to IntegrityToken, but uses SHA-256 and the URL
// safe alternate base64 encoding (RFC 4648) with no padding.
//
// It is important that the length of the resulting string is less than 64.
// This means that integrity tokens can be used in GCP labels which are
// constrained to 64 bytes.
//
// Note: The tokens returned by this function is not strictly compliant with
// https://www.w3.org/TR/SRI/ which specifies using the standard Base64
// encoding.
func IntegrityURLToken(data []byte) string {
	digest := sha256.Sum256(data)
	return sriMetadataPrefix + base64.RawURLEncoding.EncodeToString(digest[:])
}

// ValidateIntegrity validates subresource integrity metadata for a given blob.
//
// Subresource Integrity is described in https://www.w3.org/TR/SRI/ with
// integrity metadata described in
// https://www.w3.org/TR/SRI/#integrity-metadata . Consult GetIntegrity() for
// information on which digest algorithm is supported.
//
// ValidateIntegrity() can be called with the output of either GetIntegrity()
// or GetIntegrityURLToken().
func ValidateIntegrity(data []byte, integrity string) error {
	if !strings.HasPrefix(integrity, sriMetadataPrefix) {
		return UnknownDigestAlgorithmError
	}

	integrity = integrity[len(sriMetadataPrefix):]
	digest := sha256.Sum256(data)
	expected_digest, err := base64.StdEncoding.DecodeString(integrity)

	// Try URL encoding as well. This way we have our bases covered regardless
	// of which encoding scheme is in use. The condings differ in their use of
	// '+' and '/' characters. All other matching characters are used to
	// represent the same base64 values.
	if err != nil {
		expected_digest, err = base64.RawURLEncoding.DecodeString(integrity)
	}
	if err != nil {
		return err
	}
	if !bytes.Equal(digest[:], expected_digest) {
		return DigestMismatchError
	}
	return nil
}
