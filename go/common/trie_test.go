// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"testing"
)

func TestTrie_Set_basic(t *testing.T) {
	var trie Trie
	p := RefPathFromComponents("a", "b", "c")
	if !trie.Set(p, "oh", false) {
		t.Fatal()
	}

	if trie.Set(p, "bye", false) {
		t.Fatal()
	}

	if !trie.Set(p, "hello", true) {
		t.Fatal()
	}

	if trie.Size() != 1 {
		t.Fatal()
	}

	if v, ok := trie.Get(p); !ok || v.(string) != "hello" {
		t.Fatal()
	}

	if !trie.Has(p) {
		t.Fatal()
	}

	trie.Unset(p)
	if v, ok := trie.Get(p); ok || v != nil {
		t.Fatal()
	}

	if trie.Has(p) {
		t.Fatal()
	}

	if trie.Size() != 0 {
		t.Fatal()
	}
}

func TestTrie_Empty(t *testing.T) {
	var trie Trie
	if !trie.Empty() {
		t.Fatal("Empty() failed for emtpy trie")
	}

	trie.Set(RefPathFromComponents("a", "b", "c"), "hello", false)
	if trie.Empty() {
		t.Fatal("Empty() returned true for non-empty trie")
	}
}

func TestTrie_GetDeepest(t *testing.T) {
	var trie Trie
	trie.Set(RefPathFromComponents("a", "b"), "hello", false)
	o, p := trie.GetDeepest(RefPathFromComponents("a", "b", "c"))
	if o == nil || o.(string) != "hello" {
		t.Fatal()
	}

	o, p = trie.GetDeepest(RefPathFromComponents("x"))
	if !p.Equals(RefPathFromComponents("x")) {
		t.Fatalf("wrong remaining path returned by GetClosest(). Got %s. Want x", p)
	}

	if o != nil {
		t.Fatalf("wrong object returned by GetClosest(). Got %#v. Want nil", o)
	}

	trie.Set(EmptyPath, "rootValue", false)
	o, p = trie.GetDeepest(RefPathFromComponents("x"))
	if o == nil || o.(string) != "rootValue" {
		t.Fatalf("wrong object returned by GetClosest(). Got %#v. Want \"rootValue\"", o)
	}
}

func TestTrie_Visit(t *testing.T) {
	var trie Trie
	trie.Set(RefPathFromComponents("a", "b", "c"), "foo", true)
	trie.Set(RefPathFromComponents("a", "b", "d"), "bar", true)
	trie.Set(RefPathFromComponents("x", "y", "z"), "baz", true)

	seen := make(map[string]bool)
	trie.Visit(func(p RefPath, o interface{}) bool {
		seen[p.String()] = true
		if _, ok := o.(string); !ok {
			t.Fatal()
		}
		return true
	})

	if _, ok := seen["a.b.c"]; !ok {
		t.Fatal()
	}
	if _, ok := seen["a.b.d"]; !ok {
		t.Fatal()
	}
	if _, ok := seen["x.y.z"]; !ok {
		t.Fatal()
	}
	if len(seen) != 3 {
		t.Fatal()
	}
}

func TestTrie_VisitFrom(t *testing.T) {
	var trie Trie
	trie.Set(RefPathFromComponents("a", "b", "c"), "foo", true)
	trie.Set(RefPathFromComponents("a", "b", "d"), "bar", true)
	trie.Set(RefPathFromComponents("x", "y", "z"), "baz", true)

	seen := make(map[string]bool)
	trie.VisitFrom(RefPathFromComponents("a", "b"), func(p RefPath, o interface{}) bool {
		seen[p.String()] = true
		return true
	})
	if _, ok := seen["a.b.c"]; !ok {
		t.Fatal()
	}
	if _, ok := seen["a.b.d"]; !ok {
		t.Fatal()
	}
	if len(seen) != 2 {
		t.Fatal()
	}
}

func TestTrie_DescendUntil_existingPath(t *testing.T) {
	var trie Trie

	sPath := RefPathMust("a.b.c.d")
	rPath := sPath.Append("e", "f", "g")

	trie.Set(sPath, "S", true)
	trie.Set(rPath, "R", true)
	if trie.Size() != 2 {
		t.Fatal("incorrect number of elements in trie")
	}

	sSeen := 0
	rSeen := 0
	uSeen := 0

	if !trie.DescendUntil(rPath, func(p RefPath, o interface{}) bool {
		// the only nodes that we should see are these two. Nodes with no
		// values are skipped.
		if p.Equals(sPath) {
			sSeen++
		} else if p.Equals(rPath) {
			rSeen++
		} else {
			uSeen++
		}
		return true
	}) {
		t.Error("DescendUntil returned false")
	}

	if sSeen != 1 || rSeen != 1 || uSeen != 0 {
		t.Error("required targets not seen:", sSeen, rSeen, uSeen)
	}
}

func TestTrie_DescendUntil_nonExistentPath(t *testing.T) {
	var trie Trie

	sPath := RefPathMust("a.b.c.d")
	rPath := sPath.Append("e", "f", "g")

	trie.Set(sPath, "S", true)
	trie.Set(rPath, "R", true)

	sSeen := 0
	rSeen := 0
	uSeen := 0

	if trie.DescendUntil(rPath.Append("one-more"), func(p RefPath, o interface{}) bool {
		// the only nodes that we should see are these two. Nodes with no
		// values are skipped.
		if p.Equals(sPath) {
			sSeen++
		} else if p.Equals(rPath) {
			rSeen++
		} else {
			uSeen++
		}
		return true
	}) {
		t.Error("DescendUntil returned true")
	}

	if sSeen != 1 || rSeen != 1 || uSeen != 0 {
		t.Error("required targets not seen:", sSeen, rSeen, uSeen)
	}
}

// Introspective test. Verifies the internal representation is what I think it
// is.
func TestTrie_Set_introspect(t *testing.T) {
	var trie Trie
	trie.Set(RefPathFromComponents("a", "b", "c"), "foo", true)
	trie.Set(RefPathFromComponents("a", "b", "d"), "bar", true)
	trie.Set(RefPathFromComponents("x", "y", "z"), "baz", true)

	if len(trie.m) != 2 {
		t.Fatal()
	}

	if len(trie.m["a"].m) != 1 {
		t.Fatal()
	}

	if len(trie.m["a"].m["b"].m) != 2 {
		t.Fatal()
	}

	trie.Unset(RefPathFromComponents("a", "b"))
	if len(trie.m) != 2 {
		t.Fatal()
	}
	trie.Unset(RefPathFromComponents("a", "b", "c"))
	if len(trie.m["a"].m["b"].m) != 1 {
		t.Fatal()
	}

	trie.Unset(RefPathFromComponents("a", "b", "d"))
	if len(trie.m) != 1 {
		t.Fatal()
	}
}
