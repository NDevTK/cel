// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"testing"
)

func TestRefPath_Basic(t *testing.T) {
	r := RefPathFromList("a", "b")
	if len(r) != 2 {
		t.Fail()
	}

	r = r.Append("c")
	if len(r) != 3 {
		t.Fail()
	}

	if r.String() != "a.b.c" {
		t.Fail()
	}
}

func TestRefPath_Append(t *testing.T) {
	rr := RefPathFromList("a", "b")
	r1 := rr.Append("c")
	r2 := rr.Append("d")

	if r1.String() != "a.b.c" {
		t.Fail()
	}
	if r2.String() != "a.b.d" {
		t.Fail()
	}

	r3 := rr.Append("e", "f", "g")
	if len(r3) != 5 {
		t.Fail()
	}
	if r3[3] != "f" {
		t.Fail()
	}
}

func TestRefPath_FromString(t *testing.T) {
	r := RefPathFromString("a.b.c.d")
	if len(r) != 4 || r[2] != "c" {
		t.Fatal(r)
	}

	r = RefPathFromString("a.(b.c).d")
	if len(r) != 3 {
		t.Fatal(r)
	}

	if r[1] != "b.c" {
		t.Fatal()
	}

	r = RefPathFromString("a.(b")
	if len(r) != 0 {
		t.Fatal()
	}

	r = RefPathFromString("a.b)")
	if len(r) != 2 {
		t.Fatal()
	}

	r = RefPathFromString("a.((b.c.d.e).f)")
	if len(r) != 3 {
		t.Fatal(r)
	}

	r = RefPathFromString("a.(b).c")
	if len(r) != 3 || r[1] != "b" {
		t.Fatal()
	}

	r = RefPathFromString("")
	if len(r) != 0 {
		t.Fatal()
	}
}

func TestRefPath_Equals(t *testing.T) {
	r1 := RefPath{}
	r2 := RefPath{}
	if !r1.Equals(r2) {
		t.Fail()
	}

	if !r1.Equals(r1) {
		t.Fail()
	}

	r1 = RefPathFromString("a.b.c")
	r2 = RefPathFromString("a.b.c")
	if !r1.Equals(r2) {
		t.Fail()
	}
	r1 = r1.Append("b")
	if r1.Equals(r2) {
		t.Fail()
	}
}

func TestRefPath_Contains(t *testing.T) {
	if !RefPathFromString("").Contains(RefPathFromString("")) {
		t.Fail()
	}

	if !RefPathFromString("a.b.c").Contains(RefPathFromString("a.b.c")) {
		t.Fail()
	}

	if !RefPathFromString("a.b.c").Contains(RefPathFromString("a.b.c.d")) {
		t.Fail()
	}

	if RefPathFromString("a.b.c").Contains(RefPathFromString("a.b")) {
		t.Fail()
	}

	if RefPathFromString("a.b.c").Contains(RefPathFromString("a.c")) {
		t.Fail()
	}
}

func TestRefPath_After(t *testing.T) {
	if a, ok := RefPathFromString("a.b.c").After(RefPathFromString("a.b")); !ok || !a.Equals(RefPathFromString("c")) {
		t.Fail()
	}
}
