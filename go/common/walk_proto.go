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

// NamedProto is a proto.Message that has a GetName() method. This is a
// well-known pattern for identifying self naming objects.
type NamedProto interface {
	proto.Message
	GetName() string
}

// WalkProtoFunc is invoked for each field found in a data structure
// representing a proto.Message. See WalkProto for more details.
type WalkProtoFunc func(reflect.Value, RefPath, *pd.FieldDescriptorProto) error

// WalkProtoMessage is a convenience function for invoking WalkProto using a
// protom.Message directly. It only exists so that call sites don't need a
// direct dependency on "reflect", which in retrospect is kind of silly since
// the WalkProtoFunc has more dependencies than WalkProto.
func WalkProtoMessage(m proto.Message, p RefPath, f WalkProtoFunc) error {
	return WalkProto(reflect.ValueOf(m), p, f)
}

// WalkProto walks through the fields of a proto.Message recursively invoking
// WalkProtoFunc at each field.
//
// For each field, the WalkProtoFunc is invoked with the following:
//
//     reflect.Value : The value of the field.
//     RefPath       : A reference path leading up to and including the field.
//                     See RefPath for more information on reference paths.
//     *pd.FieldDescriptorProto : The field descriptor for the field.
//
// Take care to examine the reflect.Value that's passed in since the function
// can be invoked for pointer indirection levels in addition to per-field.
//
// E.g.: Say we have the following proto message:
//
//     message Foo {
//       string name = 1;
//       repeated Bar baz = 2;
//     }
//
//     message Bar {
//       string name = 1;
//       string quux = 2;
//     }
//
// ... which generates code that looks like:
//
//     type Foo struct {
//       Name string
//       Baz []*Bar
//     }
//
//     type Bar struct {
//       Name string
//       Quux string
//     }
//
// Assuming |f| is a WalkProtoFunc, when WalkProto is invoked on |v| which is a
// *Foo object with one Bar object in the Baz field, the following invocations
// will happen (in order):
//
//     f(value_of(v), RefPath{v.Name}, nil)
//     f(value_of(v.Name), RefPath{v.Name, "name"}, descriptor_of(Foo.Name))
//     f(value_of(v.Baz), RefPath{v.Name, "baz"}, descriptor_of(Foo.Baz))
//     f(value_of(v.Baz[0]), RefPath{v.Name, "baz", b.Baz[0].Name}, nil)
//     f(value_of(v.Baz[0].Name), RefPath{v.Name, "baz", b.Baz[0].Name, "name"}, descriptor_of(Bar.Name))
//     f(value_of(v.Baz[0].Quux), RefPath{v.Name, "baz", b.Baz[0].Name, "quux"}, descriptor_of(Bar.Quux))
//
// For convenience value_of(x) is assumed to mean reflect.ValueOf(x),
// descriptor_of(x) is the descriptor.FieldDescriptorProto for the specified
// field.
func WalkProto(av reflect.Value, p RefPath, f WalkProtoFunc) error {
	err_list := []error{}

	if !av.IsValid() {
		return nil
	}

	switch av.Kind() {
	case reflect.Array:
		fallthrough

	case reflect.Slice:
		// The field is expected to be of the form:
		//     FieldName []*InnerType
		for i := 0; i < av.Len(); i++ {
			np := p
			if name, ok := nameOfNamedProto(av.Index(i)); ok {
				np = np.Append(name)
			}
			err_list = AppendErrorList(err_list, WalkProto(av.Index(i), np, f))
		}

	case reflect.Ptr:
		// Nothing to do if field is Nil. This effectively considers all
		// embedded messages as optional. Any required fields should be checked
		// for in the Validate() method of the enclosing type.
		if av.IsNil() {
			return nil
		}
		if av.Elem().Kind() == reflect.Struct {
			err_list = AppendErrorList(err_list, walkPtrToStruct(av, p, f))
		}
		err_list = AppendErrorList(err_list, WalkProto(av.Elem(), p, f))

	case reflect.Struct:
		err_list = walkStruct(av, p, f)

	case reflect.Interface:
		if av.IsNil() || !av.CanInterface() {
			return nil
		}

		// For a non-empty field, this also covers the 'oneof' case.
		return WalkProto(av.Elem(), p, f)

	default:
		return nil
	}

	return errors.Wrapf(WrapErrorList(err_list), "\ntype \"%s\"", getTypeString(av.Type()))
}

