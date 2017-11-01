// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"flag"
	"github.com/google/subcommands"
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
	if p.instance == "" || p.username == "" {
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
