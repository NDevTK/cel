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

var UnknownDigestError = errors.New("unknown digest while validating integrity string")
var DigestMismatchError = errors.New("integrity check failure")

func GetIntegrity(data []byte) string {
	digest := sha256.Sum256(data)
	return "sha256-" + base64.RawURLEncoding.EncodeToString(digest[:])
}

func ValidateIntegrity(data []byte, integrity string) error {
	if !strings.HasPrefix("sha256-", integrity) {
		return UnknownDigestError
	}
	digest := sha256.Sum256(data)
	expected_digest, err := base64.RawURLEncoding.DecodeString(integrity[7:])
	if err != nil {
		return err
	}
	if !bytes.Equal(digest[:], expected_digest) {
		return DigestMismatchError
	}
	return nil
}
