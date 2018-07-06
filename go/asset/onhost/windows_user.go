// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

import (
	"chromium.googlesource.com/enterprise/cel/go/asset"
	"chromium.googlesource.com/enterprise/cel/go/common"
	"github.com/pkg/errors"
)

type windowsUser struct{}

func (*windowsUser) ResolveOnHost(ctx common.Context, u *asset.WindowsUser) error {
	d, ok := ctx.(*deployer)
	if !ok {
		return errors.New("ctx is not Deployer")
	}

	adContainer, ok := u.Container.Container.(*asset.WindowsContainer_AdDomain)
	if ok {
		adAsset, err := d.getAdDomainAsset(adContainer.AdDomain)
		if err != nil {
			return err
		}

		if adAsset.DomainController[0].WindowsMachine == d.instanceName {
			return createUser(d, adAsset, u)
		} else {
			return nil
		}
	} else {
		return common.NewNotImplementedError("non ADDomain container is not supported")
	}
}

func createUser(d *deployer, ad *asset.ActiveDirectoryDomain, u *asset.WindowsUser) error {
	fileToRun := ""
	if d.IsWindows2012() || d.IsWindows2016() {
		fileToRun = d.GetSupportingFilePath("create_user_win2012.ps1")
	} else if d.IsWindows2008() {
		fileToRun = d.GetSupportingFilePath("create_user_win2008.ps1")
	} else {
		return errors.New("unsupported windows version")
	}

	return d.RunConfigCommand("powershell.exe", "-File", fileToRun,
		"-userName", u.Name,
		"-password", string(u.Password.Final),
		"-description", u.Description,
		"-domainName", ad.Name,
		"-adminPassword", string(ad.SafeModeAdminPassword.Final))
}

func init() {
	common.RegisterResolverClass(&windowsUser{})
}