// validatePtrToStruct invokes |f| on |av| if |av| is a proto.Message. This is
// where |f| is invoked at the message level. See walkStruct() for where |f| is
// invoked at the field level.
func walkPtrToStruct(av reflect.Value, p RefPath, f WalkProtoFunc) error {
	// Ignore if *InnerType can't be converted to a proto.Message. There are
	// additional types that are relevant, but those are handled in WalkProto().
	if !av.Type().Implements(reflect.TypeOf((*proto.Message)(nil)).Elem()) {
		return nil
	}

	return f(av, p, nil)
}

// walkStruct iterates over the fields of a struct and recursively invokes
// WalkProto() on them.
func walkStruct(av reflect.Value, p RefPath, f WalkProtoFunc) (err_list []error) {
	fpm := constructFieldToDescriptorMap(av)

	for i := 0; i < av.NumField(); i++ {
		fp := p
		fd, ok := fpm[av.Type().Field(i).Name]
		if ok {
			fp = fp.Append(fd.GetName())
			err := f(av.Field(i), fp, fd)
			err_list = AppendErrorList(err_list,
				errors.Wrapf(err, "field \"%s\" (proto field \"%s\")",
					av.Type().Field(i).Name, fd.GetName()))
		}

		err_list = AppendErrorList(err_list, WalkProto(av.Field(i), fp, f))
	}
	return
}

// constructFieldToDescriptorMap constructs a map from Go field names to
// *FieldDescriptorProto. On entry |av| is a value representing a *Struct. As a
// special case, if a oneof field is encountered that is not nil, this function
// elevates the concrete underlying oneof field descriptor to the top level one
// of field.
//
// In case that wasn't clear, let's look at an example:
//
// Let's say we have a Go struct corresponding to a proto message with a one-of
// field. One such example is TestMessageWithTypes which is partially
// reproduced below:
//
//     type TestMessageWithTypes struct {
//         ...
//     	   // Types that are valid to be assigned to Optional:
//     	   //	*TestMessageWithTypes_OptionalField
//     	   Optional  isTestMessageWithTypes_Optional `protobuf_oneof:"optional"`
//         ...
//     }
//
// The TestMessageWithTypes_OptionalField struct looks like this:
//
//     type TestMessageWithTypes_OptionalField struct {
//     	   OptionalField *TestGoodProto `protobuf:"bytes,6,opt,name=optional_field,json=optionalField,oneof"`
//     }
//
// If we are looking at an instance of TestMessageWithTypes where the
// 'Optional' field is non-nil, this function creates a mapping from "Optional"
// to the FieldDescriptorProto for "optional_field".
//
// If 'Optional' were nil, then in the general case we wouldn't know what type
// 'Optional' is. Therefore no mapping would be created.
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

	// proto_m is a map from proto field name to FieldDescriptorProto
	proto_m := make(map[string]*pd.FieldDescriptorProto)
	for _, f := range md.Field {
		proto_m[f.GetName()] = f
	}

	for i := 0; i < av.NumField(); i++ {
		ps, ok := getProtobufTagForField(av.Field(i), av.Type().Field(i))
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
				m[av.Type().Field(i).Name] = fd
			}
			break
		}
	}

	return
}

// getProtobufTagForField returns the "protobuf" field annotation tag
// corresponding to the struct field value specified by |fv| and |ft|.
func getProtobufTagForField(fv reflect.Value, ft reflect.StructField) (string, bool) {
	ps, ok := ft.Tag.Lookup("protobuf")
	if ok {
		return ps, true
	}

	// could be a oneof field.
	_, ok = ft.Tag.Lookup("protobuf_oneof")
	if !ok {
		return "", false
	}

	// Ignore nil one-of fields. We don't know what type they are.
	if fv.IsNil() {
		return "", false
	}

	// Also don't know what to do if av.Field(i) is not an interface
	// that's implemented by a Ptr to a Struct.
	if fv.Elem().Type().Kind() != reflect.Ptr ||
		fv.Elem().Type().Elem().Kind() != reflect.Struct {
		return "", false
	}

	opt_t := fv.Elem().Type().Elem()
	// Generated optional structs should only have one field.
	if opt_t.NumField() != 1 {
		panic("oneof field points to struct that doesn't have a unique field")
	}

	return opt_t.Field(0).Tag.Lookup("protobuf")
}

