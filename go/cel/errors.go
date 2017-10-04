// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cel

import (
	"fmt"
)

// Convenince function for generating arbitrary errors.
func NewError(format string, v ...interface{}) error {
	return fmt.Errorf(format, v...)
}

// DependencyCycleError is returned when there is a cycle in the dependency graph.
type DependencyCycleError struct {
	// List of DependencyNodes that form a cycle. This is not necessarily the
	// only cycle in the graph, nor is there anything special about this
	// particular cycle.
	Cycle []*DependencyNode
}

func (d *DependencyCycleError) Error() string {
	cycle_string := ""
	for _, node := range d.Cycle {
		if len(cycle_string) != 0 {
			cycle_string += " -> "
		}
		cycle_string += node.Asset.FullName()
	}
	cycle_string += fmt.Sprintf(" -> %s", d.Cycle[0].Asset.FullName())
	return "dependency cycle between assets : " + cycle_string
}

func NewDependencyCycleError(cycle []*DependencyNode) *DependencyCycleError {
	if len(cycle) == 0 {
		return nil
	}

	var cycle_copy []*DependencyNode
	copy(cycle_copy, cycle)
	return &DependencyCycleError{cycle_copy}
}

// InvalidDependencyError is returned when an asset has an illegal dependency
// on another asset. See AllowedToDependOn() for more details on allowed
// dependencies.
type InvalidDependencyError struct {
	// From asset depends on To asset.
	From Asset

	// From asset depends on To asset.
	To Asset
}

func (d *InvalidDependencyError) Error() string {
	if _, ok := d.From.(PermanentAsset); ok {
		return fmt.Sprintf("permanent asset %s not allowed to depend on non-permanent asset %s",
			d.From.FullName(), d.To.FullName())
	}

	if _, ok := d.From.(ResolvableAsset); ok {
		return fmt.Sprintf("resolvable asset %s not allowed to depend on script asset %s",
			d.From.FullName(), d.To.FullName())
	}
	return ""
}

func NewInvalidDependencyError(from, to Asset) *InvalidDependencyError {
	return &InvalidDependencyError{from, to}
}

// SkippedResolutionError is returned during asset resolution when an asset was
// not resolved due to there being unreolved (and often unresolvable)
// dependencies. Such is the result of one asset failing to resolve due to some
// reason, and consequently all assets that depend on it are not resolved.
type SkippedResolutionError struct {
	// Node on which error was observed.
	Node *DependencyNode

	// Unresolved dependency count.
	UnresolvedDependencies int32
}

func (e *SkippedResolutionError) Error() string {
	return fmt.Sprintf("asset %s was not resolved and still has %d unresolved dependencies",
		e.Node.Asset.FullName(), e.UnresolvedDependencies)
}

func NewSkippedResolutionError(node *DependencyNode) *SkippedResolutionError {
	return &SkippedResolutionError{node, node.UnresolvedDependencies}
}

// UnknownAssetError is returned when an attempt was made to use an asset that
// was not added via the relevant Assets.Add() method. I.e. the asset was not
// found in the asset catalog.
type UnknownAssetError struct {
	A Asset
}

func (e *UnknownAssetError) Error() string {
	return fmt.Sprintf("asset %s was not found in the global asset catalog", e.A.FullName())
}

func NewUnknownAssetError(a Asset) *UnknownAssetError {
	return &UnknownAssetError{a}
}
