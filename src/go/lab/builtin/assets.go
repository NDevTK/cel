// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package builtin

import (
	"go/lab/config"
)

// Describes the builtin assets that are required by the lab.

const (
	GreeterAccountName     = "greeter"
	GreeterHostName        = "greeter-host"
	GreeterIpName          = "greeter-ip"
	DeployerAccountName    = "deployer"
	DeployerHostName       = "deployer-host"
	NetworkName            = "astral-plane"
	NetworkAddressRange    = "10.128.0.0/16"
	SubnetworkAddressRange = "10.128.1.0/24"
	ImageName              = "admin-os"
	InstanceTypeName       = "admin-instances"
	RegionName             = "us-east1"
	ZoneName               = "us-east1-b"
	MachineTypeName        = "n1-standard-4"
)

var Assets = config.Assets{
	ServiceAccount: []*config.ServiceAccount{
		{
			Id:          GreeterAccountName,
			DisplayName: "Lab greeter (Builtin)"},

		{
			Id:          DeployerAccountName,
			DisplayName: "Lab deployer (Builtin)"}},

	Network: []*config.Network{{
		Name:                NetworkName,
		Description:         "Administrative network (Builtin)",
		Ipv4AddressRange:    NetworkAddressRange,
		AllowInternalAccess: true,
		Subnetwork: []*config.Network_Subnetwork{{
			Name:             RegionName,
			Region:           RegionName,
			Ipv4AddressRange: SubnetworkAddressRange}}}},

	SourceImage: []*config.SourceImage{{
		Name: ImageName,
		SourceType: &config.SourceImage_Latest_{
			Latest: &config.SourceImage_Latest{
				Family:  "windows-2016-core",
				Project: "windows-cloud"}}}},

	InstanceType: []*config.InstanceType{{
		Name: InstanceTypeName,
		CreateOptions: &config.InstanceCreateOptions{
			Zone:        ZoneName,
			Image:       ImageName,
			MachineType: MachineTypeName}}},

	Instance: []*config.Instance{
		{
			Name:        GreeterHostName,
			Description: "Greeter : Tasked with accepting incoming test requests and running them",
			Type:        InstanceTypeName,
			CreateOptions: &config.InstanceCreateOptions{
				Interface: []*config.NetworkInterface{{
					Network:           NetworkName,
					InternalIp:        "10.128.1.2",
					ExternallyVisible: true,
					ExternalIpName:    GreeterIpName}}}},
		{
			Name:        DeployerHostName,
			Description: "Deployer : Host responsible for deploying Google Cloud assets",
			Type:        InstanceTypeName,
			CreateOptions: &config.InstanceCreateOptions{
				Interface: []*config.NetworkInterface{{
					Network:           NetworkName,
					InternalIp:        "10.128.1.3",
					ExternallyVisible: false}}}}}}