// getNamedProtoField returns the Go field corresponding to the proto field
// named |name|.
func getNamedProtoField(av reflect.Value, name string) (reflect.Value, bool) {
	namefield := "name=" + name

	for i := 0; i < av.NumField(); i++ {
		ps, ok := getProtobufTagForField(av.Field(i), av.Type().Field(i))
		if !ok {
			continue
		}

		for _, s := range strings.Split(ps, ",") {
			if s == namefield {
				return av.Field(i), true
			}
		}
	}
	return av, false
}

// getTypeString returns a string identifying the underlying Go type of |ty|.
func getTypeString(ty reflect.Type) string {
	s := ty.Name()
	if s != "" {
		return s
	}
	switch ty.Kind() {
	case reflect.Array:
		return "[" + string(ty.Len()) + "]" + getTypeString(ty.Elem())

	case reflect.Ptr:
		return "*" + getTypeString(ty.Elem())

	case reflect.Slice:
		return "[]" + getTypeString(ty.Elem())

	case reflect.Map:
		return "map[" + getTypeString(ty.Key()) + "]" + getTypeString(ty.Elem())
	}
	return "(unknown)"
}

// ResolvePath descends into a proto.Message and attempts to resolve a propery
// reference chain. See RefPath for more details.
func ResolvePath(m proto.Message, root RefPath, path RefPath, mode ResolutionMode) (interface{}, error) {
	path, ok := path.After(root)
	if !ok {
		return nil, errors.Errorf("can't look up path %s in %s with root %s",
			path.String(), m.String(), root.String())
	}
	root = EmptyPath

	av, err := resolvePathInValue(reflect.ValueOf(m), path, mode)
	if err != nil {
		return nil, err
	}
	return av.Interface(), nil
}

var OutputFieldSkippedError = errors.New("unresolved output field skipped")

// nameOfNamedProto returns the "name" field of a proto.Message represented by
// a reflect.Value. If the value passed-in doesn't have a "name" field, returns
// "",false. Otherwise returns the name.
func nameOfNamedProto(av reflect.Value) (string, bool) {
	if av.Type().Implements(reflect.TypeOf((*NamedProto)(nil)).Elem()) {
		return av.Interface().(NamedProto).GetName(), true
	}
	return "", false
}

// resolvePathInValue resolves |path| relative to |av|.
func resolvePathInValue(av reflect.Value, path RefPath, mode ResolutionMode) (reflect.Value, error) {
	if len(path) == 0 {
		return av, nil
	}

	switch av.Kind() {
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		// The name refers to a named object
		name, np := path.Shift()
		path = np
		for i := 0; i < av.Len(); i++ {
			if n, ok := nameOfNamedProto(av.Index(i)); ok && n == name {
				return resolvePathInValue(av.Index(i), path, mode)
			}
		}
		return av, errors.Errorf("path not found \"%s\". no object named \"%s\" in collection",
			path.String(), name)

	case reflect.Ptr:
		if av.IsNil() {
			return av, errors.Errorf("path not found \"%s\". object is nil.", path.String())
		}
		if av.Type().Implements(reflect.TypeOf((*descriptor.Message)(nil)).Elem()) {
			_, md := descriptor.ForMessage(av.Interface().(descriptor.Message))
			for _, f := range md.Field {
				if f.GetName() == path[0] {
					if mode == ResolutionSkipOutputs && f.GetOptions() != nil {
						v := getValidationForField(f)
						if v.Type == Validation_OUTPUT {
							return av, OutputFieldSkippedError
						}
					}
				}
			}
		}
		return resolvePathInValue(av.Elem(), path, mode)

	case reflect.Interface:
		if av.IsNil() {
			return av, errors.Errorf("path not found \"%s\". oneof is nil.", path.String())
		}
		if av.Elem().Kind() != reflect.Ptr || av.Elem().IsNil() ||
			av.Elem().Elem().Kind() != reflect.Struct || av.Elem().Elem().NumField() != 1 {
			return av, errors.Errorf("path not found \"%s\". one of is invalid.", path.String())
		}
		return resolvePathInValue(av.Elem().Elem().Field(0), path, mode)

	case reflect.Struct:
		name, path := path.Shift()
		av, ok := getNamedProtoField(av, name)
		if !ok {
			return av, errors.Errorf("path not found \"%s.%s\". no field named \"%s\"",
				name, path.String(), name)
		}
		return resolvePathInValue(av, path, mode)
	}
	return av, errors.Errorf("path not found: %s", path.String())
}
