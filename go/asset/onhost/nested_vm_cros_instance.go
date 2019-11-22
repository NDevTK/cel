package onhost

import (
	"path"
	"time"
)

type nestedVMCrosInstance struct {
	instanceBase
	nestedVMInstanceBase
}

func (instance *nestedVMCrosInstance) OnBoot() bool {
	return nestedVmInitialSetup(instance, []string{
		"-net", "nic,model=virtio",
		"-vga", "virtio",
	})
}

func (instance *nestedVMCrosInstance) RunCommand(name string, arg ...string) (string, error) {
	return sshRunCommand(instance, name, arg...)
}

func (instance *nestedVMCrosInstance) RunLocalCommand(name string, arg ...string) (string, error) {
	return RunCommand(&instance.instanceBase, name, arg...)
}

func (instance *nestedVMCrosInstance) GetLocalSupportingFilePath(filename string) string {
	return path.Join(instance.d.directory, "supporting_files", filename)
}

func (instance *nestedVMCrosInstance) GetSupportingFilePath(filename string) string {
	return path.Join(workingDirectoryOnNestedVM, "supporting_files", filename)
}

func (instance *nestedVMCrosInstance) OneTimeSetup() bool {
	if err := waitUntilSshIsAlive(instance, 2*time.Minute); err != nil {
		instance.d.Logf("Cros nested VM setup failed: %+v", err)
		instance.d.setRuntimeConfigVariable(instance.machineConfigVar, statusError)
		return true
	}

	return false
}

func (instance *nestedVMCrosInstance) GetWindowsVersion() windowsVersion {
	return other
}
