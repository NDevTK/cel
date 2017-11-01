// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"github.com/golang/protobuf/descriptor"
	"reflect"
	"strings"
	"testing"
)

// TestInvokeValidate_Good goes through a list of cases where we consider the
// protobuf to be valid.
func TestInvokeValidate_Good(t *testing.T) {
	t.Run("Good", func(t *testing.T) {
		v := TestGoodProto{"foo"}
		err := InvokeValidate(&v)
		if err != nil {
			t.Fatalf("valid proto failed on InvokeValidate: %#v", err)
		}
	})

	t.Run("GoodSlice", func(t *testing.T) {
		v := TestHasGoodSlice{Name: "foo"}
		// If the slice is empty, the test will vavuously pass.
		v.Field = []*TestGoodProto{&TestGoodProto{"foo"}}
		err := InvokeValidate(&v)
		if err != nil {
			t.Fatalf("valid proto failed on InvokeValidate: %#v", err)
		}
	})

	t.Run("BadSliceIsEmpty", func(t *testing.T) {
		v := TestHasBadSlice{Name: "foo"}
		err := InvokeValidate(&v)
		if err != nil {
			t.Fatalf("valid proto failed on InvokeValidate: %#v", err)
		}
	})

	t.Run("GoodField", func(t *testing.T) {
		v := TestHasGoodField{Name: "foo"}
		// If the field is nil, the test will vavuously pass.
		v.Field = &TestGoodProto{"foo"}
		err := InvokeValidate(&v)
		if err != nil {
			t.Fatalf("valid proto failed on InvokeValidate: %#v", err)
		}
	})

	t.Run("GoodOneOf", func(t *testing.T) {
		v := TestGoodOneOf{Name: "foo"}
		v.Opt = &TestGoodOneOf_Field{&TestGoodProto{"foo"}}
		err := InvokeValidate(&v)
		if err != nil {
			t.Fatalf("valid proto failed on InvokeValidate: %#v", err)
		}
	})

	t.Run("BadOneOfIsEmpty", func(t *testing.T) {
		v := TestBadOneOf{Name: "foo"}
		err := InvokeValidate(&v)
		if err != nil {
			t.Fatalf("valid proto failed on InvokeValidate: %#v", err)
		}
	})

	t.Run("BadFieldIsEmpty", func(t *testing.T) {
		v := TestHasBadField{Name: "foo"}
		err := InvokeValidate(&v)
		if err != nil {
			t.Fatalf("valid proto failed on InvokeValidate: %#v", err)
		}
	})
}

// TestInvokeValidate_Bad goes through a list of cases where the protobuf is
// considered invalid.
func TestInvokeValidate_Bad(t *testing.T) {
	t.Run("Bad", func(t *testing.T) {
		v := TestBadProto{"foo"}
		err := InvokeValidate(&v)
		if err == nil {
			t.Fatalf("invalid proto succeeded InvokeValidate")
		}
		if !strings.Contains(err.Error(), `No "Validate" method found for type *common.TestBadProto`) {
			t.Fatalf("bad error message: %s", err.Error())
		}
	})

	t.Run("BadArgs", func(t *testing.T) {
		v := TestBadValidateArgs{Name: "foo"}
		err := InvokeValidate(&v)
		if err == nil {
			t.Fatalf("invalid proto succeeded InvokeValidate")
		}
		if !strings.Contains(err.Error(), `BadValidateArgs has an incorrect type`) {
			t.Fatalf("bad error message: %s", err.Error())
		}
	})

	t.Run("BadReturnType", func(t *testing.T) {
		v := TestBadReturnType{Name: "foo"}
		err := InvokeValidate(&v)
		if err == nil {
			t.Fatalf("invalid proto succeeded InvokeValidate")
		}
		if !strings.Contains(err.Error(), `BadReturnType has an incorrect type`) {
			t.Fatalf("bad error message: %s", err.Error())
		}
	})

	t.Run("BadField", func(t *testing.T) {
		v := TestHasBadField{Name: "foo"}
		v.Field = &TestBadProto{}
		err := InvokeValidate(&v)
		if err == nil {
			t.Fatalf("invalid proto succeeded InvokeValidate")
		}
		if !strings.Contains(err.Error(), `No "Validate" method found for type *common.TestBadProto`) {
			t.Fatalf("bad error message: %s", err.Error())
		}
	})

	t.Run("BadSlice", func(t *testing.T) {
		v := TestHasBadSlice{Name: "foo"}
		v.Field = []*TestBadProto{&TestBadProto{}}
		err := InvokeValidate(&v)
		if err == nil {
			t.Fatalf("invalid proto succeeded InvokeValidate")
		}
		if !strings.Contains(err.Error(), `No "Validate" method found for type *common.TestBadProto`) {
			t.Fatalf("bad error message: %s", err.Error())
		}
	})

	t.Run("BadOneOf", func(t *testing.T) {
		v := TestBadOneOf{Name: "foo"}
		v.Opt = &TestBadOneOf_Field{&TestBadProto{}}
		err := InvokeValidate(&v)
		if err == nil {
			t.Fatalf("invalid proto succeeded InvokeValidate")
		}
		if !strings.Contains(err.Error(), `No "Validate" method found for type *common.TestBadProto`) {
			t.Fatalf("bad error message: %s", err.Error())
		}
	})
}

