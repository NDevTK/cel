// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"github.com/spf13/cobra"
)

type DeployCommand struct{}

func init() {
	app.AddCommand(&cobra.Command{
		Use:   "deploy",
		Short: "deploy build artifacts to target lab environment",
		Long: `Deploys build artifacts to target lab environment.

Use as: deploy [target]
`,
	}, &DeployCommand{})
}

func (d *DeployCommand) Run(ctx context.Context, a *Application, cmd *cobra.Command, args []string) error {
	panic("not implemented")
}
