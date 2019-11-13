package onhost

// windowsInstance represents a GCP instance running Windows.
type windowsInstance struct {
	instanceBase
	windowsInstanceBase
}

func (instance *windowsInstance) OnBoot() bool {
	// A no-op

	// Return false to make deployment continue
	return false
}

func (instance *windowsInstance) RunCommand(name string, arg ...string) (string, error) {
	return RunCommand(&instance.instanceBase, name, arg...)
}

func (instance *windowsInstance) OneTimeSetup() bool {
	err := PrepareWindowsInstance(instance)
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
