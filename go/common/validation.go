// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

func (v Validation) IsNamedReference() bool {
	return v.Ref != ""
}

func (v Validation) IsOutput() bool {
	return v.Type == Validation_OUTPUT
}

func (v Validation) IsRuntime() bool {
	return v.Type == Validation_RUNTIME
}

func (v Validation) ReferenceRoot() (RefPath, error) {
	return RefPathFromString(v.Ref)
}
