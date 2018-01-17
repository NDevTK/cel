// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
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

type SymbolSource struct {
	Message proto.Message
	Root    RefPath
}

type References struct {
	Unresolved Trie // objects are *UnresolvedReference
	Resolved   Trie // objects are {}interface
	Sources    Trie // objects implement proto.Message
}

type DeferredString struct {
	parent   *References
	path     RefPath
	value    string
	resolved bool
}

func (d *DeferredString) Get() (string, error) {
	if d.resolved {
		return d.value, nil
	}
	is, err := d.parent.Get(d.path)
	if err != nil {
		return "", err
	}
	if s, ok := is.(string); ok {
		d.value = s
		d.resolved = true
		d.path = EmptyPath
		return s, nil
	}
	return "", errors.Errorf("value at \"%v\" is not a string. Found %#v", d.path, is)
}

func (r *References) AddSource(m proto.Message, root RefPath) error {
	ok := r.Sources.Set(root, m, false)
	if !ok {
		return errors.Errorf("a source object already exists at \"%s\"", root.String())
	}
	return nil
}

func (r *References) AddUnresolved(u *UnresolvedReference) {
	if v := r.Resolved.Get(u.To); v != nil {
		return
	}

	r.Unresolved.Set(u.To, u, false)
}

func (r *References) AddResolved(path RefPath, o interface{}) {
	r.Resolved.Set(path, o, true)
	r.Unresolved.Unset(path)
}

func (r *References) Get(p RefPath) (interface{}, error) {
	v := r.Resolved.Get(p)
	if v == nil {
		return nil, errors.Errorf("object \"%v\" not found", p)
	}
	return v, nil
}

func (r *References) GetDeferredString(p RefPath) DeferredString {
	return DeferredString{parent: r, path: p, resolved: false}
}

func (r *References) Resolve(mode ResolutionMode) error {
	var err_list []error
	r.Sources.Visit(func(p RefPath, o interface{}) bool {
		err_list = AppendErrorList(err_list, r.CollectFrom(o.(proto.Message), p))
		return true
	})
	r.Sources.Visit(func(p RefPath, o interface{}) bool {
		err_list = AppendErrorList(err_list, r.ResolveWith(o.(proto.Message), p, mode))
		return true
	})
	r.Sources.Visit(func(p RefPath, o interface{}) bool {
		err_list = AppendErrorList(err_list, r.ResolveInlineRefs(o.(proto.Message), p))
		return true
	})
	return WrapErrorList(err_list)
}

func (r *References) CollectFrom(m proto.Message, root RefPath) error {
	return WalkProtoMessage(m, root, func(av reflect.Value, p RefPath, f *pd.FieldDescriptorProto) error {
		// only interested in fields.
		if f == nil {
			return nil
		}

		if av.Kind() == reflect.String {
			refs, err := extractStringReferences(av.String())
			if err != nil {
				return err
			}
			for _, ref := range refs {
				r.AddUnresolved(&UnresolvedReference{To: ref.Ref, From: p})
			}
		}

		v := getValidationForField(f)
		// Assumes that the field is already valid. I.e. InvokeValidate() would
		// find no errors on |m|. Otherwise we can't assume that a missing
		// field is legitimate.
		if v.Ref == "" || av.Kind() != reflect.String || av.Len() == 0 {
			return nil
		}

		refpath, err := RefPathFromString(v.Ref)
		if err != nil {
			return errors.Wrapf(err, "invalid reference in validation string %#v", v.Ref)
		}

		key := refpath.Append(av.String())
		r.AddUnresolved(&UnresolvedReference{To: key, From: p})
		return nil
	})
}

func (r *References) ResolveWith(m proto.Message, root RefPath, mode ResolutionMode) error {
	done := []RefPath{}

	r.Unresolved.VisitFrom(root, func(p RefPath, o interface{}) bool {
		u := o.(*UnresolvedReference)
		iv, err := ResolvePath(m, root, u.To, mode)
		if err == nil {
			r.AddResolved(p, iv)
			done = append(done, p)
		}
		if err == OutputFieldSkippedError {
			u.IsOutput = true
		}
		return true
	})

	for _, k := range done {
		r.Unresolved.Unset(k)
	}
	return nil
}

type ResolutionMode int

const (
	ResolutionSkipOutputs ResolutionMode = iota
	ResolutionIncludeOutputs
)

func (r *References) ResolveInlineRefs(m proto.Message, p RefPath) error {
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
	refs, err := extractStringReferences(s)
	if err != nil {
		return "", err
	}
	if len(refs) == 0 {
		return s, nil
	}

	var err_list []error

	sort.SliceStable(refs, func(i, j int) bool {
		if refs[i].Ref.Equals(refs[j].Ref) {
			return refs[i].Offset >= refs[j].Offset
		}
		return !refs[i].Ref.Less(refs[j].Ref)
	})
	for _, ref := range refs {
		subst := r.Resolved.Get(ref.Ref)
		if subst == nil {
			if iu := r.Unresolved.Get(ref.Ref); iu != nil && iu.(*UnresolvedReference).IsOutput {
				// Unresolved output field. Leave it as-is.
				continue
			}
			err_list = append(err_list, &UnresolvedReferenceError{
				To:        ref.Ref,
				InlineRef: s,
				Reason:    "not found"})
			continue
		}

		subst_s, ok := subst.(string)
		if !ok {
			err_list = append(err_list, &UnresolvedReferenceError{
				To:        ref.Ref,
				InlineRef: s,
				Reason:    "target object is not a string"})
			continue
		}

		s = s[:ref.Offset] + subst_s + s[ref.Offset+ref.Length:]
	}

	return s, WrapErrorList(err_list)
}

type stringRange struct {
	Ref    RefPath
	Offset int
	Length int
}

func extractStringReferences(s string) (refs []stringRange, err error) {
	original := s
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
			return nil, errors.Errorf("mismatched object reference in string: %#v", original)
		}
		ends += 1 // consume }

		path, err := RefPathFromString(s[2 : ends-1])
		if err != nil {
			// malformed
			return nil, errors.Wrapf(err, "invalid object reference in string %#v", original)
		}

		refs = append(refs, stringRange{Ref: path, Offset: offset, Length: ends})
		offset += ends
		s = s[ends:]
	}
	return
}
