// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"encoding/json"
	"testing"
)

func TestTrie_MarshalJSON_type(t *testing.T) {
	var trie interface{} = &Trie{}
	if _, ok := trie.(json.Marshaler); !ok {
		t.Errorf("Trie type does not implement json.Marshaler")
	}
}

func TestTrie_MarshalJSON_trivial(t *testing.T) {
	var trie Trie

	if !trie.Set(EmptyPath, "hello", true) {
		t.Fatal("failed to set value")
	}

	e, err := json.Marshal(&trie)
	if err != nil {
		t.Fatal(err)
	}

	s := string(e)
	if s != "\"hello\"" {
		t.Fatalf("unexpected encoding: %s", s)
	}

	trie.Set(EmptyPath, 1, true)
	e, err = json.Marshal(&trie)
	s = string(e)
	if s != `1` {
		t.Fatalf("unexpected encoding: %s", s)
	}
}

func TestTrie_MarshalJSON_nested(t *testing.T) {
	var trie Trie
	trie.Set(RefPathMust("a.b.c"), "hello", true)
	trie.Set(RefPathMust("a.b.d"), "world", true)
	e, err := json.Marshal(&trie)
	if err != nil {
		t.Fatal(err)
	}
	s := string(e)
	if s != `{"a":{"b":{"c":"hello","d":"world"}}}` {
		t.Fatalf("unexpected encoding: %s", s)
	}
}

func TestTrie_MarshalJSON_value_overrides_subtree(t *testing.T) {
	var trie Trie
	trie.Set(RefPathMust("a.b.c"), "hello", true)
	trie.Set(RefPathMust("a.b.d"), "world", true)
	trie.Set(RefPathMust("a.b"), "the buck stops here", true)
	e, err := json.Marshal(&trie)
	if err != nil {
		t.Fatal(err)
	}
	s := string(e)
	if s != `{"a":{"b":"the buck stops here"}}` {
		t.Fatalf("unexpected encoding: %s", s)
	}
}
