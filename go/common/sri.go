// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"bytes"
	"crypto/sha256"
	"encoding/base32"
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

// Token prefix to use for label encoded SRI token. This is distinct from the
// spec compliant prefix sriMetadataPrefix in order to distinguish tokens with
// different encodings.
const labelPrefix = "lsha256-"

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

// IntegrityLabel is similar to IntegrityToken, but uses SHA-256 and uses a
// base-32 encoding that produces valid label tokens.
//
// It is important that the length of the resulting string is at most 63.  This
// means that integrity tokens can be used in GCP labels which are constrained
// to 63 bytes.
//
// Note: The tokens returned by this function is not compliant with
// https://www.w3.org/TR/SRI/ which specifies using the standard Base64
// encoding.
func IntegrityLabel(data []byte) string {
	digest := sha256.Sum256(data)
	return labelPrefix + strings.ToLower(base32.HexEncoding.WithPadding(base32.NoPadding).
		EncodeToString(digest[:]))
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
	var rawDigest []byte
	var err error
	switch {
	case strings.HasPrefix(integrity, sriMetadataPrefix):
		rawDigest, err = base64.StdEncoding.DecodeString(integrity[len(sriMetadataPrefix):])

	case strings.HasPrefix(integrity, labelPrefix):
		rawDigest, err = base32.HexEncoding.WithPadding(base32.NoPadding).
			DecodeString(strings.ToUpper(integrity[len(labelPrefix):]))

	default:
		err = UnknownDigestAlgorithmError
	}

	if err != nil {
		return err
	}

	digest := sha256.Sum256(data)

	if !bytes.Equal(digest[:], rawDigest) {
		return DigestMismatchError
	}

	return nil
}
