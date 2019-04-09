// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"syscall"
	"time"
	"unicode/utf8"

	"chromium.googlesource.com/enterprise/cel/go/host"

	"chromium.googlesource.com/enterprise/cel/go/asset"
	"chromium.googlesource.com/enterprise/cel/go/cel"
	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/gcp/onhost"
	"chromium.googlesource.com/enterprise/cel/go/lab"
	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/logging"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
	"google.golang.org/api/googleapi"
	runtimeconfig "google.golang.org/api/runtimeconfig/v1beta1"
)

const statusError = "error"
const statusReady = "ready"
const statusInProgress = "in-progress"

// the working directory on the nested VM
const workingDirectoryOnNestedVM = "c:\\cel"

type windowsVersion int

const (
	// list of supported windows versions
	win2008 windowsVersion = iota + 1
	win2012
	win2016

	// this represents unsupported windows versions
	other
)

type deployer struct {
	ctx           context.Context
	service       *compute.Service
	project       *compute.Project
	loggingClient *logging.Client
	configService *runtimeconfig.ProjectsConfigsService
	instanceName  string

	// directory where cel_agent.exe is
	directory string

	// the windows version
	winVersion windowsVersion

	configuration *cel.Configuration

	// The nested VM if this instance is a host.
	nestedVM *host.NestedVM

	// The internal IP addresses for hosted VM. E.g. 192.168.122.89
	internalIP string

	// The external IP addresses for hosted VM. E.g. 10.128.0.2
	externalIP string
}

// Implement interface common.context
func (d *deployer) Deadline() (time.Time, bool) {
	return d.ctx.Deadline()
}

func (d *deployer) Done() <-chan struct{} {
	return d.ctx.Done()
}

func (d *deployer) Err() error {
	return d.ctx.Err()
}

func (d *deployer) Value(key interface{}) interface{} {
	return nil
}

func (d *deployer) Publish(m proto.Message, field string, value interface{}) error {
	return nil
}

func (d *deployer) PublishDependency(m proto.Message, dependsOn common.RefPath) error {
	return d.configuration.GetNamespace().PublishDependency(m, dependsOn)
}

func (d *deployer) GetObjectStore() common.ObjectStore {
	return nil
}

func (d *deployer) Debug(v fmt.Stringer) {
	d.Logf(v.String())
}

func (d *deployer) Info(v fmt.Stringer) {
	d.Logf(v.String())
}

func (d *deployer) Warning(v fmt.Stringer) {
	d.Logf(v.String())
}

func (d *deployer) Error(v fmt.Stringer) {
	d.Logf(v.String())
}

func (d *deployer) Get(p common.RefPath) (interface{}, error) {
	return nil, nil
}

func (d *deployer) Indirect(m proto.Message, f string) (interface{}, error) {
	return d.configuration.GetNamespace().Indirect(m, f)
}

func (d *deployer) Close() error {
	return d.loggingClient.Close()
}

// CreateDeployer creates the deployer
func CreateDeployer() (*deployer, error) {
	timeOut := 5 * time.Minute
	for start := time.Now(); time.Since(start) < timeOut; {
		d, err := createDeployer()
		if err == nil {
			return d, nil
		}

		log.Printf("Deployer creation failed. error: %s. Retry", err)
		time.Sleep(1 * time.Minute)
	}

	return nil, errors.New("CreateDeployer failed")
}