func TestVerifyValidatableType_Good(t *testing.T) {
	t.Run("Good", func(t *testing.T) {
		v := TestGoodProto{}
		err := VerifyValidatableType(reflect.TypeOf(&v))
		if err != nil {
			t.Fatalf("valid proto failed on VerifyValidatableType: %#v", err)
		}
	})

	t.Run("GoodSlice", func(t *testing.T) {
		v := TestHasGoodSlice{}
		err := VerifyValidatableType(reflect.TypeOf(&v))
		if err != nil {
			t.Fatalf("valid proto failed on VerifyValidatableType: %#v", err)
		}
	})

	t.Run("GoodField", func(t *testing.T) {
		v := TestHasGoodField{}
		err := VerifyValidatableType(reflect.TypeOf(&v))
		if err != nil {
			t.Fatalf("valid proto failed on VerifyValidatableType: %#v", err)
		}
	})

	t.Run("GoodOneOf", func(t *testing.T) {
		v := TestGoodOneOf{}
		err := VerifyValidatableType(reflect.TypeOf(&v))
		if err != nil {
			t.Fatalf("valid proto failed on VerifyValidatableType: %#v", err)
		}
	})
}

func TestVerifyValidatableType_Bad(t *testing.T) {
	t.Run("Bad", func(t *testing.T) {
		v := TestBadProto{}
		err := VerifyValidatableType(reflect.TypeOf(&v))
		if err == nil {
			t.Fatalf("invalid proto succeeded VerifyValidatableType")
		}
		if !strings.Contains(err.Error(), `No "Validate" method found for type *common.TestBadProto`) {
			t.Fatalf("bad error message: %s", err.Error())
		}
	})

	t.Run("BadArgs", func(t *testing.T) {
		v := TestBadValidateArgs{}
		err := VerifyValidatableType(reflect.TypeOf(&v))
		if err == nil {
			t.Fatalf("invalid proto succeeded VerifyValidatableType")
		}
		if !strings.Contains(err.Error(), `BadValidateArgs has an incorrect type`) {
			t.Fatalf("bad error message: %s", err.Error())
		}
	})

	t.Run("BadReturnType", func(t *testing.T) {
		v := TestBadReturnType{}
		err := VerifyValidatableType(reflect.TypeOf(&v))
		if err == nil {
			t.Fatalf("invalid proto succeeded VerifyValidatableType")
		}
		if !strings.Contains(err.Error(), `BadReturnType has an incorrect type`) {
			t.Fatalf("bad error message: %s", err.Error())
		}
	})

	t.Run("BadField", func(t *testing.T) {
		v := TestHasBadField{}
		err := VerifyValidatableType(reflect.TypeOf(&v))
		if err == nil {
			t.Fatalf("invalid proto succeeded VerifyValidatableType")
		}
		if !strings.Contains(err.Error(), `No "Validate" method found for type *common.TestBadProto`) {
			t.Fatalf("bad error message: %s", err.Error())
		}
	})

	t.Run("BadSlice", func(t *testing.T) {
		v := TestHasBadSlice{}
		err := VerifyValidatableType(reflect.TypeOf(&v))
		if err == nil {
			t.Fatalf("invalid proto succeeded VerifyValidatableType")
		}
		if !strings.Contains(err.Error(), `No "Validate" method found for type *common.TestBadProto`) {
			t.Fatalf("bad error message: %s", err.Error())
		}
	})

	t.Run("BadOneOf", func(t *testing.T) {
		v := TestBadOneOf{}
		err := VerifyValidatableType(reflect.TypeOf(&v))
		if err == nil {
			t.Fatalf("invalid proto succeeded VerifyValidatableType")
		}
		if !strings.Contains(err.Error(), `No "Validate" method found for type *common.TestBadProto`) {
			t.Fatalf("bad error message: %s", err.Error())
		}
	})
}

