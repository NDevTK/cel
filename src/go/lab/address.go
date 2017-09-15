// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package lab

import (
	"go/lab/config"
	compute "google.golang.org/api/compute/v1"
)

type gceAddress struct {
	BaseNamedAsset
	Project        *Project
	Zone           string
	ConfigAddress  *config.ExternalIP
	ComputeAddress *compute.Address
}

const (
	kGceAddressNamespace = "gceAddresses"
	kAddressNamespace    = "addresses"
)

func (a *gceAddress) ComputeAddressTemplate() (ca *compute.Address) {
	return &compute.Address{
		Name:        a.ConfigAddress.Name,
		Description: a.ConfigAddress.Description,
		IpVersion:   a.ConfigAddress.Version.String()}
}

func (a *gceAddress) Resolve(s *Session) (err error) {
	if ca, ok := s.Config.Cloud.Addresses[a.ConfigAddress.Name]; ok {
		a.ComputeAddress = ca
		return
	}

	defer Action(&err, "reserving %s address %s", a.ConfigAddress.Version.String(),
		a.ConfigAddress.Name)

	region := s.Config.Cloud.Zones[a.Zone].Region
	op, err := s.GetComputeService().Addresses.Insert(a.Project.id, region, a.ComputeAddressTemplate()).
		Context(s.Context).Do()
	if err != nil {
		return
	}

	err = WaitForOperation(s, op)
	if err != nil {
		return err
	}

	a.ComputeAddress, err = s.GetComputeService().Addresses.Get(a.Project.id, region,
		a.ConfigAddress.Name).Context(s.Context).Do()
	return
}

func (a *gceAddress) Check(s *Session) (error, bool) {
	panic("not implemented")
}

func (a *gceAddress) Purge(s *Session) error {
	panic("not implemented")
}

type Address struct {
	BaseNamedAsset
	GceAddress     *gceAddress
	ComputeAddress *compute.Address
}

func (a Address) Resolve(s *Session) error {
	a.ComputeAddress = a.GceAddress.ComputeAddress
	return nil
}

func (a Address) Check(s *Session) (error, bool) {
	panic("not implemented")
}

func (a Address) Purge(s *Session) error {
	panic("not implemented")
}

func NewAddress(A *Assets, p *Project, zone string,
	ca *config.ExternalIP) *Address {

	gce_address := gceAddress{BaseNamedAsset{kGceAddressNamespace, ca.Name, []Asset{p}}, p, zone, ca, nil}
	A.Add(&gce_address)

	a := Address{BaseNamedAsset{kAddressNamespace, ca.Name, []Asset{&gce_address}}, &gce_address, nil}
	A.Add(&a)

	return &a
}
