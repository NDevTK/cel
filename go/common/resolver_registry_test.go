// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

type fakeContext struct{}

func (f *fakeContext) Deadline() (deadline time.Time, ok bool)                    { return time.Now(), false }
func (f *fakeContext) Done() <-chan struct{}                                      { return nil }
func (f *fakeContext) Err() error                                                 { return nil }
func (f *fakeContext) Value(key interface{}) interface{}                          { return nil }
func (f *fakeContext) Publish(proto.Message, string, interface{}) error           { return nil }
func (f *fakeContext) PublishDependency(m proto.Message, dependsOn RefPath) error { return nil }
func (f *fakeContext) Debug(v fmt.Stringer)                                       {}
func (f *fakeContext) Info(v fmt.Stringer)                                        {}
func (f *fakeContext) Warning(v fmt.Stringer)                                     {}
func (f *fakeContext) Error(v fmt.Stringer)                                       {}
func (f *fakeContext) GetObjectStore() ObjectStore                                { return nil }
func (f *fakeContext) Get(RefPath) (interface{}, error)                           { return nil, nil }
func (f *fakeContext) Indirect(proto.Message, string) (interface{}, error)        { return nil, nil }

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
