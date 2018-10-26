// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package asset

import (
	"github.com/pkg/errors"
)

func (a *AssetManifest) FindIISServer(name string) (*IISServer, error) {
	for _, iisServer := range a.IisServer {
		if iisServer.Name == name {
			return iisServer, nil
		}
	}

	return nil, errors.Errorf("failed to find IIS Server '%s'", name)
}

func (a *AssetManifest) FindWindowsMachine(name string) (*WindowsMachine, error) {
	for _, machine := range a.WindowsMachine {
		if machine.Name == name {
			return machine, nil
		}
	}

	return nil, errors.Errorf("failed to find Windows Machine '%s'", name)
}

func (a *AssetManifest) FindActiveDirectoryDomain(name string) (*ActiveDirectoryDomain, error) {
	for _, ad := range a.AdDomain {
		if ad.Name == name {
			return ad, nil
		}
	}

	return nil, errors.Errorf("failed find ActiveDirectoryDomain '%s'", name)
}

func (a *AssetManifest) FindActiveDirectoryDomainFor(m *WindowsMachine) (*ActiveDirectoryDomain, error) {
	if m != nil {
		if m.Container != nil {
			// machine joining a domain
			ad, err := a.FindActiveDirectoryDomain(m.Container.GetAdDomain())
			if err == nil {
				return ad, nil
			}
		} else {
			// machine could be the Domain Controller
			for _, ad := range a.AdDomain {
				if ad.DomainController[0].WindowsMachine == m.Name {
					return ad, nil
				}
			}
		}

	}

	return nil, errors.Errorf("failed to find ActiveDirectoryDomain for '%v'", m)
}
