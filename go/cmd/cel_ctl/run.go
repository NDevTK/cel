// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"strings"

	"chromium.googlesource.com/enterprise/cel/go/cel/deploy"
	"chromium.googlesource.com/enterprise/cel/go/gcp"
	"github.com/spf13/cobra"
)

type RunCommand struct {
	UseBuiltins bool
	Instance    string

	// The command to execute.
	// Metadata entries are limited to 256KB and have a total limit (512KB).
	Command string
}

func (p *RunCommand) Run(ctx context.Context, c *Application, cmd *cobra.Command, args []string) error {
	if p.Instance == "" || p.Command == "" {
		cmd.Usage()
		return fmt.Errorf("instance and command are required options")
	}

	session, err := c.CreateSession(ctx, args, p.UseBuiltins)
	if err != nil {
		return err
	}

	return RunCommandOnInstance(ctx, c, session, p.Instance, p.Command)
}

func RunCommandOnInstance(ctx context.Context, c *Application, session *deploy.Session, instanceName, command string) error {
	cs, err := session.GetBackend().GetComputeService()
	if err != nil {
		return err
	}
	state := gcp.CloudState{HostEnvironment: &session.GetConfiguration().HostEnvironment}
	err = state.FetchInstances(ctx, cs)
	if err != nil {
		return err
	}

	instance := state.Instances[instanceName]
	if instance == nil {
		return fmt.Errorf("instance not found: %v", instanceName)
	}

	// Instance.Zone is the URL of the zone where the instance resides; ex:
	// https://www.googleapis.com/compute/v1/projects/proj-id/zones/us-east1-b
	zoneUrl := instance.Zone
	zone := zoneUrl[strings.LastIndex(zoneUrl, "/")+1:]

	runCommand := gcp.NewRunCommand(command)
	exitCode, err := gcp.RunCommandOnInstance(ctx, c.Client,
		session.GetConfiguration().HostEnvironment.Project.Name,
		zone, instance.Name, runCommand)
	if err != nil {
		return fmt.Errorf("failed to execute command: %v", err)
	}

	if exitCode != 0 {
		return fmt.Errorf("command executed but returned: %v", exitCode)
	}

	return nil
}

func init() {
	c := &cobra.Command{
		Use:   "run",
		Short: "run a command on a Windows instance",
		Long: `Runs a command on a Windows instance via cel_agent.
The environment must exist and match the one described in the asset file.
`,
		TraverseChildren: true,
	}

	f := c.Flags()

	p := &RunCommand{}
	f.BoolVarP(&p.UseBuiltins, "builtins", "B", false, "Use builtin assets")
	f.StringVar(&p.Instance, "instance", "", "short instance name of VM")
	f.StringVar(&p.Command, "command", "", "command to execute")

	app.AddCommand(c, p)
}
