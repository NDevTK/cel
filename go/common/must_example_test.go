// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common_test

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"os"
)

func ExampleMust_basic() {
	f := common.Must(os.Open("foo")).(*os.File)
	defer f.Close()
	// ...
	// ...
}
