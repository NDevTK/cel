// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"encoding/json"
)

// MarshalJSON returns a json encoding of the trie object.
//
// The JSON serialization of a trie is done using the following method:
//
// If the trie node has a value, then the serialization of the trie is the
// serialization of the value.
//
// Otherwise, the trie serializes to a JSON map where the keys are the node
// names of the trie's children, and the value as the JSON serialization of
// each child.
//
// In CEL's use of trie, if an object is stored in a non-leaf node, the subtree
// rooted at that node must correspond to the fields of that object. In other
// words, CEL's use of Trie is to map names to structs that are grafted at
// various points in a namespace. Hence the serialization doesn't need to
// venture past a non-leaf value.
//
// Note: If it isn't otherwise unclear, Trie's MarshalJSON() is not suitable
// for general consumption. It is specifically coded around CEL's convenience
// and usage.
func (t *Trie) MarshalJSON() ([]byte, error) {
	if t.value != nil {
		if b, err := json.Marshal(t.value); err == nil {
			return b, nil
		}
	}

	return json.Marshal(t.m)
}
