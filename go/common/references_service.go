// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"github.com/golang/protobuf/proto"
)

type ReferencesService interface {
	NameOf(proto.Message) (RefPath, error)
	ObjectAt(RefPath) (proto.Message, error)
	Get(RefPath, interface{}) error
}
