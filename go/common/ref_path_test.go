// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"testing"
)

func TestRefPath_Basic(t *testing.T) {
	r := RefPathFromStrings("a", "b")
	if len(r) != 2 {
		t.Errorf("RefPathFromStrings() returned a path with length %d. Want 2.", len(r))
	}

	r = r.Append("c")
	if len(r) != 3 {
		t.Errorf("RefPath.Append() returned a path with length %d. Want 3", len(r))
	}

	if r.String() != "a.b.c" {
		t.Errorf("RefPath.String() returned %s. Want \"a.b.c\"", r.String())
	}

	if !r.Equals(RefPathFromStrings("a", "b", "c")) {
		t.Error("RefPath.Equals() failed")
	}

	r2 := RefPathFromStrings("a", "b.c", "d")
	if r2.String() != "a.(b.c).d" {
		t.Errorf("RefPath.String() returned %v. Want \"a.(b.c).d\"", r2.String())
	}
}

// It's safe to call Append() multiple times on the same RefPath.
func TestRefPath_Append_Multiple(t *testing.T) {
	rr := RefPathFromStrings("a", "b")
	r1 := rr.Append("c")
	r2 := rr.Append("d")

	if r1.String() != "a.b.c" {
		t.Errorf("Append() returned %v. Want \"a.b.c\"", r1.String())
	}
	if r2.String() != "a.b.d" {
		t.Errorf("Append() returned %v. Want \"a.b.d\"", r2.String())
	}
}

func TestRefPath_FromString(t *testing.T) {
	r, err := RefPathFromString("a.b.c.d")
	if err != nil {
		t.Error(err)
	}
	want := RefPath{"a", "b", "c", "d"}
	if !want.Equals(r) {
		t.Errorf("RefPathFromString() returned %#v. Want %#v", r, want)
	}

	r, err = RefPathFromString("a.(b.c).d")
	if err != nil {
		t.Error(err)
	}
	want = RefPath{"a", "b.c", "d"}
	if !want.Equals(r) {
		t.Errorf("RefPathFromString() returned %#v. Want %#v", r, want)
	}

	r, err = RefPathFromString("a.(b")
	if err == nil {
		t.Error("RefPath string with mismatched parens successfuly parsed")
	}

	r, err = RefPathFromString("a.b)")
	if err != nil {
		t.Error(err)
	}
	want = RefPath{"a", "b)"}
	if !want.Equals(r) {
		t.Errorf("RefPathFromString() returned %#v. Want %#v", r, want)
	}

	r, err = RefPathFromString("a.((b.c.d.e).f)")
	if err != nil {
		t.Error(err)
	}
	want = RefPath{"a", "(b.c.d.e", "f)"}
	if !want.Equals(r) {
		t.Errorf("RefPathFromString() returned %#v. Want %#v", r, want)
	}

	r, err = RefPathFromString("")
	if err != nil {
		t.Error(err)
	}
	if len(r) != 0 {
		t.Error("RefPathFromString() returned %#v for empty string.", r)
	}
}

func TestRefPath_Equals(t *testing.T) {
	r1 := RefPath{}
	r2 := RefPath{}
	if !r1.Equals(r2) {
		t.Errorf("Equals() failed for empty paths %#v and %#v", r1, r2)
	}

	if !r1.Equals(r1) {
		t.Errorf("Equals() failed on self for %#v", r1)
	}

	r1 = RefPathFromStrings("a", "b", "c")
	r2 = RefPathFromStrings("a", "b", "c")
	if !r1.Equals(r2) {
		t.Errorf("Equals() failed for %#v and %#v", r1, r2)
	}
}

func TestRefPath_Contains(t *testing.T) {
	cases := []struct {
		base, contains string
		expected       bool
	}{
		{"", "", true},
		{"a.b.c", "a.b.c", true},
		{"a.b.c", "a.b.c.d", true},
		{"a.b.c", "a.b", false},
		{"a.b.c", "a.c", false},
	}

	for _, c := range cases {
		base, err := RefPathFromString(c.base)
		if err != nil {
			t.Error(err)
		}

		contains, err := RefPathFromString(c.contains)
		if err != nil {
			t.Error(err)
		}

		if base.Contains(contains) != c.expected {
			t.Errorf("(%#v).Contains(%#v) failed. Expected %v", base, contains, c.expected)
		}
	}
}

func TestRefPath_After(t *testing.T) {
	a, ok := RefPathFromStrings("a", "b", "c").After(RefPathFromStrings("a", "b"))
	if !ok {
		t.Error("RefPathFromString().After() failed")
	}
	if !a.Equals(RefPathFromStrings("c")) {
		t.Errorf("RefPathFromString().After() returned %#+v. Wanted \"c\"", a)
	}
}
