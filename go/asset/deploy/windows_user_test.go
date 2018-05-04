// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"testing"
)

func TestWindowsUser_GeneratePassword_length(t *testing.T) {
	p, err := generatePassword()
	if err != nil {
		t.Fatal(err)
	}

	if len(p) != 32 {
		t.Errorf("unexpected length for password. Got %d. Want 32", len(p))
	}
}
