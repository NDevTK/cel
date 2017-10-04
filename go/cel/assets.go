// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cel

import ()

// Assets is a collection of Asset objects. It allows quick lookups by
// namespace and ID, as well as enumeration of all assets. In addition to
// keeping track of assets, an Assets collection also creates a DependencyNode
// for each managed asset.
type Assets struct {
	ByNamespaceAndId map[string]map[string]*DependencyNode
	All              []*DependencyNode
}

// Add is the only way an Asset should be added to an Assets collection.
func (A *Assets) Add(a Asset) error {
	ns := a.Namespace()

	if A.ByNamespaceAndId == nil {
		A.ByNamespaceAndId = make(map[string]map[string]*DependencyNode)
	}

	t, _ := A.ByNamespaceAndId[ns]
	if t == nil {
		t = make(map[string]*DependencyNode)
		A.ByNamespaceAndId[ns] = t
	}

	_, ok := t[a.Id()]
	if ok {
		return NewError("asset %s of type %s already exists", a.Id(), ns)
	}

	n := &DependencyNode{Asset: a}
	t[a.Id()] = n
	A.All = append(A.All, n)
	return nil
}

// Get gets an asset based on its namespace and ID. Returns nil if the asset
// doesn't exist.
func (A *Assets) Get(namespace, id string) Asset {
	if A.ByNamespaceAndId == nil {
		return nil
	}

	t, _ := A.ByNamespaceAndId[namespace]
	if t == nil {
		return nil
	}

	d, _ := t[id]
	if d == nil {
		return nil
	}

	return t[id].Asset
}

// GetNodeForAsset returns the DependencyNode corresponding to an asset.
func (A *Assets) GetNodeForAsset(a Asset) *DependencyNode {
	if t, ok := A.ByNamespaceAndId[a.Namespace()]; ok {
		if d, ok := t[a.Id()]; ok {
			return d
		}
	}
	return nil
}
