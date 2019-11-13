// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package schema

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"reflect"
)

// Validate methods returns nil if the object is valid. Otherwise returns
// an error that describes why the object is not valid.
//
// Guidelines for writing validate methods for a proto.Message type:
// ------------------------------------------------------------
//
// *  DO Check if required fields are present.
//
// *  DO Enforce applicable formatting restrictions for string fields, and
//    range checks for numerics. Use IsRFC1035Label() for validating names.
//
// *  DONT recursively validate embedded fields. The toolchain will do that
//    for you.
//
// *  DONT check formatting for external references. The references will
//    fail in due course if invalid.
//
// *  DONT be overly aggressive when enforcing restrictions. The validator
//    methods are merely a convenience and shouldn't be considered a
//    comprehensive set of tests for a valid configuation. E.g. Let Windows
//    decide what a valid username looks like. Don't enforce username
//    restrictions in your validator.
//
// There are examples in //go/asset/validate.go and //go/host/validate.go.

var validateRegistry = make(map[reflect.Type]reflect.Value)

var ErrorType = reflect.TypeOf((*error)(nil)).Elem()
var ProtoMessageType = reflect.TypeOf((*proto.Message)(nil)).Elem()

func GetValidateFunction(t reflect.Type) (reflect.Value, error) {
	if function, ok := validateRegistry[t]; ok {
		return function, nil
	} else {
		return function, fmt.Errorf("can't find Validate function for type %v", t)
	}
}

func RegisterValidateFunction(function interface{}) {
	value := reflect.ValueOf(function)
	err := addValidateFunction(value)
	if err != nil {
		panic(errors.Errorf("invalid validate function %s: %s", value.Type().String(), err))
	}
}

func RegisterAllValidateFunctions(functions []interface{}) {
	for _, f := range functions {
		RegisterValidateFunction(f)
	}
}

func VerifyValidateFunction(function reflect.Value) error {
	fType := function.Type()
	if fType.NumIn() != 1 {
		return errors.New("incorrect number of arguments")
	}
	targetType := fType.In(0)

	if !targetType.Implements(ProtoMessageType) {
		return errors.New("first argument doesn't implement ProtoMessageType")
	}

	if fType.NumOut() != 1 {
		return errors.New("incorrect number of return values")
	}

	if fType.Out(0) != reflect.TypeOf((*error)(nil)).Elem() {
		return errors.New("incorrect return type")
	}

	return nil
}

func addValidateFunction(function reflect.Value) error {
	if err := VerifyValidateFunction(function); err != nil {
		return err
	}

	fType := function.Type()
	targetType := fType.In(0)

	if _, ok := validateRegistry[targetType]; ok {
		return errors.New("already has a validate method")
	} else {
		validateRegistry[targetType] = function
	}

	return nil
}
