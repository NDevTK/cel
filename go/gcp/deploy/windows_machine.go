// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"chromium.googlesource.com/enterprise/cel/go/asset"
	"chromium.googlesource.com/enterprise/cel/go/common"
)

var startupScriptPath = common.RefPathMust("host.resources.startup.win_startup")
var agentX64Path = common.RefPathMust("host.resources.startup.win_agent_x64")

// WindowsMachine_AddStartupDependencies adds depedencies on the windows
// startup scripts and windows agent binaries to a machine.
func windowsMachine_AddStartupDeps(ctx common.Context, m *asset.WindowsMachine) error {
	err := ctx.PublishDependency(m, startupScriptPath)
	if err != nil {
		return err
	}

	return ctx.PublishDependency(m, agentX64Path)
}

func init() {
	common.RegisterResolverFunc(common.AdditionalDependencyResolverKind,
		windowsMachine_AddStartupDeps)
}
