// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"chromium.googlesource.com/enterprise/cel/go/schema"
	commonpb "chromium.googlesource.com/enterprise/cel/go/schema/common"
	"github.com/golang/protobuf/descriptor"
	"reflect"
	"strings"
	"testing"
)

// TestInvokeValidate_Good goes through a list of cases where we consider the
// protobuf to be valid.
func TestInvokeValidate_Good(t *testing.T) {
	p := RefPath{}
	t.Run("Good", func(t *testing.T) {
		v := commonpb.TestGoodProto{Name: "foo"}
		err := ValidateProto(&v, p)
		if err != nil {
			t.Fatalf("valid proto failed on InvokeValidate: %#v", err)
		}
	})

	t.Run("GoodSlice", func(t *testing.T) {
		v := commonpb.TestHasGoodSlice{Name: "foo"}
		// If the slice is empty, the test will vavuously pass.
		v.Field = []*commonpb.TestGoodProto{&commonpb.TestGoodProto{Name: "foo"}}
		err := ValidateProto(&v, p)
		if err != nil {
			t.Fatalf("valid proto failed on InvokeValidate: %#v", err)
		}
	})

	t.Run("BadSliceIsEmpty", func(t *testing.T) {
		v := commonpb.TestHasBadSlice{Name: "foo"}
		err := ValidateProto(&v, p)
		if err != nil {
			t.Fatalf("valid proto failed on InvokeValidate: %#v", err)
		}
	})

	t.Run("GoodField", func(t *testing.T) {
		v := commonpb.TestHasGoodField{Name: "foo"}
		// If the field is nil, the test will vavuously pass.
		v.Field = &commonpb.TestGoodProto{Name: "foo"}
		err := ValidateProto(&v, p)
		if err != nil {
			t.Fatalf("valid proto failed on InvokeValidate: %#v", err)
		}
	})

	t.Run("GoodOneOf", func(t *testing.T) {
		v := commonpb.TestGoodOneOf{Name: "foo"}
		v.Opt = &commonpb.TestGoodOneOf_Field{Field: &commonpb.TestGoodProto{Name: "foo"}}
		err := ValidateProto(&v, p)
		if err != nil {
			t.Fatalf("valid proto failed on InvokeValidate: %#v", err)
		}
	})

	t.Run("BadOneOfIsEmpty", func(t *testing.T) {
		v := commonpb.TestBadOneOf{Name: "foo"}
		err := ValidateProto(&v, p)
		if err != nil {
			t.Fatalf("valid proto failed on InvokeValidate: %#v", err)
		}
	})

	t.Run("BadFieldIsEmpty", func(t *testing.T) {
		v := commonpb.TestHasBadField{Name: "foo"}
		err := ValidateProto(&v, p)
		if err != nil {
			t.Fatalf("valid proto failed on InvokeValidate: %#v", err)
		}
	})
}

