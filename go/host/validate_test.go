// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package host

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"reflect"
	"testing"
)

func TestAssetManifest_validateFields(t *testing.T) {
	var h HostEnvironment
	err := common.InvokeValidate(&h, common.EmptyPath)
	if err != nil {
		t.Fatal("unexpected error ", err)
	}

	err = common.VerifyValidatableType(reflect.TypeOf(&h))
	if err != nil {
		t.Fatal("unexpected error ", err)
	}
}
