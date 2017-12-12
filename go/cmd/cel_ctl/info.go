// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
)

func init() {
	app.AddCommand(&cobra.Command{
		Use:   "info [configuration files]",
		Short: "Show information about current enterprise lab configuration",
		Long: `Shows information about current enterprise lab configuration.

Includes information about the desired state, and also the current state of the target Google Compute Engine project. This will spew a *lot* of information in JSON format.
`,
		Args: cobra.MinimumNArgs(1),
	}, &InfoCommand{})
}

type InfoCommand struct {
}

func (i *InfoCommand) Run(ctx context.Context, a *Application, cmd *cobra.Command, args []string) error {
	err := a.LoadConfigFiles(ctx, args)
	if err != nil {
		return err
	}
	bytes, err := json.MarshalIndent(a.Configuration, " ", "  ")
	if err != nil {
		return err
	}
	_, err = cmd.OutOrStdout().Write(bytes)
	return err
}
