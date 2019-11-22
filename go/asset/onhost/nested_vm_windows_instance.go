package onhost

import (
	"fmt"
	"path"
	"strings"
	"time"

	hostpb "chromium.googlesource.com/enterprise/cel/go/schema/host"
	"github.com/pkg/errors"
)

type nestedVMWindowsInstance struct {
	instanceBase
	nestedVMInstanceBase
	windowsInstanceBase
}

func (instance *nestedVMWindowsInstance) OnBoot() bool {
	return nestedVmInitialSetup(instance, []string{
		"-net", "nic,model=e1000",
		"-vga", "std",
	})
}

func (instance *nestedVMWindowsInstance) RunCommand(name string, arg ...string) (string, error) {
	return sshRunCommand(instance, name, arg...)
}

func (instance *nestedVMWindowsInstance) RunLocalCommand(name string, arg ...string) (string, error) {
	return RunCommand(&instance.instanceBase, name, arg...)
}

func (instance *nestedVMWindowsInstance) GetLocalSupportingFilePath(filename string) string {
	return path.Join(instance.d.directory, "supporting_files", filename)
}

func (instance *nestedVMWindowsInstance) GetSupportingFilePath(filename string) string {
	return path.Join(workingDirectoryOnNestedVM, "supporting_files", filename)
}

func (instance *nestedVMWindowsInstance) OneTimeSetup() bool {
	if err := setupWindowsNestedVM(instance); err != nil {
		instance.d.Logf("Common setup on nested VM failed: %+v", err)
		instance.d.setRuntimeConfigVariable(instance.machineConfigVar, statusError)
		return true
	}

	err := prepareInstance(instance)
	if err != nil {
		if IsRestarting(instance.d) {
			instance.d.Logf("Ignoring failure during restart: %s", err)
			return true
		}

		instance.d.Logf("Preparing the CEL instance failed: %s", err)
		instance.d.setRuntimeConfigVariable(instance.machineConfigVar, statusError)
		return true
	}

	return false
}

func prepareInstance(instance *nestedVMWindowsInstance) error {
	if instance.GetOs() != hostpb.OperatingSystem_WINDOWS {
		return nil
	}

	ver, err := getWindowsVersionByVerCommand(instance)
	if err != nil {
		return errors.WithStack(err)
	}
	instance.winVersion = ver

	if !(instance.d.IsWindows2008() || instance.d.IsWindows2012() || instance.d.IsWindows2016()) {
		return errors.New("unsupported windows version")
	}

	// Run shared Windows image preparation script.
	_, err = instance.RunCommand("powershell.exe", "-ExecutionPolicy", "ByPass",
		"-File", instance.GetSupportingFilePath("prepare_windows.ps1"))

	// Windows 2008 has extra setup steps - run them now.
	if instance.d.IsWindows2008() {
		_, err = instance.RunCommand("powershell.exe", "-ExecutionPolicy", "ByPass",
			"-File", instance.GetSupportingFilePath("prepare_win2008.ps1"))
		return err
	}

	return err
}

func setupWindowsNestedVM(instance *nestedVMWindowsInstance) error {
	err := waitUntilSshIsAlive(instance, 5*time.Minute)
	if err != nil {
		return errors.WithStack(err)
	}

	err = uploadSupportingFilesToNestedVM(instance)
	if err != nil {
		return errors.WithStack(err)
	}

	err = disableWindowsUpdate(instance)
	if err != nil {
		return errors.WithStack(err)
	}

	output, err := instance.RunCommand("hostname")
	if err != nil {
		return errors.WithStack(err)
	}
	hostname := strings.TrimRight(string(output), "\r\n")

	// Rename nestedVM if needed. Windows hostnames are not case-sentive.
	if !strings.EqualFold(hostname, instance.GetInstanceName()) {
		instance.Logf("Renaming nested VM from %v to %v.", hostname, instance.GetInstanceName())
		return renameNestedVM(instance, instance.GetInstanceName())
	}

	instance.Logf("Skip renaming nested VM because it's already named %v.", instance.GetInstanceName())
	return nil
}

func renameNestedVM(instance *nestedVMWindowsInstance, newName string) error {
	cmd := fmt.Sprintf(
		"powershell.exe -Command \"Rename-Computer -NewName %s -Force -PassThru -Restart\"",
		newName)
	_, err := sshRunCommand(instance, cmd)
	if err != nil {
		return err
	}

	return waitForNestedVMRebootComplete(instance)
}
