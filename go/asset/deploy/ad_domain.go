// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"chromium.googlesource.com/enterprise/cel/go/asset"
	"chromium.googlesource.com/enterprise/cel/go/common"
)

type adDomain struct{}

func (*adDomain) ResolveGeneratedContent(ctx common.Context, d *asset.ActiveDirectoryDomain) error {
	p, err := generatePassword()
	if err != nil {
		return err
	}

	return ctx.Publish(d, "safe_mode_admin_password", &common.Secret{Final: []byte(p)})
}

func init() {
	common.RegisterResolverClass(&adDomain{})
}
