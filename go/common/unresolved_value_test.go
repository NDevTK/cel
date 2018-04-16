// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"testing"
)

func TestRefListToString(t *testing.T) {
	testCases := []struct {
		input    []RefPath
		expected string
	}{
		{nil, "(empty)"},
		{[]RefPath{}, "(empty)"},
		{[]RefPath{RefPathMust("a.b.c")}, "\"a.b.c\""},
		{[]RefPath{RefPathMust("a.b.c"), RefPathMust("c.d.e")}, "\"a.b.c\", \"c.d.e\""},
	}

	for _, test := range testCases {
		actual := refListToString(test.input)
		if actual != test.expected {
			t.Errorf("Got \"%s\". Want \"%s\"", actual, test.expected)
		}
	}
}
