// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"strings"
	"testing"
)

func TestIsRFC1035Label(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		if err := IsRFC1035Label(""); err == nil || !strings.Contains(err.Error(), "Valid labels can't be empty.") {
			t.Fatal("unexpected error", err)
		}
	})

	t.Run("too long", func(t *testing.T) {
		t.Parallel()

		if err := IsRFC1035Label("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"); err == nil || !strings.Contains(err.Error(), "can't be longer than 63 characters") {
			t.Fatal("unexpected error", err)
		}
	})

	t.Run("starts with a digit", func(t *testing.T) {
		t.Parallel()

		if err := IsRFC1035Label("0abc"); err == nil || !strings.Contains(err.Error(), "character '0' at position 1 is not valid") {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("has an invalid character", func(t *testing.T) {
		t.Parallel()

		if err := IsRFC1035Label("ab$c"); err == nil || !strings.Contains(err.Error(), "character '$' at position 3 is not valid") {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("starts with a dash", func(t *testing.T) {
		t.Parallel()

		if err := IsRFC1035Label("-abc"); err == nil || !strings.Contains(err.Error(), "character '-' at position 1 is not valid") {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("ends with a dash", func(t *testing.T) {
		t.Parallel()

		if err := IsRFC1035Label("abc-"); err == nil || !strings.Contains(err.Error(), "Cannot end with a '-'") {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("ends with a digit", func(t *testing.T) {
		t.Parallel()

		if err := IsRFC1035Label("abc0"); err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}
	})
}
