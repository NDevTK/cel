// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

// Validator is an interface implemented by things that can be validated.
type Validator interface {
	// Validate returns nil if the object is valid. Otherwise returns an error
	// that describes why the object is not valid.
	//
	// Guidelines for writing a validator for a proto.Message type:
	// ------------------------------------------------------------
	//
	// *  DO Check if required fields are present.
	//
	// *  DO Enforce applicable formatting restrictions for string fields, and
	//    range checks for numerics. Use IsRFC1035Label() for validating names.
	//
	//
	// *  DONT recursively validate embedded fields. The toolchain will do that
	//    for you.
	//
	// *  DONT check formatting for external references.
	//
	// *  DONT be overly aggressive when enforcing restrictions. The validators
	//    are merely a convenience and shouldn't be considered a comprehensive
	//    set of tests for a valid configuation. E.g. Let Windows decide what a
	//    valid username looks like. Don't enforce username restrictions in
	//    your validator.
	//
	// There are plenty of examples of validators in //go/asset/validate.go and
	// //go/host/validate.go.
	Validate() error
}
