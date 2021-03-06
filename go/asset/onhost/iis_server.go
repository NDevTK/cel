// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

import (
	"fmt"

	"chromium.googlesource.com/enterprise/cel/go/common"
	assetpb "chromium.googlesource.com/enterprise/cel/go/schema/asset"
	"github.com/pkg/errors"
)

type IisServerResolver struct{}

func (*IisServerResolver) ResolveOnHost(ctx common.Context, iis *assetpb.IISServer) error {
	d, ok := ctx.(*deployer)
	if !ok {
		return errors.New("ctx is not Deployer")
	}

	if iis.WindowsMachine == d.instanceName {
		if err := createIisServer(d, iis); err != nil {
			return err
		}

		// contains the list of IIS Sites hosted on this instance
		iisSites := make(map[string]int)

		for _, iisSite := range d.configuration.AssetManifest.IisSite {
			if iisSite.IisServer == iis.Name {
				iisSites[iisSite.Name] = 1
				if err := createIisSite(d, iisSite); err != nil {
					return err
				}
			}
		}

		for _, iisApp := range d.configuration.AssetManifest.IisApplication {
			_, ok := iisSites[iisApp.IisSite]
			if ok {
				// TODO: add IISApplication support
				d.Logf("IIS App on this instance: %v", iisApp)
			}
		}
		return nil
	} else {
		return nil
	}
}

func createIisServer(d *deployer, iis *assetpb.IISServer) error {
	fileToRun := ""
	if d.IsWindows2008() || d.IsWindows2012() || d.IsWindows2016() {
		fileToRun = d.GetSupportingFilePath("create_iis_server.ps1")
	} else {
		return errors.New("unsupported windows version")
	}

	if err := d.RunConfigCommand("powershell.exe", "-File", fileToRun); err != nil {
		return err
	}

	d.Logf("IISServer created.")
	return nil
}

func createIisSite(d *deployer, iisSite *assetpb.IISSite) error {
	fileToRun := ""
	if d.IsWindows2008() || d.IsWindows2012() || d.IsWindows2016() {
		fileToRun = d.GetSupportingFilePath("create_iis_site.ps1")
	} else {
		return errors.New("unsupported windows version")
	}

	port := iisSite.Bindings.Port

	if port == 0 {
		switch iisSite.Bindings.Protocol {
		case assetpb.Protocol_HTTP:
			port = 80
		case assetpb.Protocol_HTTPS:
			port = 443
		default:
			return errors.New("can't find default port for unsupported protocol")
		}
	}

	if err := d.RunConfigCommand("powershell.exe", "-File", fileToRun,
		"-Name", iisSite.Name,
		"-Protocol", iisSite.Bindings.Protocol.String(),
		"-Port", fmt.Sprintf("%d", port),
		"-Authentication", iisSite.AuthType.String()); err != nil {
		return err
	}

	d.Logf("IISSite created.")
	return nil
}

func init() {
	common.RegisterResolverClass(&IisServerResolver{})
}
