// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package asset

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	assetpb "chromium.googlesource.com/enterprise/cel/go/schema/asset"
	"reflect"
	"testing"
)

func TestAssetManifest_validateFields(t *testing.T) {
	var a assetpb.AssetManifest
	err := common.ValidateProto(&a, common.EmptyPath)
	if err != nil {
		t.Fatal("unexpected error ", err)
	}

	err = common.VerifyValidatableType(reflect.TypeOf(&a))
	if err != nil {
		t.Fatal("unexpected error ", err)
	}
}

// Tests that all methods in Machine are present in WindowsMachine.
//
// Note: If any new machine types are added to the schema, they should also
// have a <new-machine-type>IsMachine test like this one.
func TestActiveDirectory_windowsMachineIsMachine(t *testing.T) {
	mt := reflect.TypeOf(&assetpb.Machine{})
	wmt := reflect.TypeOf(&assetpb.WindowsMachine{})

	for i := 0; i < mt.NumMethod(); i++ {
		m := mt.Method(i)

		wm, ok := wmt.MethodByName(m.Name)
		if !ok {
			t.Fatalf("Method \"%s\" is present in Machine, but not in WindowsMachine", m.Name)
		}

		if m.Type.NumIn() != wm.Type.NumIn() {
			t.Fatalf("Method \"%s\" in Machine has a different number of arguments than WindowsMachine (%d vs %d, respectively)", m.Name, m.Type.NumIn(), wm.Type.NumIn())
		}

		if m.Type.NumOut() != wm.Type.NumOut() {
			t.Fatalf("Method \"%s\" in Machine has a different number of outputs than WindowsMachine (%d vs %d, respectively)", m.Name, m.Type.NumOut(), wm.Type.NumOut())
		}

		// Starting at 1 since the 0th argument is the receiver which is always different.
		for j := 1; j < m.Type.NumIn(); j++ {
			if !m.Type.In(j).AssignableTo(wm.Type.In(j)) {
				t.Fatalf("Argument #%d in method \"%s\" of Machine is different from WindowsMachine",
					j+1, m.Name)
			}
		}

		for j := 0; j < m.Type.NumOut(); j++ {
			if !m.Type.Out(j).AssignableTo(wm.Type.Out(j)) {
				t.Fatalf("Output #%d in method \"%s\" of Machine is different from WindowsMachine (%#v vs %#v)",
					j+1, m.Name, m.Type.Out(j), wm.Type.Out(j))
			}
		}
	}
}
