// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"chromium.googlesource.com/enterprise/cel/go/cel/deploy"
	"context"
	"github.com/spf13/cobra"
	"log"
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
	log.Printf("Start of `cel_ctl deploy` - version %s", version)

	session, err := a.CreateSession(ctx, args, d.UseBuiltins)
	if err != nil {
		return err
	}

	return deploy.Deploy(session)
}
