// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common_test

import (
	"fmt"

	"chromium.googlesource.com/enterprise/cel/go/common"
)

// Example of basic Trie usage.
func ExampleTrie_basic() {
	var t common.Trie
	ref, _ := common.RefPathFromString("a.b.c")
	t.Set(ref, "hello", false)
	v, _ := t.Get(ref)
	fmt.Printf("Value at %s is %s", ref.String(), v.(string))
	// Output: Value at a.b.c is hello
}

// Example of visiting a subtree.
func ExampleTrie_visit() {
	var t common.Trie
	t.Set(common.RefPathMust("a.b.c"), "hello", false)
	t.Set(common.RefPathMust("a.b.d"), "world", false)
	t.Set(common.RefPathMust("x.y.z"), "not visited", false)
	t.VisitFrom(common.RefPathMust("a.b"), func(p common.RefPath, v interface{}) bool {
		fmt.Printf("Value at %s is %s\n", p, v)
		return true
	})
	// Unordered output:
	// Value at a.b.c is hello
	// Value at a.b.d is world
}
