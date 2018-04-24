// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"chromium.googlesource.com/enterprise/cel/go/cel"
	"context"
	"github.com/spf13/cobra"
)

func init() {
	dc := &DepsCommand{}
	cmd := &cobra.Command{
		Use:   "deps [configuration files]",
		Short: "Show a dependency graph of all assets in .dot format",
		Long: `Shows a dependency graph of all assets in .dot format

All the assets named in the configuration file will be included. Doesn't list
assets from the live environment.

`,
	}
	cmd.Flags().BoolVarP(&dc.Prune, "prune", "p", false, "Prune dependencies")
	cmd.Flags().BoolVarP(&dc.UseBuiltins, "builtins", "B", false, "Use builtin assets")
	app.AddCommand(cmd, dc)
}

type DepsCommand struct {
	Prune       bool
	UseBuiltins bool
}

func (d *DepsCommand) Run(ctx context.Context, a *Application, cmd *cobra.Command, args []string) error {
	session, err := a.CreateSession(ctx, args, d.UseBuiltins)
	if err != nil {
		return err
	}

	err = cel.InvokeAdditionalDependencyResolvers(session)
	if err != nil {
		return err
	}

	if d.Prune {
		err = cel.Prune(session)
		if err != nil {
			return err
		}
	}

	g, err := session.GetConfiguration().GetNamespace().AsSerializedDOTGraph()
	if err != nil {
		return err
	}

	_, err = cmd.OutOrStdout().Write([]byte(g))
	return err
}
