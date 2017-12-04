// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"net/http"

	"chromium.googlesource.com/enterprise/cel/go/cel"
	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/gcp"
	"github.com/spf13/cobra"
)

type App struct {
	Configuration cel.Configuration
	Session       *gcp.Session
	Client        *http.Client

	GenericFiles []string
	Verbose      bool
}

var root = &cobra.Command{
	Use:              "cel_ctl",
	Short:            "Tools for managing your Chrome Enterprise Lab",
	Long:             "",
	TraverseChildren: true,
}

var app = &App{}

func init() {
	app.SetFlags(root)
}

func (c *App) SetFlags(cmd *cobra.Command) {
	cmd.LocalFlags().BoolVarP(&c.Verbose, "verbose", "v", false, `verbose output`)
}

func (c *App) GetSession(ctx context.Context) (session *gcp.Session, err error) {
	if c.Session != nil {
		return c.Session, nil
	}
	c.Session, err = gcp.NewSession(ctx, c.Client, &c.Configuration.HostEnvironment)
	if err != nil {
		return nil, err
	}
	return c.Session, nil
}

func (c *App) Load(ctx context.Context, files []string) (err error) {
	c.GenericFiles = files

	for _, f := range c.GenericFiles {
		err := c.Configuration.Merge(f)
		if err != nil {
			return err
		}
	}

	err = c.Configuration.Validate()
	if err != nil {
		return err
	}

	c.Client, err = common.GetDefaultClient(ctx)
	if err != nil {
		return err
	}

	return nil
}

type Executor interface {
	Execute(context.Context, *App, *cobra.Command, []string) error
}

func (c *App) InvokeCommand(e Executor, cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	return e.Execute(ctx, c, cmd, args)
}
