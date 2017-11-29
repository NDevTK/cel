// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package asset

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"context"
)

func (w *WindowsMachine) Resolve(ctx context.Context, resolver common.ResolverService) (err error) {
	defer common.Action(&err, "Resolving WindowsMachine \"%s\"", w.Name)
	m, err := MachineProviderServiceFromContext(ctx)
	if err != nil {
		return
	}

	j := common.NewJobWaiter(w)

	var channel MachineChannel
	m.Provision(ctx, w, &channel, j)
	resolver.Resolve(ctx, w.Container, j)
	resolver.Resolve(ctx, w.ConfigurationFile, j)

	err = j.Join()
	if err != nil {
		return
	}

	// At this point the machine is running. We should be able to send commands to it.
	return
}
