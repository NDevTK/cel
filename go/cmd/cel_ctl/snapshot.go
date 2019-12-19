// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strings"
	"time"

	"chromium.googlesource.com/enterprise/cel/go/cel"
	"chromium.googlesource.com/enterprise/cel/go/cel/deploy"
)

type CreateSnapshotCommand struct {
	UseBuiltins  bool
	SnapshotName string
}

type RestoreSnapshotCommand struct {
	UseBuiltins   bool
	SourceProject string
	SnapshotName  string
}

type ListSnapshotCommand struct {
	Project string
}

type DeleteSnapshotCommand struct {
	Project      string
	SnapshotName string
}

func init() {
	csc := &CreateSnapshotCommand{}
	cmd := &cobra.Command{
		Use:   "create-snapshot [--name name] [configuration files]",
		Short: "create snapshots of instances in target lab environment",
		Long: `Creates snapshots of compute instances in target lab environment.
`,
	}
	cmd.Flags().BoolVarP(&csc.UseBuiltins, "builtins", "B", false, "Use builtin assets")
	cmd.Flags().StringVar(&csc.SnapshotName, "name", time.Now().Format("20060102-1504"), "name of the snapshot (default is the current time (YYYYMMDD-HHMM))")
	app.AddCommand(cmd, csc)

	rsc := &RestoreSnapshotCommand{}
	cmd = &cobra.Command{
		Use:   "restore-snapshot --source project --name name [configuration files]",
		Short: "restore snapshots of instances in target lab environment",
		Long: `Restores snapshots of compute instances in target lab environment.
`,
	}
	cmd.Flags().BoolVarP(&rsc.UseBuiltins, "builtins", "B", false, "Use builtin assets")
	cmd.Flags().StringVar(&rsc.SourceProject, "source", "", "project where the snapshot lives (default to host file project)")
	cmd.Flags().StringVar(&rsc.SnapshotName, "name", "", "name of the snapshot to restore")
	cmd.MarkFlagRequired("name")
	app.AddCommand(cmd, rsc)

	lsc := &ListSnapshotCommand{}
	cmd = &cobra.Command{
		Use:   "list-snapshot --project project",
		Short: "list snapshots in target lab environment",
		Long: `Lists snapshots in target lab environment.
`,
	}
	cmd.Flags().StringVar(&lsc.Project, "project", "", "project to list snapshots")
	cmd.MarkFlagRequired("project")
	app.AddCommand(cmd, lsc)

	dsc := &DeleteSnapshotCommand{}
	cmd = &cobra.Command{
		Use:   "delete-snapshot --project project --name name",
		Short: "delete snapshots of instances in target lab environment",
		Long: `Deletes snapshots of compute instances in target lab environment.
`,
	}
	cmd.Flags().StringVar(&dsc.Project, "project", "", "project to delete a snapshot from")
	cmd.Flags().StringVar(&dsc.SnapshotName, "name", "", "name of the snapshot to delete")
	cmd.MarkFlagRequired("project")
	cmd.MarkFlagRequired("name")
	app.AddCommand(cmd, dsc)
}

func (csc *CreateSnapshotCommand) Run(ctx context.Context, a *Application, cmd *cobra.Command, args []string) error {
	log.Printf("Start of `cel_ctl create-snapshot` - version %s", version)

	// Validate arguments (e.g. snapshot name)
	if strings.Contains(csc.SnapshotName, "--") {
		return fmt.Errorf("snapshot name can't contain two consecutive dashes (\"--\")")
	}

	// Create a cel session from the host file passed as argument
	session, err := a.CreateSession(ctx, args, csc.UseBuiltins)
	if err != nil {
		return err
	}

	// Check that no snapshot with this name already exists
	backend := session.GetBackend()
	_, err = cel.FindEnvironmentSnapshot(ctx, backend.GetHttpClient(), backend.GetProject(), csc.SnapshotName)
	if err == nil {
		return fmt.Errorf("A snapshot with the name '%s' already exists", csc.SnapshotName)
	} else if err != cel.ErrSnapshotNotFound {
		return err
	}

	// Get RUNNING instances (to start them back up later)
	instances, err := cel.GetRunningInstances(session.GetContext(), backend)
	if err != nil {
		return err
	}

	// Running compute instances must be stopped before creating images from disk:
	// From Compute: Filesystem integrity can't be guaranteed while the instance
	//               is running which may create a corrupted image.
	err = cel.StopInstances(session.GetContext(), backend, instances)
	if err != nil {
		return err
	}

	// Create the environment snapshot
	err = cel.CreateSnapshots(session.GetContext(), backend, csc.SnapshotName)
	if err != nil {
		return err
	}

	// Restart the instances we previously stopped.
	return cel.StartInstances(session.GetContext(), backend, instances)
}

