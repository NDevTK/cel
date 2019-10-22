// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"fmt"

	"chromium.googlesource.com/enterprise/cel/go/asset"
	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/gcp"
	"chromium.googlesource.com/enterprise/cel/go/gcp/compute"
)

type network struct{}

func (*network) ResolveConstructedAssets(ctx common.Context, n *asset.Network) error {
	d := GetDeploymentManifest()

	if n.AddressRange != nil {
		return common.NewNotImplementedError("asset.network.{}.address_range")
	}

	s, err := gcp.SessionFromContext(ctx)
	if err != nil {
		return err
	}

	if s.AllowExternalRdpSsh {
		s.Logger.Info(common.MakeStringer("Creating rule %s.", n.Name+"-allow-rdp-ssh"))
		if err := d.Emit(nil, &compute.Firewall{
			Name:      n.Name + "-allow-rdp-ssh",
			Network:   fmt.Sprintf("$(ref.%s.selfLink)", n.Name),
			Direction: "INGRESS",
			Allowed: []*compute.Firewall_Allowed{
				&compute.Firewall_Allowed{
					IPProtocol: "tcp",
					Ports:      []string{"3389", "22"},
				},
			},
		}); err != nil {
			return err
		}
	} else {
		s.Logger.Info(common.MakeStringer("Skipping rule %s.", n.Name+"-allow-rdp-ssh"))
	}

	if err := d.Emit(nil, &compute.Firewall{
		Name:         n.Name + "-allow-internal",
		Description:  "Allow internal traffic on the network",
		Network:      fmt.Sprintf("$(ref.%s.selfLink)", n.Name),
		Direction:    "INGRESS",
		SourceRanges: []string{"10.128.0.0/9"},
		Allowed: []*compute.Firewall_Allowed{
			&compute.Firewall_Allowed{
				IPProtocol: "tcp",
			},
			&compute.Firewall_Allowed{
				IPProtocol: "udp",
			},
			&compute.Firewall_Allowed{
				IPProtocol: "icmp",
			},
		},
	}); err != nil {
		return err
	}

	return d.Emit(n, &compute.Network{
		Name:                  n.Name,
		Description:           "",
		AutoCreateSubnetworks: true,
	})
}

func init() {
	common.RegisterResolverClass(&network{})
}
