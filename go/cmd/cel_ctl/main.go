// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"chromium.googlesource.com/enterprise/cel/go/common"
)

func showIndentedLine(i int, s string, w io.Writer) {
	fmt.Fprintf(w, "%s%s\n", strings.Repeat(" ", i*2), s)
}

func showIndentedError(i int, err error, w io.Writer) {
	type causer interface {
		Cause() error
	}

	switch e := err.(type) {
	case common.ErrorUnpacker:
		el := e.UnpackErrors()
		showIndentedLine(i, "{", w)
		for _, ei := range el {
			showIndentedError(i+1, ei, w)
		}
		showIndentedLine(i, "}", w)

	case causer:
		cl := strings.Split(err.Error(), ": ")
		for _, c := range cl {
			showIndentedLine(i, c, w)
			i++
		}

	default:
		showIndentedLine(i, err.Error(), w)
	}
}

func showStructuredError(err error, w io.Writer) {
	if err == nil {
		return
	}

	showIndentedError(0, err, w)
}

func main() {
	err := app.Run()
	if err != nil {
		showStructuredError(err, os.Stderr)
		os.Exit(1)
	}
}
