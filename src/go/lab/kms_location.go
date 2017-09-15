// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package lab

import ()

type KmsLocation struct {
	BaseNamedAsset
	Project *Project
}

func (l *KmsLocation) Resolve(s *Session) error {
	// TODO(asanka): Check if location exists.
	return nil
}

func (l *KmsLocation) Check(s *Session) (error, bool) {
	panic("not implemented")
}

func (l *KmsLocation) Purge(s *Session) error {
	panic("not implemented")
}

func (l *KmsLocation) ResourcePath() string {
	return l.Project.ResourcePath() + "/locations/" + l.id
}

const kDefaultKmsKeyLocation = "global"

func NewKmsLocation(A *Assets, p *Project, id string) *KmsLocation {
	const kNamespace = "locations"
	if a := A.Get(kNamespace, id); a != nil {
		return a.(*KmsLocation)
	}

	l := KmsLocation{BaseNamedAsset{kNamespace, id, []Asset{p}}, p}
	A.Add(&l)
	return &l
}
