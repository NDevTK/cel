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

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: cel_agent manifest_file")
		return
	}

	flag.Parse()
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Printf("Start of cel_agent")

	d, err := onhost.CreateDeployer()
	if err != nil {
		log.Printf("Deployer creation failed. error: %s", err)
		return
	}

	defer d.Close()
	d.Deploy(os.Args[1])

	// Keep running and watch for command signals
	c, err := onhost.CreateCommander()
	if err != nil {
		log.Printf("Commander creation failed. error: %s", err)
		return
	}

	defer c.Close()
	c.WatchForCommands()
}
