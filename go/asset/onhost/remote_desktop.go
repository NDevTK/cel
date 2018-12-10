// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

import (
	"chromium.googlesource.com/enterprise/cel/go/asset"
	"chromium.googlesource.com/enterprise/cel/go/common"
	"github.com/pkg/errors"
)

type RemoteDesktopHostResolver struct{}

func (*RemoteDesktopHostResolver) ResolveOnHost(ctx common.Context, rd *asset.RemoteDesktopHost) error {
	d, ok := ctx.(*deployer)
	if !ok {
		return errors.New("ctx is not Deployer")
	}

	if rd.WindowsMachine == d.instanceName {
		ad := d.getActiveDirectoryDomain()
		if ad == nil {
			return errors.New("RDS only supports domain-joined servers.")
		}

		return setupRemoteDesktopHost(d, ad, rd)
	} else {
		return nil
	}
}

func setupRemoteDesktopHost(d *deployer, ad *asset.ActiveDirectoryDomain, rd *asset.RemoteDesktopHost) error {
	fileToRun := ""
	if d.IsWindows2012() || d.IsWindows2016() {
		if len(rd.CollectionName) <= 0 {
			return errors.New("collection_name is required for RDS on Windows Server >= 2012.")
		}
		fileToRun = d.GetSupportingFilePath("create_remote_desktop_win2012.ps1")
	} else if d.IsWindows2008() {
		fileToRun = d.GetSupportingFilePath("create_remote_desktop_win2008.ps1")
	} else {
		return errors.New("unsupported windows version")
	}

	if err := d.RunConfigCommand("powershell.exe", "-File", fileToRun,
		"-adminName", ad.Name+"\\administrator",
		"-adminPassword", string(ad.SafeModeAdminPassword.Final),
		"-collectionName", rd.CollectionName,
		"-collectionDescription", rd.CollectionDescription); err != nil {
		return err
	}

	d.Logf("Remote Desktop Host config finished")
	return nil
}

func init() {
	common.RegisterResolverClass(&RemoteDesktopHostResolver{})
}