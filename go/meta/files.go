// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package meta

import (
	"sort"
)

// Canonicalize ensures that the encoding of the Tree object is stable by
// ordering the file references it contains.
func (t *Tree) Canonicalize() {
	if len(t.File) == 0 {
		return
	}
	sort.Slice(t.File, func(i, j int) bool {
		return t.File[i].Name < t.File[j].Name
	})
}
