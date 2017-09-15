// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/maruel/subcommands"
	"go/lab"
	"net/http"
	"strings"
)

type ArrayValueFlag struct {
	Array *[]string
}

func (a ArrayValueFlag) String() string {
	if a.Array == nil {
		return "nil"
	}
	return strings.Join(*a.Array, ":")
}

func (a ArrayValueFlag) Set(s string) error {
	if a.Array == nil {
		return fmt.Errorf("array not initialized")
	}
	*a.Array = append(*a.Array, s)
	return nil
}

type CommonFlags struct {
	Session *lab.Session

	HostConfig   string
	AssetConfigs []string
	UseBuiltins  bool
}

func (c *CommonFlags) Configure(f *flag.FlagSet) {
	f.StringVar(&c.HostConfig, "host", "",
		`Hosting environment configuration. This should be a .textpb file conforming to config.proto:HostEnvironment.`)
	f.Var(ArrayValueFlag{&c.AssetConfigs}, "assets",
		`Asset configuration. This should be a .textpb file conforming to config.proto:Assets. Specify multiple times to add additional asset configuration files.`)
	f.BoolVar(&c.UseBuiltins, "builtin", false, `Include builtin assets.`)
}

func (c *CommonFlags) GetSession(a subcommands.Application, ctx context.Context, client *http.Client) *lab.Session {
	if c.Session != nil {
		return c.Session
	}

	if c.HostConfig == "" {
		fmt.Fprintln(a.GetErr(), "the --host argument is required")
		return nil
	}

	if len(c.AssetConfigs) == 0 {
		fmt.Fprintln(a.GetErr(), "the --assets argument is required")
		return nil
	}

	var err error
	c.Session, err = lab.NewSession(ctx, client, c.HostConfig, c.AssetConfigs)
	if err != nil {
		fmt.Fprintln(a.GetErr(), "failed to initialize session: ", err)
		c.Session = nil
	}
	return c.Session
}

//------------------------------------------------------------------------------
// Dump State command
//
var InfoCommand = subcommands.Command{
	UsageLine: "dump_state [options]",
	ShortDesc: "Show information about current enterprise lab configuration",
	LongDesc: `Shows information about current enterprise lab configuration.

Includes information about the desired state, and also the current state of the
target Google Compute Engine project. This will spew a *lot* of information in
JSON format.`,
	CommandRun: func() subcommands.CommandRun {
		r := &InfoCommandRun{}
		r.Configure(r.GetFlags())
		return r
	}}

type InfoCommandRun struct {
	subcommands.CommandRunBase
	CommonFlags
}

func (i *InfoCommandRun) Run(a subcommands.Application, args []string, evn subcommands.Env) int {
	ctx := context.Background()

	client, err := lab.GetDefaultClient(ctx)
	if err != nil {
		fmt.Fprintln(a.GetErr(), "failed to initialize http client: ", err)
		return -2
	}

	s := i.GetSession(a, ctx, client)
	if s == nil {
		return -1
	}

	b, err := json.MarshalIndent(s.Config, "", "  ")
	if err != nil {
		fmt.Fprintln(a.GetErr(), "failed to marshall JSON: ", err)
		return -3
	}

	fmt.Fprintf(a.GetOut(), "%s", string(b))
	return 0
}

//------------------------------------------------------------------------------
// Dependency graph command
//
var DepsCommand = subcommands.Command{
	UsageLine: "deps [options]",
	ShortDesc: "Show a dependency graph of all assets in .dot format",
	LongDesc: `Shows a dependency graph of all assets in .dot format

All the assets named in the configuration file will be included. Doesn't list
assets from the live environment.`,
	CommandRun: func() subcommands.CommandRun {
		r := &DepsCommandRun{}
		r.Configure(r.GetFlags())
		return r
	}}

type DepsCommandRun struct {
	subcommands.CommandRunBase
	CommonFlags
}

