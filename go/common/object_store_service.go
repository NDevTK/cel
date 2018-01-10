// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"context"
)

type ObjectStore interface {
	PutObject(ctx context.Context, content []byte, reference *string, result chan<- error)
	GetObject(ctx context.Context, reference string, target *[]byte, result chan<- error)
}

type objectStoreServiceKeyType int

const objectStoreServiceKey objectStoreServiceKeyType = 0

func ObjectStoreFromContext(ctx context.Context) (ObjectStore, error) {
	o, ok := ctx.Value(objectStoreServiceKey).(ObjectStore)
	if !ok {
		return nil, &ServiceNotFoundError{Service: "ObjectStore"}
	}
	return o, nil
}

func ContextWithObjectStore(ctx context.Context, o ObjectStore) context.Context {
	return context.WithValue(ctx, objectStoreServiceKey, o)
}
