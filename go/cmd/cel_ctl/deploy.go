// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"chromium.googlesource.com/enterprise/cel/go/cel"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
)

type DeployCommand struct{}

func init() {
	app.AddCommand(&cobra.Command{
		Use:   "deploy [configuration files]",
		Short: "deploy build artifacts to target lab environment",
		Long: `Deploys build artifacts to target lab environment.
`,
	}, &DeployCommand{})
}

func (d *DeployCommand) Run(ctx context.Context, a *Application, cmd *cobra.Command, args []string) error {
	session, err := a.CreateSession(ctx, args)
	if err != nil {
		return err
	}

	err = cel.Start(session)
	if err != nil {
		if bytes, err := json.MarshalIndent(session.GetConfiguration().GetNamespace(), " ", "  "); err == nil {
			cmd.OutOrStderr().Write(bytes)
		}
		fmt.Fprintln(cmd.OutOrStderr(), "\n----")
	}
	return err
}
