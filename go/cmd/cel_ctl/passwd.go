// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"fmt"

	"chromium.googlesource.com/enterprise/cel/go/gcp"
	"github.com/spf13/cobra"
)

func init() {
	(&PasswdCommand{}).Register(root)
}

type PasswdCommand struct {
	Instance string
	Username string
	Email    string
}

func (p *PasswdCommand) Register(root *cobra.Command) {
	c := &cobra.Command{
		Use:   "passwd",
		Short: "reset password on a Windows instance",
		Long: `Resets the password for a local user on a Windows instance.
`,
		TraverseChildren: true,
		Run: func(c *cobra.Command, args []string) {
		}}

	f := c.LocalFlags()
	f.StringVar(&p.Instance, "instance", "", "short instance name of VM")
	f.StringVar(&p.Username, "username", "", "username of account to reset password")
	f.StringVar(&p.Email, "email", "", "email address to associate with account")

	c.RunE = func(c *cobra.Command, args []string) error {
		return app.InvokeCommand(p, c, args)
	}

	root.AddCommand(c)
}

func (p *PasswdCommand) Execute(ctx context.Context, c *App, cmd *cobra.Command, args []string) error {
	if p.Instance == "" || p.Username == "" {
		cmd.Usage()
		return fmt.Errorf("instance and username are required options")
	}

	err := c.Load(ctx, args)
	if err != nil {
		return err
	}

	session, err := c.GetSession(ctx)
	if err != nil {
		return err
	}

	instance := session.Cloud.Instances[p.Instance]
	if instance == nil {
		return fmt.Errorf("instance not found: ", p.Instance)
	}

	password, err := gcp.ResetWindowsPassword(ctx, c.Client, c.Configuration.HostEnvironment.Project.Name, instance.Zone, instance.Name, p.Username, p.Email)
	if err != nil {
		return fmt.Errorf("failed to reset password: ", err)
	}

	fmt.Fprintln(cmd.OutOrStdout(), password)
	return nil
}