func (rsc *RestoreSnapshotCommand) Run(ctx context.Context, a *Application, cmd *cobra.Command, args []string) error {
	log.Printf("Start of `cel_ctl restore-snapshot` - version %s", version)

	// Create a cel session from the host file passed as argument
	session, err := a.CreateSession(ctx, args, rsc.UseBuiltins)
	if err != nil {
		return err
	}

	// Find the requested snapshot in the target host project or overriden in cmd args
	backend := session.GetBackend()
	sourceProject := rsc.SourceProject
	if sourceProject == "" {
		sourceProject = backend.GetProject()
	}

	snapshot, err := cel.FindEnvironmentSnapshot(ctx, backend.GetHttpClient(), sourceProject, rsc.SnapshotName)
	if err != nil {
		return err
	}

	// Verify that the images in the EnvironmentSnapshot match the machines in the asset file
	assets := session.GetConfiguration().AssetManifest
	for _, machine := range assets.WindowsMachine {
		_, ok := snapshot.Instances[machine.Name]
		if !ok {
			return fmt.Errorf("Couldn't find image for machine '%s' in snapshot '%s'", machine.Name, snapshot.Name)
		}
	}

	// Prepare the standard deployment configuration for these host/asset configs.
	err = deploy.PrepareDeploymentConfiguration(session)
	if err != nil {
		return err
	}

	// Modify the configuration to use images from the environment snapshot found earlier.
	cel.RestoreDeploymentConfigurationFromSnapshot(session.GetContext(), snapshot)
	if err != nil {
		return err
	}

	// Perform the actual deployment with updated deployment configuration.
	return deploy.InvokeDeploymentManager(session)
}

func (lsc *ListSnapshotCommand) Run(ctx context.Context, a *Application, cmd *cobra.Command, args []string) error {
	client, err := GetDefaultClient(ctx)
	if err != nil {
		return err
	}

	snapshots, err := cel.GetAllEnvironmentSnapshots(ctx, client, lsc.Project)
	if err != nil {
		return err
	}

	if len(snapshots) == 0 {
		fmt.Println("No celab environment snapshot found.")
		return nil
	}

	lineFormat := "%-20s%-25s%s\n"
	fmt.Printf(lineFormat, "SNAPSHOT", "CREATED ON", "MACHINES")
	for _, snapshot := range snapshots {
		var machines []string = make([]string, 0, len(snapshot.Instances))
		for name, _ := range snapshot.Instances {
			machines = append(machines, name)
		}

		// Turn the RFC3339 time (2019-02-01T15:52:34.289-08:00) to a more readable format.
		creationTime, err := time.Parse(time.RFC3339, snapshot.CreationTimestamp)
		if err != nil {
			return err
		}
		readableTime := creationTime.Format("02/01/2006 15:04 MST")

		fmt.Printf(lineFormat, snapshot.Name, readableTime, machines)
	}

	return nil
}

func (dsc *DeleteSnapshotCommand) Run(ctx context.Context, a *Application, cmd *cobra.Command, args []string) error {
	client, err := GetDefaultClient(ctx)
	if err != nil {
		return err
	}

	snapshot, err := cel.FindEnvironmentSnapshot(ctx, client, dsc.Project, dsc.SnapshotName)
	if err != nil {
		return err
	}

	fmt.Println("Deleting environment snapshot...")
	err = cel.DeleteEnvironmentSnapshot(ctx, client, dsc.Project, snapshot)
	if err != nil {
		return err
	}

	fmt.Println("Environment snapshot deleted.")

	return nil
}
