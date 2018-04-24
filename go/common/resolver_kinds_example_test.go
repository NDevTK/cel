// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common_test

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/host"
)

// ExampleResolver is an example resolver.
//
// Resolver functions can be named anything. But by convention, they are named
// as <Something>Resolver.
//
// Note how the second argument is not a proto.Message? When the second
// argument is a pointer to a concrete type that implements proto.Message, the
// RegisterResolverFunc() invocation correctly deduces the tyep of resources
// that the resolver is expected to handle.
func ExampleResolver(ctx common.Context, i *host.Image) error {
	// Do stuff
	return nil
}

// Don't forget to call RegisterResolverFunc in the init() function or any time
// before the resolver is run.
func init() {
	common.RegisterResolverFunc(common.ImmediateResolverKind, ExampleResolver)
}
