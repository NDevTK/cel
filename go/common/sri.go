// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"bytes"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"strings"
)

// UnknownDigestError indicates that the subresource integrity metadata uses a
// digest algorithm that is not supported by this implementation.
var UnknownDigestError = errors.New("unknown digest while validating integrity string. Only SHA-384 is supported.")

// DigestMismatchError indicates that the subresource integrity metadata does
// not match the data.
var DigestMismatchError = errors.New("integrity check failure")

const sriMetadataPrefix = "sha384-"

// GetIntegrity calculates the subresource integrity metadata for a given blob
// of data.
//
// Subresource Integrity is described in https://www.w3.org/TR/SRI/ with
// integrity metadata described in
// https://www.w3.org/TR/SRI/#integrity-metadata . The calculations made by
// this function is limited to SHA-384 and does not attempt to capture anything
// beyond a content hash.
//
// The generated value can later be validated by the ValidateIntegrity()
// function, or any other SRI compliant mechanism as long as that mechanism
// supports the digest algorithm used when calculating the metadata.
func GetIntegrity(data []byte) string {
	digest := sha512.Sum384(data)
	return sriMetadataPrefix + base64.StdEncoding.EncodeToString(digest[:])
}

// ValidateIntegrity validates subresource integrity metadata for a given blob.
//
// Subresource Integrity is described in https://www.w3.org/TR/SRI/ with
// integrity metadata described in
// https://www.w3.org/TR/SRI/#integrity-metadata . This function is limited to
// verifying metadata that use SHA-384. Currently that's the only cryptographic
// hash function we support for the toolchain.
//
// TODO(asanka): Update these comments if we support any other hash functions.
func ValidateIntegrity(data []byte, integrity string) error {
	if !strings.HasPrefix(integrity, sriMetadataPrefix) {
		return UnknownDigestError
	}

	integrity = integrity[len(sriMetadataPrefix):]
	digest := sha512.Sum384(data)
	expected_digest, err := base64.StdEncoding.DecodeString(integrity)
	if err != nil {
		return err
	}
	if !bytes.Equal(digest[:], expected_digest) {
		return DigestMismatchError
	}
	return nil
}
