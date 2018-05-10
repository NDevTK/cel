// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"github.com/golang/protobuf/proto"
)

// ApplyResolvers applies resolvers of a given kind over the entire namespace.
//
// Resolvers are typically registered for a single message type. However, there
// can be more than one resolver of a given type registered for a message. This
// function traverses the namespace in topological order. For each message
// encountered, it looks up all the resolvers of the specified type registered
// for that message class, and invokes them in an unspecified order.
//
// The operation may fail at an arbitrary point if any of the resolvers return
// a non-nil error. In this case, the returned error may be an aggregate of all
// encountered errors.
//
// All applicable resolvers for a message must return a nil error in order for
// that message to be considered successfully resolved. A messages will only be
// visited if all messages messages that it depends on were successfully
// resolved.
func ApplyResolvers(ctx Context, ns *Namespace, kind ResolverKind) error {
	return ns.TopoVisit(func(m proto.Message) error {
		return resolverWorker(ctx, kind, m)
	})
}

func resolverWorker(ctx Context, kind ResolverKind, m proto.Message) error {
	rs, err := getResolvers(m, kind)

	if err == ResolverNotFoundError {
		return nil
	}

	if err != nil {
		return err
	}

	var el []error
	for _, resolver := range rs {
		err := resolver(ctx, m)
		el = AppendErrorList(el, err)
	}
	return WrapErrorList(el)
}
