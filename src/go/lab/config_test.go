// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package lab

import (
	"regexp"
	"testing"
)

func expectMatch(m *regexp.Regexp, s string, t *testing.T) {
	if !m.MatchString(s) {
		t.Errorf("%s should match '%s'", m.String(), s)
	}
}

func expectNoMatch(m *regexp.Regexp, s string, t *testing.T) {
	if m.MatchString(s) {
		t.Errorf("%s should not match '%s'", m.String(), s)
	}
}

func TestLoadConfigFile(t *testing.T) {
	config, err := LoadConfigFiles("testdata/basic-host.textpb", []string{"testdata/basic-assets.textpb"})
	if err != nil {
		t.Fatal("Can't load data file.", err)
	}

	if config.GetProject() != "test:google.com:chrome-auth-lab" {
		t.Fatal("Project: ", config.GetProject())
	}

	machines := config.Instance
	if len(machines) != 1 {
		t.Fatal("Machines: ", len(config.Instance))
	}

	machine := config.Instance[0]
	if machine.GetName() != "a" {
		t.Fatal("Machine name:", machine.GetName())
	}
}

func TestMatcher_Passthrough(t *testing.T) {
	m := matcher("/foo/bar")
	expectMatch(m, "/foo/bar", t)
	expectNoMatch(m, "foo/bar", t)
}

func TestMatcher_Infix(t *testing.T) {
	m := matcher("a/*/b")
	expectMatch(m, "a/c/b", t)
	expectMatch(m, "a/asdfasdfasdf23423asdf/b", t)
	expectMatch(m, "a/example.com:foo/b", t)
	expectNoMatch(m, "a//b", t)
	expectNoMatch(m, "a/0/b", t)
}

func TestMatcher_Postfix(t *testing.T) {
	m := matcher("a/*/b/*")
	expectMatch(m, "a/b/b/c", t)
	expectNoMatch(m, "a/b/b/c/", t)
	expectMatch(m, "a/b/b/asdfsdfdsf", t)
}

func TestMatcher_Label(t *testing.T) {
	m := matcher("*")
	expectMatch(m, "asdf", t)
	expectNoMatch(m, "asdf/basdf", t)
	expectNoMatch(m, "A", t)
}
