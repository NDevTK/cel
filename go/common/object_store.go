// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

type ObjectStore interface {
	PutObject(payload []byte) (reference string, result error)
	PutNamedObject(name string, content []byte) (reference string, result error)
	GetObject(reference string) (name string, payload []byte, result error)
}
