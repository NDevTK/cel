// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

import hostpb "chromium.googlesource.com/enterprise/cel/go/schema/host"

// instanceInterface is the interface that represents an instance.
// The instance can be
// - a GCP instance. E.g. a GCP instance running Windows 2012R2, or a GCP instance running Debian;
// - an instance that is run as a nested VM by a GCP host instance.
type instanceInterface interface {
	// Runs a command on the instance. Returns the tuple of output and error.
	RunCommand(name string, arg ...string) (string, error)

	// Setup of the instance after boot/restart. This step always runs after boot/restart,
	// no matter what the status is.
	// Return value indicates whether deployer should stop or continue.
	// True to stop, false to continue.
	OnBoot() bool

	// Setup before OnHost configuration starts. This step is a one time setup, so it
	// will not run again after a reboot if status is ready or error.
	OneTimeSetup() bool

	// GetSupportingFilePath returns the full path of the supporting file on the instance.
	GetSupportingFilePath(filename string) string

	// GetOs returns the operating system of the instance.
	GetOs() hostpb.OperatingSystem

	// GetWindowsVersion returns the Windows version of the instance.
	// For non-Windows instances, it returns 'other'.
	GetWindowsVersion() windowsVersion

	Logf(format string, arg ...interface{})

	// GetCurrentDirectory gets the current directory that is used by the onhost deployer.
	GetCurrentDirectory() string

	// GetStatus gets the status of the machine config variable associated with this instance.
	GetStatus() string

	// SetStatus sets the status of the machine config variable associated with this instance.
	SetStatus(status string)

	// GetInstanceName gets the name of the instance.
	GetInstanceName() string
}
