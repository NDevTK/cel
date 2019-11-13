// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	commonpb "chromium.googlesource.com/enterprise/cel/go/schema/common"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"reflect"
	"strings"
	"testing"
)

// TestWalkProto_Order verifies that the order in which fields are visited is
// consistent with what's documented. It also verifies that no node is visited
// twice.
func TestWalkProto_Order(t *testing.T) {
	type Invocation struct {
		V reflect.Value
		P RefPath
		D *descriptor.FieldDescriptorProto
	}

	type Result struct {
		value_type string
		ref_path   string
		field_name string
	}

	Validate := func(t *testing.T, expected []Result, accumulator []Invocation) {
		if len(accumulator) != len(expected) {
			s := ""
			for i := 0; i < len(accumulator); i++ {
				type_name := accumulator[i].V.Type().Name()
				field_name := "(object)"
				if accumulator[i].D != nil {
					field_name = accumulator[i].D.GetName()
				}
				path := accumulator[i].P.String()
				s = s + "Result{\"" + type_name + "\", \"" + path + "\", \"" + field_name + "\"},\n"
			}
			t.Fatalf("unexpected number of invocations: \n%s", s)
		}

		for i := 0; i < len(accumulator); i++ {
			type_name := accumulator[i].V.Type().Name()
			field_name := "(object)"
			if accumulator[i].D != nil {
				field_name = accumulator[i].D.GetName()
			}

			refpath, err := RefPathFromString(expected[i].ref_path)
			if err != nil {
				t.Fatal(err)
			}
			if !accumulator[i].P.Equals(refpath) {
				t.Errorf("unexpected refpath %s at index %d (expected %s)",
					accumulator[i].P.String(), i+1, expected[i].ref_path)
			}

			if expected[i].field_name != field_name {
				t.Errorf("unexpected invocation  at index [%d] field:%s, type:%s", i+1, field_name, type_name)
			}
		}
	}

	t.Run("simple", func(t *testing.T) {
		var accumulator []Invocation

		v := commonpb.TestHasGoodSlice{}
		v.Name = "root"
		v.Field = []*commonpb.TestGoodProto{&commonpb.TestGoodProto{Name: "child1"}, &commonpb.TestGoodProto{Name: "child2"}}

		err := WalkProtoMessage(&v, EmptyPath,
			func(v reflect.Value, p RefPath, d *descriptor.FieldDescriptorProto) (bool, error) {
				accumulator = append(accumulator, Invocation{v, p, d})
				return true, nil
			})

		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}

		if len(accumulator) == 0 {
			t.Fatalf("no invocations recorded")
		}

		expected := []Result{
			Result{"*TestHasGoodSlice", "", "(object)"},
			Result{"string", "name", "name"},
			Result{"[]*TestGoodProto", "field", "field"},
			Result{"*TestGoodProto", "field.child1", "(object)"},
			Result{"string", "field.child1.name", "name"},
			Result{"*TestGoodProto", "field.child2", "(object)"},
			Result{"string", "field.child2.name", "name"},
		}
		Validate(t, expected, accumulator)
	})

	t.Run("skipSome", func(t *testing.T) {
		var accumulator []Invocation

		v := commonpb.TestHasGoodSlice{}
		v.Name = "root"
		v.Field = []*commonpb.TestGoodProto{&commonpb.TestGoodProto{Name: "child1"}, &commonpb.TestGoodProto{Name: "child2"}}

		err := WalkProtoMessage(&v, EmptyPath,
			func(v reflect.Value, p RefPath, d *descriptor.FieldDescriptorProto) (bool, error) {
				accumulator = append(accumulator, Invocation{v, p, d})
				if p.Equals(RefPathMust("field.child2")) {
					return false, nil
				}
				return true, nil
			})

		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}

		if len(accumulator) == 0 {
			t.Fatalf("no invocations recorded")
		}

		expected := []Result{
			Result{"*TestHasGoodSlice", "", "(object)"},
			Result{"string", "name", "name"},
			Result{"[]*TestGoodProto", "field", "field"},
			Result{"*TestGoodProto", "field.child1", "(object)"},
			Result{"string", "field.child1.name", "name"},
			Result{"*TestGoodProto", "field.child2", "(object)"},
		}
		Validate(t, expected, accumulator)
	})

	t.Run("complicated", func(t *testing.T) {
		var accumulator []Invocation

		v := commonpb.TestMessageWithTypes{
			Name:      "myname",
			BoolValue: true,
			IntValue:  3,
			Field:     &commonpb.TestGoodProto{Name: "field-value"},
			RepeatedField: []*commonpb.TestGoodProto{&commonpb.TestGoodProto{Name: "repeated-one"},
				&commonpb.TestGoodProto{Name: "repeated-two"}},
			Optional:  &commonpb.TestMessageWithTypes_OptionalField{OptionalField: &commonpb.TestGoodProto{Name: "opt"}},
			MapField:  map[string]*commonpb.TestGoodProto{"map-key": &commonpb.TestGoodProto{Name: "map-value"}},
			MapString: map[string]string{"map-string-key": "map-string-value"}}

		root := RefPath{"abc", "def"}
		err := WalkProtoMessage(&v, root,
			func(v reflect.Value, p RefPath, d *descriptor.FieldDescriptorProto) (bool, error) {
				accumulator = append(accumulator, Invocation{v, p, d})
				return true, nil
			})

		if err != nil {
			t.FailNow()
		}

		if len(accumulator) == 0 {
			t.FailNow()
		}

		expected := []Result{
			Result{"*TestMessageWithTypes", "abc.def", "(object)"},
			Result{"string", "abc.def.name", "name"},
			Result{"bool", "abc.def.bool_value", "bool_value"},
			Result{"int32", "abc.def.int_value", "int_value"},
			Result{"*TestGoodProto", "abc.def.field", "field"},
			Result{"*TestGoodProto", "abc.def.field", "(object)"},
			Result{"string", "abc.def.field.name", "name"},
			Result{"[]*TestGoodProto", "abc.def.repeated_field", "repeated_field"},
			Result{"*TestGoodProto", "abc.def.repeated_field.repeated-one", "(object)"},
			Result{"string", "abc.def.repeated_field.repeated-one.name", "name"},
			Result{"*TestGoodProto", "abc.def.repeated_field.repeated-two", "(object)"},
			Result{"string", "abc.def.repeated_field.repeated-two.name", "name"},
			Result{"isTestMessageWithTypes_Optional", "abc.def.optional_field", "optional_field"},
			Result{"*TestGoodProto", "abc.def.optional_field", "(object)"},
			Result{"string", "abc.def.optional_field.name", "name"},
			Result{"map[string]*TestGoodProto", "abc.def.map_field", "map_field"},
			Result{"map[string]string", "abc.def.map_string", "map_string"},
		}
		Validate(t, expected, accumulator)
	})
}

