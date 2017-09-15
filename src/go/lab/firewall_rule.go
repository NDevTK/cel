// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package lab

import (
	"go/lab/config"
	compute "google.golang.org/api/compute/v1"
)

type GceVpcFirewallRule struct {
	BaseNamedAsset
	Network  *GceVpcNetwork
	Firewall *config.FirewallRule
}

func (f *GceVpcFirewallRule) Resolve(s *Session) (err error) {
	if fmap, ok := s.Config.Cloud.Firewalls[f.Network.id]; ok {
		if _, ok := fmap[f.id]; ok {
			// The named firewall rule already exists. Assume it is correct.
			return
		}
	}

	defer Action(&err, "creating firewall rule %s for network %s", f.id, f.Network.id)

	a := []*compute.FirewallAllowed{}
	for _, al := range f.Firewall.Allowed {
		if al.Protocol == "*" {
			a = append(a, &compute.FirewallAllowed{
				IPProtocol: "tcp",
				Ports:      al.Ports}, &compute.FirewallAllowed{
				IPProtocol: "udp",
				Ports:      al.Ports}, &compute.FirewallAllowed{
				IPProtocol: "icmp"})
		} else {
			a = append(a, &compute.FirewallAllowed{
				IPProtocol: al.Protocol,
				Ports:      al.Ports})
		}
	}

	fr := compute.Firewall{
		Name:        f.Firewall.Name,
		Description: f.Firewall.Description,
		Network:     f.Network.ComputeNetwork.SelfLink,
		SourceTags:  f.Firewall.SourceTag,
		TargetTags:  f.Firewall.TargetTag,
		Allowed:     a}
	op, err := s.GetComputeService().Firewalls.Insert(f.Network.Project.id, &fr).Context(s.Context).Do()
	if err != nil {
		return
	}

	return WaitForOperation(s, op)
}

func (g *GceVpcFirewallRule) Check(s *Session) (error, bool) {
	panic("not implemented")
}

func (g *GceVpcFirewallRule) Purge(s *Session) error {
	panic("not implemented")
}

func NewFirewallRule(A *Assets, n *GceVpcNetwork, cf *config.FirewallRule) *GceVpcFirewallRule {
	const kNamespace = "firewallRules"
	if a := A.Get(kNamespace, cf.Name); a != nil {
		return a.(*GceVpcFirewallRule)
	}

	f := GceVpcFirewallRule{BaseNamedAsset{kNamespace, cf.Name, []Asset{n}}, n, cf}
	A.Add(&f)
	return &f
}
