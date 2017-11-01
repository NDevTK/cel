// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
)

type DepsCommand struct {
	CommonFlags
}

func (d *DepsCommand) Name() string { return "deps" }

func (d *DepsCommand) Synopsis() string { return "Show a dependency graph of all assets in .dot format" }

func (d *DepsCommand) Usage() string {
	return `Shows a dependency graph of all assets in .dot format

All the assets named in the configuration file will be included. Doesn't list
assets from the live environment.`
}

func (d *DepsCommand) SetFlags(f *flag.FlagSet) {
	d.CommonFlags.SetFlags(f)
}

func (d *DepsCommand) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	client, err := lab.GetDefaultClient(ctx)
	if err != nil {
		fmt.Fprintln(a.GetErr(), "failed to initialize http client: ", err)
		return subcommands.ExitFailure
	}

	s := d.GetSession(a, ctx, client)
	if s == nil {
		return subcommands.ExitFailure
	}

	A := lab.Assets{}
	err = lab.ConstructAssets(&A, s.Config)
	if err != nil {
		fmt.Fprintln(a.GetErr(), "failed to construct assets: ", err)
		return subcommands.ExitFailure
	}

	_, err = lab.PrepareToResolve(&A)
	if err != nil {
		fmt.Fprintln(a.GetErr(), "failed to resolve assets: ", err)
		return subcommands.ExitFailure
	}

	err = lab.DumpAssetDepsInDotFormat(&A, a.GetOut())
	if err != nil {
		fmt.Fprintln(a.GetErr(), "failed to generate graph: ", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
