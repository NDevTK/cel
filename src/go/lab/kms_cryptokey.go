// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package lab

import (
	"go/lab/config"
	cloudkms "google.golang.org/api/cloudkms/v1"
)

type KmsCryptoKey struct {
	BaseNamedAsset
	KeyRing *KmsKeyRing
	UsedBy  []*ServiceAccount

	cryptokey *cloudkms.CryptoKey
}

func (k *KmsCryptoKey) Resolve(s *Session) (err error) {
	defer Action(&err, "ensuring Google Cloud KMS key %s in %s",
		k.id, k.KeyRing.ResourcePath())

	// Fastpath is to first check for the existence of the key. If it does,
	// verify that the service accounts have access to it.
	k.cryptokey, err = s.GetCloudKmsService().Projects.Locations.KeyRings.CryptoKeys.Get(
		k.ResourcePath()).Context(s.Context).Do()
	if err == nil {
		return
	}

	defer Action(&err, "creating Google Cloud KMS key %s in %s",
		k.id, k.KeyRing.ResourcePath())

	// The keyRing should already exist since we've listed it as a dependency.
	// So we should be able to construct the key now.
	k.cryptokey, err = s.GetCloudKmsService().Projects.Locations.KeyRings.CryptoKeys.Create(
		k.KeyRing.ResourcePath(), &cloudkms.CryptoKey{Purpose: "ENCRYPT_DECRYPT"}).
		CryptoKeyId(k.id).Context(s.Context).Do()

	// Not sure what to do about this case.
	if err == nil && k.cryptokey.Primary.State != "ENABLED" {
		return NewError(
			"cryptokey %s in %s created, but not enabled. Primary key version state is %s",
			k.id, k.KeyRing.id, k.cryptokey.Primary.State)
	}

	return
}

func (k *KmsCryptoKey) Check(s *Session) (error, bool) {
	panic("not implemented")
}

func (k *KmsCryptoKey) Purge(s *Session) error {
	panic("not implemented")
}

func (k *KmsCryptoKey) ResourcePath() string {
	return k.KeyRing.ResourcePath() + "/cryptoKeys/" + k.id
}

func NewCryptoKey(A *Assets, p *Project, ck *config.Key, used_by []*ServiceAccount) *KmsCryptoKey {
	const kNamespace = "cryptoKeys"
	if a := A.Get(kNamespace, ck.Cryptokey); a != nil {
		return a.(*KmsCryptoKey)
	}

	l := NewKmsLocation(A, p, kDefaultKmsKeyLocation)
	r := NewKmsKeyRing(A, l, ck.Keyring)
	u := []Asset{r}
	for _, ub := range used_by {
		u = append(u, ub)
	}
	k := KmsCryptoKey{BaseNamedAsset{kNamespace, ck.Cryptokey, u}, r, used_by, nil}

	A.Add(&k)
	return &k
}

func ConstructCryptoKeyAndPermissionAssets(A *Assets, c *Config) error {
	project := LookupProject(A, c.Project)
	usage_map := make(map[*config.Key]map[*ServiceAccount]bool)

	for _, inst := range c.effective_instances {
		inst_sa := LookupServiceAccount(A, inst.CreateOptions.ServiceAccount.Id)
		inst_key := inst.CreateOptions.Cryptokey
		// Assuming |c.validate()| succeeded, both inst_sa and inst_key should be valid.
		if inst_sa == nil || inst_key == nil {
			panic(NewError(`instance %s does not have a service account or does not have a crypto key.

This case should've been handled in Config.validate(). I.e. Validation
should've failed and execution shouldn't have made this far.`, inst.Name))
		}

		if p, ok := usage_map[inst_key]; ok {
			p[inst_sa] = true
		} else {
			usage_map[inst_key] = map[*ServiceAccount]bool{inst_sa: true}
		}
	}

	for k, used_by_set := range usage_map {
		used_by := []*ServiceAccount{}
		for sa, _ := range used_by_set {
			used_by = append(used_by, sa)
		}

		ck := NewCryptoKey(A, project, k, used_by)
		NewKmsCryptoKeyPermissions(A, ck, used_by)
	}
	return nil
}
