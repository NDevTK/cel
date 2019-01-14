// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"chromium.googlesource.com/enterprise/cel/go/cel/deploy"
	"context"
	"github.com/spf13/cobra"
	"time"
)

type WaitForCommand struct {
	UseBuiltins bool
	Timeout     int
}

func init() {
	wfc := &WaitForCommand{}
	cmd := &cobra.Command{
		Use:   "waitfor [configuration files]",
		Short: "Wait for onhost configuration on target lab environment",
		Long: `Waits for onhost configuration to finish for all assets on target lab environment.
`,
	}
	cmd.Flags().BoolVarP(&wfc.UseBuiltins, "builtins", "B", false, "Use builtin assets")
	cmd.Flags().IntVarP(&wfc.Timeout, "timeout", "T", 3600, "Timeout in seconds to wait for onhost configuration")
	app.AddCommand(cmd, wfc)
}

func (d *WaitForCommand) Run(ctx context.Context, a *Application, cmd *cobra.Command, args []string) error {
	session, err := a.CreateSession(ctx, args, d.UseBuiltins)
	if err != nil {
		return err
	}

	return deploy.WaitForAllAssetsReady(session.GetBackend(), time.Duration(d.Timeout)*time.Second)
}
