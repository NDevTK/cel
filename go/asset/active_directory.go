// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package asset

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"context"
)

func (w *WindowsMachine) Resolve(ctx context.Context, p common.RefPath, host HostService) (err error) {
	common.Action(&err, "Resolving WindowsMachine \"%s\"", w.Name)
	j := common.NewJobWaiter()

	host.ProvisionMachine(ctx, p, &Machine{
		Name:             w.Name,
		MachineType:      w.MachineType,
		NetworkInterface: w.NetworkInterface}, j.Collect())

	err = j.Join()

	if err != nil {
		return err
	}

	// Ensure network parameters.
	// Ensure machine type is usable.
	return nil
}

func (w *WindowsMachine) Purge(ctx context.Context, p common.RefPath, host HostService) (err error) {
	common.Action(&err, "Purging WindowsMachine \"%s\"", w.Name)
	j := common.NewJobWaiter()
	host.PurgeMachine(ctx, p, &Machine{
		Name:             w.Name,
		MachineType:      w.MachineType,
		NetworkInterface: w.NetworkInterface}, j.Collect())
	return j.Join()
}
