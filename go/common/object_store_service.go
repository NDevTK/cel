// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"context"
)

type ObjectStore interface {
	PutObject(ctx context.Context, content []byte) (reference string, result error)
	PutNamedObject(ctx context.Context, name string, content []byte) (reference string, result error)
	GetObject(ctx context.Context, reference string) (name string, target []byte, result error)
}
