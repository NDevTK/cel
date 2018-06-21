// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

import "fmt"

const CelConfigName = "cel-config"

// The type of RuntimeConfig config for deployment manager
const RuntimeconfigConfigType = "runtimeconfig.v1beta1.config"

// The type of RuntimeConfig config for deployment manager
const RuntimeconfigVariableType = "runtimeconfig.v1beta1.variable"

// The cel-config name for deployment manager
const RuntimeconfigConfigName = "runtimeconfig-cel-config"

// The parent of all cel RuntimeConfig variables
const RuntimeconfigVariableParent = "$(ref.runtimeconfig-cel-config.name)"

// RuntimeConfigConfig & RuntimeConfigVariable used to generate deployment manager
// YAML are defined in this file.

// Note that we cannot use classes generated from runtimeconfig-api.json,
// since the property names of the generated classes are not the names required
// by the deployment manager.
// E.g. the properties of RuntimeConfig config are "name" & "description"
// (see // https://cloud.google.com/deployment-manager/runtime-configurator/reference/rest/v1beta1/projects.configs),
// while the properties required by deployment manager are "config" & "description".
// (See https://cloud.google.com/deployment-manager/runtime-configurator/create-and-delete-runtimeconfig-resources)

type RuntimeConfigConfig struct {
	Name        string `json:"-"`
	Config      string `json:"config,omitempty"`
	Description string `json:"description,omitempty"`
}

func (RuntimeConfigConfig) Reset()             {}
func (RuntimeConfigConfig) String() string     { return "" }
func (RuntimeConfigConfig) ProtoMessage()      {}
func (r *RuntimeConfigConfig) GetName() string { return r.Name }

type RuntimeConfigConfigVariable struct {
	Name     string `json:"-"`
	Parent   string `json:"parent,omitempty"`
	Variable string `json:"variable,omitempty"`
	Text     string `json:"text,omitempty"`
}

func (RuntimeConfigConfigVariable) Reset()             {}
func (RuntimeConfigConfigVariable) String() string     { return "" }
func (RuntimeConfigConfigVariable) ProtoMessage()      {}
func (r *RuntimeConfigConfigVariable) GetName() string { return r.Name }

// Returns the runtimeconfig variable's name for ActiveDirectory
func GetActiveDirectoryRuntimeConfigVariableName(adName string) string {
	return fmt.Sprintf("asset/ad_domain/%s/status", adName)
}

// Returns the runtimeconfig variable's name for a machine
func GetWindowsMachineRuntimeConfigVariableName(machineName string) string {
	return fmt.Sprintf("asset/windows_machine/%s/status", machineName)
}
