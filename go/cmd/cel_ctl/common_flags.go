// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"flag"
)

type CommonFlags struct {
	HostConfigs  []string
	AssetConfigs []string
	Verbose      bool
}

func (c *CommonFlags) SetFlags(f *flag.FlagSet) {
	f.Var(ArrayValueFlag{&c.HostConfigs}, "host",
		`hosting environment configuration. This should be a .textpb file conforming to config.proto:HostEnvironment.`)
	f.Var(ArrayValueFlag{&c.HostConfigs}, "-h", "alias for \"--host\"")
	f.Var(ArrayValueFlag{&c.AssetConfigs}, "assets",
		`asset configuration. This should be a .textpb file conforming to config.proto:Assets. Specify multiple times to add additional asset configuration files.`)
	f.Var(ArrayValueFlag{&c.AssetConfigs}, "-a", "alias for \"--assets\"")
	f.BoolVar(&c.Verbose, "verbose", false, `verbose output`)
}
