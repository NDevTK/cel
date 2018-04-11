// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

// Trie implements a trie where the keys are RefPaths and values are
// interface{}. The depth of the trie is the number of components of the
// longest RefPath contained therein.
//
// More information about generic tries can be found at
// https://en.wikipedia.org/wiki/Trie
//
// Values are not restricted to leaves. Internal nodes can contain values as
// well.
//
// Trie is not thread safe. It should only be accessed from the main thread, or
// its access should be synchronized by its owner.
//
// Each node in a Trie is also a Trie.
//
// For example, the following trie represents an object `o` with the RefPath
// "a.b.c", and an object 'p' with the RefPath "a.b.d".
//
//                                   "a"
//                                    |
//                                   "b"
//                                  /   \
//                                "c"   "d"
//                                 o     p
//
// In order to look up the object at path "a.b.c", one must visit the trie
// labeled "a", find the child with label "b", and then visit that to find the
// child with label "c".
//
// This Trie implementation is conceptually equivalent to an associative array
// (or a map) from RefPath to an arbitrary value. However, Trie provides more
// efficient implementations for dealing with subtrees than an associative
// array. This property is important in the context of CEL where subtrees
// represent scopes. Hence we can quickly run a predicate over a scope via a
// Trie.
//
// For example, in the Trie above, the subtree rooted at "a.b" is readily
// determined by a Trie, and all the nodes can be efficiently visited because
// the subtree under "a.b" is isolated within a sub-trie. I.e. the cost of
// visiting the subtree does not depend on how many nodes there are outside the
// subtree. This wouldn't be the case for an associative array.
//
// The types for values are not restricted. However nil is not a valid value.
type Trie struct {
	m     map[string]*Trie
	value interface{}
}

// Set adds value |o| to the trie at path |p|. If |replace| is true, then the
// value replaces any existing value at the same path. Otherwise the operation
// fails if a value already exists at the path.
//
// Adding a value at path |p| does not affect any values that already exist
// deeper in the tree that share the same prefix.
//
// The return value is true if the operation succeeded and the value was added
// to the trie.
//
// A path with length 0 can be used to set a value at the root of the trie.
//
// The value |o| cannot be nil.
func (t *Trie) Set(p RefPath, o interface{}, replace bool) bool {
	if o == nil {
		return false
	}

	if len(p) == 0 {
		if !replace && t.value != nil {
			return false
		}
		t.value = o
		return true
	}

	if t.m == nil {
		t.m = make(map[string]*Trie)
	}

	head, tail := p.Shift()
	if _, ok := t.m[head]; !ok {
		t.m[head] = &Trie{}
	}
	return t.m[head].Set(tail, o, replace)
}

// Get retrieves the value at path |p|.
//
// Returns two values. The first is the value stored at |p|, and the second is
// a boolean which is true iff there was a value stored at |p|. Obviously the
// first return value is valid only if the second value is true.
func (t *Trie) Get(p RefPath) (interface{}, bool) {
	if len(p) == 0 {
		return t.value, t.value != nil
	}

	if t.m == nil {
		return nil, false
	}

	head, tail := p.Shift()
	if c, ok := t.m[head]; ok {
		return c.Get(tail)
	}
	return nil, false
}

// GetDeepest finds the deepest node in the trie whose path is a prefix of the
// given path `p`, and returns that node's value. The return value is the tuple
// consisting of the value of that node, and the remainder of the path.
//
// If the entire path could be traversed, the returned path segment will be
// empty.
//
// If the none of the path could be traversed, the returend path will be the
// same as the input. The value would be the value at the root of the trie
// since that is the bifurcation point.
func (t *Trie) GetDeepest(p RefPath) (interface{}, RefPath) {
	if len(p) == 0 || t.m == nil {
		return t.value, p
	}

	head, tail := p.Shift()
	if c, ok := t.m[head]; ok {
		return c.GetDeepest(tail)
	}

	return t.value, p
}

