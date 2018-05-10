// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"net/http"

	"chromium.googlesource.com/enterprise/cel/go/cel/deploy"
	"chromium.googlesource.com/enterprise/cel/go/gcp"
	"github.com/spf13/cobra"
)

// Runner is an interface used for invoking a command with a Context, an
// Application, a Command, and a list of free arguments.
type Runner interface {
	// Run should do whatever the command wants to do. |args| is a list of
	// arguments that were left over after the known options  and arguments
	// were parsed.
	Run(ctx context.Context, a *Application, cmd *cobra.Command, args []string) error
}

// Application is the central hub for all the commands and is also houses the
// logic for common operations like loading configuration files.
//
// Unlike usual cobra applications, ours has a non-public internal rootCmd
// (called rootCommand).  This is important because we want to control how
// commands are added and invoked. Cobra's Command object is a struct and
// doesn't directly allow attaching auxiliary objects and doesn't allow dynamic
// flag discovery. In addition, the Run() methods in cobra.Command doesn't
// allow passing in a context or additional arguments that are determined
// *after* the flags were added.
//
// To mitigate these we are using a separate AddCommand() method that takes as
// input a Runner interface that supports all the parameters we care about.
type Application struct {
	Session *deploy.Session
	Client  *http.Client

	GenericFiles []string
	Verbose      bool

	rootCommand *cobra.Command
}

// setFlags is called during init() to register common flags. These are
// available to all commands and must be exposed via public fields in the
// Application struct.
func (a *Application) setFlags() {
	a.rootCommand.LocalFlags().BoolVarP(&a.Verbose, "verbose", "v", false, `verbose output`)
}

// CreateSession creates a DeployerSession based on a set of configuration
// files.
func (a *Application) CreateSession(ctx context.Context, inputs []string, useBuiltins bool) (session *deploy.Session, err error) {
	a.Client, err = gcp.GetDefaultClient(ctx)
	if err != nil {
		return nil, err
	}
	a.Session, err = deploy.NewSession(ctx, a.Client, inputs, useBuiltins)
	if err != nil {
		return nil, err
	}
	return a.Session, nil
}

// AddCommand should be used by init() functions to add new commands to the
// application. Each command consists of a cobra.Command and a Runner.
func (a *Application) AddCommand(cmd *cobra.Command, e Runner) {
	if cmd.Run != nil || cmd.RunE != nil {
		panic("AddCommand called with Command that has a Run or RunE field specified.")
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return a.invoke(e, cmd, args)
	}
	a.rootCommand.AddCommand(cmd)
}

// invoke executes |e|.Run() with a Context.
func (a *Application) invoke(e Runner, cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	return e.Run(ctx, a, cmd, args)
}

// Run runs the application using os.Args[1:] as the list of arguments.
func (a *Application) Run() error {
	return a.rootCommand.Execute()
}

var app = &Application{
	rootCommand: &cobra.Command{
		Use:              "cel_ctl",
		Short:            "Tools for managing your Chrome Enterprise Lab",
		Long:             "",
		TraverseChildren: true,
		SilenceErrors:    true,
		SilenceUsage:     true,
	},
}

func init() {
	app.setFlags()
}
