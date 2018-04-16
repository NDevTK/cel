// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/host"
	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
	"time"
)

type StorageResolver struct{}

func (StorageResolver) ResolveImmediate(ctx common.Context, s *host.Storage) error {
	session, err := SessionFromContext(ctx)
	if err != nil {
		return err
	}

	client, err := storage.NewClient(ctx, option.WithHTTPClient(session.client))
	if err != nil {
		return err
	}

	bh := client.Bucket(s.GetBucket())
	attrs, err := bh.Attrs(session.ctx)
	if err != nil {
		return err
	}
	ctime := attrs.Created.Format(time.RFC3339)
	return ctx.Publish(s, "created_on", ctime)
}

func init() {
	common.RegisterResolverClass(&StorageResolver{})
}
