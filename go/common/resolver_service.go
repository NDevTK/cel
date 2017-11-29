// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"context"
	"github.com/golang/protobuf/proto"
)

type ResolverService interface {
	// ResolvePath resolves the asset identified by |path| and notifies |waiter|.
	ResolvePath(ctx context.Context, path RefPath, waiter *JobWaiter)

	// Resolve resolves the asset described by |target| and notifies |waiter|
	// when the operation completes.
	//
	// The |waiter| also identies the source object, if known and will be used
	// for resolving deadlocks due to circular references.
	//
	// If |target| is nil, then the call should return without doing anything.
	// |waiter| should be unmodified in this case.
	Resolve(ctx context.Context, target proto.Message, waiter *JobWaiter)
}
