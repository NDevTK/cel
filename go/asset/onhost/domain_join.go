// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

import (
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

	retries := 0
	for {
		if err := d.RunConfigCommand("powershell.exe", "-File", fileToRun,
			"-domainName", ad.Name,
			"-dnsServer", ad.DomainController[0].WindowsMachine,
			"-adminName", ad.Name+"\\administrator", "-adminPassword",
			string(ad.SafeModeAdminPassword.Final)); err != nil {
			if err == ErrTransient && retries <= maxRetries {
				retries++
				d.Logf("Script returned a transient error. Will wait a minute and try again.")
				time.Sleep(1 * time.Minute)
				continue
			}
			return err
		}
		break
	}

	d.Logf("Domain join finished")
	return nil
}

func init() {
	common.RegisterResolverClass(&DomainJoinResolver{})
}
