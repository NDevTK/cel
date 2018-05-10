// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"chromium.googlesource.com/enterprise/cel/go/asset"
	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/gcp/compute"
)

type network struct{}

func (*network) ResolveConstructedAssets(ctx common.Context, n *asset.Network) error {
	d := GetDeploymentManifest()

	if n.AddressRange != nil {
		return common.NewNotImplementedError("asset.network.{}.address_range")
	}

	return d.Emit(n, &compute.Network{
		Name:        n.Name,
		Description: "",
	})
}

func init() {
	common.RegisterResolverClass(&network{})
}
