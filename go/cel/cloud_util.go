// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cel

import (
	"fmt"
	"github.com/pkg/errors"
	googleapi "google.golang.org/api/googleapi"
	"strings"
)

// LastPathComponent extracts the zone out of a full URL. Many APIs require the
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

// MachineTypeResource returns the partial URL for a machine type resource.
func MachineTypeResource(project, zone, machtype string) string {
	return fmt.Sprintf("projects/%s/zones/%s/machineTypes/%s", project, zone, machtype)
}

// KmsLocationResource returns the partial URL for a KMS Location resource.
func KmsLocationResource(project string) string {
	return fmt.Sprintf("projects/%s/locations/global", project)
}

// KeyringResource return the partial URL for a KMS Keyring resource.
func KeyringResource(project, keyring string) string {
	return fmt.Sprintf("projects/%s/locations/global/keyRings/%s", project, keyring)
}

// CryptoKeyResource returns the partial URL for a KMS crypto key resource.
func CryptoKeyResource(project, keyring, cryptokey string) string {
	return fmt.Sprintf("projects/%s/locations/global/keyRings/%s/cryptoKeys/%s", project, keyring, cryptokey)
}

// IsNotFoundError returns true if |err| indicates that a requested Google
// cloud resource was not found.
func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	inner := errors.Cause(err)
	if e, ok := inner.(*googleapi.Error); ok {
		return e.Code == 404
	}
	return false
}
