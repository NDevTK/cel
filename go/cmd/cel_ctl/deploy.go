// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"chromium.googlesource.com/enterprise/cel/go/cel"
	"context"
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
)

type DeployCommand struct {
	UseBuiltins bool
}

func init() {
	dc := &DeployCommand{}
	cmd := &cobra.Command{
		Use:   "deploy [configuration files]",
		Short: "deploy build artifacts to target lab environment",
		Long: `Deploys build artifacts to target lab environment.
`,
	}
	cmd.Flags().BoolVarP(&dc.UseBuiltins, "builtins", "B", false, "Use builtin assets")
	app.AddCommand(cmd, dc)
}

func (d *DeployCommand) Run(ctx context.Context, a *Application, cmd *cobra.Command, args []string) error {
	session, err := a.CreateSession(ctx, args, d.UseBuiltins)
	if err != nil {
		return err
	}

	err = cel.Deploy(session)
	if err != nil {
		if bytes, err := yaml.Marshal(session.GetConfiguration().GetNamespace()); err == nil {
			cmd.OutOrStderr().Write(bytes)
		}
		fmt.Fprintln(cmd.OutOrStderr(), "\n----")
	}
	return err
}
