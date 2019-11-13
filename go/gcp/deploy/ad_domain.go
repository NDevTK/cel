// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/gcp/onhost"
	assetpb "chromium.googlesource.com/enterprise/cel/go/schema/asset"
)

type adDomain struct{}

func (*adDomain) ResolveConstructedAssets(ctx common.Context, d *assetpb.ActiveDirectoryDomain) error {
	m := GetDeploymentManifest()
	variableName := onhost.GetActiveDirectoryRuntimeConfigVariableName(d.Name)
	return m.Emit(nil,
		&onhost.RuntimeConfigConfigVariable{
			Name:     "runtimeconfigVariable_" + variableName,
			Parent:   onhost.RuntimeconfigVariableParent,
			Variable: variableName,
			Text:     "",
		})
}

func init() {
	common.RegisterResolverClass(&adDomain{})
}
