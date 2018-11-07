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
	"chromium.googlesource.com/enterprise/cel/go/gcp/onhost"
	"chromium.googlesource.com/enterprise/cel/go/host"
	google_iam_admin_v1 "google.golang.org/genproto/googleapis/iam/admin/v1"
)

var winStartupScriptPath = common.RefPathMust("host.resources.startup.win_startup")
var winAgentX64Path = common.RefPathMust("host.resources.startup.win_agent_x64")
var linuxStartupScriptPath = common.RefPathMust("host.resources.startup.linux_startup")
var linuxAgentX64Path = common.RefPathMust("host.resources.startup.linux_agent_x64")
var projectPath = common.RefPathMust("host.project")
var serviceAccountPath = common.RefPathMust("host.resources.service_account")
var storageBucketPath = common.RefPathMust("host.storage.bucket")

const windowsStartupScriptUrlMetadataKey = "windows-startup-script-url"
const linuxStartupScriptUrlMetadataKey = "startup-script-url"
const linuxShutdownScriptMetadataKey = "shutdown-script"

type windowsMachine struct{}

// ResolveAdditionalDependencies adds depedencies on the windows startup
// scripts and windows agent binaries to a machine.
func (*windowsMachine) ResolveAdditionalDependencies(ctx common.Context, m *asset.WindowsMachine) error {
	mt := common.Must(ctx.Indirect(m, "machine_type")).(*host.MachineType)

	err := ctx.PublishDependency(m, projectPath)
	if err != nil {
		return err
	}

	_, ok := mt.Base.(*host.MachineType_NestedVm)
	if ok {
		// nested VM
		err = ctx.PublishDependency(m, linuxStartupScriptPath)
		if err != nil {
			return err
		}

		return ctx.PublishDependency(m, linuxAgentX64Path)
	} else {
		// normal Windows machine
		err = ctx.PublishDependency(m, winStartupScriptPath)
		if err != nil {
			return err
		}

		return ctx.PublishDependency(m, winAgentX64Path)
	}
}

func (*windowsMachine) ResolveConstructedAssets(ctx common.Context, m *asset.WindowsMachine) error {
	d := GetDeploymentManifest()

	// add runtime config variable for this windows machine
	variableName := onhost.GetWindowsMachineRuntimeConfigVariableName(m.Name)
	if err := d.Emit(nil, &onhost.RuntimeConfigConfigVariable{
		Name:     "runtimeconfigVariable_" + variableName,
		Parent:   onhost.RuntimeconfigVariableParent,
		Variable: variableName,
		Text:     "",
	}); err != nil {
		return err
	}

	mt := common.Must(ctx.Indirect(m, "machine_type")).(*host.MachineType)

	_, ok := mt.Base.(*host.MachineType_NestedVm)
	if ok {
		return resolveNestedVM(ctx, m)
	}
	return resolveNormalMachineType(ctx, m)
}

func resolveNestedVM(ctx common.Context, m *asset.WindowsMachine) error {
	d := GetDeploymentManifest()
	p := common.Must(ctx.Get(projectPath)).(*host.Project)
	si := common.Must(ctx.Get(serviceAccountPath)).(*google_iam_admin_v1.ServiceAccount)

	if err := d.Emit(nil, &compute.Disk{
		Name:        m.Name + "-disk",
		Zone:        p.Zone,
		SourceImage: "projects/ubuntu-os-cloud/global/images/family/ubuntu-1604-lts",
		SizeGb:      "70",
		Licenses: []string{
			"https://www.googleapis.com/compute/v1/projects/vm-options/global/licenses/enable-vmx",
		},
	}); err != nil {
		return err
	}

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

			// Add the IPAlias. The alias IP will be assigned to the nested VM.
			AliasIpRanges: []*compute.AliasIpRange{
				{
					IpCidrRange: "/32",
				},
			},
		})
	}

	ss := gcp.AbsoluteReference(
		common.Must(ctx.Get(storageBucketPath)).(string),
		common.Must(ctx.Get(linuxStartupScriptPath)).(*common.FileReference).ObjectReference)

	// script to tell KVM to shutdown the VM
	shutdownScript := `#!/bin/bash

# Send monitor command to kvm to shutdown the VM
nc 127.0.0.1 25555 <<JSON
{ "execute": "qmp_capabilities" }
{ "execute": "system_powerdown" }
JSON

# wait until kvm stops
while true; do
  kvm_count=$(ps ax | grep kvm | wc -l)
  echo "kvm process count: $kvm_count"
  if [ "$kvm_count" = "1" ]; then
	break
  fi

  sleep 1
done
`
	md := &compute.Metadata{
		Items: []*compute.Metadata_Items{
			{
				Key:   linuxStartupScriptUrlMetadataKey,
				Value: ss,
			},
			{
				Key:   linuxShutdownScriptMetadataKey,
				Value: shutdownScript,
			},
			{
				// without this, SSH to the instance won't work
				Key:   "enable-oslogin",
				Value: "TRUE",
			},
		},
	}

	return d.Emit(m, &compute.Instance{
		Name:              m.Name,
		Description:       "CEL VM",
		MachineType:       fmt.Sprintf("projects/%s/zones/%s/machineTypes/n1-standard-2", p.Name, p.Zone),
		Zone:              p.Zone,
		CanIpForward:      true,
		NetworkInterfaces: cni,
		Disks: []*compute.AttachedDisk{
			{
				AutoDelete: true,
				Boot:       true,
				DeviceName: m.Name + "-disk",
				Source:     fmt.Sprintf("$(ref.%s.selfLink)", m.Name+"-disk"),
			},
		},
		Metadata: md,
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
	})
}

func resolveNormalMachineType(ctx common.Context, m *asset.WindowsMachine) error {
	d := GetDeploymentManifest()
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
		common.Must(ctx.Get(winStartupScriptPath)).(*common.FileReference).ObjectReference)

	md := &compute.Metadata{
		Items: []*compute.Metadata_Items{
			&compute.Metadata_Items{
				Key:   windowsStartupScriptUrlMetadataKey,
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
