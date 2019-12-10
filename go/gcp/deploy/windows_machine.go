// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"fmt"
	"strconv"

	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/gcp"
	"chromium.googlesource.com/enterprise/cel/go/gcp/onhost"
	assetpb "chromium.googlesource.com/enterprise/cel/go/schema/asset"
	commonpb "chromium.googlesource.com/enterprise/cel/go/schema/common"
	computepb "chromium.googlesource.com/enterprise/cel/go/schema/gcp/compute"
	runtimeconfigpb "chromium.googlesource.com/enterprise/cel/go/schema/gcp/runtimeconfig"
	hostpb "chromium.googlesource.com/enterprise/cel/go/schema/host"
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
func (*windowsMachine) ResolveAdditionalDependencies(ctx common.Context, m *assetpb.WindowsMachine) error {
	mt := common.Must(ctx.Indirect(m, "machine_type")).(*hostpb.MachineType)

	err := ctx.PublishDependency(m, projectPath)
	if err != nil {
		return err
	}

	_, ok := mt.Base.(*hostpb.MachineType_NestedVm)
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

func (*windowsMachine) ResolveConstructedAssets(ctx common.Context, m *assetpb.WindowsMachine) error {
	d := GetDeploymentManifest()

	// add runtime config variable for this windows machine
	variableName := onhost.GetWindowsMachineRuntimeConfigVariableName(m.Name)
	if err := d.Emit(nil, &runtimeconfigpb.Variable{
		Name:     "runtimeconfigVariable_" + variableName,
		Parent:   onhost.RuntimeconfigVariableParent,
		Variable: variableName,
		Text:     "",
	}); err != nil {
		return err
	}

	mt := common.Must(ctx.Indirect(m, "machine_type")).(*hostpb.MachineType)

	nestedVm, ok := mt.Base.(*hostpb.MachineType_NestedVm)
	if ok {
		return resolveNestedVM(ctx, m, nestedVm)
	}
	return resolveNormalMachineType(ctx, m)
}

func resolveNestedVM(ctx common.Context, m *assetpb.WindowsMachine, nestedVm *hostpb.MachineType_NestedVm) error {
	d := GetDeploymentManifest()
	p := common.Must(ctx.Get(projectPath)).(*hostpb.Project)
	si := common.Must(ctx.Get(serviceAccountPath)).(*google_iam_admin_v1.ServiceAccount)

	diskSizeGb := nestedVm.NestedVm.DiskSizeGb
	if diskSizeGb == 0 {
		diskSizeGb = 70
	}

	if err := d.Emit(nil, &computepb.Disk{
		Name:        m.Name + "-disk",
		Zone:        p.Zone,
		SourceImage: "projects/ubuntu-os-cloud/global/images/family/ubuntu-1604-lts",
		SizeGb:      strconv.FormatUint(diskSizeGb, 10),
		Licenses: []string{
			"https://www.googleapis.com/compute/v1/projects/vm-options/global/licenses/enable-vmx",
		},
	}); err != nil {
		return err
	}

	var cni []*computepb.NetworkInterface
	for _, ni := range m.NetworkInterface {
		if ni.FixedAddress != nil {
			return common.NewNotImplementedError("support for fixed_address in assetpb.Network")
		}
		np := common.Must(ctx.Indirect(ni, "network")).(*assetpb.Network)
		cni = append(cni, &computepb.NetworkInterface{
			Network: fmt.Sprintf("$(ref.%s.selfLink)", d.Ref(np)),

			// Enables external IP. For some reason, this is needed for instances
			// to download start up scripts.
			AccessConfigs: []*computepb.AccessConfig{
				{
					Type: "ONE_TO_ONE_NAT",
					Name: "External NAT",
				},
			},

			// Add the IPAlias. The alias IP will be assigned to the nested VM.
			AliasIpRanges: []*computepb.AliasIpRange{
				{
					IpCidrRange: "/32",
				},
			},
		})
	}

	ss := gcp.AbsoluteReference(
		common.Must(ctx.Get(storageBucketPath)).(string),
		common.Must(ctx.Get(linuxStartupScriptPath)).(*commonpb.FileReference).ObjectReference)

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
	md := &computepb.Metadata{
		Items: []*computepb.Metadata_Items{
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

	machineType := nestedVm.NestedVm.MachineType
	if machineType == "" {
		machineType = fmt.Sprintf("projects/%s/zones/%s/machineTypes/n1-standard-2", p.Name, p.Zone)
	}

	return d.Emit(m, &computepb.Instance{
		Name:              m.Name,
		Description:       "CEL VM",
		MachineType:       machineType,
		Zone:              p.Zone,
		CanIpForward:      true,
		NetworkInterfaces: cni,
		Disks: []*computepb.AttachedDisk{
			{
				AutoDelete: true,
				Boot:       true,
				DeviceName: m.Name + "-disk",
				Source:     fmt.Sprintf("$(ref.%s.selfLink)", m.Name+"-disk"),
			},
		},
		Metadata: md,
		ServiceAccounts: []*computepb.ServiceAccount{
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

func resolveNormalMachineType(ctx common.Context, m *assetpb.WindowsMachine) error {
	d := GetDeploymentManifest()
	p := common.Must(ctx.Get(projectPath)).(*hostpb.Project)
	mt := common.Must(ctx.Indirect(m, "machine_type")).(*hostpb.MachineType)
	si := common.Must(ctx.Get(serviceAccountPath)).(*google_iam_admin_v1.ServiceAccount)

	var cni []*computepb.NetworkInterface
	for _, ni := range m.NetworkInterface {
		if ni.FixedAddress != nil {
			return common.NewNotImplementedError("support for fixed_address in assetpb.Network")
		}
		np := common.Must(ctx.Indirect(ni, "network")).(*assetpb.Network)
		cni = append(cni, &computepb.NetworkInterface{
			Network: fmt.Sprintf("$(ref.%s.selfLink)", d.Ref(np)),

			// Enables external IP. For some reason, this is needed for instances
			// to download start up scripts.
			AccessConfigs: []*computepb.AccessConfig{
				{
					Type: "ONE_TO_ONE_NAT",
					Name: "External NAT",
				},
			},
		})
	}

	ss := gcp.AbsoluteReference(
		common.Must(ctx.Get(storageBucketPath)).(string),
		common.Must(ctx.Get(winStartupScriptPath)).(*commonpb.FileReference).ObjectReference)

	md := &computepb.Metadata{
		Items: []*computepb.Metadata_Items{
			&computepb.Metadata_Items{
				Key:   windowsStartupScriptUrlMetadataKey,
				Value: ss,
			},
		},
	}

	for _, i := range mt.GetInstanceProperties().Metadata.GetItems() {
		md.Items = append(md.Items, i)
	}

	return d.Emit(m, &computepb.Instance{
		Name:              m.Name,
		Description:       "CEL VM",
		MachineType:       mt.GetInstanceProperties().MachineType,
		Zone:              p.Zone,
		CanIpForward:      mt.GetInstanceProperties().CanIpForward,
		NetworkInterfaces: cni,
		Disks:             mt.GetInstanceProperties().Disks,
		Metadata:          md,
		ServiceAccounts: []*computepb.ServiceAccount{
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
