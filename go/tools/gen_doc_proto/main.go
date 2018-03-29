// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func usage(f *flag.FlagSet) {
	fmt.Fprintln(os.Stderr, `gen_doc_proto: Generate Markdown documentation from .proto files.

This tool is intended to be used with the Chrome Enterprise Lab toolchain for
producing Markdown documentation for ProtoBuf schema.

Usage:
    gen_doc_proto -out <output-file> <proto-file>...

The <proto-file>s should contain a binary encoded FileDescriptorProtoSet
protocol buffer. This could be generated using the --descriptor_set_out option
to the ProtoBuf compiler ('protoc'). It is required that the SourceCodeInfo
fields be intact in the supplied FileDescriptorProtoSet. The latter can be
ensured via the --include_source_info argument to protoc.

The MarkDown documentation will be written to <output-file>.

Within the Chrome Enterprise Lab build environment, you can generate the
documentation via the build.py build tool by invoking it as "build.py gen".

Each invocation of gen_doc_proto must be limited to a single package.`)
	os.Exit(2)
}

func fatal(action string) {
	fatalf(nil, action)
}

func fatalf(err error, action string, v ...interface{}) {
	p_name := filepath.Base(os.Args[0])
	if err == nil {
		log.Printf("%s (%s)", p_name, fmt.Sprintf(action, v...))
	} else {
		log.Printf("%s (%s): %s", p_name, fmt.Sprintf(action, v...), err.Error())
	}
	os.Exit(1)
}

func main() {
	var d generator
	var showHelp bool
	var outputPath string

	flags := flag.NewFlagSet("gen_doc_proto", flag.ExitOnError)
	flags.SetOutput(os.Stderr)
	flags.StringVar(&outputPath, "out", "", "")
	flags.BoolVar(&showHelp, "h", false, "")
	flags.BoolVar(&showHelp, "help", false, "")

	err := flags.Parse(os.Args[1:])

	if showHelp || err == flag.ErrHelp || outputPath == "" || len(flags.Args()) == 0 {
		usage(flags)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while parsing command line options: ", err.Error())
		usage(flags)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		fatalf(err, "writing output file to %s", outputPath)
	}
	defer file.Close()

	for _, f := range flags.Args() {
		err = d.MergeFileDescriptorSet(f)
		if err != nil {
			fatalf(err, "while attempting to parse %s", f)
		}
	}

	err = d.Gen(file)
	if err != nil {
		fatalf(err, "while running generator")
	}
}