func createDeployer() (*deployer, error) {
	ex, err := os.Executable()
	if err != nil {
		return nil, err
	}

	exPath := filepath.Dir(ex)
	log.Printf("Running in directory %s\n", exPath)

	projectId, err := metadata.ProjectID()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ctx := context.Background()
	c, err := google.DefaultClient(ctx, compute.ComputeScope)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	project, err := computeService.Projects.Get(projectId).Context(ctx).Do()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	hostname, err := metadata.InstanceName()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	loggingClient, err := logging.NewClient(ctx, projectId)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	runtimeConfigService, err := runtimeconfig.New(c)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &deployer{
		ctx:           ctx,
		service:       computeService,
		project:       project,
		loggingClient: loggingClient,
		configService: runtimeConfigService.Projects.Configs,
		instanceName:  hostname,
		directory:     exPath,
		winVersion:    other,
	}, nil
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

// getWindowsVersionByVerCommand parses the output of the "ver" command to get the Windows version string.
// The output of ver looks like this:
//   Microsoft Windows [Version 6.1.7601]
// The return value will be "6.1.7601"
func (d *deployer) getWindowsVersionByVerCommand() (windowsVersion, error) {
	out, err := d.RunCommandWithOutput("cmd", "ver")
	if err != nil {
		return other, err
	}

	return getWindowsVersion(out)
}

// Deploy assets on the current CEL instance, either locally or on a NestedVM.
// manifestFilePath is the path to the CEL manifest file.
func (d *deployer) Deploy(manifestFile string) {
	log.Printf("Start on-host deployment")

	machineConfigVar := onhost.GetWindowsMachineRuntimeConfigVariableName(d.instanceName)
	status := d.getRuntimeConfigVariableValue(machineConfigVar)
	if status == statusError {
		d.Logf("Status of %s is %s. Nothing needs to be done.", machineConfigVar, status)
		return
	}

	if err := d.getCelConfiguration(manifestFile); err != nil {
		d.Logf("Error getting CEL configuration: %s", err)
		d.setRuntimeConfigVariable(machineConfigVar, statusError)
		return
	}

	d.configuration.AssetManifest = *d.configuration.Lab.AssetManifest
	d.configuration.HostEnvironment = *d.configuration.Lab.HostEnvironment
	d.configuration.Lab = lab.Lab{}

	if err := d.configuration.Validate(); err != nil {
		d.Logf("Error validating configuration : %s", err)
		d.setRuntimeConfigVariable(machineConfigVar, statusError)
		return
	}

	m := d.getWindowsMachine()
	if mt, ok := d.getMachineType(m.MachineType).Base.(*host.MachineType_NestedVm); ok {
		d.nestedVM = mt.NestedVm
	}

	if err := d.SaveSupportingFilesToDisk(); err != nil {
		if IsRestarting(d) {
			d.Logf("Ignoring failure during restart: %s", err)
			return
		}

		d.Logf("Saving supporting files to disk failed: %s", err)
		d.setRuntimeConfigVariable(machineConfigVar, statusError)
		return
	}

	if d.IsNestedVM() {
		if err := d.setupNestedVM(manifestFile); err != nil {
			d.Logf("Error setting up the nested VM: %s.", err)

			// if the status is already 'ready', then we shouldn't change it.
			// The reason is that if it is changed to error, then the nested VM cannot
			// work after host reboot since setupNestedVM() will never be called.
			if status != statusReady {
				d.setRuntimeConfigVariable(machineConfigVar, statusError)
			}
			return
		}
	}

	if status == statusReady {
		d.Logf("Status of %s is %s. We're done.", machineConfigVar, status)
		return
	}

	d.setRuntimeConfigVariable(machineConfigVar, statusInProgress)

	if d.IsNestedVM() {
		if err := d.commonSetupOnNestedVM(); err != nil {
			d.Logf("Common setup on nested VM failed: %s", err)
			d.setRuntimeConfigVariable(machineConfigVar, statusError)
			return
		}
	}

	err := d.PrepareInstance()
	if err != nil {
		if IsRestarting(d) {
			d.Logf("Ignoring failure during restart: %s", err)
			return
		}

		d.Logf("Preparing the CEL instance failed: %s", err)
		d.setRuntimeConfigVariable(machineConfigVar, statusError)
		return
	}

	d.createOnHostAssets()
}

// create on-host assets
func (d *deployer) createOnHostAssets() {
	machineConfigVar := onhost.GetWindowsMachineRuntimeConfigVariableName(d.instanceName)

	// Add additional dependency
	if err := common.ApplyResolvers(d, d.configuration.GetNamespace(), common.AdditionalDependencyResolverKind); err != nil {
		d.Logf("Error adding additonal dependency : %s", err)
		d.setRuntimeConfigVariable(machineConfigVar, statusError)
		return
	}
	if err := d.configuration.Validate(); err != nil {
		d.Logf("Error validating configuration after adding additional dependency : %s", err)
		d.setRuntimeConfigVariable(machineConfigVar, statusError)
		return
	}

	// Call on-host resolvers to create on-host assets
	d.Logf("Call resolvers")
	if err := common.ApplyResolvers(d, d.configuration.GetNamespace(),
		common.OnHostResolverKind); err != nil {
		if err == ErrRebootNeeded {
			// Reboot is needed. In this case, we don't change the config status.
			// After reboot, the configuration will be applied again, and this time since
			// the asset is already configured, RebootNeeded will not be returned.
			d.Logf("Reboot needed. Continue configuration after reboot.")
			if err := d.Reboot(); err != nil {
				d.setRuntimeConfigVariable(machineConfigVar, statusError)
				d.Logf("Failed to reboot. error: %s", err)
				return
			}
		} else {
			if IsRestarting(d) {
				d.Logf("Ignoring failure during restart: %s", err)
				return
			}

			d.setRuntimeConfigVariable(machineConfigVar, statusError)
			d.Logf("Setup Instance failed. error: %s", err)
			return
		}
	} else {
		d.setRuntimeConfigVariable(machineConfigVar, statusReady)
		d.Logf("Everything is OK.")
	}
}

// Setup nested VM, mainly the virtual network setup. This step is always needed.
func (d *deployer) setupNestedVM(manifestFile string) error {
	imageFile := path.Join(d.directory, path.Base(d.nestedVM.Image))

	// download the image file if it does not exist
	if _, err := os.Stat(imageFile); os.IsNotExist(err) {
		if err := d.RunLocalCommand("gsutil", "cp", d.nestedVM.Image, d.directory); err != nil {
			return err
		}
	}

	fileToRun := d.GetLocalSupportingFilePath("setup_vm_host.sh")
	if err := d.RunLocalCommand("bash", fileToRun); err != nil {
		return err
	}

	// start the VM
	d.RunLocalCommandWithoutWait("sudo", "kvm", "-m", "5120", "-net", "nic",
		"-net", "tap,ifname=tap0,script=no", "-usbdevice", "tablet",
		// the default CPU is qemu64, which cannot run Win10. So we need to
		// change it to "host"
		"-cpu", "host",
		// the monitor so that we can tell kvm to cleanly shutdown the VM.
		"-qmp", "tcp:127.0.0.1:25555,server,nowait",
		"-vnc", ":20100", imageFile)

	internalIP, err := d.waitForVMToStart()
	if err != nil {
		return err
	}

	d.internalIP = internalIP

	aliasIP, err := metadata.Get("instance/network-interfaces/0/ip-aliases/0")
	if err != nil {
		return errors.Errorf("Error getting ip alias: %s", err)
	}

	// aliasIP is CIDR such as 10.128.0.3/32 . We need to get the IP address.
	externalIP := strings.Split(aliasIP, "/")[0]
	d.externalIP = externalIP
	return d.RunLocalCommand("bash", d.GetLocalSupportingFilePath("setup_iptables.sh"), externalIP, internalIP)
}

// The output will look like this:
// Expiry Time          MAC address        Protocol  IP address                Hostname        Client ID or DUID
// -------------------------------------------------------------------------------------------------------------------
// 2018-07-27 19:26:29  52:54:00:12:34:56  ipv4      192.168.122.89/24         win7            01:52:54:00:12:34:56
// returns internalIp
func (d *deployer) waitForVMToStart() (string, error) {
	// wait until the VM gets its IP address
	for i := 0; i < 10; i++ {
		output, err := d.RunLocalCommandWithOutput("sudo", "virsh", "net-dhcp-leases", "default")
		if err != nil {
			return "", err
		}

		lines := strings.Split(output, "\n")
		if len(lines) == 5 {
			fields := strings.Fields(lines[2])

			// the format of field 4 is 192.168.122.89/24, so we need another split
			// to get the IP address.
			return strings.Split(fields[4], "/")[0], nil
		}

		time.Sleep(1 * time.Minute)
	}

	return "", errors.New("Time out")
}

func (d *deployer) getWindowsMachine() *asset.WindowsMachine {
	// find the instance
	for _, m := range d.configuration.AssetManifest.WindowsMachine {
		if m.Name == d.instanceName {
			return m
		}
	}
	return nil
}

func (d *deployer) getActiveDirectoryDomain() *asset.ActiveDirectoryDomain {
	m := d.getWindowsMachine()

	ad, err := d.configuration.AssetManifest.FindActiveDirectoryDomainFor(m)
	if err == nil {
		return ad
	}

	return nil
}

func (d *deployer) getMachineType(machineType string) *host.MachineType {
	for _, mt := range d.configuration.HostEnvironment.MachineType {
		if mt.Name == machineType {
			return mt
		}
	}
	return nil
}

func (d *deployer) SaveSupportingFilesToDisk() error {
	for filename, file := range _escData {
		if !file.isDir {
			outputFileName := path.Join(d.directory, filename)
			log.Printf("Save file %s", outputFileName)
			fileContent := FSMustByte(false, filename)
			os.Mkdir(path.Dir(outputFileName), os.ModePerm)
			if err := ioutil.WriteFile(outputFileName, fileContent, os.ModePerm); err != nil {
				return errors.Wrapf(err, "error saving file %s", outputFileName)
			}
		}
	}

	return nil
}

func (d *deployer) PrepareInstance() error {
	ver, err := d.getWindowsVersionByVerCommand()
	if err != nil {
		return errors.WithStack(err)
	}
	d.winVersion = ver

	if d.IsWindows2008() {
		return d.RunCommand("powershell.exe", "-ExecutionPolicy", "ByPass",
			"-File", d.GetSupportingFilePath("prepare_win2008.ps1"))
	} else if d.IsWindows2016() || d.IsWindows2012() {
		return d.RunCommand("powershell.exe", "-ExecutionPolicy", "ByPass",
			"-File", d.GetSupportingFilePath("prepare_win2012.ps1"))
	}

	return errors.New("unsupported windows version")
}

// Reboot the instance.
func (d *deployer) Reboot() error {
	d.Logf("Execute shutdown to reboot")

	err := d.RunCommand("shutdown", "/r", "/t", "0")
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus, ok := exitError.Sys().(syscall.WaitStatus)
			if !ok {
				return err
			}

			// Some exit codes should be treated as success:
			// ERROR_SHUTDOWN_IN_PROGRESS (1115): "A system shutdown is in progress."
			// ERROR_SHUTDOWN_IS_SCHEDULED (1190): "A system shutdown has already been scheduled."
			// RPC_S_UNKNOWN_IF (1717): "The interface is unknown."
			codesToIgnore := []int{1115, 1190, 1717}
			for _, code := range codesToIgnore {
				if waitStatus.ExitStatus() == code {
					d.Logf("Ignoring shutdown error with exit code %v : %v", code, err)
					return nil
				}
			}
		}
		return err
	}

	return nil
}

