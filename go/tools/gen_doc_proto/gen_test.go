// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerator_Basic(t *testing.T) {
	var g generator
	var test_data = filepath.Join("testdata", "test.pb")

	err := g.MergeFileDescriptorSet(test_data)
	if err != nil {
		t.Errorf("Unexpected error while reading test data set: %#v", err)
	}

	var b bytes.Buffer
	err = g.Gen(&b)
	if err != nil {
		t.Errorf("Unexpected error while generating documentation: %#v", err)
	}

	contents := b.String()

	// We are going to spot test the documentaiton to make sure it includes
	// comments from testdata/test.proto. These tests should not depend on the
	// exact format of the output.
	must_haves := []string{
		"Package docs for testdata",
		"A Message",
		"A's name",
		"An annotated field",
		"B Message",
		"Another annotated field",
	}

	for _, s := range must_haves {
		if !strings.Contains(contents, s) {
			t.Errorf(`Expected string \"%s\" was not found in generated contents.
Generated contents:
%s`, s, contents)
		}
	}
}
