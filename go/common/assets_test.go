// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"testing"
)

type testAsset struct {
	BaseNamedAsset
	Resolved bool
}

const kTestNamespace = "testing"

func (t *testAsset) Resolve() error {
	t.Resolved = true
	return nil
}

func (t *testAsset) Check() error {
	panic("not implemented")
}

func (t *testAsset) Purge() error {
	panic("not implemented")
}

type testScriptAsset struct {
	BaseNamedAsset
	Resolved bool
}

func (t *testScriptAsset) GenerateScript() error {
	t.Resolved = true
	return nil
}

type testPermanentAsset struct {
	BaseNamedAsset
	Resolved bool
}

func (t *testPermanentAsset) Check() error {
	t.Resolved = true
	return nil
}

func (t *testPermanentAsset) IsPermanentAsset() {
}

func TestType(t *testing.T) {
	var a Asset
	a = &testAsset{}
	if _, ok := a.(Asset); !ok {
		t.Errorf("test asset does not satisfy Asset interface")
	}

	if _, ok := a.(ResolvableAsset); !ok {
		t.Errorf("test asset does not satisfy ResolvableAsset interface")
	}

	if _, ok := a.(ScriptAsset); ok {
		t.Errorf("test asset satisfies ScriptAsset. it should not")
	}

	if _, ok := a.(PermanentAsset); ok {
		t.Errorf("test asset satisfies PermanentAsset interface. it should not")
	}

	a = &testScriptAsset{}
	if _, ok := a.(ScriptAsset); !ok {
		t.Errorf("test script asset does not satisfy ScriptAsset")
	}

	if _, ok := a.(PermanentAsset); ok {
		t.Errorf("test script asset satisfies PermanentAsset. it should not")
	}

	a = &testPermanentAsset{}
	if _, ok := a.(PermanentAsset); !ok {
		t.Errorf("test permanent asset does not satisfy PermanentAsset")
	}
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
