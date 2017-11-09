// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"github.com/pkg/errors"
)

type NamespaceAndId struct {
	Namespace string
	Id        string
}

// Assets is a collection of Asset objects. It allows quick lookups by
// namespace and ID, as well as enumeration of all assets. In addition to
// keeping track of assets, an Assets collection also creates a DependencyNode
// for each managed asset.
type Assets struct {
	All map[NamespaceAndId]*DependencyNode
}

// Add is the only way an Asset should be added to an Assets collection.
func (A *Assets) Add(a Asset) error {
	if A.All == nil {
		A.All = make(map[NamespaceAndId]*DependencyNode)
	}

	nid := NamespaceAndId{a.Namespace(), a.Id()}
	if _, ok := A.All[nid]; ok {
		return errors.Errorf("asset named \"%s\" already exists", a.FullName())
	}
	A.All[nid] = &DependencyNode{Asset: a}

	return nil
}

// Get gets an asset based on its namespace and ID. Returns nil if the asset
// doesn't exist.
func (A *Assets) Get(namespace, id string) Asset {
	if A.All == nil {
		return nil
	}
	if d, ok := A.All[NamespaceAndId{namespace, id}]; ok {
		return d.Asset
	}
	return nil
}

// GetNodeForAsset returns the DependencyNode corresponding to an asset.
func (A *Assets) GetNodeForAsset(a Asset) *DependencyNode {
	if A.All == nil {
		return nil
	}
	if d, ok := A.All[NamespaceAndId{a.Namespace(), a.Id()}]; ok {
		return d
	}
	return nil
}
