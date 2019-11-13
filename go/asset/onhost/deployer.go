// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

import (
	"context"
	"fmt"
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

	"chromium.googlesource.com/enterprise/cel/go/asset"
	"chromium.googlesource.com/enterprise/cel/go/cel"
	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/gcp/onhost"
	assetpb "chromium.googlesource.com/enterprise/cel/go/schema/asset"
	hostpb "chromium.googlesource.com/enterprise/cel/go/schema/host"
	labpb "chromium.googlesource.com/enterprise/cel/go/schema/lab"
	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/logging"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
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

	configuration *cel.Configuration

	instance instanceInterface
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
	}, nil
}

func createInstance(d *deployer, machineType *hostpb.MachineType, machineConfigVar string) instanceInterface {
	nestedVM := machineType.GetNestedVm()
	os := machineType.GetOs()
	switch os {
	case hostpb.OperatingSystem_WINDOWS:
		if nestedVM == nil {
			return &windowsInstance{
				instanceBase{
					d:                d,
					operatingSystem:  os,
					machineConfigVar: machineConfigVar,
				},
				windowsInstanceBase{
					winVersion: other,
				},
			}
		}

		return &nestedVMWindowsInstance{
			instanceBase{
				d:                d,
				operatingSystem:  os,
				machineConfigVar: machineConfigVar,
			},
			nestedVMInstanceBase{
				nestedVM: nestedVM,
			},
			windowsInstanceBase{
				winVersion: other,
			},
		}
	case hostpb.OperatingSystem_CHROMEOS:
		if nestedVM == nil {
			d.Logf("ChromeOS can only run as nested VM")
			d.setRuntimeConfigVariable(machineConfigVar, statusError)
			return nil
		}

		return &nestedVMCrosInstance{
			instanceBase{
				d:                d,
				operatingSystem:  os,
				machineConfigVar: machineConfigVar,
			},
			nestedVMInstanceBase{
				nestedVM: nestedVM,
			},
		}
	default:
		d.Logf("Unsupported Operating System: %v", machineType.GetOs())
		d.setRuntimeConfigVariable(machineConfigVar, statusError)
		return nil
	}
}

// Deploy assets on the current CEL instance, either locally or on a NestedVM.
// manifestFilePath is the path to the CEL manifest file.
func (d *deployer) Deploy(manifestFile string) {
	log.Printf("Start on-host deployment")

	machineConfigVar := onhost.GetWindowsMachineRuntimeConfigVariableName(d.instanceName)
	status := d.getRuntimeConfigVariableValue(machineConfigVar)
	d.Logf("Status of %s is %s.", machineConfigVar, status)
	if status == statusError {
		d.Logf("Status is %s. Nothing needs to be done.", status)
		return
	}

	if err := d.getCelConfiguration(manifestFile); err != nil {
		d.Logf("Error getting CEL configuration: %s", err)
		d.setRuntimeConfigVariable(machineConfigVar, statusError)
		return
	}

	d.configuration.AssetManifest = *d.configuration.Lab.AssetManifest
	d.configuration.HostEnvironment = *d.configuration.Lab.HostEnvironment
	d.configuration.Lab = labpb.Lab{}

	if err := d.configuration.Validate(); err != nil {
		d.Logf("Error validating configuration : %s", err)
		d.setRuntimeConfigVariable(machineConfigVar, statusError)
		return
	}

	m := d.getWindowsMachine()
	machineType := d.getMachineType(m.MachineType)
	d.instance = createInstance(d, machineType, machineConfigVar)
	if d == nil {
		return
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

	if d.instance.OnBoot() {
		return
	}

	if status == statusReady {
		d.Logf("Status of %s is %s. We're done.", machineConfigVar, status)
		return
	}

	d.setRuntimeConfigVariable(machineConfigVar, statusInProgress)

	if d.instance.OneTimeSetup() {
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

func (d *deployer) getWindowsMachine() *assetpb.WindowsMachine {
	// find the instance
	for _, m := range d.configuration.AssetManifest.WindowsMachine {
		if m.Name == d.instanceName {
			return m
		}
	}
	return nil
}

func (d *deployer) getActiveDirectoryDomain() *assetpb.ActiveDirectoryDomain {
	m := d.getWindowsMachine()

	ad, err := asset.FindActiveDirectoryDomainFor(&d.configuration.AssetManifest, m)
	if err == nil {
		return ad
	}

	return nil
}

func (d *deployer) getMachineType(machineType string) *hostpb.MachineType {
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

// Reboot the instance.
func (d *deployer) Reboot() error {
	d.Logf("Execute shutdown to reboot")

	_, err := d.instance.RunCommand("shutdown", "/r", "/t", "0")
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
			d.Logf("config variable %s doesn't exist, so we'll create it.", variable)
			v := &runtimeconfig.Variable{
				Name: d.getFullRuntimeConfigVariableName(variable),
				Text: value}
			parent := d.getRuntimeConfigVariableParentName()
			d.configService.Variables.Create(parent, v).Context(d.ctx).Do()
			return
		}

		if ok && (apiError.Code == 503 || apiError.Code == 500) {
			// service unavailable / backend error. In both cases, retry.
			d.Logf("Error updating config variable %s: %s. Retry", variable, err)
			time.Sleep(15 * time.Second)
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
	return d.instance.GetWindowsVersion() == win2016
}

// Returns true if we are running on Windows Server 2012 R2 or Windows 8
func (d *deployer) IsWindows2012() bool {
	return d.instance.GetWindowsVersion() == win2012
}

// Returns true if we are running on Windows Server 2008 R2 or Windows 7
func (d *deployer) IsWindows2008() bool {
	return d.instance.GetWindowsVersion() == win2008
}

// getAdDomainAsset returns the ActiveDirectoryDomain asset of the given domain.
func (d *deployer) getAdDomainAsset(domainName string) (*assetpb.ActiveDirectoryDomain, error) {
	return asset.FindActiveDirectoryDomain(&d.configuration.AssetManifest, domainName)
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

// Requirement for ConfigCommand:
// exit code 100 indicating that the failure is fatal
// exit code 150 indicating that the failure is transient/retryable
// exit code 200 indicating that reboot is needed
func (d *deployer) RunConfigCommand(name string, arg ...string) error {
	if _, err := d.instance.RunCommand(name, arg...); err != nil {
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

func (d *deployer) GetSupportingFilePath(filename string) string {
	return d.instance.GetSupportingFilePath(filename)
}

func (d *deployer) GetOs() hostpb.OperatingSystem {
	return d.instance.GetOs()
}

func (d *deployer) RunCommand(name string, arg ...string) error {
	_, err := d.instance.RunCommand(name, arg...)
	return err
}

func (d *deployer) IsNestedVM() bool {
	_, ok := d.instance.(nestedVMInstanceInterface)
	return ok
}
