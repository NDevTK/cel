// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"github.com/pkg/errors"
	"strings"
)

// RefPath or a reference path is list of labels that represent the path to an
// asset or property of an asset relative to another asset.
//
// Usually reference paths are constructed relative to the root of the combined
// asset and host environment namespace.
//
// Each label in a reference path identifies the next step down the path and
// may represent one of the following:
//
// * A field of an object.
//
//   The label is simply the name of the field and refers to the field value.
//
// * A named object in a collection.
//
//   To refer to an object in a collection where the object has a 'name' field,
//   use the value of the 'name' field as the label.
//
// * A value in a map whose key is a string.
//
//   To refer to an object stored in a key-value pair where the object is the
//   value and the key is a string, use the key as the label.
//
// * An object in a list.
//
//   To refer to an object contained in an ordered list, use an "@" followed by
//   the decimal representation of the index (0-based) as the label. E.g. the
//   label "@2" can be used to rerfer to the 3rd object in an array.
//
// By convention, all asset references start with the label "asset", while all
// host environment references start with the label "host".
//
// Serializing a reference path is achieved by concatenating the list of labels
// in order separated by ".". If a label contains the separator, then that
// label is enclosed in parenthesis ( "(" and ")" ) in order to disambiguate
// between internal "." and the separator "."s.
//
// For example, take the following asset:
//
//     asset {
//       windows_machine {
//         name: "ad"
//         windows_feature: "Web-Server"
//         windows_feature: "Coffee-Maker"
//         configuration_file {
//           source: "configs/ad.config"
//         }
//       }
//     }
//
// The entire windows_machine asset has the ref path ["asset",
// "windows_machine", "ad"]. Equivalently, the single serialized string
// "asset.windows_machine.ad" can also be used to refer to the asset.
//
// The source of the configuration file can be referred to as:
//
//     "asset.windows_machine.ad.configuration_file.source"
//
// The second windows_feature can be referred to as:
//
//     "asset.windows_machine.ad.windows_feature.@1"
//
// If the name of the asset was "foo.bar" (i.e. it contains a "." in the name),
// then the serialized reference path becomes:
//
//     "asset.windows_machine.(foo.bar)"
//
// The RefPathFromString() and RefPathFromComponents() functions can be used to
// construct RefPath objects. Meanwhile, the convenient global EmptyPath can be
// used in place of an empty path.
type RefPath []string

// RefPathFromComponents constructs a RefPath from its arguments which are assumed
// to be all strings.
//
// For example, RefPathFromComponents("a", "b", "c") will construct a RefPath with
// three components, "a", "b", and "c".
func RefPathFromComponents(s ...string) RefPath {
	return s
}

// RefPathFromString constructs a RefPath from a string containing a serialized
// RefPath.
//
// See the documentation for RefPath for details on how a RefPath is
// serialized.
//
// For example, RefPathFromString("a.b.c") will construct a RefPath with three
// components, "a", "b", and "c".
func RefPathFromString(s string) (RefPath, error) {
	if s == "" {
		return EmptyPath, nil
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
				return EmptyPath, ErrInvalidRefPathString
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
	return r, nil
}

// RefPathMust parses a string and returns a RefPath. Unlike
// RefPathFromString(), this version panics if the supplied string is not
// valid.
//
// Only use this function with hardcoded or trusted strings.
func RefPathMust(s string) RefPath {
	r, err := RefPathFromString(s)
	if err != nil {
		panic(err)
	}
	return r
}

// Append returns a new RefPath that represents a path with one or more
// additional labels. The receiver is not modified.
//
// For example:
//
//     a := RefPathFromComponents("a", "b")
//     b := a.Append("c", "d")
//
// |a| now contains the path "a.b", while |b| contains the path "a.b.c.d".
func (r RefPath) Append(s ...string) RefPath {
	rn := make([]string, len(r)+len(s))
	copy(rn, r)
	copy(rn[len(r):], s)
	return rn
}

// String returns the serialized form of the reference path.
//
// See RefPath type documentation for serialization rules.
//
// For example:
//
//     a := RefPathFromComponents("a", "b", "c.d", "e")
//
// a.String() returns "a.b.(c.d).e"
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

// Equals returns true of the receiver and |o| contain equivalent reference
// paths.
//
// Two reference paths are considered equivalent if they contain the same
// sequence of labels. Label comparisons are always case sensitive.
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

// Less returns true of r should be considered lexicographically less than o.
func (r RefPath) Less(o RefPath) bool {
	count := len(r)
	if len(o) < count {
		count = len(o)
	}

	for i := 0; i < count; i++ {
		c := strings.Compare(r[i], o[i])
		if c != 0 {
			return c < 0
		}
	}

	return len(r) < len(o)
}

// Contains returns true iff the reference path in |o| is falls within the
// hierarchy of the receiver.
//
// In other words, the reference path in |o| is reachable from the reference
// path in the receiver.
//
// In other words, the receiving reference path is a prefix of |o|.
//
// E.g.:
//
//     a := RefPathFromComponents("a", "b")
//
// These predicates evaluate to true:
//
//     a.Contains(RefPathFromComponents("a", "b", "c"))
//     a.Contains(RefPathFromComponents("a", "b"))
//
// These predicates evaluate to false:
//
//     a.Contains(RefPathFromComponents("a", "x")) returns false.
//     a.Contains(RefPathFromComponents("a")) returns false.
//     a.Contains(RefPathFromComponents("x")) returns false.
func (r RefPath) Contains(o RefPath) bool {
	if len(r) > len(o) {
		return false
	}
	return r.Equals(o[:len(r)])
}

// After returns the portion of the reference path following the path in |o|.
//
// If |o| is not a prefix of |r| (or equivalently, if !o.Contains(r)), it is
// considered an error. In this case, a nil RefPath is returned along with
// the boolean false.
//
// E.g.: RefPathFromString("a.b.c").After(RefPathFromString("a.b")) returns the
// equivalent of RefPathFromString("c")
func (r RefPath) After(o RefPath) (RefPath, bool) {
	if !o.Contains(r) {
		return nil, false
	}
	return r[len(o):], true
}

// Shift returns the head and the tail of a reference path. The receiver is not
// modified. Returns an empty string and a nil path if there the receiver is
// empty.
func (r RefPath) Shift() (string, RefPath) {
	if len(r) == 0 {
		return "", nil
	}

	return r[0], r[1:]
}

// TopLevel returns the top-level reference corresponding to this RefPath.
//
// For CEL, top level resources are always located three levels deep in the
// hierarchy. E.g. "asset.ad_domain.my-domain.some_property" belongs to the
// top-level resource "asset.ad_domain.my-domain".
//
// Returns EmptyPath if there's no top-level resource corresponding to `r`.
func (r RefPath) TopLevel() RefPath {
	if len(r) < 3 {
		return EmptyPath
	}
	return r[:3]
}

// Parent returns a RefPath that excludes the final component.
func (r RefPath) Parent() RefPath {
	if len(r) == 0 {
		return EmptyPath
	}
	return r[:len(r)-1]
}

// Empty returns true if the path is empty.
func (r RefPath) Empty() bool {
	return len(r) == 0
}

var (
	// EmptyPath is a convenient RefPath instance that contains an empty path.
	EmptyPath = RefPath{}

	// ErrInvalidRefPathString is returned by RefPathFromString if the input was malformed.
	ErrInvalidRefPathString = errors.New("invalid RefPath string")
)

const (
	RefPathSeperator = "."
)
