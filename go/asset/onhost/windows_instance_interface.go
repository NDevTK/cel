// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

// windowsInstanceInterface is the interface that represents an instance running Windows.
type windowsInstanceInterface interface {
	instanceInterface

	SetWindowsVersion(ver windowsVersion)
}

// Disables Windows Update and makes sure wuauserv doesn't run.
func disableWindowsUpdate(instance windowsInstanceInterface) error {
	// Disable auto-start.
	_, err := instance.RunCommand("sc", "config", "wuauserv", "start=", "disabled")
	if err != nil {
		return err
	}

	// Stop the service if it's running.
	output, err := instance.RunCommand("net", "stop", "wuauserv")
	if err != nil {
		// Ignore ExitError 2: The Windows Update service is not started
		exitError, ok := err.(*ssh.ExitError)
		if !ok || exitError.Waitmsg.ExitStatus() != 2 {
			return err
		}
	}

	// `net stop` returns 0/SUCCESS even when "service could not be stopped".
	// In that case, we'll kill WindowsUpdate (wuauserv) and TrustedInstaller.
	if strings.Contains(output, "The Windows Update service could not be stopped.") {
		if err = killServiceProcess(instance, "wuauserv"); err != nil {
			return err
		}

		return killServiceProcess(instance, "TrustedInstaller")
	}

	return nil
}

// Kills the service process on a CEL instance.
func killServiceProcess(instance windowsInstanceInterface, service string) error {
	// Get service information.
	output, err := instance.RunCommand("sc", "queryex", service)
	if err != nil {
		return err
	}

	// Find the service PID.
	for _, line := range strings.Split(output, "\n") {
		if strings.Contains(line, "PID") {
			parts := strings.Split(line, ": ")
			if len(parts) != 2 {
				return fmt.Errorf("The output of 'sc queryex' cannot be parsed: %s", parts)
			}

			pid := strings.TrimSpace(parts[1])

			// Kill the process if it has a PID. `sc queryex` can return 0 if
			// the service has stopped by the time we get here.
			if pid != "0" {
				_, err = instance.RunCommand("taskkill", "/f", "/pid", pid)
				return err
			}

			instance.Logf("Skip killing service process because PID is zero (%v).", pid)
			return nil
		}
	}

	return nil
}

func PrepareWindowsInstance(instance windowsInstanceInterface) error {
	ver, err := getWindowsVersionByVerCommand(instance)
	if err != nil {
		return errors.WithStack(err)
	}
	instance.SetWindowsVersion(ver)

	if !(ver == win2008 || ver == win2012 || ver == win2016) {
		return errors.New("unsupported windows version")
	}

	// Run shared Windows image preparation script.
	_, err = instance.RunCommand("powershell.exe", "-ExecutionPolicy", "ByPass",
		"-File", instance.GetSupportingFilePath("prepare_windows.ps1"))

	// Windows 2008 has extra setup steps - run them now.
	if ver == win2008 {
		_, err = instance.RunCommand("powershell.exe", "-ExecutionPolicy", "ByPass",
			"-File", instance.GetSupportingFilePath("prepare_win2008.ps1"))
		return err
	}

	return err
}

// getWindowsVersionByVerCommand parses the output of the "ver" command to get the Windows version string.
// The output of ver looks like this:
//   Microsoft Windows [Version 6.1.7601]
// The return value will be "6.1.7601"
func getWindowsVersionByVerCommand(instance windowsInstanceInterface) (windowsVersion, error) {
	out, err := instance.RunCommand("cmd", "ver")
	if err != nil {
		return other, err
	}

	return getWindowsVersion(out)
}

// versionString is the output of running "ver"
func getWindowsVersion(verOutput string) (windowsVersion, error) {
	ver := other

	// remove newlines from the output
	outStr := strings.Replace(verOutput, "\r\n", "", -1)
	tmp1 := strings.Index(outStr, "[Version ")
	tmp2 := strings.Index(outStr, "]")
	if tmp1 == -1 || tmp2 == -1 {
		return ver, fmt.Errorf("The output of 'ver' cannot be parsed: %s", outStr)
	}

	verString := outStr[tmp1+9 : tmp2]
	if strings.HasPrefix(verString, "10.0.") {
		ver = win2016
	} else if strings.HasPrefix(verString, "6.3.") {
		ver = win2012
	} else if strings.HasPrefix(verString, "6.1.") {
		ver = win2008
	}
	return ver, nil
}
