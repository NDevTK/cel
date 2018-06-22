// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

import (
	"chromium.googlesource.com/enterprise/cel/go/asset"
	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/gcp/onhost"
	"github.com/pkg/errors"
)

type AdDomainResolver struct{}

func (*AdDomainResolver) ResolveOnHost(ctx common.Context, ad *asset.ActiveDirectoryDomain) error {
	d, ok := ctx.(*deployer)
	if !ok {
		return errors.New("ctx is not Deployer")
	}

	if ad.DomainController[0].WindowsMachine == d.instanceName {
		return setupADDomain(d, ad)
	} else {
		return nil
	}
}

func setupADDomain(d *deployer, ad *asset.ActiveDirectoryDomain) error {
	d.Logf("Configuring Ad Domain: %+v\n", *ad)

	configVar := onhost.GetActiveDirectoryRuntimeConfigVariableName(ad.Name)
	status := d.getRuntimeConfigVariableValue(configVar)
	d.Logf("Status of asset %s is [%s]", configVar, status)
	if status == statusReady {
		d.Logf("AD Domain already configured")
		return nil
	}
	if status == statusError {
		return errors.New("AD Domain configuration failed")
	}

	d.setRuntimeConfigVariable(configVar, statusInProgress)

	// There are 3 cases:
	// No ancestor -> The root domain of a new forest
	// Ancestor is ParentName -> The child domain of the parent domain
	// Ancestor is Forest -> Adding a new tree domain to an existing root domain
	if ad.Ancestor == nil {
		return createRootDomain(d, ad)
	} else {
		// TODO(feiling): add support for child & tree AD
		return common.NewNotImplementedError("Support for Child and Tree AD")
	}
}

func createRootDomain(d *deployer, ad *asset.ActiveDirectoryDomain) error {
	configVar := onhost.GetActiveDirectoryRuntimeConfigVariableName(ad.Name)
	fileToRun := ""
	if d.IsWindows2012() || d.IsWindows2016() {
		fileToRun = d.GetSupportingFilePath("create_ad_win2012.ps1")
	} else if d.IsWindows2008() {
		fileToRun = d.GetSupportingFilePath("create_ad_win2008.ps1")
	} else {
		return errors.New("unsupported windows version")
	}

	// normal domain
	if err := d.RunConfigCommand("powershell.exe", "-File", fileToRun, "-domainName", ad.Name,
		"-netbiosName", ad.NetbiosName, "-adminName", "administrator",
		"-adminPassword", string(ad.SafeModeAdminPassword.Final)); err != nil {
		return err
	}

	// TODO(feiling): create DNS forwarder if there are tree domains underneath it.
	d.Logf("AD Domain config is finished.")
	d.setRuntimeConfigVariable(configVar, statusReady)
	return nil
}

func init() {
	common.RegisterResolverClass(&AdDomainResolver{})
}
