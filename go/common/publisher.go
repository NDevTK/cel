// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"github.com/golang/protobuf/proto"
)

// Publisher is an interface that can be used to publish OUTPUT and RUNTIME
// values. This interface is provided for use by resolvers.
//
// Both of these methods are implemented by the *Namespace type.
type Publisher interface {
	// Publish sets the OUTPUT field named "field" of Message specified by
	// "message" to "object".
	Publish(message proto.Message, field string, object interface{}) error

	// PublishDependency records an explicit dependency.
	//
	// The message specified by "message" will henceforth depend on the object
	// at path "dependsOn".  It is valid for the path at "dependsOn" to not
	// exist at the time this call is made.
	//
	// This is typically invoked by AdditionalDependencyResolver type
	// resolvers.
	//
	// Note that once a dependency is added, it cannot be removed.
	//
	// Dependencies that are added during deployment do not persist in the
	// expanded asset manifest, and don't make it through to on-host
	// configuration.
	PublishDependency(message proto.Message, dependsOn RefPath) error
}
