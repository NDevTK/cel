// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"testing"
)

func TestTrie_SetGetUnset(t *testing.T) {
	var trie Trie
	p := RefPathFromString("a.b.c")
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

	if trie.Get(p).(string) != "hello" {
		t.Fatal()
	}

	trie.Unset(p)
	if trie.Get(p) != nil {
		t.Fatal()
	}

	if trie.Size() != 0 {
		t.Fatal()
	}

}

func TestTrie_Visit(t *testing.T) {
	var trie Trie
	trie.Set(RefPathFromString("a.b.c"), "foo", true)
	trie.Set(RefPathFromString("a.b.d"), "bar", true)
	trie.Set(RefPathFromString("x.y.z"), "baz", true)

	if trie.Size() != 3 {
		t.Fatal()
	}

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

	seen = make(map[string]bool)
	trie.VisitFrom(RefPathFromString("a.b"), func(p RefPath, o interface{}) bool {
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

func TestTrie_SetUnset_Load(t *testing.T) {
	var trie Trie
	trie.Set(RefPathFromString("a.b.c"), "foo", true)
	trie.Set(RefPathFromString("a.b.d"), "bar", true)
	trie.Set(RefPathFromString("x.y.z"), "baz", true)

	if len(trie.m) != 2 {
		t.Fatal()
	}

	if len(trie.m["a"].m) != 1 {
		t.Fatal()
	}

	if len(trie.m["a"].m["b"].m) != 2 {
		t.Fatal()
	}

	trie.Unset(RefPathFromString("a.b"))
	if len(trie.m) != 2 {
		t.Fatal()
	}
	trie.Unset(RefPathFromString("a.b.c"))
	if len(trie.m["a"].m["b"].m) != 1 {
		t.Fatal()
	}

	trie.Unset(RefPathFromString("a.b.d"))
	if len(trie.m) != 1 {
		t.Fatal()
	}
}