func TestExtractExtension(t *testing.T) {
	tm := TestMessageWithOptions{}
	fd, md := descriptor.ForMessage(&tm)
	if fd == nil || md == nil {
		t.Fatalf("can't determine fd and md")
	}

	if len(md.Field) == 0 || md.Field[1].GetName() != "key" {
		t.Fatalf(`unexpected TestMessageWithOptions proto.

This test was written based on the expectation that the message was:

    message TestMessageWithOptions {
	  string name = 1;
	  string key = 2 [(common.v)={type:ASSET, key:"foreign-key"}];
	  ... // other fields
	}
`)
	}

	v := getValidationForField(md.Field[1])
	if v.Type == Validation_REQUIRED {
		t.Fatalf("failed to query validation information for field")
	}
	if v.Type != Validation_ASSET {
		t.Fatalf("unexpected validation type")
	}
	if v.Key != "foreign-key" {
		t.Fatalf("unexpected key field")
	}
}

func TestValidateOptions(t *testing.T) {
	// TestMessageWithOptions is annotated with various requirements. First
	// populate it with valid values and see what happens when we replace each
	// valid value with an invalid one, one value at a time.
	v := TestMessageWithOptions{
		Name:        "Foo",
		Key:         "Key",
		Label:       "Label",
		OptionalKey: "Key",
		Fqdn:        "foo.bar.baz",
		Reqd:        "S"}
	err := InvokeValidate(&v)
	if err != nil {
		t.Fatalf("unexpected error %#v", err)
	}

	// Label is annotated to require validation.
	w := v
	w.Label = "?"
	err = InvokeValidate(&w)
	if err == nil {
		t.Fatalf("invalid value succeeded validation")
	}
	if !strings.Contains(err.Error(), "'?' at position 1 is not valid") {
		t.Fatalf("bad error message %s", err.Error())
	}

	// It can't be empty either.
	w = v
	w.Label = ""
	err = InvokeValidate(&w)
	if err == nil {
		t.Fatalf("invalid value succeeded validation")
	}
	if !strings.Contains(err.Error(), "labels can't be empty") {
		t.Fatalf("bad error message %s", err.Error())
	}

	// Key is a foreign key and hence is required.
	w = v
	w.Key = ""
	err = InvokeValidate(&w)
	if err == nil {
		t.Fatalf("invalid value succeeded validation")
	}
	if !strings.Contains(err.Error(), "field \"key\" is a reference and cannot be empty") {
		t.Fatalf("bad error message %s", err.Error())
	}

	// OptionalKey is also a foreign key, but is annotated to be optional.
	w = v
	w.OptionalKey = ""
	err = InvokeValidate(&w)
	if err != nil {
		t.Fatalf("unexpected error %#v", err)
	}

	// Name is not annotated, but is required and must validate as a label by
	// virtue of it being called 'name'.
	w = v
	w.Name = ""
	err = InvokeValidate(&w)
	if err == nil {
		t.Fatalf("invalid value succeeded validation")
	}
	if !strings.Contains(err.Error(), "labels can't be empty") {
		t.Fatalf("bad error message %s", err.Error())
	}

	// FQDN is annotated as requiring validation as a domain name.
	w = v
	w.Fqdn = "a.b.c.?"
	err = InvokeValidate(&w)
	if err == nil {
		t.Fatalf("invalid value succeeded validation")
	}
	if !strings.Contains(err.Error(), "'?' at position 7 is not valid") {
		t.Fatalf("bad error message %s", err.Error())
	}

	// Reqd is annotated as required.
	w = v
	w.Reqd = ""
	err = InvokeValidate(&w)
	if err == nil {
		t.Fatalf("invalid value succeeded validation")
	}
	if !strings.Contains(err.Error(), "required field 'reqd' was not specified") {
		t.Fatalf("bad error message %s", err.Error())
	}
}
