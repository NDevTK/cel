// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"context"
	"reflect"
)

// Context is the variant of context.Context that's passed down to resolvers.
//
// In addition to the methods available on context.Context, the common.Context
// interface defines several interfaces that provide access to base CEL
// toolchain services.
type Context interface {
	context.Context
	Publisher
	Logger
	GetObjectStore() ObjectStore
}

var ContextType = reflect.TypeOf((*Context)(nil)).Elem()