func (d *deployer) WaitForNestedVMRebootComplete() error {
	// Wait a while to give shutdown enough time to finish.
	time.Sleep(1 * time.Minute)

	return d.waitUntilSshIsAlive()
}

func (d *deployer) Logf(format string, arg ...interface{}) {
	text := fmt.Sprintf(format, arg...)
	log.Output(3, text)
	d.loggingClient.Logger("cel").Log(
		logging.Entry{
			Payload: strings.Map(removeInvalidUTF8, text),
		},
	)
}

// This is used as a strings mapping function: https://golang.org/pkg/strings/#Map
// Negative values are dropped from the string with no replacement.
func removeInvalidUTF8(char rune) rune {
	if char == utf8.RuneError {
		return -1
	} else {
		return char
	}
}

func (d *deployer) getCelConfiguration(manifestFilePath string) error {
	content, err := ioutil.ReadFile(manifestFilePath)
	if err != nil {
		return err
	}

	d.configuration = &cel.Configuration{}
	err = proto.UnmarshalText(string(content), &d.configuration.Lab)
	if err != nil {
		return errors.Wrapf(err, "error when parsing the configuration: %s", string(content))
	}

	return nil
}

// returns "" if the variable does not exist or error
// Note that errors are ignored
func (d *deployer) getRuntimeConfigVariableValue(variable string) string {
	v, err := d.configService.Variables.Get(
		d.getFullRuntimeConfigVariableName(variable)).Context(d.ctx).Do()

	if err == nil {
		return v.Text
	} else {
		return ""
	}
}

