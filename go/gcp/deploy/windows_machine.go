// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"chromium.googlesource.com/enterprise/cel/go/asset"
	"chromium.googlesource.com/enterprise/cel/go/common"
)

var startup_script_path = common.RefPathMust("host.resources.startup.win_startup")
var agent_x64_path = common.RefPathMust("host.resources.startup.win_agent_x64")

// WindowsMachine_AddStartupDependencies adds depedencies on the windows
// startup scripts and windows agent binaries to a machine.
func windowsMachine_AddStartupDeps(ctx common.Context, m *asset.WindowsMachine) error {
	err := ctx.PublishDependency(m, startup_script_path)
	if err != nil {
		return err
	}

	return ctx.PublishDependency(m, agent_x64_path)
}

func init() {
	common.RegisterResolverFunc(common.AdditionalDependencyResolverKind,
		windowsMachine_AddStartupDeps)
}
