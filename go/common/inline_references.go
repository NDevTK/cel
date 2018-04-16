// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"github.com/pkg/errors"
	"sort"
	"strings"
)

type stringRange struct {
	Ref    RefPath
	Offset int
	Length int
}

// extractInlineReferences finds internal string references in `s`. If any are
// illformed, returns an error. Returns a nil list and nil error if `s` does
// not contain any references.
//
// The returned list is sorted in reverse order of by position.
func extractInlineReferences(s string) (refs []stringRange, err error) {
	original := s
	offset := 0
	for idx := strings.Index(s, "${"); idx != -1; idx = strings.Index(s, "${") {
		if idx > 0 && s[idx-1] == '\\' {
			// escaped $. I.e. the real token is "\${" which doesn't expand to anything.
			offset += idx + 2
			s = s[idx+2:]
			continue
		}

		offset += idx
		s = s[idx:]
		ends := strings.Index(s, "}")
		if ends == -1 {
			return nil, errors.Errorf("mismatched braces (i.e. {}) in %#v", original)
		}
		ends += 1 // consume }

		path, err := RefPathFromString(s[2 : ends-1])
		if err != nil {
			// malformed
			return nil, errors.Wrapf(err, "can't extract reference from %#v", original)
		}

		refs = append(refs, stringRange{Ref: path, Offset: offset, Length: ends})
		offset += ends
		s = s[ends:]
	}

	sort.SliceStable(refs, func(i, j int) bool {
		if refs[i].Ref.Equals(refs[j].Ref) {
			return refs[i].Offset >= refs[j].Offset
		}
		return !refs[i].Ref.Less(refs[j].Ref)
	})

	return
}

// expandSingleReference expands inline references to a single RefPath.
func expandSingleReference(target string, p RefPath, v string) (string, error) {
	refs, err := extractInlineReferences(target)
	if err != nil {
		return target, err
	}

	ok := false
	for _, ref := range refs {
		if ref.Ref.Equals(p) {
			target = target[:ref.Offset] + v + target[ref.Offset+ref.Length:]
			ok = true
		}
	}

	if !ok {
		return target, errors.Errorf("target reference %s was not found in string \"%s\"",
			p.String(), target)
	}

	return target, nil
}

// ContainsInlineReferences returns true if the input string contains marked up
// references as understood by References.CollectFrom.
func ContainsInlineReferences(s string) bool {
	return strings.Index(s, "${") != -1
}
