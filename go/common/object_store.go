// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

// ObjectStore describes an abstract content indexed blob storage service.
//
// See the description of "Object Storage" [1] in the deployment documentation.
// This interface describes the serviced sketched out in that section.
//
// As the interface implies, the object reference returned by the Put* methods
// must encode a secure digest of the contents -- i.e. the content index
// portion of context indexed storage.
//
// [1]: https://chromium.googlesource.com/enterprise/cel/+/HEAD/docs/deployment.md#object-storage
type ObjectStore interface {
	// PutObject stores a blob of data and returns an object reference that can
	// be used to retrieve it.
	//
	// The returned string is an object reference that must have an implicit
	// content digest. It can be used with GetObject to later retrieve the
	// object from storage.
	//
	// Any errors are indicated via "result".
	PutObject(payload []byte) (objectReference string, result error)

	// PutNamedObject stores a named blob of data and returns an object
	// reference that can be used to retrieve both the data and name at a later
	// time.
	//
	// The returned string is an object reference that must have an implicit
	// content digest. It can be used with GetObject to later retrieve the object
	// and the name. The name doesn't necessarily need to have the name
	// embedded inside it.
	PutNamedObject(name string, payload []byte) (objectReference string, result error)

	// GetObject retrieves the data and optionally the name of the data that
	// was stored via a prior call to PutObject or PutNamedObject.
	//
	// The returned "name" string would be non-empty only if the object
	// reference is the result of calling PutNamedObject.
	GetObject(objectReference string) (name string, payload []byte, result error)
}
