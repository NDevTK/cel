// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"reflect"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

func TestWrapGenericResolverFunction_identity(t *testing.T) {
	expectedErr := errors.New("mock")
	gotCalled := false
	f := wrapGenericResolverFunction(reflect.ValueOf(func(ctx Context, m proto.Message) error {
		gotCalled = true
		return expectedErr
	}), ProtoMessageType)

	err := f(&fakeContext{}, &TestMessageWithOptions{})

	if !gotCalled {
		t.Error("wrapped function did not get called")
	}

	if err != expectedErr {
		t.Errorf("expected error message did not get passed through: %#v", err)
	}
}

func TestWrapGenericResolverFunction_conversion(t *testing.T) {
	expectedErr := errors.New("mock")
	gotCalled := false
	gotLabel := ""
	f := wrapGenericResolverFunction(reflect.ValueOf(func(ctx Context, t *TestMessageWithOptions) error {
		gotCalled = true
		gotLabel = t.Label
		return expectedErr
	}), ProtoMessageType)

	err := f(&fakeContext{}, &TestMessageWithOptions{Label: "knock-knock"})

	if !gotCalled {
		t.Error("wrapped function did not get called")
	}

	if gotLabel != "knock-knock" {
		t.Errorf("message did not get passed through")
	}

	if err != expectedErr {
		t.Errorf("expected error message did not get passed through: %#v", err)
	}
}
