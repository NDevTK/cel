// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"chromium.googlesource.com/enterprise/cel/go/gcp/deploy"
	"context"
	"github.com/spf13/cobra"
)

type PurgeCommand struct {
	UseBuiltins bool
}

func init() {
	pc := &PurgeCommand{}
	cmd := &cobra.Command{
		Use:   "purge [configuration files]",
		Short: "purge deployed lab artifacts",
		Long: `Purges artifacts that were previously deployed via a 'Deploy' command.
`}
	cmd.Flags().BoolVarP(&pc.UseBuiltins, "builtins", "B", false, "Use builtin assets")
	app.AddCommand(cmd, pc)
}

func (p *PurgeCommand) Run(ctx context.Context, a *Application, cmd *cobra.Command, args []string) error {
	session, err := a.CreateSession(ctx, args, p.UseBuiltins)
	if err != nil {
		return err
	}

	return deploy.DeleteObsoleteDeployments(session.GetContext(), session.GetBackend())
}
