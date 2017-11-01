// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"reflect"
	"testing"
)

func TestWalkProto_Order(t *testing.T) {
	v := TestHasGoodSlice{}
	v.Name = "root"
	v.Field = []*TestGoodProto{&TestGoodProto{"child1"}, &TestGoodProto{"child2"}}

	type Invocation struct {
		V reflect.Value
		D *descriptor.FieldDescriptorProto
	}

	var accumulator []Invocation

	err := WalkProto(reflect.ValueOf(&v), func(v reflect.Value, d *descriptor.FieldDescriptorProto) error {
		accumulator = append(accumulator, Invocation{v, d})
		return nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}

	if len(accumulator) == 0 {
		t.Fatalf("no invocations recorded")
	}

	type Result struct {
		value_type string
		field_name string
	}
	expected := []Result{
		Result{"string", "name"},
		Result{"string", "name"},
		Result{"*TestGoodProto", "(object)"},
		Result{"string", "name"},
		Result{"*TestGoodProto", "(object)"},
		Result{"[]*TestGoodProto", "field"},
		Result{"*TestHasGoodSlice", "(object)"}}

	if len(accumulator) != len(expected) {
		t.Fatalf("unexepcted number of invocations")
	}

	for i := 0; i < len(accumulator); i++ {
		type_name := getTypeString(accumulator[i].V.Type())
		field_name := "(object)"
		if accumulator[i].D != nil {
			field_name = accumulator[i].D.GetName()
		}

		if expected[i].field_name != field_name {
			t.Errorf("unexpected invocation  at index [%d] field:%s, type:%s", i+1, field_name, type_name)
		}
	}
}