// TestInvokeValidate_Bad goes through a list of cases where the protobuf is
// considered invalid.
func TestInvokeValidate_Bad(t *testing.T) {
	p := RefPath{}
	t.Run("Bad", func(t *testing.T) {
		v := commonpb.TestBadProto{Name: "foo"}
		err := ValidateProto(&v, p)
		if err == nil {
			t.Fatalf("invalid proto succeeded InvokeValidate")
		}
		if !strings.Contains(err.Error(), `No "Validate" method found for type *common.TestBadProto`) {
			t.Fatalf("bad error message: %s", err.Error())
		}
	})

	t.Run("BadArgs", func(t *testing.T) {
		v := commonpb.TestBadValidateArgs{Name: "foo"}
		err := ValidateProto(&v, p)
		if err == nil {
			t.Fatalf("invalid proto succeeded InvokeValidate")
		}
		if !strings.Contains(err.Error(), `No "Validate" method found for type *common.TestBadValidateArgs`) {
			t.Fatalf("bad error message: %s", err.Error())
		}
	})

	t.Run("BadReturnType", func(t *testing.T) {
		v := commonpb.TestBadReturnType{Name: "foo"}
		err := ValidateProto(&v, p)
		if err == nil {
			t.Fatalf("invalid proto succeeded InvokeValidate")
		}
		if !strings.Contains(err.Error(), `No "Validate" method found for type *common.TestBadReturnType`) {
			t.Fatalf("bad error message: %s", err.Error())
		}
	})

	t.Run("BadField", func(t *testing.T) {
		v := commonpb.TestHasBadField{Name: "foo"}
		v.Field = &commonpb.TestBadProto{}
		err := ValidateProto(&v, p)
		if err == nil {
			t.Fatalf("invalid proto succeeded InvokeValidate")
		}
		if !strings.Contains(err.Error(), `No "Validate" method found for type *common.TestBadProto`) {
			t.Fatalf("bad error message: %s", err.Error())
		}
	})

	t.Run("BadSlice", func(t *testing.T) {
		v := commonpb.TestHasBadSlice{Name: "foo"}
		v.Field = []*commonpb.TestBadProto{&commonpb.TestBadProto{}}
		err := ValidateProto(&v, p)
		if err == nil {
			t.Fatalf("invalid proto succeeded InvokeValidate")
		}
		if !strings.Contains(err.Error(), `No "Validate" method found for type *common.TestBadProto`) {
			t.Fatalf("bad error message: %s", err.Error())
		}
	})

	t.Run("BadOneOf", func(t *testing.T) {
		v := commonpb.TestBadOneOf{Name: "foo"}
		v.Opt = &commonpb.TestBadOneOf_Field{Field: &commonpb.TestBadProto{}}
		err := ValidateProto(&v, p)
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
		v := commonpb.TestGoodProto{}
		err := VerifyValidatableType(reflect.TypeOf(&v))
		if err != nil {
			t.Fatalf("valid proto failed on VerifyValidatableType: %#v", err)
		}
	})

	t.Run("GoodSlice", func(t *testing.T) {
		v := commonpb.TestHasGoodSlice{}
		err := VerifyValidatableType(reflect.TypeOf(&v))
		if err != nil {
			t.Fatalf("valid proto failed on VerifyValidatableType: %#v", err)
		}
	})

	t.Run("GoodField", func(t *testing.T) {
		v := commonpb.TestHasGoodField{}
		err := VerifyValidatableType(reflect.TypeOf(&v))
		if err != nil {
			t.Fatalf("valid proto failed on VerifyValidatableType: %#v", err)
		}
	})

	t.Run("GoodOneOf", func(t *testing.T) {
		v := commonpb.TestGoodOneOf{}
		err := VerifyValidatableType(reflect.TypeOf(&v))
		if err != nil {
			t.Fatalf("valid proto failed on VerifyValidatableType: %#v", err)
		}
	})
}

