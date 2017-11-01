// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/pkg/errors"
	"reflect"
)

const kValidateMethodName = "Validate"

const kNoMethodFoundError = `No "Validate" method found for type %s.

See documentation at //docs/schema-guidelines.md for more information about
adding a 'Validate' method for this type.
`

const kInvalidSignatureError = `The "Validate" method for type %s has an incorrect type.

Expected:
    func (%s) Validate() error

Actual:
    %s
`

// InvokeValidate analyzes the data structure at |a| and invokes the Validate()
// method (see Validator interface) on objects that implement both
// proto.Message and Validator.
//
// In other words, it recursively invokes Validate() on proto objects rooted at
// |a|.
//
// If |a| is a pointer to a structure, (say, *Foo), and |a| implements
// message.Proto, then Foo is assumed to be a protobuf generated message. By
// convention there should be a Validate() method for *Foo (i.e. a function
// matching the signature 'func (*Foo) Validate() error'). This function is
// expected to do basic validation on the data as a means of providing an early
// warning if something is out of line.
//
// *Foo -- or Foo -- might contain fields that are themselves protobuf
// messages, in which case InvokeValidate() recurses into each field to
// validate the entire data structure.
//
// See //docs/schema-guidelines.md for more details on why we are doing things
// this way.
func InvokeValidate(a interface{}) error {
	return WalkProto(reflect.ValueOf(a), invokeValidateOnValue)
}

// invokeValidateOnValue is a WalkProtoFunc that invokes the Validate() method
// if one is found for a proto.Message object. It returns an error if the class
// is missing a Validate() method as specified in the Validator interface.
func invokeValidateOnValue(av reflect.Value, d *descriptor.FieldDescriptorProto) error {
	// A field
	if d != nil {
		return validateAnnotatedField(av, d)
	}

	// The only field type we are interested in is of the form:
	//     FooField *FooType
	// ... where *FooType implements proto.Message.
	if av.Kind() != reflect.Ptr || !av.Type().Implements(reflect.TypeOf((*proto.Message)(nil)).Elem()) {
		return nil
	}

	// Skip nil fields. If the field was required, then the prior validation
	// step would've caught it.
	if av.IsNil() {
		return nil
	}

	mv := av.MethodByName(kValidateMethodName)
	if !mv.IsValid() {
		return fmt.Errorf(kNoMethodFoundError, av.Type().String())
	}

	if mv.Type().NumIn() != 0 || mv.Type().NumOut() != 1 || mv.Type().Out(0).Kind() != reflect.Interface {
		return fmt.Errorf(kInvalidSignatureError, av.Type().String(), av.Type().String(), mv.Type().String())
	}

	result := mv.Call([]reflect.Value{})
	if len(result) != 1 {
		return fmt.Errorf(kInvalidSignatureError, av.Type().String(), av.Type().String(), mv.Type().String())
	}

	if result[0].IsNil() {
		return nil
	}

	err, ok := result[0].Interface().(error)
	if !ok {
		return fmt.Errorf(kInvalidSignatureError, av.Type().String(), av.Type().String(), mv.Type().String())
	}
	return err
}

// getValidationForField queries the extensions for |fd| and returns a
// Validation object if one is found. If there is no Validation extensino for
// this field, returns an object representing the default set of validation
// rules that should be applied to |fd|.
func getValidationForField(fd *descriptor.FieldDescriptorProto) Validation {
	v := Validation{}
	v.Optional = true

	if fd.Options != nil {
		ex, _ := proto.GetExtension(fd.Options, E_V)
		if e, ok := ex.(*Validation); ok {
			v = *e
		}
	}

	if fd.GetName() == "name" && v.Type == Validation_REQUIRED {
		v.Type = Validation_LABEL
		v.Optional = false
	}
	return v
}

