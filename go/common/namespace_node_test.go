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

func TestNamespaceNode_bind_availability_nonNil(t *testing.T) {
	var v *testFoo = &testFoo{a: 1}

	r := &namespaceNode{location: RefPathMust("a.b.c"), isPlaceholder: true}

	err := r.bind(reflect.ValueOf(&v).Elem(), &Validation{Type: Validation_OUTPUT})
	if err != nil {
		t.Fatal(err)
	}

	if !r.isValueAvailable {
		t.Fatal("value is not available after binding")
	}
}

func TestNamespaceNode_bind_availability_nil(t *testing.T) {
	var v *testFoo

	r := &namespaceNode{location: RefPathMust("a.b.c"), isPlaceholder: true}

	err := r.bind(reflect.ValueOf(&v).Elem(), &Validation{Type: Validation_OUTPUT})
	if err != nil {
		t.Fatal(err)
	}

	if r.isValueAvailable {
		t.Fatal("value is available after binding a nil object")
	}
}

func TestNamespaceNode_bind_emptyString(t *testing.T) {
	r := &namespaceNode{location: RefPathMust("a.b.c"), isPlaceholder: true}
	err := r.bind(reflect.ValueOf(""), &Validation{Type: Validation_OUTPUT})
	if err != nil {
		t.Fatal(err)
	}
	if r.isValueAvailable {
		t.Fatal("value is available after binding empty string")
	}
}

func TestNamespaceNode_bind_nonEmptyString(t *testing.T) {
	r := &namespaceNode{location: RefPathMust("a.b.c"), isPlaceholder: true}
	err := r.bind(reflect.ValueOf("foo"), &Validation{Type: Validation_OUTPUT})
	if err != nil {
		t.Fatal(err)
	}
	if !r.isValueAvailable {
		t.Fatal("value is not available after binding non-empty string")
	}
}

func TestNamespaceNode_bind_int(t *testing.T) {
	r := &namespaceNode{location: RefPathMust("a.b.c"), isPlaceholder: true}
	err := r.bind(reflect.ValueOf(3), &Validation{Type: Validation_OUTPUT})
	if err != nil {
		t.Fatal(err)
	}
	if !r.isValueAvailable {
		t.Fatal("value is not available after binding non-zero int")
	}
}
