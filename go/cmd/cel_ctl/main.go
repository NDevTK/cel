// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"github.com/google/subcommands"
)

func main() {
	subcommands.Register(subcommands.CommandsCommand(), "help")
	subcommands.Register(subcommands.HelpCommand(), "help")
	subcommands.Register(subcommands.FlagsCommand(), "help")

	subcommands.Execute(context.Background())
}
