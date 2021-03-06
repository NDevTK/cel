// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"context"
	"fmt"
	"time"

	"chromium.googlesource.com/enterprise/cel/go/common"
	"github.com/golang/protobuf/proto"
)

// deployerContext implements common.Context. It's used by DeployerSession and
// wraps the various interfaces that are required by common.Context.
//
// TODO(asanka): Move towards model of exposing the interfaces rather than
// wrapping each one.
type deployerContext struct {
	ctx         context.Context // Layered context
	objectStore common.ObjectStore
	publisher   common.Publisher
	getter      common.Getter
	logger      common.Logger
}

func (d *deployerContext) Deadline() (time.Time, bool) {
	return d.ctx.Deadline()
}

func (d *deployerContext) Done() <-chan struct{} {
	return d.ctx.Done()
}

func (d *deployerContext) Err() error {
	return d.ctx.Err()
}

func (d *deployerContext) Value(key interface{}) interface{} {
	return d.ctx.Value(key)
}

func (d *deployerContext) Publish(m proto.Message, field string, value interface{}) error {
	return d.publisher.Publish(m, field, value)
}

func (d *deployerContext) PublishDependency(m proto.Message, dependsOn common.RefPath) error {
	return d.publisher.PublishDependency(m, dependsOn)
}

func (d *deployerContext) GetObjectStore() common.ObjectStore {
	return d.objectStore
}

func (d *deployerContext) Debug(v fmt.Stringer) {
	d.logger.Debug(v)
}

func (d *deployerContext) Info(v fmt.Stringer) {
	d.logger.Info(v)
}

func (d *deployerContext) Warning(v fmt.Stringer) {
	d.logger.Warning(v)
}

func (d *deployerContext) Error(v fmt.Stringer) {
	d.logger.Error(v)
}

func (d *deployerContext) Get(p common.RefPath) (interface{}, error) {
	return d.getter.Get(p)
}

func (d *deployerContext) Indirect(m proto.Message, f string) (interface{}, error) {
	return d.getter.Indirect(m, f)
}
