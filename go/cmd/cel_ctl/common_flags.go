// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"chromium.googlesource.com/enterprise/cel/go/cel"
	"context"
	"flag"
)

type CommonFlags struct {
	Configuration cel.Configuration

	AssetFiles   []string
	HostFiles    []string
	GenericFiles []string
	Verbose      bool
}

func (c *CommonFlags) SetFlags(f *flag.FlagSet) {
	f.Var(ArrayValueFlag{&c.HostFiles}, "host",
		`hosting environment configuration. This should be a .textpb file conforming to config.proto:HostEnvironment.`)
	f.Var(ArrayValueFlag{&c.HostFiles}, "h", "alias for \"--host\"")
	f.Var(ArrayValueFlag{&c.AssetFiles}, "assets",
		`asset configuration. This should be a .textpb file conforming to config.proto:Assets. Specify multiple times to add additional asset configuration files.`)
	f.Var(ArrayValueFlag{&c.AssetFiles}, "a", "alias for \"--assets\"")
	f.BoolVar(&c.Verbose, "verbose", false, `verbose output`)
	f.BoolVar(&c.Verbose, "v", false, "alias for \"verbose\"")
}

func (c *CommonFlags) Load(ctx context.Context, f *flag.FlagSet) {
}
