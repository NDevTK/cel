// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
)

type DeployCommand struct {
	CommonFlags
	Source string
}

func (d *DeployCommand) Name() string { return "deploy" }

func (d *DeployCommand) Synopsis() string { return "deploy build artifacts to target lab environment" }

func (d *DeployCommand) Usage() string {
	return `Deploys build artifacts to target lab environment.

Use as: deploy [target]
`
}

func (d *DeployCommand) SetFlags(f *flag.FlagSet) {
	d.CommonFlags.SetFlags(f)
	f.StringVar(&d.Source, "source", "", "path to root of source directory")
	f.StringVar(&d.Source, "s", "", "alias for \"--source\"")
}

func (d *DeployCommand) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	panic("not implemented")
	return subcommands.ExitFailure
}