// validateAnnotatedField performs validator for a field which has one or more
// of the annotations in //schema/common/options.proto.
func validateAnnotatedField(af reflect.Value, fd *descriptor.FieldDescriptorProto) error {
	v := getValidationForField(fd)

	// Skip over optional fields if empty.
	if v.Optional {
		switch {
		case af.Kind() == reflect.String && af.Len() == 0:
			fallthrough
		case af.Kind() == reflect.Slice && af.Len() == 0:
			fallthrough
		case af.Kind() == reflect.Ptr && af.IsNil():
			fallthrough
		case af.Kind() == reflect.Interface && af.IsNil():
			return nil
		}
	}

	switch v.Type {
	case Validation_REQUIRED:
		switch {
		case af.Kind() == reflect.String && af.Len() == 0:
			return fmt.Errorf("required field '%s' was not specified", fd.GetName())
		case (af.Kind() == reflect.Interface || af.Kind() == reflect.Ptr) && af.IsNil():
			return fmt.Errorf("'%s' is required", fd.GetName())
		case (af.Kind() == reflect.Map || af.Kind() == reflect.Slice) && af.Len() == 0:
			return fmt.Errorf("at least one of '%s' should be specified", fd.GetName())
		}

	case Validation_OUTPUT:
		if af.Kind() != reflect.String {
			return fmt.Errorf("field \"%s\" is marked as output, but is not a string", fd.GetName())
		}
		if af.Len() != 0 {
			return fmt.Errorf("field \"%s\" is marked as output, but is not empty", fd.GetName())
		}

	case Validation_ASSET, Validation_HOST:
		if af.Kind() != reflect.String {
			return fmt.Errorf("field \"%s\" is a foreign key, but is not a string", fd.GetName())
		}
		if af.Len() == 0 {
			return fmt.Errorf("field \"%s\" is a reference and cannot be empty", fd.GetName())
		}

	case Validation_LABEL:
		if af.Kind() != reflect.String {
			return fmt.Errorf("unexpected type for label field: %#v", af)
		}

		return errors.Wrapf(IsRFC1035Label(af.String()), "validating field \"%s\"", fd.GetName())

	case Validation_FQDN:
		if af.Kind() != reflect.String {
			return fmt.Errorf("unexpected kind for validatable field: %#v", af)
		}
		return errors.Wrapf(IsRFC1035Domain(af.String()), "validating field \"%s\"", fd.GetName())
	}
	return nil
}

func VerifyValidatableType(at reflect.Type) error {
	err_list := []error{}

	switch at.Kind() {
	case reflect.Slice:
		return VerifyValidatableType(at.Elem())

	case reflect.Ptr:
		// |at| is a Type for *Foo. First dive down and examine Foo.
		err_list = AppendErrorList(err_list, VerifyValidatableType(at.Elem()))

		// In case there were 'oneof' fields...
		err_list = AppendErrorList(err_list, verifyOneOfTypes(at))

		// If *Foo doesn't implement proto.Message, assume this is something
		// else and don't look any further.
		if !at.Implements(reflect.TypeOf((*proto.Message)(nil)).Elem()) {
			break
		}

		// *Foo *should* have a Validate method since it implements proto.Message.
		mt, ok := at.MethodByName(kValidateMethodName)
		if !ok {
			err_list = AppendErrorList(err_list, fmt.Errorf(kNoMethodFoundError, at.String()))
			break
		}

		// We can only check the method signature since we don't have a value
		// to work with.
		if mt.Type.NumIn() != 1 || mt.Type.NumOut() != 1 || mt.Type.Out(0) != reflect.TypeOf((*error)(nil)).Elem() {
			err_list = AppendErrorList(err_list,
				fmt.Errorf(kInvalidSignatureError, at.Elem().String(), at.Elem().String(), mt.Type.String()))
			break
		}

	case reflect.Struct:
		for i := 0; i < at.NumField(); i++ {
			sf := at.Field(i)
			err_list = AppendErrorList(err_list, VerifyValidatableType(sf.Type))
		}
	}

	return WrapErrorList(err_list)
}

// verifyOneOfTypes attempts to verify related subtypes that are used for
// implementing 'oneof' support in ProtoBufs.
//
// Here's where things get tricky. The only connection we have from |at| to the
// candidates oneof field subtypes is via the XXX_OneofFuncs() internal proto
// method. It is likely that this will break in the future. There are panics
// scattered around to catch minor deviations, but won't catch major
// deviations.
//
// |pt| is a |Ptr| type which we suspect may contain 'oneof' fields.
func verifyOneOfTypes(pt reflect.Type) error {
	mt, ok := pt.MethodByName("XXX_OneofFuncs")
	if !ok {
		// I guess it wasn't a oneof. Or, y'know, proto decided to change
		// this implementation detail.
		return nil
	}

	results := mt.Func.Call([]reflect.Value{reflect.Zero(pt)})
	if len(results) != 4 {
		panic("XXX_OneofFuncs behaved in an unexpected way")
	}

	// clist is a reflect.Value representing a []interface{} where each element
	// is nil cast to *Type, where Type is a oneof substrate for |at|'s
	// underlying type.
	clist := results[3]
	if clist.Kind() != reflect.Slice {
		panic("Unexpected return type for interface list")
	}

	err_list := []error{}
	for i := 0; i < clist.Len(); i++ {
		ci := clist.Index(i)
		if ci.Kind() != reflect.Interface {
			panic("Unexpected return type for interface list")
		}
		ct := reflect.TypeOf(ci.Interface())
		// conveniently, ci is now a reflect.Value whose Type() is the subtype we want to examine.
		err_list = AppendErrorList(err_list, VerifyValidatableType(ct))
	}
	return WrapErrorList(err_list)
}
