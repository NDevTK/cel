// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"github.com/golang/protobuf/proto"
)

type Getter interface {
	// Get returns the value a location specified as a RefPath. If the location
	// doesn't exist or if the value is not known, the function returns a
	// non-nil error. Otherwise the returned `interface{}` is the value found
	// at the location.
	Get(path RefPath) (value interface{}, err error)

	// Indirect indirects (or follows) a named reference.
	//
	// It's function is best explained via an example. Suppose you have a field
	// that is a named reference to another as below:
	//
	//   windows_user {
	//     ...
	//     container {
	//       ad_domain: "mydomain.example"
	//     }
	//   }
	//
	// A resolver will aquire a reference to the outer WindowsUser message, and
	// in turn the inner WindowsContainer message. However, it will also likely
	// need access to the ActiveDirectoryDomain object that corresponds to the
	// "ad_domain" field in WindowsContainer.
	//
	// This can be resolved as follows:
	//
	//    v, err := namespace.Indirect(u.Container, "ad_domain")
	//    if err != nil {
	//      // handle error
	//    }
	//    domain, ok := v.(*ActiveDirectoryDomain)
	//
	// This pattern effectively indirects or follows a reference across objects
	// in a single namespace.
	//
	// If there's an error, the returned "err" will be non-nil.
	Indirect(m proto.Message, fieldName string) (value interface{}, err error)
}
