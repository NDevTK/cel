// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	pd "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

// WalkProtoFunc is invoked for each field found in a data structure
// representing a proto.Message. Take care to examine the reflect.Value that's
// passed in since the function can be invoked for pointer indirection levels
// in addition to per-field.
//
// E.g.: Say we have the following proto message:
//
//     message Foo {
//       repeated Bar baz = 1;
//     }
//
//     message Bar {
//       string quux = 1;
//     }
//
// ... which generates code that looks like:
//
//     type Foo struct {
//       Baz []*Bar
//     }
//
//     type Bar struct {
//       Quux string
//     }
//
// Assuming |f| is a WalkProtoFunc, when WalkProto is invoked on |v| which is a
// *Foo object with one Bar object in the Baz field, the following invocations
// will happen (in order):
//
//     f(value_of(v.Baz[0].Quux), descriptor_of(Bar.Quux))
//     f(value_of(v.Baz[0]), nil)
//     f(value_of(v), nil)
//
// For convenience value_of(x) means reflect.ValueOf(x), descriptor_of(x) is
// the descriptor.FieldDescriptorProto for the specified field.
type WalkProtoFunc func(reflect.Value, *pd.FieldDescriptorProto) error

// WalkProto walks through the fields of a proto.Message recursively invoking
// WalkProtoFunc at each field. See WalkProtoFunc() for more details.
func WalkProto(av reflect.Value, f WalkProtoFunc) error {
	err_list := []error{}

	if !av.IsValid() {
		return nil
	}

	switch av.Kind() {
	case reflect.Slice:
		// The field is of the form:
		//     FieldName []*InnerType
		for i := 0; i < av.Len(); i++ {
			err_list = AppendErrorList(err_list, WalkProto(av.Index(i), f))
		}

	case reflect.Ptr:
		// Nothing to do if field is Nil. This effectively considers all
		// embedded messages as optional. Any required fields should be checked
		// for in the Validate() method of the enclosing type.
		if av.IsNil() {
			return nil
		}
		// |av| is a *Foo value. Recurse on Foo prior to inspecting *Foo. |f|
		// will see the two as two distinct types.
		err_list = AppendErrorList(err_list, WalkProto(av.Elem(), f))
		if av.Elem().Kind() == reflect.Struct {
			err_list = AppendErrorList(err_list, walkPtrToStruct(av, f))
		}

	case reflect.Struct:
		err_list = walkStruct(av, f)

	case reflect.Interface:
		if av.IsNil() || !av.CanInterface() {
			return nil
		}

		// For a non-empty field, this also covers the 'oneof' case.
		return WalkProto(av.Elem(), f)

	default:
		return nil
	}

	return errors.Wrapf(WrapErrorList(err_list), "type \"%s\"", getTypeString(av.Type()))
}

// validatePtrToStruct performs validation on a reflect.Value on the assumption
// that it is a Ptr to a Struct. In practical terms |av| is a Value
// representing a *FooProto where FooProto is a generated protubuf structure.
func walkPtrToStruct(av reflect.Value, f WalkProtoFunc) error {
	// Ignore if *InnerType can't be converted to a proto.Message. There are
	// additional types that are relevant, but those are handled in WalkProto().
	if !av.Type().Implements(reflect.TypeOf((*proto.Message)(nil)).Elem()) {
		return nil
	}

	return f(av, nil)
}

// walkStruct iterates over the fields of a struct and recursively invokes
// WalkProto() on them.
func walkStruct(av reflect.Value, f WalkProtoFunc) (err_list []error) {
	fpm := constructFieldToDescriptorMap(av)

	for i := 0; i < av.NumField(); i++ {
		err_list = AppendErrorList(err_list, WalkProto(av.Field(i), f))

		fd, ok := fpm[av.Type().Field(i).Name]
		if !ok {
			continue
		}

		err := f(av.Field(i), fd)
		err_list = AppendErrorList(err_list,
			errors.Wrapf(err, "field \"%s\" (proto field \"%s\")",
				av.Type().Field(i).Name, fd.GetName()))
	}
	return
}

// constructFieldToDescriptorMap constructs a map from field names to
// *FieldDescriptorProto. On entry |av| is a value representing a *Struct*.
func constructFieldToDescriptorMap(av reflect.Value) (m map[string]*pd.FieldDescriptorProto) {
	// m is a map from generated Go field name to FieldDescriptoProto
	m = make(map[string]*pd.FieldDescriptorProto)

	// First make sure we are dealing with a proto message. pd.Message is
	// proto.Message + Descriptor().
	msg, ok := av.Addr().Interface().(descriptor.Message)
	if !ok {
		return
	}

	// This invocation isn't super efficient since it extracts and unzips the
	// file descriptor for every field and message in a file. But we aren't
	// given a public(ish) API that's any more efficient than this :-(.
	_, md := descriptor.ForMessage(msg)

	// fm is a map from proto field name to FieldDescriptorProto
	proto_m := make(map[string]*pd.FieldDescriptorProto)
	for _, f := range md.Field {
		proto_m[f.GetName()] = f
	}

	for i := 0; i < av.NumField(); i++ {
		ft := av.Type().Field(i)

		ps, ok := ft.Tag.Lookup("protobuf")
		if !ok {
			continue
		}

		// The proto: tag for generated Go fields looks like:
		//     "varint,3,opt,name=number"
		fl := strings.Split(ps, ",")
		for _, s := range fl {
			if !strings.HasPrefix(s, "name=") {
				continue
			}

			if fd, ok := proto_m[s[5:]]; ok {
				m[ft.Name] = fd
			}
			break
		}
	}

	return
}

func getTypeString(ty reflect.Type) string {
	s := ty.Name()
	if s != "" {
		return s
	}
	switch ty.Kind() {
	case reflect.Ptr:
		return "*" + getTypeString(ty.Elem())

	case reflect.Slice:
		return "[]" + getTypeString(ty.Elem())

	case reflect.Map:
		return "map[" + getTypeString(ty.Key()) + "]" + getTypeString(ty.Elem())
	}
	return "(unknown)"
}
