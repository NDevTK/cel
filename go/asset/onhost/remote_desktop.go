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
		return setupRemoteDesktopHost(d, rd)
	} else {
		return nil
	}
}

func setupRemoteDesktopHost(d *deployer, rd *asset.RemoteDesktopHost) error {
	fileToRun := ""
	if d.IsWindows2012() || d.IsWindows2016() {
		fileToRun = d.GetSupportingFilePath("create_remote_desktop_win2012.ps1")
	} else if d.IsWindows2008() {
		// TODO: The powershell script, create_remote_desktop_win2008.ps1, doesn't work because
		// module RemoteDesktop is not available on Windows 2008R2. The solution is to
		// use the steps documented in
		// https://blogs.technet.microsoft.com/ptsblog/2011/12/09/extending-remote-desktop-services-using-powershell/
		return errors.New("unsupported windows version")
	} else {
		return errors.New("unsupported windows version")
	}

	if err := d.RunConfigCommand("powershell.exe", "-File", fileToRun,
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
