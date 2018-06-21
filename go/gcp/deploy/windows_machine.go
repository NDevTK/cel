// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"chromium.googlesource.com/enterprise/cel/go/asset"
	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/gcp"
	"chromium.googlesource.com/enterprise/cel/go/gcp/compute"
	"chromium.googlesource.com/enterprise/cel/go/gcp/onhost"
	"chromium.googlesource.com/enterprise/cel/go/host"
	"fmt"
	google_iam_admin_v1 "google.golang.org/genproto/googleapis/iam/admin/v1"
)

var startupScriptPath = common.RefPathMust("host.resources.startup.win_startup")
var agentX64Path = common.RefPathMust("host.resources.startup.win_agent_x64")
var projectPath = common.RefPathMust("host.project")
var serviceAccountPath = common.RefPathMust("host.resources.service_account")
var storageBucketPath = common.RefPathMust("host.storage.bucket")

const windowsStartupScriptMetadataKey = "windows-startup-script-url"

type windowsMachine struct{}

// ResolveAdditionalDependencies adds depedencies on the windows startup
// scripts and windows agent binaries to a machine.
func (*windowsMachine) ResolveAdditionalDependencies(ctx common.Context, m *asset.WindowsMachine) error {
	err := ctx.PublishDependency(m, startupScriptPath)
	if err != nil {
		return err
	}

	err = ctx.PublishDependency(m, projectPath)
	if err != nil {
		return err
	}

	return ctx.PublishDependency(m, agentX64Path)
}

func (*windowsMachine) ResolveConstructedAssets(ctx common.Context, m *asset.WindowsMachine) error {
	d := GetDeploymentManifest()

	// add runtime config variable for this windows machine
	variableName := onhost.GetActiveDirectoryRuntimeConfigVariableName(m.Name)
	if err := d.Emit(nil, &onhost.RuntimeConfigConfigVariable{
		Name:     "runtimeconfigVariable_" + variableName,
		Parent:   onhost.RuntimeconfigVariableParent,
		Variable: variableName,
		Text:     "",
	}); err != nil {
		return err
	}

	p := common.Must(ctx.Get(projectPath)).(*host.Project)
	mt := common.Must(ctx.Indirect(m, "machine_type")).(*host.MachineType)
	si := common.Must(ctx.Get(serviceAccountPath)).(*google_iam_admin_v1.ServiceAccount)

	var cni []*compute.NetworkInterface
	for _, ni := range m.NetworkInterface {
		if ni.FixedAddress != nil {
			return common.NewNotImplementedError("support for fixed_address in asset.Network")
		}
		np := common.Must(ctx.Indirect(ni, "network")).(*asset.Network)
		cni = append(cni, &compute.NetworkInterface{
			Network: fmt.Sprintf("$(ref.%s.selfLink)", d.Ref(np)),

			// Enables external IP. For some reason, this is needed for instances
			// to download start up scripts.
			AccessConfigs: []*compute.AccessConfig{
				{
					Type: "ONE_TO_ONE_NAT",
					Name: "External NAT",
				},
			},
		})
	}

	ss := gcp.AbsoluteReference(
		common.Must(ctx.Get(storageBucketPath)).(string),
		common.Must(ctx.Get(startupScriptPath)).(*common.FileReference).ObjectReference)

	md := &compute.Metadata{
		Items: []*compute.Metadata_Items{
			&compute.Metadata_Items{
				Key:   windowsStartupScriptMetadataKey,
				Value: ss,
			},
		},
	}

	for _, i := range mt.GetInstanceProperties().Metadata.GetItems() {
		md.Items = append(md.Items, i)
	}

	return d.Emit(m, &compute.Instance{
		Name:              m.Name,
		Description:       "CEL VM",
		MachineType:       mt.GetInstanceProperties().MachineType,
		Zone:              p.Zone,
		CanIpForward:      mt.GetInstanceProperties().CanIpForward,
		NetworkInterfaces: cni,
		Disks:             mt.GetInstanceProperties().Disks,
		Metadata:          md,
		ServiceAccounts: []*compute.ServiceAccount{
			{
				Email: si.Email,
				Scopes: []string{
					"https://www.googleapis.com/auth/devstorage.read_only",
					"https://www.googleapis.com/auth/logging.write",
					"https://www.googleapis.com/auth/compute.readonly",
					"https://www.googleapis.com/auth/cloudruntimeconfig",
				},
			},
		},
		Scheduling: mt.GetInstanceProperties().Scheduling,
	})
}

func init() {
	common.RegisterResolverClass(&windowsMachine{})
}
