// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cel

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"github.com/pkg/errors"
	"sync"
)

// kMaxParallelOperations is the maximum number of resolver jobs that are going
// to be allowed to be active at the same time. There's no host specific reason
// for the limit. It's just there to throttle API requests that are going out
// to GCE.
const kMaxParallelOperations = 10

// checkForCycles returns the largest cycle found in graph A or nil if A is
// acyclic. Should only be called after the depedents of each node has been
// identified.
//
// The returned list of DependencyNode is ordered such that node at position 𝒊
// depends on node at position 𝒊+1, and finally the last node in the list
// depends on the first.
func checkForCycles(A *Assets) []*DependencyNode {
	// We are going to use Tarjan's Strongly Connected Components algorithm and
	// look for SCCs of order larger than 1.
	//
	// https://en.wikipedia.org/wiki/Tarjan%27s_strongly_connected_components_algorithm

	type vertex struct {
		index    int
		lowlink  int
		on_stack bool
	}

	type state struct {
		vertices map[*DependencyNode]*vertex
		stack    []*DependencyNode
		index    int
		cycle    []*DependencyNode
	}

	var strong_connect func(s *state, n *DependencyNode)
	strong_connect = func(s *state, node *DependencyNode) {
		v, _ := s.vertices[node]
		v.index = s.index
		v.lowlink = s.index
		s.index += 1
		s.stack = append(s.stack, node)
		v.on_stack = true

		for _, wnode := range node.Dependents {
			w, _ := s.vertices[wnode]
			if w.index == -1 {
				strong_connect(s, wnode)
				v.lowlink = min(v.lowlink, w.lowlink)
			} else if w.on_stack {
				v.lowlink = min(v.lowlink, w.index)
			}
		}

		// Not a root node. Leave it on the stack.
		if v.lowlink != v.index {
			return
		}

		for i := len(s.stack) - 1; i >= 0; i -= 1 {
			s.vertices[s.stack[i]].on_stack = false
			if s.stack[i] != node {
				continue
			}

			cycle := make([]*DependencyNode, len(s.stack)-i)
			copy(cycle, s.stack[i:])
			s.stack = s.stack[:i]

			// Cycles exist only if we find a component of size greater than 1.
			// If one is found, only keep it around if it is the largest we've
			// seen so far.
			if len(cycle) > 1 && len(cycle) > len(s.cycle) {
				s.cycle = cycle
			}
			break
		}
	}

	s := state{vertices: make(map[*DependencyNode]*vertex)}
	for _, node := range A.All {
		s.vertices[node] = &vertex{index: -1, lowlink: -1, on_stack: false}
	}

	for _, node := range A.All {
		v, _ := s.vertices[node]
		if v.index == -1 {
			strong_connect(&s, node)
		}

		if len(s.cycle) > 1 {
			return s.cycle
		}
	}
	return nil
}

// AllowedToDependOn determines whether |child| is allowed to depend on
// |parent|. If so, returns a nil error. Otherwise returns an error that
// explains why |child| isn't allowed to depend on |parent|.
//
// In this |child| ⇔ |parent| relationship, |parent| would be an element of
// |child.DependsOn()|.
//
// The dependencies that are allowed between different type of assets is as
// follows:
//                    ┌─┐                         ┌──┐
//                    │ ↓                         │  ↓
//               ╔════╧══════╗               ╔════╧═══════╗
//               ║           ║               ║            ║
//               ║ Permanent ║←──────────────╢ Resolvable ║
//               ║           ║               ║            ║
//               ╚═══════════╝               ╚════════════╝
//                     ↑            ┌─┐            ↑
//                     │            │ ↓            │
//                     │       ╔════╧═════╗        │
//                     │       ║          ║        │
//                     └───────╢  Script  ╟────────┘
//                             ║          ║
//                             ╚══════════╝
// I.e.:
//   * A Permanent asset must only depend on another permanent asset.
//   * A Resolvable asset must only depend on a permanent or resolvable asset.
//   * A Script asset can depend on any other type of asset.
func AllowedToDependOn(child Asset, parent Asset) error {
	if _, child_is_dsc := child.(ScriptAsset); child_is_dsc {
		return nil
	}

	if _, child_is_permanent := child.(PermanentAsset); child_is_permanent {
		if _, parent_is_permanent := parent.(PermanentAsset); parent_is_permanent {
			return nil
		}
		return NewInvalidDependencyError(parent, child)
	}

	if _, parent_is_script := parent.(ScriptAsset); parent_is_script {
		return NewInvalidDependencyError(parent, child)
	}
	return nil
}

