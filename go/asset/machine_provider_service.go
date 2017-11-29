// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package asset

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/meta"
	"context"
)

type BaseMachine interface {
	GetName() string
	GetMachineType() string
	GetNetworkInterface() []*NetworkInterface
}

type MachineProviderService interface {
	// ProvisionMachineAsync provisions a virtual or physcal machine.
	//
	// Upon successful completion, the following conditions must hold:
	//
	// * The specified machine must exist.
	// * It must conform to its machine type.
	// * It must conform to the network parameters as specified in its list of interfaces.
	// * It must have the latest lab on-host tools and scripts available.
	// * It must be powered on and ready to accept further commands. I.e. be running the host agent.
	// * All output parameters of Machine must be published within ctx.
	Provision(ctx context.Context, machine BaseMachine, channel *MachineChannel, waiter *common.JobWaiter)
}

type machineProviderServiceKeyType int

const machineProviderServiceKey machineProviderServiceKeyType = 0

func MachineProviderServiceFromContext(ctx context.Context) (MachineProviderService, error) {
	o, ok := ctx.Value(machineProviderServiceKey).(MachineProviderService)
	if !ok {
		return nil, &common.ServiceNotFoundError{Service: "MachineProvider"}
	}
	return o, nil
}

func ContextWithMachineProviderService(ctx context.Context, o MachineProviderService) context.Context {
	return context.WithValue(ctx, machineProviderServiceKey, o)
}

type MachineChannel interface {
	Invoke(ctx context.Context, command *meta.Command, response *meta.Response, result chan<- error)
}
