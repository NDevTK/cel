// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
)

type fakeContext struct {
	objectStore ObjectStore
}

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
func (f *fakeContext) GetObjectStore() ObjectStore                                { return f.objectStore }
func (f *fakeContext) Get(RefPath) (interface{}, error)                           { return nil, nil }
func (f *fakeContext) Indirect(proto.Message, string) (interface{}, error)        { return nil, nil }