// returns the full config variable name required by the API
func (d *deployer) getFullRuntimeConfigVariableName(variable string) string {
	return fmt.Sprintf("projects/%s/configs/%s/variables/%s",
		d.project.Name,
		// hard coded config name
		"cel-config",
		variable)
}

func (d *deployer) getRuntimeConfigVariableParentName() string {
	return fmt.Sprintf("projects/%s/configs/%s",
		d.project.Name,
		// hard coded config name
		"cel-config")
}

// sets the value of a runtime config variable.
// Note that errors are ignored
func (d *deployer) setRuntimeConfigVariable(variable string, value string) {
	for i := 0; i < 3; i++ {
		_, err := d.configService.Variables.Update(
			d.getFullRuntimeConfigVariableName(variable),
			&runtimeconfig.Variable{Text: value}).Context(d.ctx).Do()
		if err == nil {
			d.Logf("config variable %s is set to %s", variable, value)
			return
		}

		apiError, ok := err.(*googleapi.Error)
		if ok && apiError.Code == 404 {
			// the variable does not exist. So we create it instead
			v := &runtimeconfig.Variable{
				Name: d.getFullRuntimeConfigVariableName(variable),
				Text: value}
			parent := d.getRuntimeConfigVariableParentName()
			d.configService.Variables.Create(parent, v).Context(d.ctx).Do()
			return
		}

		if ok && apiError.Code == 503 {
			// service unavailable. In this case, retry
			d.Logf("Error updating config variable %s: %s. Retry", variable, err)
			time.Sleep(5 * time.Second)
		} else if strings.Contains(err.Error(), "no such host") {
			// Ideally, `errNoSuchHost` would be exposed, but it's not:
			// https://github.com/golang/go/issues/28635
			d.Logf("Error updating config variable %s: %s. Retry", variable, err)
			time.Sleep(1 * time.Minute)
		} else {
			// Log the error and return
			d.Logf("Error updating config variable %s: %s.", variable, err)
			return
		}

	}
}

