// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package lab

import (
	"fmt"
	cloudkms "google.golang.org/api/cloudkms/v1"
)

// kKmsCryptoKeyDecryptorRole is the IAM role used to grant service account
// permission to use a KMS key for decryption.
//
// https://developers.google.com/apis-explorer/#search/iam.roles.querygrantableroles/m/iam/v1/iam.roles.queryGrantableRoles
// with a fullResourceName of
// "//cloudkms.googleapis.com/projects/*/locations/*/keyRings/*/cryptoKey/*"
// to view the applicable roles for a CryptoKey.
const kKmsCryptoKeyDecryptorRole = "roles/cloudkms.cryptoKeyDecrypter"

type KmsCryptoKeyPermissions struct {
	BaseNamedAsset
	Key    *KmsCryptoKey
	UsedBy []*ServiceAccount
}

func (kp *KmsCryptoKeyPermissions) Resolve(s *Session) (err error) {
	defer Action(&err, "ensuring service accounts can use Google Cloud KMS key %s", kp.Key.ResourcePath())

	p, err := s.GetCloudKmsService().Projects.Locations.KeyRings.CryptoKeys.GetIamPolicy(
		kp.Key.ResourcePath()).Context(s.Context).Do()
	if err != nil {
		return
	}

	var m []string
	for _, s := range kp.UsedBy {
		m = append(m, fmt.Sprintf("serviceAccount:%s", s.IamServiceAccount.Email))
	}

	found := false
	modified := false

	for _, b := range p.Bindings {
		if b.Role == kKmsCryptoKeyDecryptorRole {

			// Note that the logic here makes it impossible to revoke access to
			// a key by a service account since all modifications are additive.
			// Doing it this way for now because we don't want to clobber
			// manually configured data.
			b.Members, modified = unionString(b.Members, m)
			found = true
		}
	}

	// The IamPolicy already has everything we need.
	if found && !modified {
		return nil
	}

	if !found {
		p.Bindings = append(p.Bindings, &cloudkms.Binding{Role: kKmsCryptoKeyDecryptorRole, Members: m})
	}

	defer Action(&err, "updating policy for %s", kp.Key.ResourcePath())
	_, err = s.GetCloudKmsService().Projects.Locations.KeyRings.CryptoKeys.SetIamPolicy(
		kp.Key.ResourcePath(),
		&cloudkms.SetIamPolicyRequest{
			Policy:     p,
			UpdateMask: "bindings, etag"}).Context(s.Context).Do()
	return
}

func (kp *KmsCryptoKeyPermissions) Check(s *Session) (error, bool) {
	panic("not implemented")
}

func (kp *KmsCryptoKeyPermissions) Purge(s *Session) error {
	panic("not implemented")
}

func NewKmsCryptoKeyPermissions(A *Assets, k *KmsCryptoKey, u []*ServiceAccount) *KmsCryptoKeyPermissions {
	const kNamespace = "cryptoKey-Permissions"
	if a := A.Get(kNamespace, k.id); a != nil {
		return a.(*KmsCryptoKeyPermissions)
	}

	deps := []Asset{k}
	for _, a := range u {
		deps = append(deps, a)
	}

	p := KmsCryptoKeyPermissions{BaseNamedAsset{kNamespace, k.id, deps}, k, u}
	A.Add(&p)
	return &p
}

// unionString sticks |neeedles| into |haystack| if the latter doesn't already
// contain the elements in |needles|.
//
// The returend array is a union of the strings in |haystack| and |needles|. If
// no changes were required, then the bool return value would be false. Any
// strings from |needles| that were added to |haystack| will show up at the end
// of the returned slice.
func unionString(haystack []string, needles []string) ([]string, bool) {
	hm := make(map[string]bool)
	modified := false
	for _, h := range haystack {
		hm[h] = true
	}
	for _, n := range needles {
		if _, ok := hm[n]; ok {
			continue
		}
		haystack = append(haystack, n)
		modified = true
	}

	return haystack, modified
}
