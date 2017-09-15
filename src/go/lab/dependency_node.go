// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package lab

import (
	"sync/atomic"
)

type DependencyNode struct {
	Asset Asset // The Asset contained in this node.

	UnresolvedDependencies int32 // Volatile. Final value is zero. int32 because it can be incremented atomically.
	Processed              bool  // Flips once from false -> true. Final value is true.
	Result                 error // Only valid if Resolved == true. Final value when Resolved == true.

	// Other dependecy nodes that depend on this. I.e. this node must be
	// resolved before Dependents can be resolved.
	Dependents []*DependencyNode
}

// DependencyResolved is called when one dependency of this node has been
// resolved. Once all dependencies are resolved (i.e.
// UnresolvedDependencies==0), then this node can also resolve.
//
// Returns true if all outstanding dependencies have been resolved. Subsequent
// calls to d.Ready() will return true.
func (d *DependencyNode) DependencyResolved() bool {
	r := atomic.AddInt32(&d.UnresolvedDependencies, -1)
	if r < 0 {
		panic("dependency count underrun")
	}
	return r == 0
}

// Ready returns true if this node is ready to be resolved. I.e. this node is
// not waiting for anything. Note that this function will continue to return
// true even after |d.Asset| is resolved.
func (d *DependencyNode) Ready() bool {
	return d.UnresolvedDependencies == 0
}

// Resolve resolves the asset. It's an error to call this method more than
// once, or call it before Ready() returns true. If the resolution was
// successful, it also returns a list of DependencyNodes that are now ready to
// be resolved.
func (d *DependencyNode) Resolve(S *Session) (ready []*DependencyNode, err error) {
	if !d.Ready() {
		return nil, NewError("asset %s has %d unmet dependencies, but Resolve() was called to resolve it.",
			d.Asset.FullName(),
			d.UnresolvedDependencies)
	}

	if d.Processed {
		return nil, NewError("asset %s was already resolved.", d.Asset.FullName())
	}

	switch a := d.Asset.(type) {
	case ResolvableAsset:
		d.Result = a.Resolve(S)

	case PermanentAsset:
		d.Result = a.Check(S)

	case ScriptAsset:
		d.Result = a.GenerateScript(S)

	default:
		return nil, NewError("unknown asset type for %s", d.Asset.FullName())
	}

	d.Processed = true

	if d.Result != nil {
		return nil, d.Result
	}

	for _, downstream := range d.Dependents {
		if downstream.DependencyResolved() {
			ready = append(ready, downstream)
		}
	}
	return
}
