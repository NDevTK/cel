// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(&cobra.Command{
		Use:   "deploy",
		Short: "deploy build artifacts to target lab environment",
		Long: `Deploys build artifacts to target lab environment.

Use as: deploy [target]
`,
		Run: func(c *cobra.Command, args []string) {
			panic("not implemented")
		}})
}
