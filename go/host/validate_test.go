// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package host

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	compute "chromium.googlesource.com/enterprise/cel/go/schema/gcp/compute"
	hostpb "chromium.googlesource.com/enterprise/cel/go/schema/host"
	"reflect"
	"strings"
	"testing"

	// The following must be imported here for side-effects (e.g. loading package validators).
	_ "chromium.googlesource.com/enterprise/cel/go/asset"
	_ "chromium.googlesource.com/enterprise/cel/go/gcp/compute"
)

func TestHostEnvironment_validateFields(t *testing.T) {
	var h hostpb.HostEnvironment
	h.Project = &hostpb.Project{Name: "T", Zone: "Z"}
	h.LogSettings = &hostpb.LogSettings{AdminLog: "A"}
	h.Storage = &hostpb.Storage{Bucket: "x"}
	err := common.ValidateProto(&h, common.EmptyPath)
	if err != nil {
		t.Fatal("unexpected error ", err)
	}

	err = common.VerifyValidatableType(reflect.TypeOf(&h))
	if err != nil {
		t.Fatal("unexpected error: ", err)
	}
}

func TestHostEnvironment_validateStorage(t *testing.T) {
	var h hostpb.HostEnvironment
	h.Project = &hostpb.Project{Name: "T", Zone: "Z"}
	h.LogSettings = &hostpb.LogSettings{AdminLog: "A"}
	h.Storage = &hostpb.Storage{Bucket: "x", Prefix: "x/"}
	err := common.ValidateProto(&h, common.EmptyPath)
	if err == nil {
		t.Fatal()
	}

	if !strings.Contains(err.Error(), "must not end with a slash") {
		t.Fatal("unexpected error", err)
	}

	h.Storage.Prefix = "/x"
	err = common.ValidateProto(&h, common.EmptyPath)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMachineType_Validate(t *testing.T) {
	t.Run("noInstanceProperties", func(t *testing.T) {
		m := &hostpb.MachineType{Name: "foo"}
		if err := common.ValidateProto(m, common.EmptyPath); err == nil || !strings.Contains(err.Error(), "instance_properties") {
			t.Fatal("failed to catch missing instance properties")
		}
	})

	t.Run("hasNetworkInterfaces", func(t *testing.T) {
		m := &hostpb.MachineType{Name: "foo", Base: &hostpb.MachineType_InstanceProperties{
			InstanceProperties: &compute.InstanceProperties{
				NetworkInterfaces: []*compute.NetworkInterface{
					&compute.NetworkInterface{},
				},
				Disks: []*compute.AttachedDisk{&compute.AttachedDisk{
					AutoDelete: true,
					Boot:       true,
					InitializeParams: &compute.AttachedDiskInitializeParams{
						SourceImage: "some source image",
					},
				}},
			},
		}}
		if err := common.ValidateProto(m, common.EmptyPath); err == nil || !strings.Contains(err.Error(), "network_interfaces") {
			t.Fatal("failed to catch missing instance properties")
		}
	})

	t.Run("noDisks", func(t *testing.T) {
		m := &hostpb.MachineType{Name: "foo", Base: &hostpb.MachineType_InstanceProperties{
			InstanceProperties: &compute.InstanceProperties{},
		}}
		if err := common.ValidateProto(m, common.EmptyPath); err == nil || !strings.Contains(err.Error(), "at least one disk") {
			t.Fatal("failed to catch missing instance properties")
		}
	})

	t.Run("noAutoDelete", func(t *testing.T) {
		m := &hostpb.MachineType{Name: "foo", Base: &hostpb.MachineType_InstanceProperties{
			InstanceProperties: &compute.InstanceProperties{
				Disks: []*compute.AttachedDisk{&compute.AttachedDisk{
					Boot: true,
					InitializeParams: &compute.AttachedDiskInitializeParams{
						SourceImage: "some source image",
					},
				}},
			},
		}}
		if err := common.ValidateProto(m, common.EmptyPath); err == nil || !strings.Contains(err.Error(), "auto_delete") {
			t.Fatal("failed to catch missing instance properties")
		}
	})

	t.Run("isValid", func(t *testing.T) {
		m := &hostpb.MachineType{Name: "foo", Base: &hostpb.MachineType_InstanceProperties{
			InstanceProperties: &compute.InstanceProperties{
				Disks: []*compute.AttachedDisk{&compute.AttachedDisk{
					AutoDelete: true,
					Boot:       true,
					InitializeParams: &compute.AttachedDiskInitializeParams{
						SourceImage: "some source image",
					},
				}},
			},
		}}
		if err := common.ValidateProto(m, common.EmptyPath); err != nil {
			t.Fatal("unexpected error: ", err)
		}
	})
}
