// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package lab

import (
	cloudkms "google.golang.org/api/cloudkms/v1"
)

type KmsKeyRing struct {
	BaseNamedAsset

	Location *KmsLocation
	key_ring *cloudkms.KeyRing
}

func (r *KmsKeyRing) Resolve(s *Session) (err error) {
	location := r.Location.ResourcePath()
	rn := r.ResourcePath()

	defer Action(&err, "creating Google Cloud KMS keyring %s in location %s", r.id, location)
	r.key_ring, err = s.GetCloudKmsService().Projects.Locations.KeyRings.Get(rn).Context(s.Context).Do()
	if err == nil {
		return
	}

	r.key_ring, err = s.GetCloudKmsService().Projects.Locations.KeyRings.Create(location, &cloudkms.KeyRing{}).
		KeyRingId(r.id).Context(s.Context).Do()
	return err
}

func (r *KmsKeyRing) Check(s *Session) (error, bool) {
	panic("not implemented")
}

func (r *KmsKeyRing) Purge(s *Session) error {
	panic("not implemented")
}

func (r *KmsKeyRing) ResourcePath() string {
	return r.Location.ResourcePath() + "/keyRings/" + r.id
}

func NewKmsKeyRing(A *Assets, l *KmsLocation, id string) *KmsKeyRing {
	const kNamespace = "keyRings"

	a := A.Get(kNamespace, id)
	if a != nil {
		return a.(*KmsKeyRing)
	}

	k := KmsKeyRing{
		BaseNamedAsset: BaseNamedAsset{namespace: kNamespace,
			id:         id,
			depends_on: []Asset{l}},
		Location: l}

	A.Add(&k)
	return &k
}
