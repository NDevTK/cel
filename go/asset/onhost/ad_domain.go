// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

import (
	"strings"
	"time"

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
	var err error
	if ad.Ancestor == nil {
		err = createRootDomain(d, ad)
	} else {
		parent, ok := ad.Ancestor.(*asset.ActiveDirectoryDomain_ParentName)
		if ok {
			err = createChildDomain(d, ad, parent.ParentName)
		} else {
			// TODO(feiling): add support for child & tree AD
			err = common.NewNotImplementedError("Support for Child and Tree AD")
		}
	}

	if err == nil {
		d.setRuntimeConfigVariable(configVar, statusReady)
	} else if err != ErrRebootNeeded {
		d.setRuntimeConfigVariable(configVar, statusError)
	}
	return err
}

func createRootDomain(d *deployer, ad *asset.ActiveDirectoryDomain) error {
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
	return nil
}

// Returns the root domain of the forest
func getRootDomain(d *deployer, ad *asset.ActiveDirectoryDomain) (*asset.ActiveDirectoryDomain, error) {
	for {
		if ad.Ancestor == nil {
			return ad, nil
		}

		parent, ok := ad.Ancestor.(*asset.ActiveDirectoryDomain_ParentName)
		if ok {
			parentAd, err := d.getAdDomainAsset(parent.ParentName)
			if err != nil {
				return nil, err
			}

			ad = parentAd
		} else {
			forest := ad.Ancestor.(*asset.ActiveDirectoryDomain_Forest)
			forestAd, err := d.getAdDomainAsset(forest.Forest)
			if err != nil {
				return nil, err
			}

			ad = forestAd
		}
	}
}

func createChildDomain(d *deployer, ad *asset.ActiveDirectoryDomain, parentAdName string) error {
	parentAd, err := d.getAdDomainAsset(parentAdName)
	if err != nil {
		return err
	}

	rootAd, err := getRootDomain(d, ad)
	if err != nil {
		return err
	}

	// name check
	if !(strings.HasSuffix(ad.Name, parentAd.Name) &&
		ad.Name[len(ad.Name)-len(parentAd.Name)-1] == '.') {
		return errors.Errorf(
			"Parent and child domain names are not valid. Parent %s, child %s",
			parentAd.Name, ad.Name)
	}

	// wait for parent domain to be ready
	depVar := onhost.GetActiveDirectoryRuntimeConfigVariableName(parentAd.Name)
	err = d.waitForDependency(depVar, time.Duration(60)*time.Minute)
	if err != nil {
		return err
	}

	fileToRun := ""
	if d.IsWindows2012() || d.IsWindows2016() {
		fileToRun = d.GetSupportingFilePath("create_child_domain_win2012.ps1")
	} else if d.IsWindows2008() {
		fileToRun = d.GetSupportingFilePath("create_child_domain_win2008.ps1")
	} else {
		return errors.New("unsupported windows version")
	}

	for count := 1; count <= 3; count++ {
		err = d.RunConfigCommand("powershell.exe", "-File", fileToRun,
			"-domainName", ad.Name[0:len(ad.Name)-len(parentAd.Name)-1],
			"-netbiosName", ad.NetbiosName,
			"-parentDomainName", parentAd.Name,

			// use parent domain controller as DNS, though,
			// it can be any DC in the forest where DNS is installed.
			"-dnsServer", parentAd.DomainController[0].WindowsMachine,

			// note that the admin is the admin on the forest root domain, not on the parent domain.
			"-adminName", rootAd.Name+"\\administrator",
			"-adminPassword", string(rootAd.SafeModeAdminPassword.Final),
			"-localAdminPassword", string(ad.SafeModeAdminPassword.Final))

		if err == nil || err == ErrRebootNeeded {
			// success or reboot needed
			return err
		}

		// retry
		d.Logf("Child domain creation failed. This might be caused " +
			"by DNS replication not done yet. Retry")
		time.Sleep(3 * time.Minute)
	}

	return err
}

func init() {
	common.RegisterResolverClass(&AdDomainResolver{})
}
