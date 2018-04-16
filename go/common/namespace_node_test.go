// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"reflect"
	"testing"
)

type testFoo struct {
	a int
}

// Test that assign() works with a pointer to a structure.
func TestNamespaceNode_assign_obj(t *testing.T) {
	var v *testFoo

	r := &namespaceNode{location: RefPathMust("a.b.c"), isPlaceholder: true}

	err := r.bind(reflect.ValueOf(&v).Elem(), &Validation{Type: Validation_OUTPUT})
	if err != nil {
		t.Fatal(err)
	}

	err = r.assign(&testFoo{a: 1})
	if err != nil {
		t.Fatal(err)
	}

	if r.value.Interface().(*testFoo).a != 1 {
		t.Fatal()
	}
}