// Returns true if we are running on Windows Server 2016 or Windows 10
func (d *deployer) IsWindows2016() bool {
	return d.winVersion == win2016
}

// Returns true if we are running on Windows Server 2012 R2 or Windows 8
func (d *deployer) IsWindows2012() bool {
	return d.winVersion == win2012
}

// Returns true if we are running on Windows Server 2008 R2 or Windows 7
func (d *deployer) IsWindows2008() bool {
	return d.winVersion == win2008
}

// Returns the path to a supporting file on the current CEL instance.
// It can be a local path or a path inside a CEL NestedVM.
// In both cases, the path is safe to pass to RunCommand or RunConfigCommand,
// so this is the method that should be used most of the time.
func (d *deployer) GetSupportingFilePath(filename string) string {
	if d.IsNestedVM() {
		return filepath.Join(workingDirectoryOnNestedVM, filename)
	} else {
		return d.GetLocalSupportingFilePath(filename)
	}
}

// Returns the path to a local supporting file, on the current Compute instance
// running this code. The current instance can be a Linux host or a NestedVM.
// This path is only safe to pass to RunLocalCommand and should only be used
// for NestedVM-specific code that needs to refer to the Linux host.
func (d *deployer) GetLocalSupportingFilePath(filename string) string {
	return path.Join(d.directory, "supporting_files", filename)
}

func (d *deployer) IsNestedVM() bool {
	return d.nestedVM != nil
}

// Runs a command on the current CEL instance, either locally or a Nested VM.
func (d *deployer) RunCommandWithOutput(name string, arg ...string) (string, error) {
	if d.IsNestedVM() {
		return d.sshRunCommand(name, arg...)
	} else {
		return d.RunLocalCommandWithOutput(name, arg...)
	}
}

// Runs a command on the current CEL instance, either locally or a Nested VM.
func (d *deployer) RunCommand(name string, arg ...string) error {
	_, err := d.RunCommandWithOutput(name, arg...)
	return err
}

func (d *deployer) RunLocalCommandWithOutput(name string, arg ...string) (string, error) {
	d.Logf("Run command: %s, args: %s", name, arg)
	output, err := exec.Command(name, arg...).CombinedOutput()
	if output != nil {
		d.Logf("Output of command %s, args %s is: %s", name, arg, output)
	}

	return string(output), err
}

func (d *deployer) RunLocalCommand(name string, arg ...string) error {
	_, err := d.RunLocalCommandWithOutput(name, arg...)
	return err
}

func (d *deployer) RunLocalCommandWithoutWait(name string, arg ...string) error {
	d.Logf("Run command: %s, args: %s", name, arg)
	err := exec.Command(name, arg...).Start()
	return errors.Wrap(err, "run command")
}

