// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"time"

	"chromium.googlesource.com/enterprise/cel/go/gcp/deploy"
)

type CloneCommand struct {
	UseBuiltins bool
	SourceHost  string
	TargetHost  string
}

func init() {
	clone := &CloneCommand{}
	cmd := &cobra.Command{
		Use:   "clone --source project [configuration files]",
		Short: "clone compute instances in a lab environment to another environment",
		Long: `Clones a lab environment to another target lab environment.
`,
	}
	cmd.Flags().BoolVarP(&clone.UseBuiltins, "builtins", "B", false, "Use builtin assets")
	cmd.Flags().StringVar(&clone.SourceHost, "source", "", "host file for the source environment")
	cmd.Flags().StringVar(&clone.TargetHost, "target", "", "host file for the target environment")
	cmd.MarkFlagRequired("source")
	app.AddCommand(cmd, clone)
}

func (c *CloneCommand) Run(ctx context.Context, a *Application, cmd *cobra.Command, args []string) (err error) {
	log.Printf("Start of `cel_ctl clone` - version %s", version)

	// Create a cel session for both hosts passed as argument
	source, err := a.CreateSession(ctx, append(args, c.SourceHost), c.UseBuiltins)
	if err != nil {
		return err
	}
	srcBackend := source.GetBackend()
	srcProject := srcBackend.GetProject()

	target, err := a.CreateSession(ctx, append(args, c.TargetHost), c.UseBuiltins)
	if err != nil {
		return err
	}

	// Allow TargetHost DeploymentManager to read SourceHost's images.
	saEmail, err := deploy.GetDeploymentManagerServiceAccount(target.GetContext(), target.GetBackend())
	if err != nil {
		return err
	}
	err = deploy.AddServiceAccountBinding(source.GetContext(), srcBackend, saEmail, "roles/compute.imageUser")
	if err != nil {
		return err
	}

	// Create the environment snapshot in the source host
	snapshotName := fmt.Sprintf("clone-%s", time.Now().Format("20060102-1504"))
	create := CreateSnapshotCommand{UseBuiltins: c.UseBuiltins, SnapshotName: snapshotName}
	err = create.Run(ctx, a, cmd, append(args, c.SourceHost))
	if err != nil {
		return err
	}

	// Delete the temporary snapshot we created when we return.
	defer func() {
		delete := DeleteSnapshotCommand{Project: srcProject, SnapshotName: snapshotName}
		errDelete := delete.Run(ctx, a, cmd, nil)

		// Only replace err if it hasn't been set yet
		if err == nil && errDelete != nil {
			err = errDelete
		}
	}()

	// Restore that environment in the target host
	restore := RestoreSnapshotCommand{UseBuiltins: c.UseBuiltins, SourceProject: srcProject, SnapshotName: snapshotName}
	err = restore.Run(ctx, a, cmd, append(args, c.TargetHost))
	if err != nil {
		return err
	}

	return nil
}
