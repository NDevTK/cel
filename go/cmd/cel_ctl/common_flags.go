// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"chromium.googlesource.com/enterprise/cel/go/cel"
	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/gcp"
	"context"
	"flag"
	"net/http"
)

type CommonFlags struct {
	Configuration cel.Configuration
	Session       *gcp.Session
	Client        *http.Client

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

func (c *CommonFlags) Load(ctx context.Context, f *flag.FlagSet) error {
	c.GenericFiles = f.Args()

	for _, f := range c.AssetFiles {
		err := c.Configuration.MergeAssets(f)
		if err != nil {
			return err
		}
	}

	for _, f := range c.HostFiles {
		err := c.Configuration.MergeHost(f)
		if err != nil {
			return err
		}
	}

	for _, f := range c.GenericFiles {
		err := c.Configuration.Merge(f)
		if err != nil {
			return err
		}
	}

	err := c.Configuration.Validate()
	if err != nil {
		return err
	}

	c.Client, err = common.GetDefaultClient(ctx)
	if err != nil {
		return err
	}

	c.Session, err = gcp.NewSession(ctx, c.Client, &c.Configuration.HostEnvironment)
	if err != nil {
		return err
	}

	return nil
}
