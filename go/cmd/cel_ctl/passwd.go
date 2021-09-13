// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"strings"

	"chromium.googlesource.com/enterprise/cel/go/gcp"
	"github.com/spf13/cobra"
)

type PasswdCommand struct {
	UseBuiltins bool
	Instance    string
	Username    string
	Email       string
}

func (p *PasswdCommand) Run(ctx context.Context, c *Application, cmd *cobra.Command, args []string) error {
	if p.Instance == "" || p.Username == "" {
		cmd.Usage()
		return fmt.Errorf("instance and username are required options")
	}

	session, err := c.CreateSession(ctx, args, p.UseBuiltins)
	if err != nil {
		return err
	}

	cs, err := session.GetBackend().GetComputeService()
	if err != nil {
		return err
	}
	state := gcp.CloudState{HostEnvironment: session.GetConfiguration().HostEnvironment}
	err = state.FetchInstances(ctx, cs)
	if err != nil {
		return err
	}

	instance := state.Instances[p.Instance]
	if instance == nil {
		return fmt.Errorf("instance not found: %v", p.Instance)
	}

	// Instance.Zone is the URL of the zone where the instance resides; ex:
	// https://www.googleapis.com/compute/v1/projects/proj-id/zones/us-east1-b
	zoneUrl := instance.Zone
	zone := zoneUrl[strings.LastIndex(zoneUrl, "/")+1:]

	password, err := gcp.ResetWindowsPassword(ctx, c.Client,
		session.GetConfiguration().HostEnvironment.Project.Name,
		zone, instance.Name, p.Username, p.Email)
	if err != nil {
		return fmt.Errorf("failed to reset password: %v", err)
	}

	fmt.Fprintln(cmd.OutOrStdout(), password)
	return nil
}

func init() {
	c := &cobra.Command{
		Use:   "passwd",
		Short: "reset password on a Windows instance",
		Long: `Resets the password for a local user on a Windows instance.
The environment must exist and match the one described in the asset file.
`,
		TraverseChildren: true,
	}

	f := c.Flags()

	p := &PasswdCommand{}
	f.BoolVarP(&p.UseBuiltins, "builtins", "B", false, "Use builtin assets")
	f.StringVar(&p.Instance, "instance", "", "short instance name of VM")
	f.StringVar(&p.Username, "username", "", "username of account to reset password")
	f.StringVar(&p.Email, "email", "", "email address to associate with account")

	app.AddCommand(c, p)
}
