// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

import (
	"fmt"
	"path/filepath"
	"time"

	"chromium.googlesource.com/enterprise/cel/go/asset"
	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/gcp/onhost"
	"github.com/pkg/errors"
)

type DomainJoinResolver struct{}

func (*DomainJoinResolver) ResolveOnHost(ctx common.Context, m *asset.WindowsMachine) error {
	d, ok := ctx.(*deployer)
	if !ok {
		return errors.New("ctx is not Deployer")
	}

	if m.Name == d.instanceName && m.Container != nil {
		adContainer, ok := m.Container.Container.(*asset.WindowsContainer_AdDomain)
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

func joinDomain(d *deployer, ad *asset.ActiveDirectoryDomain) error {
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

	dnsServerAddress, err := getInstanceAddress(ad.DomainController[0].WindowsMachine)
	if err != nil {
		return err
	}

	retries := 0
	for {
		if d.nestedVM == nil {
			err = d.RunConfigCommand("powershell.exe", "-File", fileToRun,
				"-domainName", ad.Name,
				"-dnsServerAddress", dnsServerAddress,
				"-adminName", ad.Name+"\\administrator", "-adminPassword",
				string(ad.SafeModeAdminPassword.Final))

		} else {
			err = joinNestedVMtoDomain(d, ad, fileToRun, dnsServerAddress)
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

func joinNestedVMtoDomain(d *deployer, ad *asset.ActiveDirectoryDomain, fileToRun string, dnsServerAddress string) error {
	fileToRun = filepath.Join(workingDirectoryOnNestedVM, filepath.Base(fileToRun))
	err := d.runConfigCommandOnNestedVM("powershell.exe",
		"-File", fileToRun,
		"-domainName", ad.Name,
		"-dnsServerAddress", dnsServerAddress,
		"-adminName", ad.Name+"\\administrator",
		"-adminPassword",
		fmt.Sprintf("\"%s\"", ad.SafeModeAdminPassword.Final))

	if err != nil && err != ErrRebootNeeded {
		return err
	}

	// for nested VM, reboot is handled here.
	d.Logf("Reboot needed. Continue configuration after reboot.")
	if err := d.Reboot(); err != nil {
		return err
	}

	// wait a while to give shutdown enough time to finish.
	time.Sleep(1 * time.Minute)

	// After domain join, on Win7, local user login thru ssh stops working.
	// So switch to domain admin account for log in.
	d.nestedVM.UserName = ad.Name + "\\administrator"
	d.nestedVM.Password = string(ad.SafeModeAdminPassword.Final)
	if err := d.waitUntilSshIsAlive(); err != nil {
		return err
	}

	// fix dns record. For some reason, this can fail (e.g. on Win7) so we
	// need retry logic here
	fileToRun = filepath.Join(workingDirectoryOnNestedVM, "add_dns_entry.ps1")
	return d.runConfigCommandOnNestedVM("powershell.exe",
		"-File", fileToRun,
		"-adminName", ad.Name+"\\administrator",
		"-adminPassword", fmt.Sprintf("\"%s\"", ad.SafeModeAdminPassword.Final),
		"-dnsServerName", ad.DomainController[0].WindowsMachine,
		"-domainName", ad.Name,
		"-computerName", d.instanceName,
		"-ipAddress", d.externalIP)
}

func init() {
	common.RegisterResolverClass(&DomainJoinResolver{})
}