func (d *DepsCommandRun) Run(a subcommands.Application, args []string, env subcommands.Env) int {
	ctx := context.Background()

	client, err := lab.GetDefaultClient(ctx)
	if err != nil {
		fmt.Fprintln(a.GetErr(), "failed to initialize http client: ", err)
		return -2
	}

	s := d.GetSession(a, ctx, client)
	if s == nil {
		return -1
	}

	A := lab.Assets{}
	err = lab.ConstructAssets(&A, s.Config)
	if err != nil {
		fmt.Fprintln(a.GetErr(), "failed to construct assets: ", err)
		return -3
	}

	_, err = lab.PrepareToResolve(&A)
	if err != nil {
		fmt.Fprintln(a.GetErr(), "failed to resolve assets: ", err)
		return -5
	}

	err = lab.DumpAssetDepsInDotFormat(&A, a.GetOut())
	if err != nil {
		fmt.Fprintln(a.GetErr(), "failed to generate graph: ", err)
		return -4
	}

	return 0
}

//------------------------------------------------------------------------------
// Set Windows Password command
//
var SetPasswordCommand = subcommands.Command{
	UsageLine: "passwd [options]",
	ShortDesc: "Reset Windows password for a user account",
	CommandRun: func() subcommands.CommandRun {
		r := &SetPasswordCommandRun{}
		r.Configure(r.GetFlags())
		return r
	}}

type SetPasswordCommandRun struct {
	subcommands.CommandRunBase
	CommonFlags

	instance string
	username string
	email    string
}

func (s *SetPasswordCommandRun) Configure(f *flag.FlagSet) {
	f.StringVar(&s.instance, "instance", "", "short instance name of VM")
	f.StringVar(&s.username, "username", "", "username of account to reset password")
	f.StringVar(&s.email, "email", "", "email address to associate with account")
	s.CommonFlags.Configure(f)
}

func (s *SetPasswordCommandRun) Run(a subcommands.Application, args []string, evn subcommands.Env) int {
	if s.instance == "" || s.username == "" {
		fmt.Fprintln(a.GetErr(), "instance and username are required options")
		return -1
	}

	ctx := context.Background()
	client, err := lab.GetDefaultClient(ctx)
	if err != nil {
		fmt.Fprintln(a.GetErr(), "failed to get GCE client: ", err)
		return -2
	}

	S := s.GetSession(a, ctx, client)
	if S == nil {
		return -1
	}

	instance := S.Cloud.Instances[s.instance]
	if instance == nil {
		fmt.Fprintln(a.GetErr(), "instance not found: ", s.instance)
		return -3
	}

	password, err := lab.ResetWindowsPassword(ctx, client, S.Config.Project, instance.Zone, instance.Name, s.username, s.email)
	if err != nil {
		fmt.Fprintln(a.GetErr(), "failed to reset password: ", err)
		return -4
	}

	fmt.Fprintln(a.GetOut(), password)
	return 0
}

//------------------------------------------------------------------------------
// Deploy
//
var DeployCommand = subcommands.Command{
	UsageLine: "deploy [options]",
	ShortDesc: "Deploy build artifacts to target lab environment",
	CommandRun: func() subcommands.CommandRun {
		r := &DeployCommandRun{}
		r.Configure(r.GetFlags())
		return r
	}}

type DeployCommandRun struct {
	subcommands.CommandRunBase
	CommonFlags

	source string
}

func (d *DeployCommandRun) Configure(f *flag.FlagSet) {
	f.StringVar(&d.source, "source", "", "path to root of source directory")
	d.CommonFlags.Configure(f)
}

func (d *DeployCommandRun) Run(a subcommands.Application, args []string, evn subcommands.Env) int {
	ctx := context.Background()

	client, err := lab.GetDefaultClient(ctx)
	if err != nil {
		fmt.Fprintln(a.GetErr(), "failed to initialize http client: ", err)
		return -2
	}

	s := d.GetSession(a, ctx, client)
	if s == nil {
		return -1
	}

	A := lab.Assets{}
	err = lab.ConstructAssets(&A, s.Config)
	if err != nil {
		fmt.Fprintln(a.GetErr(), "failed to construct assets: ", err)
		return -3
	}

	err = lab.ResolveAssets(&A, s)
	if err != nil {
		fmt.Fprintln(a.GetErr(), "failed to deploy: ", err)
		return -4
	}

	return 0
}

func main() {
	a := subcommands.DefaultApplication{
		Name:  "cel_admin",
		Title: "Chrome Enterprise Lab Administrative Utility",
		Commands: []*subcommands.Command{
			&InfoCommand,
			&SetPasswordCommand,
			&DeployCommand,
			&DepsCommand,
			subcommands.CmdHelp}}
	subcommands.Run(&a, nil)
}
