// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	assetpb "chromium.googlesource.com/enterprise/cel/go/schema/asset"
)

type windowsContainer struct{}

// Add dependency from ADDomain WindowsContainer to the ADDomain. This
// is needed because the dependency is not represented in the completed asset manifest.
func (*windowsContainer) ResolveAdditionalDependencies(ctx common.Context, u *assetpb.WindowsContainer) (err error) {
	if u == nil {
		return nil
	}

	adContainer, ok := u.Container.(*assetpb.WindowsContainer_AdDomain)
	if ok {
		ad := common.RefPathFromComponents("asset", "ad_domain", adContainer.AdDomain)
		return ctx.PublishDependency(u, ad)
	} else {
		return common.NewNotImplementedError("non ADDomain container is not supported")
	}
}

func init() {
	common.RegisterResolverClass(&windowsContainer{})
}
