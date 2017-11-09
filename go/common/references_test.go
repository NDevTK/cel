// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"github.com/pkg/errors"
	"strings"
	"testing"
)

func TestReferences_CollectFrom(t *testing.T) {
	m := TestMessageWithOptions{
		Name:        "from.here",
		Key:         "foo",
		OptionalKey: "bar",
		Fqdn:        "${with_types.xyz} ${with_types.abc}"}
	var r References

	err := r.CollectFrom(&m, RefPathFromString("with_options"))
	if err != nil {
		t.Fatal(err)
	}

	if u := r.Unresolved.Get(RefPathFromString("with_types.repeated_field.foo")).(*UnresolvedReference); u != nil {
		if u.From.String() != "with_options.key" {
			t.Fatal(u.From)
		}

		if u.To.String() != "with_types.repeated_field.foo" {
			t.Fatal(u.To)
		}
	} else {
		t.Fatal(r)
	}

	if u := r.Unresolved.Get(RefPathFromString("with_types.repeated_field.bar")).(*UnresolvedReference); u != nil {
		if u.From.String() != "with_options.optional_key" {
			t.Fatal(u.From)
		}
	} else {
		t.Fatal(r)
	}

	if u := r.Unresolved.Get(RefPathFromString("with_types.xyz")).(*UnresolvedReference); u != nil {
		if u.From.String() != "with_options.fqdn" {
			t.Fatal(u.From)
		}
	} else {
		t.Fatal(r)
	}

	if u := r.Unresolved.Get(RefPathFromString("with_types.abc")).(*UnresolvedReference); u != nil {
		if u.From.String() != "with_options.fqdn" {
			t.Fatal(u.From)
		}
	} else {
		t.Fatal(r)
	}
}

func TestReferences_ResolveWith(t *testing.T) {
	m := TestMessageWithOptions{
		Name:        "from.here",
		Key:         "foo",
		OptionalKey: "bar",
		Fqdn:        "${with_types.repeated_field.foo.name}${with_types.repeated_field.bar.name}"}
	w := TestMessageWithTypes{
		Name: "to",
		RepeatedField: []*TestGoodProto{
			&TestGoodProto{"foo"}}}

	var r References
	r.AddSource(&m, RefPathFromString("with_options"))
	r.AddSource(&w, RefPathFromString("with_types"))

	err := r.Resolve(ResolutionSkipOutputs)
	e, ok := errors.Cause(err).(*UnresolvedReferenceError)
	if !ok {
		t.Fatal(err)
	}
	if e.From.String() != "with_options.fqdn" || e.To.String() != "with_types.repeated_field.bar.name" {
		t.Fatal(e)
	}

	if r.Resolved.Size() != 2 {
		t.Fatal(r)
	}

	if v := r.Resolved.Get(RefPathFromString("with_types.repeated_field.foo")); v == nil || v.(*TestGoodProto).Name != "foo" {
		t.Fatal()
	}

	if v := r.Resolved.Get(RefPathFromString("with_types.repeated_field.foo.name")); v == nil || v.(string) != "foo" {
		t.Fatal()
	}

	if r.Unresolved.Size() != 2 {
		t.Fatal()
	}

	if v := r.Unresolved.Get(RefPathFromString("with_types.repeated_field.bar")); v == nil {
		t.Fatal()
	}
}

func TestReferences_ResolveWith_Output(t *testing.T) {
	path := RefPathFromString("with_options")
	m := TestMessageWithOptions{Output: "x", Fqdn: "${with_options.output}"}
	var r References
	r.AddSource(&m, path)

	err := r.Resolve(ResolutionSkipOutputs)
	if err != nil {
		t.Fatal(err)
	}

	if m.Fqdn != "${with_options.output}" {
		t.Fatal(m)
	}

	err = r.Resolve(ResolutionIncludeOutputs)
	if err != nil {
		t.Fatal(err)
	}

	if m.Fqdn != "x" {
		t.Fatal(m)
	}
}

func TestReferences_ExpandString(t *testing.T) {
	m := TestMessageWithOptions{Fqdn: "${repeated_field.a.name}${repeated_field.b.name}${repeated_field.c}${repeated_field.c.name}"}
	w := TestMessageWithTypes{
		Name: "to",
		RepeatedField: []*TestGoodProto{
			&TestGoodProto{"a"},
			&TestGoodProto{"b"},
			&TestGoodProto{"c"},
		}}

	var r References
	r.AddSource(&m, EmptyPath.Append("foo"))
	r.AddSource(&w, EmptyPath)

	err := r.Resolve(ResolutionIncludeOutputs)
	if err == nil {
		t.Fatal()
	}
	ure, ok := errors.Cause(err).(*UnresolvedReferenceError)
	if !ok {
		t.Fatal(err)
	}

	if ure.To.String() != "repeated_field.c" || ure.Reason != "target object is not a string" {
		t.Fatal(ure)
	}

	if !r.Unresolved.Empty() {
		t.Fatal(r)
	}

	if s, err := r.ExpandString(""); err != nil || s != "" {
		t.Fatal(s, err)
	}

	if s, err := r.ExpandString("abc"); err != nil || s != "abc" {
		t.Fatal(s, err)
	}

	if s, err := r.ExpandString("\\${abc}"); err != nil || s != "\\${abc}" {
		t.Fatal(s, err)
	}

	if s, err := r.ExpandString("abc${abc"); err != nil || s != "abc${abc" {
		t.Fatal(s, err)
	}

	if s, err := r.ExpandString("abc\\${abc}"); err != nil || s != "abc\\${abc}" {
		t.Fatal(s, err)
	}

	if s, err := r.ExpandString("${abc}"); err == nil || !strings.Contains(err.Error(), "could not be resolved") {
		t.Fatal(s, err)
	}

	if s, err := r.ExpandString("${repeated_field.c}"); err == nil || !strings.Contains(err.Error(), "is not a string") {
		t.Fatal(s, err)
	}

	if s, err := r.ExpandString("a${repeated_field.a.name}c"); err != nil || s != "aac" {
		t.Fatal(s, err)
	}

	if s, err := r.ExpandString("a${repeated_field.a.name}${repeated_field.b.name}${repeated_field.c.name}"); err != nil || s != "aabc" {
		t.Fatal(s, err)
	}
}

func TestReferences_ResolveInlineRefs(t *testing.T) {
	m := TestMessageWithOptions{Fqdn: "${repeated_field.a.name}${repeated_field.b.name}${repeated_field.c}${repeated_field.c.name}"}
	w := TestMessageWithTypes{
		Name: "to",
		RepeatedField: []*TestGoodProto{
			&TestGoodProto{"a"},
			&TestGoodProto{"b"},
			&TestGoodProto{"c"},
		}}
	var r References
	r.AddSource(&m, EmptyPath.Append("foo"))
	r.AddSource(&w, EmptyPath)

	err := r.Resolve(ResolutionIncludeOutputs)
	if !strings.Contains(err.Error(), "is not a string") || !strings.Contains(err.Error(), "fqdn") {
		t.Fatalf("unexpected error %#v", err)
	}
	m.Fqdn = "${repeated_field.a.name}${repeated_field.b.name}${repeated_field.c.name}"

	err = r.Resolve(ResolutionIncludeOutputs)
	if err != nil {
		t.Fatal(err)
	}
	if m.Fqdn != "abc" {
		t.Fatal(m.Fqdn)
	}
}
