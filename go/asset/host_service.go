// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package asset

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"context"
)

type HostService interface {
	ResolveAsync(context.Context, string, chan<- error)

	// ProvisionMachineAsync provisions a virtual or physcal machine.
	//
	// Upon successful completion, the following conditions must hold:
	//
	// * The specified machine must exist.
	// * It must conform to its machine type.
	// * It must conform to the network parameters as specified in its list of interfaces.
	// * It must be powered on and ready to accept further commands.
	// * All output parameters of Machine must be published within ctx.
	ProvisionMachine(ctx context.Context, p common.RefPath, machine *Machine, result chan<- error)

	PurgeMachine(ctx context.Context, p common.RefPath, machine *Machine, result chan<- error)
}
