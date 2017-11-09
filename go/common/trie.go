// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

type Trie struct {
	m       map[string]*Trie
	payload interface{}
}

func (t *Trie) Set(p RefPath, o interface{}, replace bool) bool {
	if o == nil {
		return false
	}

	if len(p) == 0 {
		if !replace && t.payload != nil {
			return false
		}
		t.payload = o
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

func (t *Trie) Get(p RefPath) interface{} {
	if len(p) == 0 {
		return t.payload
	}

	if t.m == nil {
		return nil
	}

	head, tail := p.Shift()
	if c, ok := t.m[head]; ok {
		return c.Get(tail)
	}
	return nil
}

func (t *Trie) Empty() bool {
	return len(t.m) == 0 && t.payload == nil
}

func (t *Trie) Unset(p RefPath) {
	if len(p) == 0 {
		t.payload = nil
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

func (t *Trie) Size() int {
	size := 0
	t.Visit(func(RefPath, interface{}) bool {
		size += 1
		return true
	})
	return size
}

type TrieVisitor func(RefPath, interface{}) bool

func (t *Trie) Visit(f TrieVisitor) bool {
	return t.visit(EmptyPath, EmptyPath, f)
}

func (t *Trie) VisitFrom(start RefPath, f TrieVisitor) bool {
	return t.visit(start, EmptyPath, f)
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

	if t.payload != nil {
		cont := f(sofar, t.payload)
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
