// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"fmt"
	"strings"
)

// ArrayValueFlag implements flag.Value and can be used for specifying a flag
// that can be used multiple times. For example:
//
//     var foo []string
//     flag.Var(ArrayValueFlag{&foo}, "foo", "use multiple times with an argument")
//
// Now, the '-foo' argument can take multiple values:
//
//     flagset.Parse([]string{"<exename>", "-foo", "a", "-foo", "b"})
//
// The above results in |foo| being equal to []string{"a", "b"}.
type ArrayValueFlag struct {
	Array *[]string
}

func (a ArrayValueFlag) String() string {
	if a.Array == nil {
		return "nil"
	}
	return strings.Join(*a.Array, ":")
}

func (a ArrayValueFlag) Set(s string) error {
	if a.Array == nil {
		return fmt.Errorf("array not initialized")
	}
	*a.Array = append(*a.Array, s)
	return nil
}
