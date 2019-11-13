// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	commonpb "chromium.googlesource.com/enterprise/cel/go/schema/common"
)

func IsNamedReference(v *commonpb.Validation) bool {
	return v.Ref != ""
}

func IsOutput(v *commonpb.Validation) bool {
	return v.Type == commonpb.Validation_OUTPUT
}

func IsRuntime(v *commonpb.Validation) bool {
	return v.Type == commonpb.Validation_RUNTIME
}

func IsTopLevelCollection(v *commonpb.Validation) bool {
	return v.Type == commonpb.Validation_TOPLEVEL
}

func ReferenceRoot(v *commonpb.Validation) (RefPath, error) {
	return RefPathFromString(v.Ref)
}
