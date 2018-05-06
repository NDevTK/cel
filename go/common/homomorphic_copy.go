// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"github.com/pkg/errors"
	"reflect"
)

// HomomorphicCopy makes potentially shallow copies of on object to another.
// The two objects don't need to be the same type or even cross assignable. It
// will work as long as there are corresponding field names.
//
// For example:
//
//    type A struct {
//      A int
//    }
//
//    type B struct {
//      A int
//    }
//
//    func Foo() {
//      var a A
//      var b B
//      HomomorphicCopy(&a, &b)
//    }
//
// It is not an error if there are extra fields in the `from` value that don't
// match anything on the `to` value. It is also not an error for there to be
// fields in the `to` value that doesn't have any corresponding fields in the
// `from` value.
//
// This function applies itself recursively to any slices, arrays, pointers, or
// structs within the value.
//
// If two corresponding fields in `from` and `to` happen to be assignable, then
// they are assigned rather than copied. This condition applies at the top
// level as well. This, this function isn't guaranteed to make a deep copy.
func HomomorphicCopy(from interface{}, to interface{}) error {
	return homomorphicCopy(reflect.ValueOf(from), reflect.ValueOf(to))
}

func homomorphicCopy(from reflect.Value, to reflect.Value) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case error:
				err = errors.Wrapf(e, "copying %v to %v", from, to)

			case string:
				err = errors.Errorf("copying %v to %v: %s", from, to, e)
			default:
				err = errors.Errorf("failed to copy %v to %v")
			}
		}
	}()

	if !to.CanSet() && from.Kind() == reflect.Ptr && to.Kind() == reflect.Ptr && to.Elem().CanSet() {
		return homomorphicCopy(from.Elem(), to.Elem())
	}

	if !to.CanSet() {
		return errors.Errorf("value cannot be set: %#v", to)
	}

	// Special case:
	//
	//   While generating protobuf messages from discovery document schema, it
	//   turned out that some fields that are optional in the Google API client
	//   libraries can't be marked as opetional in the generated proto files.
	//   It just so happens that there are no distinguishing attributes for
	//   such fields.
	//
	//   For now, we are going to use an ugly hack here. When copying fields
	//   from "A" to "B", if A has a field that has type *Foo and B has field
	//   of type Foo of the same name, then we'll copy the indirected object
	//   from A to B. The same applies when the indirection exists in reverse.
	if from.Kind() == reflect.Ptr && to.Kind() != reflect.Ptr {
		if !from.IsNil() {
			to.Set(from.Elem())
		}
		return nil
	}

	if from.Kind() != reflect.Ptr && to.Kind() == reflect.Ptr {
		to.Set(reflect.New(to.Type().Elem()))
		to.Elem().Set(from)
		return nil
	}

	if from.Type().AssignableTo(to.Type()) {
		to.Set(from)
		return nil
	}

	switch from.Kind() {
	case reflect.Ptr:
		to.Set(reflect.New(to.Type().Elem()))
		return homomorphicCopy(from.Elem(), to.Elem())

	case reflect.Slice:
		to.Set(reflect.MakeSlice(to.Type(), from.Len(), from.Len()))
		fallthrough

	case reflect.Array:
		for i := 0; i < from.Len(); i++ {
			if err := homomorphicCopy(from.Index(i), to.Index(i)); err != nil {
				return err
			}
		}

	case reflect.Struct:
		for i := 0; i < from.NumField(); i++ {
			fn := from.Type().Field(i).Name
			ff := from.Field(i)
			tf := to.FieldByName(fn)
			if !tf.IsValid() {
				continue
			}

			if err := homomorphicCopy(ff, tf); err != nil {
				return err
			}
		}

	default:
		return errors.Errorf("unsupported type %+v vs %+v", from.Kind(), to.Kind())
	}

	return nil
}
