// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package host

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"reflect"
	"strings"
	"testing"
)

func TestHostEnvironment_validateFields(t *testing.T) {
	var h HostEnvironment
	h.Project = &Project{Name: "T", Zone: "Z"}
	h.LogSettings = &LogSettings{AdminLog: "A"}
	h.Storage = &Storage{Bucket: "x"}
	err := common.ValidateProto(&h, common.EmptyPath)
	if err != nil {
		t.Fatal("unexpected error ", err)
	}

	err = common.VerifyValidatableType(reflect.TypeOf(&h))
	if err != nil {
		t.Fatal("unexpected error: ", err)
	}
}

func TestHostEnvironment_validateStorage(t *testing.T) {
	var h HostEnvironment
	h.Project = &Project{Name: "T", Zone: "Z"}
	h.LogSettings = &LogSettings{AdminLog: "A"}
	h.Storage = &Storage{Bucket: "x", Prefix: "x/"}
	err := common.ValidateProto(&h, common.EmptyPath)
	if err == nil {
		t.Fatal()
	}

	if !strings.Contains(err.Error(), "must not end with a slash") {
		t.Fatal("unexpected error", err)
	}

	h.Storage.Prefix = "/x"
	err = common.ValidateProto(&h, common.EmptyPath)
	if err != nil {
		t.Fatal(err)
	}
}