// PrepareToResolve resets the state of the dependency nodes in the asset set
// |A|. In preparation, this function:
//
// * Resets the Processed state of all the nodes.
// * Calculates the dependency graph based on the Asset.DependsOn() results.
//
// The function then returns the list of root nodes that can be used to start
// the resolution process.
func PrepareToResolve(A *Assets) (roots []*DependencyNode, err error) {
	for _, node := range A.All {
		node.Dependents = nil
		node.Processed = false
	}

	for _, node := range A.All {
		depends_on := node.Asset.DependsOn()
		node.UnresolvedDependencies = int32(len(depends_on))
		for _, parent_asset := range depends_on {
			err = AllowedToDependOn(node.Asset, parent_asset)
			if err != nil {
				return nil, err
			}

			parent := A.GetNodeForAsset(parent_asset)
			if parent == nil {
				return nil, NewUnknownAssetError(parent_asset)
			}
			parent.Dependents = append(parent.Dependents, node)
		}
	}

	cycle := checkForCycles(A)
	if len(cycle) != 0 {
		return nil, NewDependencyCycleError(cycle)
	}

	for _, node := range A.All {
		if node.Ready() {
			roots = append(roots, node)
		}
	}

	return
}

// ResolveAssets begins the process of resolving the assets in |A| using the
// session information in |S|.
//
// The resolution will parallelize requests up to a factor of
// kMaxParallelOperations. The limit is there mostly to prevent DDoSing the
// Google Cloud infrastructure.
//
// If any of the resolution steps fail, then the returned error will indicate
// what went wrong. The returned error object will aggregate all the errors
// that occurred during the resolution process.
//
// It is not necessary to call PrepareToResolve() prior to calling this
// function since it will call PrepareToResolve() on its own.
func ResolveAssets(A *Assets) error {
	roots, err := PrepareToResolve(A)
	if err != nil {
		return err
	}

	// The maximum number of workers we are going to allow.
	W := min(kMaxParallelOperations, len(A.All))

	// resolvable is a channel where all nodes that are ready to resolve will
	// be pushed through. When a node is resolved, it may result in a new set
	// of nodes that are now ready to be resolved. These nodes will be enqueued
	// to |resolvable| via dispatchBatch() below.
	resolvable := make(chan *DependencyNode, W)

	// wg will be incremented *before* any nodes are enqueued to resolvable,
	// and decremented *after* the node is resolved -- successfully or
	// otherwise. Every node that is enqueued here *must* be resolved one way
	// or another. The resolver is done when the counter reaches zero.
	wg := sync.WaitGroup{}

	// Here we kick off our worker pool. The pool is here only to limit the
	// number of possible outstanding network requests.
	for i := 0; i < W; i += 1 {
		go resolveNodesFromChannel(resolvable, &wg)
	}

	// Kick of the first batch.
	wg.Add(len(roots))
	go dispatchBatch(roots, resolvable)

	wg.Wait()
	// Closing |resolvable| ends the |resolveNodesFromChannel| goroutines.
	close(resolvable)

	// Once we get here, there are no more pending tasks. Any nodes left
	// unresolved didn't have their dependencies resolved. This is to be
	// expected if upstream dependent nodes failed to resolve for some reason.
	error_list := []error{}
	for _, d := range A.All {
		if !d.Processed {
			// Not really an error.
			error_list = append(error_list, NewSkippedResolutionError(d))
			continue
		}

		if d.Result != nil {
			error_list = append(error_list, errors.Wrapf(d.Result,
				"asset %s failed to resolve", d.Asset.FullName()))
		}
	}
	return common.WrapErrorList(error_list)
}

// resolveNodesFromChannel fetches DependencyNodes from |resolvable|, resolves
// them, and if successful, enqueues a batch of newly resolvable
// DependencyNodes back to |resolvable|.
//
// Terminates when the |resolvable| channel is closed.
func resolveNodesFromChannel(resolvable chan *DependencyNode, wg *sync.WaitGroup) {
	for d := range resolvable {
		next, _ := d.Resolve()

		if len(next) > 0 {
			wg.Add(len(next))
			go dispatchBatch(next, resolvable)
		}

		wg.Done()
	}
}

// dispatchBatch enqueues a batch of |DependencyNode|s to |resolvable|. Should
// be called from a goroutine in order to not block resolution.
func dispatchBatch(block []*DependencyNode, resolvable chan *DependencyNode) {
	for _, n := range block {
		resolvable <- n
	}
}

// min returns the minimum of two integers |a| and |b|. It exists because golang.
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
