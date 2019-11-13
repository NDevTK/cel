// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

import (
	"time"

	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/gcp/onhost"
	assetpb "chromium.googlesource.com/enterprise/cel/go/schema/asset"
	hostpb "chromium.googlesource.com/enterprise/cel/go/schema/host"
	"github.com/pkg/errors"
)

type DomainJoinResolver struct{}

func (*DomainJoinResolver) ResolveOnHost(ctx common.Context, m *assetpb.WindowsMachine) error {
	d, ok := ctx.(*deployer)
	if !ok {
		return errors.New("ctx is not Deployer")
	}

	if m.Name == d.instanceName && m.Container != nil {
		adContainer, ok := m.Container.Container.(*assetpb.WindowsContainer_AdDomain)
		if ok {
			adAsset, err := d.getAdDomainAsset(adContainer.AdDomain)
			if err != nil {
				return err
			}
			return joinDomain(d, adAsset)
		} else {
			return common.NewNotImplementedError("non ADDomain container is not supported")
		}
	} else {
		return nil
	}
}

const maxRetries = 5

func joinDomain(d *deployer, ad *assetpb.ActiveDirectoryDomain) error {
	if d.GetOs() != hostpb.OperatingSystem_WINDOWS {
		return errors.New("Domain join is only supported on Windows")
	}

	depVar := onhost.GetActiveDirectoryRuntimeConfigVariableName(ad.Name)

	err := d.waitForDependency(depVar, time.Duration(60)*time.Minute)
	if err != nil {
		return err
	}

	fileToRun := ""
	if d.IsWindows2012() || d.IsWindows2016() {
		fileToRun = d.GetSupportingFilePath("join_domain_win2012.ps1")
	} else if d.IsWindows2008() {
		fileToRun = d.GetSupportingFilePath("join_domain_win2008.ps1")
	} else {
		return errors.New("unsupported windows version")
	}

	dnsServerAddress, err := d.getInstanceAddress(ad.DomainController[0].WindowsMachine)

	if err != nil {
		return err
	}

	retries := 0
	for {
		err = d.RunConfigCommand("powershell.exe",
			"-File", fileToRun,
			"-domainName", ad.Name,
			"-dnsServerAddress", dnsServerAddress,
			"-adminName", ad.Name+"\\administrator",
			"-adminPassword", string(ad.SafeModeAdminPassword.Final))

		if d.IsNestedVM() && err == ErrRebootNeeded {
			// for nested VM, we need to handle the reboot ourselves.
			d.Logf("Reboot needed. Continue configuration after reboot.")
			err = d.Reboot()
			if err != nil {
				return err
			}
		}

		if err == nil {
			break
		}

		if err == ErrTransient && retries <= maxRetries {
			retries++
			d.Logf("Script returned a transient error. Will wait a minute and try again.")
			time.Sleep(1 * time.Minute)
			continue
		}
		return err
	}

	d.Logf("Domain join finished")
	return nil
}

func init() {
	common.RegisterResolverClass(&DomainJoinResolver{})
}