// Requirement for ConfigCommand:
// exit code 100 indicating that the failure is fatal
// exit code 150 indicating that the failure is transient/retryable
// exit code 200 indicating that reboot is needed
func (d *deployer) RunConfigCommand(name string, arg ...string) error {
	if err := d.RunCommand(name, arg...); err != nil {
		var exitCode int

		if exitError, ok := err.(*exec.ExitError); ok {
			if waitStatus, ok := exitError.Sys().(syscall.WaitStatus); ok {
				exitCode = waitStatus.ExitStatus()
			}
		} else if exitError, ok := err.(*ssh.ExitError); ok {
			exitCode = exitError.Waitmsg.ExitStatus()
		}

		if exitCode == 150 {
			// Exit code 150 means "failure is retryable."
			return ErrTransient
		} else if exitCode == 200 {
			// Exit code 200 means "reboot is needed."
			return ErrRebootNeeded
		}

		return err
	}

	return nil
}

// getAdDomainAsset returns the ActiveDirectoryDomain asset of the given domain.
func (d *deployer) getAdDomainAsset(domainName string) (*asset.ActiveDirectoryDomain, error) {
	return d.configuration.AssetManifest.FindActiveDirectoryDomain(domainName)
}

// waitForDependency waits for the dependency to be ready.
// depVar is the runtime configuration variable of the dependency.
func (d *deployer) waitForDependency(depVar string, timeOut time.Duration) error {
	t := time.Now()
	sleepDuration := 60 * time.Second
	for {
		if time.Now().Sub(t) > timeOut {
			d.Logf("Time out reached waiting for %s", depVar)
			return errors.Errorf("time out waiting for dependency")
		}

		status := d.getRuntimeConfigVariableValue(depVar)
		d.Logf("Status of %s is [%s]", depVar, status)

		if status == statusReady {
			break
		} else if status == statusError {
			d.Logf("Cannot continue because the dependency status is error")
			return errors.Errorf("cannot continue because the dependency status is error")
		} else {
			d.Logf("Sleep for %s", sleepDuration.String())
			time.Sleep(sleepDuration)
		}
	}

	return nil
}

// Builds a command string that is safe to pass inside a SSH session.
func sshGetCommandString(name string, arg ...string) string {
	cmd := name

	for _, argument := range arg {
		cmd += " " + sshEscapeArgument(argument)
	}

	return cmd
}

// Escapes parameters that contains special characters.
func sshEscapeArgument(argument string) string {
	if strings.ContainsAny(argument, " |&<>^\"") {
		argument = strings.Replace(argument, "\"", "\\\"", -1)
		return fmt.Sprintf("\"%s\"", argument)
	}

	return argument
}

func (d *deployer) sshRunCommand(name string, arg ...string) (string, error) {
	err := d.ensureWorkingDirOnNestedVmExists()
	if err != nil {
		return "", err
	}

	cmd := sshGetCommandString(name, arg...)

	d.Logf("Run command on nested VM: %s", cmd)
	conn, err := d.sshConnect()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		return "", err
	}

	// Set WorkingDirectory to workingDirectoryOnNestedVM
	cmd = fmt.Sprintf("cd %s && %s", workingDirectoryOnNestedVM, cmd)
	output, err := session.CombinedOutput(cmd)
	d.Logf("output from nested VM '%s': [%s], err: %v", cmd, output, err)

	return string(output), err
}

func (d *deployer) ensureWorkingDirOnNestedVmExists() error {
	cmd := fmt.Sprintf("if not exist %s mkdir %s", workingDirectoryOnNestedVM, workingDirectoryOnNestedVM)

	d.Logf("Run command on nested VM: %s", cmd)
	conn, err := d.sshConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		return err
	}

	output, err := session.CombinedOutput(cmd)
	d.Logf("output from ensureWorkingDirOnNestedVmExists: [%s], err: %v", output, err)

	return err
}

