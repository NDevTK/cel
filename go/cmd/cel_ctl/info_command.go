// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
)

type InfoCommand struct {
	CommonFlags
}

func (i *InfoCommand) Name() string { return "info" }

func (i *InfoCommand) Synopsis() string {
	return "Show information about current enterprise lab configuration"
}

func (i *InfoCommand) Usage() string {
	return `Shows information about current enterprise lab configuration.

Includes information about the desired state, and also the current state of the
target Google Compute Engine project. This will spew a *lot* of information in
JSON format.`
}

func (i *InfoCommand) SetFlags(f *flag.FlagSet) {
	i.CommonFlags.SetFlags(f)
}

func (i InfoCommand) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	client, err := lab.GetDefaultClient(ctx)
	if err != nil {
		fmt.Fprintln(a.GetErr(), "failed to initialize http client: ", err)
		return subcommands.ExitFailure
	}

	s := i.GetSession(a, ctx, client)
	if s == nil {
		return subcommands.ExitFailure
	}

	b, err := json.MarshalIndent(s.Config, "", "  ")
	if err != nil {
		fmt.Fprintln(a.GetErr(), "failed to marshall JSON: ", err)
		return subcommands.ExitFailure
	}

	fmt.Fprintf(a.GetOut(), "%s", string(b))
	return subcommands.ExitSuccess
}

func init() {
	subcommands.Register(&InfoCommand{}, "")
}
