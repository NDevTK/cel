// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

import (
	hostpb "chromium.googlesource.com/enterprise/cel/go/schema/host"
)

// base struct for nested VM instance implementations
type nestedVMInstanceBase struct {
	// The nested VM if this instance is a host.
	nestedVM *hostpb.NestedVM

	// The internal IP addresses for hosted VM. E.g. 192.168.122.89
	internalIP string

	// The external IP addresses for hosted VM. E.g. 10.128.0.2
	externalIP string
}

func (instance *nestedVMInstanceBase) GetNestedVM() *hostpb.NestedVM {
	return instance.nestedVM
}

func (instance *nestedVMInstanceBase) SetInternalIP(ip string) {
	instance.internalIP = ip
}

func (instance *nestedVMInstanceBase) GetInternalIP() string {
	return instance.internalIP
}

func (instance *nestedVMInstanceBase) SetExternalIP(ip string) {
	instance.externalIP = ip
}
