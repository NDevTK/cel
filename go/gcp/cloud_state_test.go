// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

import (
	"golang.org/x/net/context"
	"net/http"
	"testing"
)

func TestQueryGceState(t *testing.T) {
	c := http.Client{Transport: &RequestReplayer{DataPath: "testdata/cloud_state_test"}}
	_, err := QueryCloudState(context.Background(), &c, "chrome-auth-lab-dev", nil)
	if err != nil {
		t.Fatal(err)
	}
}
