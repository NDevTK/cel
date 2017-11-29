// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"chromium.googlesource.com/enterprise/cel/go/gcp"
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
	"os"
)

type PasswdCommand struct {
	CommonFlags
	Instance string
	Username string
	Email    string
}

func (p *PasswdCommand) Name() string { return "passwd" }

func (p *PasswdCommand) Synopsis() string { return "reset password on a Windows instance" }

func (p *PasswdCommand) Usage() string {
	return `Resets the password for a Windows instance.
`
}

func (p *PasswdCommand) SetFlags(f *flag.FlagSet) {
	p.CommonFlags.SetFlags(f)
	f.StringVar(&p.Instance, "instance", "", "short instance name of VM")
	f.StringVar(&p.Username, "username", "", "username of account to reset password")
	f.StringVar(&p.Email, "email", "", "email address to associate with account")
}

func (p *PasswdCommand) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if p.Instance == "" || p.Username == "" {
		fmt.Fprintln(os.Stderr, "instance and username are required options")
		return subcommands.ExitUsageError
	}

	err := p.Load(ctx, f)
	if err != nil {
		fmt.Fprintln(os.Stderr, "%s", err.Error())
		return subcommands.ExitFailure
	}

	instance := p.Session.Cloud.Instances[p.Instance]
	if instance == nil {
		fmt.Fprintln(os.Stderr, "instance not found: ", p.Instance)
		return subcommands.ExitFailure
	}

	password, err := gcp.ResetWindowsPassword(ctx, p.Client, p.Configuration.HostEnvironment.Project.Name, instance.Zone, instance.Name, p.Username, p.Email)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to reset password: ", err)
		return subcommands.ExitFailure
	}

	fmt.Fprintln(os.Stdout, password)
	return subcommands.ExitSuccess
}
