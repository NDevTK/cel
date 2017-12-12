// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"github.com/spf13/cobra"
)

type DepsCommand struct{}

func (d *DepsCommand) Run(context.Context, *Application, *cobra.Command, []string) error {
	panic("not implemented")
}

func init() {
	app.AddCommand(&cobra.Command{
		Use:   "deps",
		Short: "Show a dependency graph of all assets in .dot format",
		Long: `Shows a dependency graph of all assets in .dot format

All the assets named in the configuration file will be included. Doesn't list
assets from the live environment.

`,
	}, &DepsCommand{})
}
