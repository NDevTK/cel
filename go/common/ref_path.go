// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"strings"
)

type RefPath []string

func RefPathFromList(s ...string) RefPath {
	return s
}

func RefPathFromString(s string) RefPath {
	if s == "" {
		return EmptyPath
	}
	r := make([]string, 0, strings.Count(s, RefPathSeperator))
	for len(s) > 0 {
		if s[0] == '.' {
			s = s[1:]
			continue
		}

		if s[0] == '(' {
			e := strings.Index(s, ")")
			if e == -1 {
				// unbalanced ()
				return RefPath{}
			}
			r = append(r, s[1:e])
			s = s[e+1:]
			continue
		}

		e := strings.Index(s, RefPathSeperator)
		if e == -1 {
			r = append(r, s)
			break
		}
		r = append(r, s[:e])
		s = s[e:]
	}
	return r
}

func (r RefPath) Append(s ...string) RefPath {
	rn := make([]string, len(r), len(r)+len(s))
	copy(rn, r)
	for _, ss := range s {
		rn = append(rn, ss)
	}
	return rn
}

func (r RefPath) String() string {
	var s []string
	for _, p := range r {
		if strings.Contains(p, RefPathSeperator) {
			s = append(s, "("+p+")")
		} else {
			s = append(s, p)
		}
	}
	return strings.Join(s, RefPathSeperator)
}

func (r RefPath) Equals(o RefPath) bool {
	if len(r) != len(o) {
		return false
	}
	for i, _ := range r {
		if r[i] != o[i] {
			return false
		}
	}
	return true
}

func (r RefPath) Contains(o RefPath) bool {
	if len(r) > len(o) {
		return false
	}
	return r.Equals(o[:len(r)])
}

func (r RefPath) After(o RefPath) (RefPath, bool) {
	if !o.Contains(r) {
		return nil, false
	}
	return r[len(o):], true
}

func (r RefPath) Shift() (string, RefPath) {
	if len(r) == 0 {
		return "", nil
	}

	return r[0], r[1:]
}

var (
	EmptyPath = RefPath{}
)

const (
	RefPathSeperator = "."
)