func TestWalkProto_ResolvePath(t *testing.T) {
	v := commonpb.TestMessageWithTypes{
		Name:      "myname",
		BoolValue: true,
		IntValue:  3,
		RepeatedField: []*commonpb.TestGoodProto{&commonpb.TestGoodProto{Name: "repeated-one"},
			&commonpb.TestGoodProto{Name: "repeated-two"}},
		Optional:  &commonpb.TestMessageWithTypes_OptionalField{OptionalField: &commonpb.TestGoodProto{Name: "opt"}},
		MapField:  map[string]*commonpb.TestGoodProto{"map-key": &commonpb.TestGoodProto{Name: "map-value"}, "map-empty": nil},
		MapString: map[string]string{"map-string-key": "map-string-value"}}

	t.Run("pod-field", func(t *testing.T) {
		iv, err := Dereference(&v, RefPathFromComponents("abc", "def"), RefPathFromComponents("abc", "def", "name"))
		if err != nil {
			t.Fatal(err)
		}
		if s, ok := iv.(string); !ok || s != "myname" {
			t.Fatal()
		}

		iv, err = Dereference(&v, EmptyPath, RefPathFromComponents("repeated_field", "repeated-one", "name"))
		if err != nil {
			t.Fatal(err)
		}
		if s, ok := iv.(string); !ok || s != "repeated-one" {
			t.Fatal(iv)
		}
	})

	t.Run("regular-field", func(t *testing.T) {
		iv, err := Dereference(&v, EmptyPath, RefPathFromComponents("field"))
		if err != nil {
			t.Fatal(err)
		}
		if v, ok := iv.(*commonpb.TestGoodProto); !ok || v != nil {
			t.Fatalf("unexpected value %#v", iv)
		}
	})

	t.Run("slice-field", func(t *testing.T) {
		iv, err := Dereference(&v, EmptyPath, RefPathFromComponents("repeated_field", "repeated-one"))
		if err != nil {
			t.Fatal(err)
		}
		if s, ok := iv.(*commonpb.TestGoodProto); !ok || s.GetName() != "repeated-one" {
			t.Fatal()
		}
	})

	t.Run("oneof-field", func(t *testing.T) {
		iv, err := Dereference(&v, EmptyPath, RefPathFromComponents("optional_field", "name"))
		if err != nil {
			t.Fatal(err)
		}
		if s, ok := iv.(string); !ok || s != "opt" {
			t.Fatal()
		}
	})

	t.Run("invalid-pod", func(t *testing.T) {
		_, err := Dereference(&v, EmptyPath, RefPathFromComponents("namex"))
		if err == nil {
			t.Fatal()
		}
		if !strings.Contains(err.Error(), "no field named") {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("invalid-slice", func(t *testing.T) {
		_, err := Dereference(&v, EmptyPath, RefPathFromComponents("repeated_field", "xxx"))
		if err == nil {
			t.Fatal()
		}
		if !strings.Contains(err.Error(), "no object named \"xxx\" in collection") {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("nil-object", func(t *testing.T) {
		_, err := Dereference(&v, EmptyPath, RefPathFromComponents("field", "foo"))
		if err == nil {
			t.Fatal()
		}
		if !strings.Contains(err.Error(), "object is nil") {
			t.Fatalf("unexpected error: %#v", err)
		}
	})
}

func TestWalkProto_NamedProtoType(t *testing.T) {
	m := &commonpb.TestMessageWithTypes{}
	v := reflect.ValueOf(m)
	ty := v.Type()
	if !ty.Implements(NamedProtoType) {
		t.Error("&m should implement NamedProtoType")
	}
}
