// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"chromium.googlesource.com/enterprise/cel/go/asset"
	"chromium.googlesource.com/enterprise/cel/go/cel"
	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/gcp/onhost"
	"chromium.googlesource.com/enterprise/cel/go/lab"
	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/logging"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
	"google.golang.org/api/googleapi"
	runtimeconfig "google.golang.org/api/runtimeconfig/v1beta1"
)

const statusError = "error"
const statusReady = "ready"
const statusInProgress = "in-progress"

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
	return nil
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
	return nil, nil
}

func (d *deployer) Close() error {
	return d.loggingClient.Close()
}

func CreateDeployer() (*deployer, error) {
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

	verString, err := getWindowsVersion()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ver := other
	if strings.HasPrefix(verString, "10.0.") {
		ver = win2016
	} else if strings.HasPrefix(verString, "6.3.") {
		ver = win2012
	} else if strings.HasPrefix(verString, "6.1.") {
		ver = win2008
	}

	return &deployer{
		ctx:           ctx,
		service:       computeService,
		project:       project,
		loggingClient: loggingClient,
		configService: runtimeConfigService.Projects.Configs,
		instanceName:  hostname,
		directory:     exPath,
		winVersion:    ver,
	}, nil
}

// getWindowsVersion parses the output of the "ver" command to get the Windows version string.
// The output of ver looks like this:
//   Microsoft Windows [Version 6.1.7601]
// The return value will be "6.1.7601"
func getWindowsVersion() (string, error) {
	cmd := exec.Command("cmd", "ver")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	// remove newlines from the output
	outStr := strings.Replace(out.String(), "\r\n", "", -1)
	tmp1 := strings.Index(outStr, "[Version ")
	tmp2 := strings.Index(outStr, "]")
	if tmp1 == -1 || tmp2 == -1 {
		return "", fmt.Errorf("The output of 'ver' cannot be parsed: %s", outStr)
	}

	ver := outStr[tmp1+9 : tmp2]
	return ver, nil
}

// Deploy assets on the current instance. manifestFilePath is the path
// to the CEL manifest file.
func (d *deployer) Deploy(manifestFile string) {
	log.Printf("Start on-host deployment")

	machineConfigVar := onhost.GetWindowsMachineRuntimeConfigVariableName(d.instanceName)
	status := d.getRuntimeConfigVariableValue(machineConfigVar)
	if status == statusReady || status == statusError {
		d.Logf("Status of %s is %s. Nothing needs to be done.", machineConfigVar, status)
		return
	}

	d.setRuntimeConfigVariable(machineConfigVar, statusInProgress)

	// common setup
	if err := d.CommonSetup(); err != nil {
		d.Logf("Common setup failed: %s", err)
		d.setRuntimeConfigVariable(machineConfigVar, statusError)
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
			}
		} else {
			d.setRuntimeConfigVariable(machineConfigVar, statusError)
			d.Logf("Setup Instance failed. error: %s", err)
		}
	} else {
		d.setRuntimeConfigVariable(machineConfigVar, statusReady)
		d.Logf("Everything is OK.")
	}
}

func (d *deployer) CommonSetup() error {
	// save supporting files on disk
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

	if d.IsWindows2008() {
		return d.RunCommand("powershell.exe", "-File",
			d.GetSupportingFilePath("prepare_win2008.ps1"))
	} else if d.IsWindows2016() || d.IsWindows2012() {
		return d.RunCommand("powershell.exe", "-File",
			d.GetSupportingFilePath("prepare_win2012.ps1"))
	} else {
		return errors.New("unsupported windows version")
	}
}

// Reboot the instance.
func (d *deployer) Reboot() error {
	d.Logf("Execute shutdown to reboot")

	// Exit code 1190 means "A system shutdown has already been scheduled."
	// This case should be treated as success
	if err := d.RunCommand("shutdown", "/r", "/t", "0"); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStaus, ok := exitError.Sys().(syscall.WaitStatus)
			if ok && waitStaus.ExitStatus() == 1190 {
				return nil
			}
		}
		return err
	}

	return nil
}

func (d *deployer) Logf(format string, arg ...interface{}) {
	text := fmt.Sprintf(format, arg...)
	log.Print(text)
	d.loggingClient.Logger("cel").Log(
		logging.Entry{
			Payload: text,
		},
	)
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
	_, err := d.configService.Variables.Update(
		d.getFullRuntimeConfigVariableName(variable),
		&runtimeconfig.Variable{Text: value}).Context(d.ctx).Do()
	if err == nil {
		return
	}

	if err != nil {
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

		// Log the error
		d.Logf("Error updating config variable %s: %s", variable, err)
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

func (d *deployer) GetSupportingFilePath(filename string) string {
	return path.Join(d.directory, "supporting_files", filename)
}

func (d *deployer) RunCommand(name string, arg ...string) error {
	d.Logf("Run command: %s, args: %s", name, arg)
	output, err := exec.Command(name, arg...).CombinedOutput()
	if output != nil {
		d.Logf("Output of command %s, args %s is: %s", name, arg, output)
	}

	return err
}

// Requirement for ConfigCommand:
// exit code 200 indicating that reboot is needed
func (d *deployer) RunConfigCommand(name string, arg ...string) error {
	if err := d.RunCommand(name, arg...); err != nil {
		exitError, ok := err.(*exec.ExitError)
		if ok {
			waitStaus, ok := exitError.Sys().(syscall.WaitStatus)
			if ok && waitStaus.ExitStatus() == 200 {
				// Exit code 200 means "reboot is needed."
				return ErrRebootNeeded
			}
		}

		return err
	}

	return nil
}

// getAdDomainAsset returns the ActiveDirectoryDomain asset of the given domain.
func (d *deployer) getAdDomainAsset(domainName string) (*asset.ActiveDirectoryDomain, error) {
	for _, ad := range d.configuration.AssetManifest.AdDomain {
		if ad.Name == domainName {
			return ad, nil
		}
	}

	return nil, errors.Errorf("cannot find asset for domain: %s", domainName)
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
