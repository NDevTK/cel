// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"context"
)

type ObjectStoreService interface {
	PutObject(ctx context.Context, content []byte, reference *string, result chan<- error)
	GetObject(ctx context.Context, reference string, target *[]byte, result chan<- error)
	PutNamedObject(ctx context.Context, name string, content []byte, reference *string, result chan<- error)
	GetNamedObject(ctx context.Context, name string, target *[]byte, result chan<- error)
}

type objectStoreServiceKeyType int

const objectStoreServiceKey objectStoreServiceKeyType = 0

func ObjectStoreServiceFromContext(ctx context.Context) (ObjectStoreService, bool) {
	o, ok := ctx.Value(objectStoreServiceKey).(ObjectStoreService)
	return o, ok
}

func ContextWithObjectStoreService(ctx context.Context, o ObjectStoreService) context.Context {
	return context.WithValue(ctx, objectStoreServiceKey, o)
}
