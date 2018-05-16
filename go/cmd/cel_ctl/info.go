// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"chromium.googlesource.com/enterprise/cel/go/gcp"
	"chromium.googlesource.com/enterprise/cel/go/gcp/deploy"
	"github.com/spf13/cobra"
)

func init() {
	ic := &InfoCommand{}
	cmd := &cobra.Command{
		Use:   "info [host environment files]",
		Short: "Gathers information about current enterprise lab deployment",
		Long: `Gathers information about current enterprise lab deployment.

This command can be used to gather detailed information about a deployment that was started using the "cel_ctl deploy" command.

The information is written to a series of files in the current directory as follows:

  gcp-deployment.json : A JSON encoded deployment record. The resource representation is described in https://cloud.google.com/deployment-manager/docs/reference/v2beta/deployments
  
  gcp-manifest.json : A JSON encoded manifest record for the latest deployment. The resource representation is described in https://cloud.google.com/deployment-manager/docs/reference/v2beta/manifests

  gcp-expanded-config.yaml : The YAML encoded expanded deployment manifest.

  gcp-resources.json : A JSON encoded list of resources that were deployed. The resource representation is described in https://cloud.google.com/deployment-manager/docs/reference/v2beta/resources and include information about the state and any errors that were encountered while deploying the resource.
`,
		Args: cobra.MinimumNArgs(1),
	}

	cmd.Flags().BoolVarP(&ic.UseBuiltins, "builtins", "B", false, "Use builtin assets")
	app.AddCommand(cmd, ic)
}

type InfoCommand struct {
	UseBuiltins bool
}

func (i *InfoCommand) Run(ctx context.Context, a *Application, cmd *cobra.Command, args []string) error {
	s, err := a.CreateSession(ctx, args, i.UseBuiltins)
	if err != nil {
		return err
	}

	d, err := deploy.GetCurrentDeployment(s.GetContext(), s.GetBackend())
	if err != nil {
		if gcp.IsNotFoundError(err) {
			fmt.Fprintf(cmd.OutOrStderr(), "There are no existing CEL deployments for the project %s\n", s.GetBackend().GetProject())
			return nil
		}
		return err
	}

	const deploymentFilename = "gcp-deployment.json"
	err = writeObjectAsJson(deploymentFilename, d)
	if err != nil {
		return err
	}
	fmt.Fprintf(cmd.OutOrStdout(), "Wrote deployment record to %s\n", deploymentFilename)

	if d.Update == nil {
		return err
	}

	m, err := deploy.GetManifest(s.GetContext(), s.GetBackend(), d.Update.Manifest)
	if err != nil {
		return err
	}

	const manifestFilename = "gcp-manifest.json"
	err = writeObjectAsJson(manifestFilename, d)
	if err != nil {
		return err
	}
	fmt.Fprintf(cmd.OutOrStdout(), "Wrote manifest record to %s\n", manifestFilename)

	const expandedConfig = "gcp-expanded-config.yaml"
	err = writeObjectToFile(expandedConfig, []byte(m.ExpandedConfig))
	if err != nil {
		return err
	}
	fmt.Fprintf(cmd.OutOrStdout(), "Wrote expanded deployment manifest to %s\n", expandedConfig)

	rl, err := deploy.GetResources(s.GetContext(), s.GetBackend())
	if err != nil {
		return err
	}

	const resourcesFile = "gcp-resources.json"
	err = writeObjectAsJson(resourcesFile, rl)
	if err != nil {
		return err
	}
	fmt.Fprintf(cmd.OutOrStdout(), "Wrote resource list to %s\n", resourcesFile)

	return err
}

func writeObjectAsJson(fname string, o interface{}) error {
	j, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return err
	}

	return writeObjectToFile(fname, j)
}

func writeObjectToFile(fname string, b []byte) error {
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(b)
	return err
}
