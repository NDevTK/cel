// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

import "fmt"

const CelConfigName = "cel-config"

// The cel-config name for deployment manager
const RuntimeconfigConfigName = "runtimeconfig-cel-config"

// The parent of all cel RuntimeConfig variables
const RuntimeconfigVariableParent = "$(ref.runtimeconfig-cel-config.name)"

// Returns the runtimeconfig variable's name for ActiveDirectory
func GetActiveDirectoryRuntimeConfigVariableName(adName string) string {
	return fmt.Sprintf("asset/ad_domain/%s/status", adName)
}

// Returns the runtimeconfig variable's name for a machine
func GetWindowsMachineRuntimeConfigVariableName(machineName string) string {
	return fmt.Sprintf("asset/windows_machine/%s/status", machineName)
}
