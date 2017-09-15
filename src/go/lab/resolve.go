// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package lab

import (
	"fmt"
	"sync/atomic"
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
//
//      +---+                       +---+
//      |   |                       |   |
//      |   v                       |   v
//  +---+-------+             +-----+-------+
//  |           |             |             |
//  | Permanent +------------>| Resolvable  |
//  |           |             |             |
//  +-----+-----+   +---+     +-------+-----+
//        |         |   |             |
//        |         |   v             |
//        |     +---+--------+        |
//        |     |            |        |
//        +---->|   Script   |<-------+
//              |            |
//              +------------+
//
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
		return NewError("permanent asset %s depends on non-permanent asset %s",
			child.FullName(), parent.FullName())
	}

	if _, parent_is_script := parent.(ScriptAsset); parent_is_script {
		return NewError("resolvable asset %s depends on unresolvable asset %s",
			child.FullName(), parent.FullName())
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
			parent.Dependents = append(parent.Dependents, node)
		}
	}

	for _, node := range A.All {
		if node.Ready() {
			roots = append(roots, node)
		}
	}
	if len(roots) == 0 {
		return nil, NewError("no root nodes found")
	}

	cycle := checkForCycles(A)
	if len(cycle) != 0 {
		cycle_string := ""
		for _, node := range cycle {
			if len(cycle_string) != 0 {
				cycle_string += " -> "
			}
			cycle_string += fmt.Sprintf("%s:%s", node.Asset.Namespace(), node.Asset.Id())
		}
		cycle_string += fmt.Sprintf(" -> %s:%s", cycle[0].Asset.Namespace(), cycle[0].Asset.Id())
		return nil, NewError("asset dependencies contain cycles: %s", cycle_string)
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
func ResolveAssets(A *Assets, S *Session) error {
	roots, err := PrepareToResolve(A)
	if err != nil {
		return err
	}

	// |resolvable| needs enough of a buffer to prevent the workers from being
	// starved due to the buffer size issues. Since there are
	// kMaxParallelOperations workers, we should be able to queue that many
	// nodes for resolution.
	resolvable := make(chan *DependencyNode, kMaxParallelOperations)

	// |waiting| should have a buffer that's big enough to prevent
	// |prepareToResolve| getting blocked when writing to the channel. The
	// waiting channel accepts batches of DependencyNodes.
	//
	// These batches are created when a DependencyNode successfully resolves.
	// If the goroutine that resolved the node attempts to queue each node to
	// the resolvable channel, it'd be posssible for the queuing operation to
	// block due to the channel buffer being full. For a dependency graph that
	// is very wide and shallow, this may cause all resolver goroutines to be
	// blocked during the queuing phase leaving none to drain the resolvable
	// channel. This would lead to a stall.
	//
	// To prevent this, we use a |waiting| channel that accepts batches of
	// nodes, and the primary thread is tasked with draining the |waiting|
	// channel and populating the |resolvable| channel. The selection of the
	// queue sizes guarantees that the resolver routines won't block.
	waiting := make(chan []*DependencyNode, kMaxParallelOperations)

	// pending == count of nodes in |waiting| + |resolvable| + waiting to be
	// enqueued to resolvable + being worked on by a worker.
	var pending int32
	pending = int32(len(roots))

	waiting <- roots

	// Here we kick off the resolver routines.
	for i := 0; i < kMaxParallelOperations; i += 1 {
		go resolveNodesFromChannel(S, resolvable, waiting, &pending)
	}

	// And the |waiting| channel drainer.
	for {
		b := <-waiting
		if b == nil {
			break
		}
		if len(b) == 0 {
			panic("invalid batch")
		}

		for _, d := range b {
			resolvable <- d
		}
	}

	// Once we get here, there are no more pending tasks. Any nodes left
	// unresolved didn't have their dependencies resolved.
	errors := []error{}
	for _, d := range A.All {
		if !d.Processed {
			errors = append(errors, NewError("asset %s:%s was unresolvable. It has %d unresolved dependencies",
				d.Asset.Namespace(), d.Asset.Id(), d.UnresolvedDependencies))
			continue
		}

		if d.Result != nil {
			errors = append(errors, WrapErrorWithMessage(d.Result,
				"asset %s:%s failed to resolve", d.Asset.Namespace(), d.Asset.Id()))
		}
	}
	return WrapErrorList(errors)
}

// resolveNodesFromChannel fetches DependencyNodes from |resolvable|, resolves
// them, and if successful, enqueues a batch of newly resolvable
// DependencyNodes to |waiting|.
//
// It also atomically updates |pending| so that the latter continues to reflect
// the number of pending requests.
//
// Terminates when there are no more pending nodes, or if the |resolvable|
// channel is closed. This function is meant to be invoked as a goroutine.
func resolveNodesFromChannel(S *Session, resolvable chan *DependencyNode, waiting chan []*DependencyNode,
	pending *int32) {
	for {
		d := <-resolvable
		if d == nil {
			return
		}

		next, err := d.Resolve(S)
		if err != nil {
			LogAssetError(S, d.Asset, err)
			return
		}

		if len(next) > 0 {
			atomic.AddInt32(pending, int32(len(next)))
			waiting <- next
		}

		if atomic.AddInt32(pending, -1) == 0 {
			// d was the last DependencyNode to complete. There's nothing in the pending pipelines.
			close(resolvable)
			close(waiting)
			return
		}
	}
}

// min returns the minimum of two integers |a| and |b|. It exists because golang.
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
