// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"chromium.googlesource.com/enterprise/cel/go/schema"
	commonpb "chromium.googlesource.com/enterprise/cel/go/schema/common"
)

var validateFunctions = []interface{}{
	func(*commonpb.TestBadOneOf) error { return nil },
	func(*commonpb.TestGoodOneOf) error { return nil },
	func(*commonpb.TestGoodProto) error { return nil },
	func(*commonpb.TestHasBadField) error { return nil },
	func(*commonpb.TestHasBadSlice) error { return nil },
	func(*commonpb.TestHasGoodField) error { return nil },
	func(*commonpb.TestHasGoodSlice) error { return nil },
	func(*commonpb.TestMessageWithOptions) error { return nil },
	// Intentionally leaving out BadProto which shouldn't have a validate method.
}

func Validate_TestBadReturnType(*commonpb.TestBadReturnType) int            { return 0 }
func Validate_TestBadValidateArgs(*commonpb.TestBadValidateArgs, int) error { return nil }

func init() {
	schema.RegisterAllValidateFunctions(validateFunctions)
}
