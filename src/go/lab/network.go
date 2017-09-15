// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package lab

import (
	"go/lab/config"
	compute "google.golang.org/api/compute/v1"
	"strings"
)

const kNetworksNamespace = "networks"
const kGceVpcNetworkNamespace = "vpc-networks"

// GceVpcNetwork represents a GCE VPC network asset. See
// https://cloud.google.com/compute/docs/vpc/ for what these words mean. The
// only kind of VPC network we currently support are those with automatic
// subnetting over regions. Hence we don't allow configuring individual
// subnets. Overall, we don't intend to support tests that rely on network
// resources being on a specific subnet.
type GceVpcNetwork struct {
	BaseNamedAsset
	Project        *Project
	Network        *config.Network
	ComputeNetwork *compute.Network
}

func (n *GceVpcNetwork) Resolve(s *Session) (err error) {
	defer Action(&err, "ensuring network %s exists", n.id)

	// Network already exists.
	var ok bool
	if n.ComputeNetwork, ok = s.Config.Cloud.Networks[n.id]; ok {
		return
	}

	defer Action(&err, "creating network %s", n.id)
	op, err := s.GetComputeService().Networks.Insert(n.Project.id, &compute.Network{
		Name:                  n.Network.Name,
		Description:           n.Network.Description,
		AutoCreateSubnetworks: true}).Context(s.Context).Do()

	if err != nil {
		return
	}

	err = WaitForOperation(s, op)
	if err != nil {
		return
	}

	n.ComputeNetwork, err = s.GetComputeService().Networks.Get(n.Project.id, n.id).Context(s.Context).Do()
	return
}

func (n *GceVpcNetwork) Check(s *Session) (error, bool) {
	panic("not implemented")
}

func (n *GceVpcNetwork) Purge(s *Session) error {
	panic("not implemented")
}

// Network represents an existing, configured and running network asset.
type Network struct {
	BaseNamedAsset
	VpcNetwork     *GceVpcNetwork
	ComputeNetwork *compute.Network
}

func (n *Network) Resolve(s *Session) error {
	n.ComputeNetwork = n.VpcNetwork.ComputeNetwork
	return nil
}

func (n *Network) Check(s *Session) (error, bool) {
	panic("not implemented")
}

func (n *Network) Purge(s *Session) (err error) {
	panic("not implemented")
}

// SubnetworkForRegion returns the partial URL for the subnetwork of |n| within
// region |region|. If |n| is an auto-mode GCE VPC, then |n| should have a
// defined subnetwork for each valid region. Hence this function should always
// succeed for a successfully resolved network assuming |region| is valid.
//
// See https://cloud.google.com/compute/docs/vpc/ for details of how
// subnetworks map to regions within an auto-mode GCE VPC.
func (n *Network) SubnetworkForRegion(region string) (string, error) {
	if n.ComputeNetwork == nil {
		return "", NewError("not resolved")
	}

	needle := "/regions/" + region + "/subnetworks/"
	for _, subnet := range n.ComputeNetwork.Subnetworks {
		if strings.Contains(subnet, needle) {
			return subnet, nil
		}
	}

	return "", NewError("not found")
}

func LookupNetwork(A *Assets, id string) *Network {
	if a := A.Get(kNetworksNamespace, id); a != nil {
		return a.(*Network)
	}
	panic("network not found")
}

// ConstructNetworkAssets constructs the assets and their dependents for all the
// |Network| resources in Config |c|. The dependency structure is as follows:
//
//                 +----------------+
//                 |                |
//                 |   VPC network  |
//                 |                |
//            +----+--------+-------+------+
//            |             |              |
//            |             |              |
//            v             v              v
//  +---------+---+ +-------+------+ +-----+-------+
//  |             | |              | |             |
//  | firewallRule| | firewallRule | | firewallRule|
//  |             | |              | |             |
//  +------+------+ +-------+------+ +------+------+
//         |                |               |
//         |                |               |
//         |                |               |
//         |       +--------v-------+       |
//         |       |                |       |
//         +------>+     network    +<------+
//                 |                |
//                 +----------------+
//
// I.e. the |networks| asset depends on the |firewallRule| assets, which in
// turn depend on the |VPC network| asset. This ensures that any asset that
// depends on |network| will be resolved *after* the network is created along
// with its firewall rules.
func ConstructNetworkAssets(A *Assets, c *Config) (err error) {
	p := LookupProject(A, c.Project)
	for _, config_net := range c.Network {
		deps := []Asset{p}

		vpc_net := GceVpcNetwork{
			BaseNamedAsset{kGceVpcNetworkNamespace, config_net.Name, []Asset{p}}, p, config_net, nil}
		A.Add(&vpc_net)

		deps = append(deps, &vpc_net)

		for _, config_fr := range config_net.FirewallRule {
			fr := NewFirewallRule(A, &vpc_net, config_fr)
			deps = append(deps, fr)
		}

		n := Network{BaseNamedAsset{kNetworksNamespace, config_net.Name, deps}, &vpc_net, nil}
		err = A.Add(&n)
		if err != nil {
			return
		}
	}
	return
}
