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
	if restart, err := setupVncServerWithGnomeDesktop(instance); err != nil || restart {
		return true
	}

	return nestedVmInitialSetup(instance, []string{
		// Please see go/cros-vm on starting a CrOS VM.
		// CeLab stands up CrOS using images from the Simple Chrome SDK.
		// The Qemu flags below are obtained by obtained by copying the
		// flags that the Simple Chrome SDK uses.
		"-cpu", "SandyBridge,-invpcid,-tsc-deadline,check,vmx=on",
		"-daemonize",
		"-device", "usb-tablet",
		"-device", "virtio-scsi-pci,id=scsi",
		"-device", "virtio-net,netdev=eth0",
		"-enable-kvm",
		"-m", "8G",
		"-net", "nic,model=virtio",
		"-netdev", "user,id=eth0,net=10.0.2.0/27,hostfwd=tcp:127.0.0.1:9222-:22",
		"-smp", "8",
		"-usb",
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
