// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

func (*TestBadOneOf) Validate() error           { return nil }
func (*TestBadReturnType) Validate() int        { return 0 }
func (*TestBadValidateArgs) Validate(int) error { return nil }
func (*TestGoodOneOf) Validate() error          { return nil }
func (*TestGoodProto) Validate() error          { return nil }
func (*TestHasBadField) Validate() error        { return nil }
func (*TestHasBadSlice) Validate() error        { return nil }
func (*TestHasGoodField) Validate() error       { return nil }
func (*TestHasGoodSlice) Validate() error       { return nil }
func (*TestMessageWithOptions) Validate() error { return nil }

// Intentionally leaving out BadProto which shouldn't have a Validator().