// Has returns true if there exists a value at |p|.
//
// If trie.Has(p) returns true, then trie.Get(p) is guaranteed to succeed.
func (t *Trie) Has(p RefPath) bool {
	if len(p) == 0 {
		return t.value != nil
	}

	if t.m == nil {
		return false
	}

	head, tail := p.Shift()
	if c, ok := t.m[head]; ok {
		return c.Has(tail)
	}
	return false
}

// Empty return true if there are no values in the trie.
func (t *Trie) Empty() bool {
	return len(t.m) == 0 && t.value == nil
}

// Unset removes the value at path |p|.
func (t *Trie) Unset(p RefPath) {
	if len(p) == 0 {
		t.value = nil
		return
	}

	if t.m == nil {
		return
	}

	head, tail := p.Shift()
	if c, ok := t.m[head]; ok {
		c.Unset(tail)
		if c.Empty() {
			delete(t.m, head)
		}
	}
}

// Size returns a count of values that are stored in this trie.
func (t *Trie) Size() int {
	size := 0
	t.Visit(func(RefPath, interface{}) bool {
		size += 1
		return true
	})
	return size
}

// TrieVisitor is the type of the function used by Visit() and VisitFrom(). See
// Visit() for more details on how this function is used.
type TrieVisitor func(RefPath, interface{}) bool

// Visit calls a TrieVisitor function for each of the values that are stored in
// the trie.
//
// For every value |v| found at path |p| in the trie, |visitor| will be called
// as follows:
//
//     visitor(p, v)
//
// If |visitor| returns false, then the visit is aborted. It is guaranteed that
// once |visitor| returns false, it will not be called again.
//
// Note: It is not safe to call Unset() or Set() from |visitor|.
//
// Note: The order in which the nodes are visited is unspecified.
func (t *Trie) Visit(visitor TrieVisitor) bool {
	return t.visit(EmptyPath, EmptyPath, visitor)
}

// VisitFrom behaves like Visit but restricts the visit to a specific path.
//
// All calls to |f| are guaranteed to be made with a path |p| that satisfies
// |start.Contains(p)|.
func (t *Trie) VisitFrom(start RefPath, visitor TrieVisitor) bool {
	return t.visit(start, EmptyPath, visitor)
}

// DescendUntil invokes `visitor` on all nodes along the way until and
// including the node identified by `target`.
//
// The ordering is from the root of the trie downwards. If the path identified
// by `target` doesn't exist in the trie, visits the node in the prefix of
// `target` that exist in the trie.
//
// The returne value is true if all the nodes along the length of `target`
// existed, and `visitor` returned true for all invocations. False if any of
// those conditions are not met.
//
// `visitor` is only invoked for nodess that have a value associated with them.
func (t *Trie) DescendUntil(target RefPath, visitor TrieVisitor) bool {
	if len(target) == 0 {
		return false
	}

	here := t
	sofar := EmptyPath
	next, togo := target.Shift()

	for here != nil {
		if here.value != nil && !visitor(sofar, here.value) {
			return false
		}

		if next == "" {
			return true
		} else if n, ok := here.m[next]; ok {
			sofar = sofar.Append(next)
			here = n
			next, togo = togo.Shift()
		} else {
			// Found node in the target path doesn't exist.
			return false
		}
	}

	return true
}

func (t *Trie) visit(start RefPath, sofar RefPath, f TrieVisitor) bool {
	if len(start) != 0 {
		if t.m == nil {
			return false
		}

		head, tail := start.Shift()
		if c, ok := t.m[head]; ok {
			return c.visit(tail, sofar.Append(head), f)
		}
		return false
	}

	if t.value != nil {
		cont := f(sofar, t.value)
		if !cont {
			return false
		}
	}

	if t.m == nil {
		return true
	}

	for n, c := range t.m {
		cont := c.visit(EmptyPath, sofar.Append(n), f)
		if !cont {
			return false
		}
	}
	return true
}