func TestVerifyValidatableType_Bad(t *testing.T) {
	t.Run("Bad", func(t *testing.T) {
		v := commonpb.TestBadProto{}
		err := VerifyValidatableType(reflect.TypeOf(&v))
		if err == nil {
			t.Fatalf("invalid proto succeeded VerifyValidatableType")
		}
		if !strings.Contains(err.Error(), `No "Validate" method found for type *common.TestBadProto`) {
			t.Fatalf("bad error message: %s", err.Error())
		}
	})

	t.Run("BadArgs", func(t *testing.T) {
		err := schema.VerifyValidateFunction(reflect.ValueOf(Validate_TestBadValidateArgs))
		if err == nil {
			t.Fatalf("invalid proto succeeded VerifyValidatableType")
		}
		if !strings.Contains(err.Error(), `incorrect number of arguments`) {
			t.Fatalf("bad error message: %s", err.Error())
		}
	})

	t.Run("BadReturnType", func(t *testing.T) {
		err := schema.VerifyValidateFunction(reflect.ValueOf(Validate_TestBadReturnType))
		if err == nil {
			t.Fatalf("invalid proto succeeded VerifyValidatableType")
		}
		if !strings.Contains(err.Error(), `incorrect return type`) {
			t.Fatalf("bad error message: %s", err.Error())
		}
	})

	t.Run("BadField", func(t *testing.T) {
		v := commonpb.TestHasBadField{}
		err := VerifyValidatableType(reflect.TypeOf(&v))
		if err == nil {
			t.Fatalf("invalid proto succeeded VerifyValidatableType")
		}
		if !strings.Contains(err.Error(), `No "Validate" method found for type *common.TestBadProto`) {
			t.Fatalf("bad error message: %s", err.Error())
		}
	})

	t.Run("BadSlice", func(t *testing.T) {
		v := commonpb.TestHasBadSlice{}
		err := VerifyValidatableType(reflect.TypeOf(&v))
		if err == nil {
			t.Fatalf("invalid proto succeeded VerifyValidatableType")
		}
		if !strings.Contains(err.Error(), `No "Validate" method found for type *common.TestBadProto`) {
			t.Fatalf("bad error message: %s", err.Error())
		}
	})

	t.Run("BadOneOf", func(t *testing.T) {
		v := commonpb.TestBadOneOf{}
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
	tm := commonpb.TestMessageWithOptions{}
	fd, md := descriptor.ForMessage(&tm)
	if fd == nil || md == nil {
		t.Fatalf("can't determine fd and md")
	}

	if len(md.Field) <= 1 || md.Field[0].GetName() != "name" || md.Field[1].GetName() != "key" {
		t.Fatalf(`unexpected TestMessageWithOptions proto.

This test was written based on the expectation that the message was:

    message TestMessageWithOptions {
	  string name = 1;
	  string key = 2 [(common.v).ref="a.b.with_types.repeated_field"];
	  ... // other fields
	}
`)
	}

	v := GetValidationForField(md.Field[1])
	if v.Type != commonpb.Validation_REQUIRED {
		t.Fatalf("failed to query validation information for field \"%s\"", md.Field[1].GetName())
	}
	if v.Ref != "a.b.with_types.repeated_field" {
		t.Fatalf("unexpected key field")
	}
}

func TestValidateOptions(t *testing.T) {
	p := RefPath{}

	// TestMessageWithOptions is annotated with various requirements. First
	// populate it with valid values and see what happens when we replace each
	// valid value with an invalid one, one value at a time.
	v := commonpb.TestMessageWithOptions{
		Name:           "Foo",
		Key:            "Key",
		Label:          "Label",
		OptionalKey:    "Key",
		Fqdn:           "foo.bar.baz",
		Reqd:           "S",
		OptionalString: "x"}
	err := ValidateProto(&v, p)
	if err != nil {
		t.Fatalf("unexpected error %#v", err)
	}

	t.Run("LabelWithInvalidChar", func(t *testing.T) {
		// Label is annotated to require validation.
		w := v
		w.Label = "?"
		err = ValidateProto(&w, p)
		if err == nil {
			t.Fatalf("invalid value succeeded validation")
		}
		if !strings.Contains(err.Error(), "'?' at position 1 is not valid") {
			t.Fatalf("bad error message %s", err.Error())
		}
	})

	t.Run("EmptyLabel", func(t *testing.T) {
		// It can't be empty either.
		w := v
		w.Label = ""
		err = ValidateProto(&w, p)
		if err == nil {
			t.Fatalf("invalid value succeeded validation")
		}
		if !strings.Contains(err.Error(), "labels can't be empty") {
			t.Fatalf("bad error message %s", err.Error())
		}
	})

	t.Run("ForeignKeyRequired", func(t *testing.T) {
		// Key is a foreign key and hence is required.
		w := v
		w.Key = ""
		err = ValidateProto(&w, p)
		if err == nil {
			t.Fatalf("invalid value succeeded validation")
		}
		if !strings.Contains(err.Error(), "required field 'key' was not specified") {
			t.Fatalf("bad error message %s", err.Error())
		}
	})

	t.Run("OptionalForeignKey", func(t *testing.T) {
		// OptionalKey is also a foreign key, but is annotated to be optional.
		w := v
		w.OptionalKey = ""
		err = ValidateProto(&w, p)
		if err != nil {
			t.Fatalf("unexpected error %#v", err)
		}
	})

	t.Run("OptionalString", func(t *testing.T) {
		// Optional string has an extended field option, but not a validation option.
		w := v
		w.OptionalString = ""
		err = ValidateProto(&w, p)
		if err != nil {
			t.Fatalf("unexpected error %#v", err)
		}
	})

	t.Run("FieldCalledName", func(t *testing.T) {
		// Name is not annotated, but is required and must validate as a label by
		// virtue of it being called 'name'.
		w := v
		w.Name = ""
		err = ValidateProto(&w, p)
		if err == nil {
			t.Fatalf("invalid value succeeded validation")
		}
		if !strings.Contains(err.Error(), "labels can't be empty") {
			t.Fatalf("bad error message %s", err.Error())
		}
	})

	t.Run("FqdnInvalid", func(t *testing.T) {
		// FQDN is annotated as requiring validation as a domain name.
		w := v
		w.Fqdn = "a.b.c.?"
		err = ValidateProto(&w, p)
		if err == nil {
			t.Fatalf("invalid value succeeded validation")
		}
		if !strings.Contains(err.Error(), "'?' at position 7 is not valid") {
			t.Fatalf("bad error message %s", err.Error())
		}
	})

	t.Run("ReqdEmpty", func(t *testing.T) {
		// Reqd is annotated as required.
		w := v
		w.Reqd = ""
		err = ValidateProto(&w, p)
		if err == nil {
			t.Fatalf("invalid value succeeded validation")
		}
		if !strings.Contains(err.Error(), "required field 'reqd' was not specified") {
			t.Fatalf("bad error message %s", err.Error())
		}
	})

}
