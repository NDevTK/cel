// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pd "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/pkg/errors"
	"reflect"
	"sort"
	"strings"
)

type UnresolvedReference struct {
	To       RefPath
	From     RefPath
	IsOutput bool
}

type UnresolvedReferenceError struct {
	To        RefPath
	From      RefPath
	InlineRef string
	Reason    string
}

func (u *UnresolvedReferenceError) Error() string {
	if len(u.From) != 0 && len(u.To) != 0 {
		return fmt.Sprintf("reference from \"%s\" to \"%s\" could not be resolved: ",
			u.From.String(), u.To.String()) + u.Reason
	}
	return fmt.Sprintf("reference to \"%s\" could not be resolved: ", u.To.String()) + u.Reason
}

type References struct {
	Unresolved map[string]*UnresolvedReference
	Resolved   map[string]interface{}
}

func (r *References) Collect(m proto.Message, root RefPath) error {
	if r.Unresolved == nil {
		r.Unresolved = make(map[string]*UnresolvedReference)
	}

	return WalkProtoMessage(m, root, func(av reflect.Value, p RefPath, f *pd.FieldDescriptorProto) error {
		// only interested in fields.
		if f == nil {
			return nil
		}

		if av.Kind() == reflect.String {
			refs := ExtractStringReferences(av.String())
			for _, ref := range refs {
				r.Unresolved[ref.Ref] = &UnresolvedReference{To: RefPathFromString(ref.Ref), From: p}
			}
		}

		v := getValidationForField(f)
		// Assumes that the field is already valid. I.e. InvokeValidate() would
		// find no errors on |m|. Otherwise we can't assume that a missing
		// field is legitimate.
		if v.Ref == "" || av.Kind() != reflect.String || av.Len() == 0 {
			return nil
		}

		key := RefPathFromString(v.Ref).Append(av.String())
		r.Unresolved[key.String()] = &UnresolvedReference{To: key, From: p}
		return nil
	})
}

func (r *References) ResolveWith(m proto.Message, root RefPath, mode ResolutionMode) error {
	if r.Unresolved == nil {
		return nil
	}

	if r.Resolved == nil {
		r.Resolved = make(map[string]interface{})
	}

	done := make(map[string]bool)

	for k, u := range r.Unresolved {
		iv, err := ResolvePath(m, root, u.To, mode)
		if err == nil {
			r.Resolved[k] = iv
			done[k] = true
		}
		if err == OutputFieldSkippedError {
			u.IsOutput = true
		}
	}

	for k, _ := range done {
		delete(r.Unresolved, k)
	}
	return nil
}

type ResolutionMode int

const (
	ResolutionSkipOutputs ResolutionMode = iota
	ResolutionIncludeOutputs
)

func (r *References) ResolveInlineRefs(m proto.Message, p RefPath) error {
	if r.Resolved == nil {
		return errors.Errorf("ResolveInlineRefs called before calling ResolveWith")
	}
	return WalkProtoMessage(m, p, func(av reflect.Value, p RefPath, f *pd.FieldDescriptorProto) error {
		if f == nil || av.Kind() != reflect.String {
			return nil
		}

		s, err := r.ExpandString(av.String())
		if err != nil {
			el := UnpackErrorList(err)
			for _, e := range el {
				if u, ok := e.(*UnresolvedReferenceError); ok {
					u.From = p
				}
			}
			return WrapErrorList(el)
		}
		av.SetString(s)
		return nil
	})
}

func (r *References) ExpandString(s string) (string, error) {
	refs := ExtractStringReferences(s)
	if len(refs) == 0 {
		return s, nil
	}

	var err_list []error

	sort.SliceStable(refs, func(i, j int) bool {
		if refs[i].Ref == refs[j].Ref {
			return refs[i].Offset >= refs[j].Offset
		}
		return strings.Compare(refs[i].Ref, refs[j].Ref) >= 0
	})
	for _, ref := range refs {
		subst, ok := r.Resolved[ref.Ref]
		if !ok {
			if u, ok := r.Unresolved[ref.Ref]; ok && u.IsOutput {
				// Unresolved output field. Leave it as-is.
				continue
			}
			err_list = append(err_list, &UnresolvedReferenceError{
				To:        RefPathFromString(ref.Ref),
				InlineRef: s,
				Reason:    "not found"})
			continue
		}

		subst_s, ok := subst.(string)
		if !ok {
			err_list = append(err_list, &UnresolvedReferenceError{
				To:        RefPathFromString(ref.Ref),
				InlineRef: s,
				Reason:    "target object is not a string"})
			continue
		}

		s = s[:ref.Offset] + subst_s + s[ref.Offset+ref.Length:]
	}

	return s, WrapErrorList(err_list)
}

type stringRange struct {
	Ref    string
	Offset int
	Length int
}

func ExtractStringReferences(s string) (refs []stringRange) {
	offset := 0
	// no delimeters?
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
			// malformed
			return
		}
		ends += 1 // consume }

		refs = append(refs, stringRange{Ref: s[2 : ends-1], Offset: offset, Length: ends})
		offset += ends
		s = s[ends:]
	}
	return
}