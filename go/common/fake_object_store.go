// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"fmt"
)

type fakeObjectStore struct {
	Log []string
}

func (f *fakeObjectStore) PutObject(payload []byte) (objectReference string, result error) {
	f.Log = append(f.Log, fmt.Sprintf("PutObject:%s", string(payload)))
	return "foo", nil
}

func (f *fakeObjectStore) PutNamedObject(name string, payload []byte) (objectReference string, result error) {
	f.Log = append(f.Log, fmt.Sprintf("PutNamedObject:%s:%s", name, string(payload)))
	return "foo", nil
}

func (f *fakeObjectStore) GetObject(objectReference string) (name string, payload []byte, result error) {
	panic("not implemented")
}
