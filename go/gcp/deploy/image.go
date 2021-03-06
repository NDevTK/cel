// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/gcp"
	hostpb "chromium.googlesource.com/enterprise/cel/go/schema/host"
	"github.com/pkg/errors"
)

func ImageResolver(ctx common.Context, i *hostpb.Image) error {
	session, err := gcp.SessionFromContext(ctx)
	if err != nil {
		return err
	}

	// Fixed URL. Nothing more to do. If the URL was invalid, the deployment
	// will fail. Not a whole lot we are going to do about it.
	if fixed, ok := i.GetSource().(*hostpb.Image_Fixed); ok {
		return ctx.Publish(i, "url", fixed.Fixed)
	}

	family := i.GetLatest()
	if family == nil {
		return errors.New("invalid input")
	}

	compute, err := session.GetComputeService()
	if err != nil {
		return err
	}
	image, err := compute.Images.GetFromFamily(family.GetProject(), family.GetFamily()).Context(ctx).Do()
	if err != nil {
		return err
	}

	return ctx.Publish(i, "url", image.SelfLink)
}

func init() {
	common.RegisterResolverFunc(common.ImmediateResolverKind, ImageResolver)
}
