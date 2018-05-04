// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy_test

import (
	"chromium.googlesource.com/enterprise/cel/go/cel"
	"testing"
)

func TestConfiguration_ValidateBuiltins(t *testing.T) {
	var l cel.Configuration

	err := l.Merge("../../../resources/deployment/gcp-builtins.host.textpb")
	if err != nil {
		t.Fatal(err)
	}

	// This shouldn't be necessary, but by merging these two, we get a well
	// formed set of project settings. Otherwise validation will fail since
	// project and log parameters are required.
	err = l.Merge("../../../examples/schema/ad/one-domain.host.textpb")
	if err != nil {
		t.Fatal(err)
	}

	err = l.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
