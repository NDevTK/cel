// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package lab

import (
	"testing"
)

type testAsset struct {
	BaseNamedAsset
	Resolved bool
}

const kTestNamespace = "testing"

func (t *testAsset) Resolve(S *Session) error {
	t.Resolved = true
	return nil
}

func (t *testAsset) Check(s *Session) (error, bool) {
	panic("not implemented")
}

func (t *testAsset) Purge(s *Session) error {
	panic("not implemented")
}

func TestAdd(t *testing.T) {
	A := &Assets{}
	err := A.Add(&testAsset{BaseNamedAsset{kTestNamespace, "a", nil}, false})
	if err != nil {
		t.Errorf("failed to Add: %s", err.Error())
	}
}

func TestGetDependency(t *testing.T) {
	A := &Assets{}
	a_asset := &testAsset{BaseNamedAsset{kTestNamespace, "a", nil}, false}
	if err := A.Add(a_asset); err != nil {
		t.Error(err)
	}

	b_asset := &testAsset{BaseNamedAsset{kTestNamespace, "b", nil}, false}
	if err := A.Add(b_asset); err != nil {
		t.Error(err)
	}

	d := A.GetNodeForAsset(a_asset)
	if d == nil {
		t.Error("failed to get dependency for 'a'")
	}

	if d.Asset.Id() != "a" {
		t.Error("got incorrect dependency")
	}

	if A.GetNodeForAsset(b_asset).Asset.Id() != "b" {
		t.Error("got incorrect dependency")
	}
}