// sshConnect connects to the instance and returns the ssh client.
func (d *deployer) sshConnect() (*ssh.Client, error) {
	sshConfig := &ssh.ClientConfig{
		User: d.nestedVM.UserName,
		Auth: []ssh.AuthMethod{
			ssh.Password(d.nestedVM.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}

	return ssh.Dial("tcp", d.internalIP+":22", sshConfig)
}

// waitUntilSshIsAlive waits until ssh connection can be made.
func (d *deployer) waitUntilSshIsAlive() error {
	d.Logf("Try ssh connecting to nested VM")
	timeOut := 20 * time.Minute
	for start := time.Now(); time.Since(start) < timeOut; {
		conn, err := d.sshConnect()
		if err == nil {
			d.Logf("Connected")
			conn.Close()
			return nil
		}

		d.Logf("Retry")
		time.Sleep(1 * time.Minute)
	}

	return errors.New("ssh timeout")
}

// Disables Windows Update and makes sure wuauserv doesn't run.
func (d *deployer) disableWindowsUpdate() error {
	// Disable auto-start.
	err := d.RunCommand("sc", "config", "wuauserv", "start=", "disabled")
	if err != nil {
		return err
	}

	// Stop the service if it's running.
	output, err := d.RunCommandWithOutput("net", "stop", "wuauserv")
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
		if err = d.killServiceProcess("wuauserv"); err != nil {
			return err
		}

		return d.killServiceProcess("TrustedInstaller")
	}

	return nil
}

// Kills the service process on a CEL instance.
func (d *deployer) killServiceProcess(service string) error {
	// Get service information.
	output, err := d.RunCommandWithOutput("sc", "queryex", service)
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

			// Kill the process.
			return d.RunCommand("taskkill", "/f", "/pid", parts[1])
		}
	}

	return nil
}

func (d *deployer) commonSetupOnNestedVM() error {
	err := d.waitUntilSshIsAlive()
	if err != nil {
		return errors.WithStack(err)
	}

	err = d.uploadSupportingFilesToNestedVM()
	if err != nil {
		return errors.WithStack(err)
	}

	err = d.disableWindowsUpdate()
	if err != nil {
		return errors.WithStack(err)
	}

	output, err := d.RunCommandWithOutput("hostname")
	if err != nil {
		return errors.WithStack(err)
	}
	hostname := strings.TrimRight(string(output), "\r\n")

	// Rename nestedVM if needed. Windows hostnames are not case-sentive.
	if !strings.EqualFold(hostname, d.instanceName) {
		d.Logf("Renaming nested VM from %v to %v.", hostname, d.instanceName)
		return d.renameNestedVM(d.instanceName)
	} else {
		d.Logf("Skip renaming nested VM because it's already named %v.", d.instanceName)
	}

	return nil
}

func (d *deployer) renameNestedVM(newName string) error {
	_, err := d.sshRunCommand(
		fmt.Sprintf("powershell.exe -Command \"Rename-Computer -NewName %s -Force -PassThru -Restart\"",
			newName))
	if err != nil {
		return err
	}

	return d.WaitForNestedVMRebootComplete()
}

func (d *deployer) UploadFileToNestedVM(srcFilePath string, destDirectory string) error {
	fileName := filepath.Base(srcFilePath)
	destFilePath := filepath.Join(destDirectory, fileName)

	d.Logf("Upload file %s -> %s to nested VM", srcFilePath, destFilePath)
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		return errors.WithStack(err)
	}

	conn, err := d.sshConnect()
	if err != nil {
		return errors.WithStack(err)
	}
	defer conn.Close()

	client, err := sftp.NewClient(conn)
	if err != nil {
		return errors.WithStack(err)
	}

	destFile, err := client.Create(destFilePath)
	if err != nil {
		return errors.WithStack(err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (d *deployer) uploadSupportingFilesToNestedVM() error {
	err := d.ensureWorkingDirOnNestedVmExists()
	if err != nil {
		return errors.WithStack(err)
	}

	// save supporting files on disk
	for filename, file := range _escData {
		if !file.isDir {
			supportingFile := path.Join(d.directory, filename)
			err := d.UploadFileToNestedVM(supportingFile, workingDirectoryOnNestedVM)
			if err != nil {
				return errors.WithStack(err)
			}
		}
	}

	return nil
}

// getInstanceAddress gets the IP address of the given instance and retries on transient errors.
func (d *deployer) getInstanceAddress(instanceName string) (string, error) {
	var dnsServerAddress string

	const maxRetries = 5
	retries := 0
	for {
		addrs, err := net.LookupHost(instanceName)

		if err == nil {
			dnsServerAddress = addrs[0]
			break
		}

		// `LookupHost` can fail when the instance is restarting.
		// Ideally, `errNoSuchHost` would be exposed so we don't have to compare strings, but it's not:
		// https://github.com/golang/go/issues/28635
		if strings.Contains(err.Error(), "no such host") && retries <= maxRetries {
			retries++
			d.Logf("net.LookupHost returned a transient error. Will wait a minute and try again.")
			time.Sleep(1 * time.Minute)
			continue
		}

		return "", errors.WithStack(err)
	}

	return dnsServerAddress, nil
}
