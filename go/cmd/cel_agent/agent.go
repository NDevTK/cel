// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"chromium.googlesource.com/enterprise/cel/go/asset/onhost"
)

var version = "unknown"

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage of cel_agent:")
		fmt.Println("  cel_agent manifest_file")
		fmt.Println("  cel_agent --nocommander manifest_file")
		fmt.Println("  cel_agent --version")
		return
	}

	show_version := flag.Bool("version", false, "version for cel_agent")
	skip_commander := flag.Bool("nocommander", false, "skip the commander and return after deployment")
	flag.Parse()

	if *show_version {
		fmt.Println("cel_agent version", version)
		return
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Printf("Start of cel_agent version %s", version)

	d, err := onhost.CreateDeployer()
	if err != nil {
		log.Printf("Deployer creation failed. error: %s", err)
		return
	}

	defer d.Close()
	d.Deploy(os.Args[len(os.Args)-1])

	if *skip_commander {
		log.Printf("Skipping Commander's WatchForCommands (nocommander).")
		return
	}

	// Keep running and watch for command signals
	c, err := onhost.CreateCommander(d)
	if err != nil {
		log.Printf("Commander creation failed. error: %s", err)
		return
	}

	defer c.Close()
	c.WatchForCommands()
}
