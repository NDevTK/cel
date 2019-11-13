// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

import (
	"os/exec"
	"path"

	hostpb "chromium.googlesource.com/enterprise/cel/go/schema/host"
)

// base struct for instance implementations
type instanceBase struct {
	// the deployer
	d *deployer

	// the machine config variable associated with this instance
	machineConfigVar string

	// the operating system of the instance
	operatingSystem hostpb.OperatingSystem
}

func (instance *instanceBase) Logf(format string, arg ...interface{}) {
	instance.d.Logf(format, arg...)
}

func (instance *instanceBase) GetCurrentDirectory() string {
	return instance.d.directory
}

func (instance *instanceBase) GetInstanceName() string {
	return instance.d.instanceName
}

func (instance *instanceBase) GetOs() hostpb.OperatingSystem {
	return instance.operatingSystem
}

func (instance *instanceBase) GetStatus() string {
	return instance.d.getRuntimeConfigVariableValue(instance.machineConfigVar)
}

func (instance *instanceBase) SetStatus(status string) {
	instance.d.setRuntimeConfigVariable(instance.machineConfigVar, status)
}

// RunCommand runs a command on the instance.
func RunCommand(instance *instanceBase, name string, arg ...string) (string, error) {
	instance.Logf("Run command: %s, args: %s", name, arg)
	output, err := exec.Command(name, arg...).CombinedOutput()
	if output != nil {
		instance.Logf("Output of command %s, args %s, err %v, is:\n%s", name, arg, err, output)
	}

	return string(output), err
}

func (instance *instanceBase) GetSupportingFilePath(filename string) string {
	return path.Join(instance.GetCurrentDirectory(), "supporting_files", filename)
}
