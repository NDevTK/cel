// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"strings"
	"testing"
)

func TestReferences_Collect(t *testing.T) {
	m := TestMessageWithOptions{Name: "from.here", Key: "foo", OptionalKey: "bar", Fqdn: "${with_types.xyz} ${with_types.abc}"}
	var r References

	err := r.Collect(&m, RefPathFromString("with_options"))
	if err != nil {
		t.Fatal(err)
	}

	if u, ok := r.Unresolved["with_types.repeated_field.foo"]; ok {
		if u.From.String() != "with_options.key" {
			t.Fatal(u.From)
		}

		if u.To.String() != "with_types.repeated_field.foo" {
			t.Fatal(u.To)
		}
	} else {
		t.Fatal(r)
	}

	if u, ok := r.Unresolved["with_types.repeated_field.bar"]; ok {
		if u.From.String() != "with_options.optional_key" {
			t.Fatal(u.From)
		}
	} else {
		t.Fatal(r)
	}

	if u, ok := r.Unresolved["with_types.xyz"]; ok {
		if u.From.String() != "with_options.fqdn" {
			t.Fatal(u.From)
		}
	} else {
		t.Fatal(r)
	}

	if u, ok := r.Unresolved["with_types.abc"]; ok {
		if u.From.String() != "with_options.fqdn" {
			t.Fatal(u.From)
		}
	} else {
		t.Fatal(r)
	}
}

func TestReferences_ResolveWith(t *testing.T) {
	m := TestMessageWithOptions{Name: "from.here", Key: "foo", OptionalKey: "bar",
		Fqdn: "${with_types.repeated_field.foo.name}${with_types.repeated_field.bar.name}"}
	var r References

	err := r.Collect(&m, RefPathFromString("with_options"))
	if err != nil {
		t.Fatal(err)
	}

	w := TestMessageWithTypes{Name: "to", RepeatedField: []*TestGoodProto{
		&TestGoodProto{"foo"}}}
	err = r.ResolveWith(&w, RefPathFromString("with_types"), ResolutionIncludeOutputs)
	if err != nil {
		t.Fatal(err)
	}

	if len(r.Resolved) != 2 {
		t.Fatal(r)
	}

	if v, ok := r.Resolved["with_types.repeated_field.foo"]; !ok || v.(*TestGoodProto).Name != "foo" {
		t.Fatal()
	}

	if v, ok := r.Resolved["with_types.repeated_field.foo.name"]; !ok || v.(string) != "foo" {
		t.Fatal()
	}

	if len(r.Unresolved) != 2 {
		t.Fatal()
	}

	if _, ok := r.Unresolved["with_types.repeated_field.bar"]; !ok {
		t.Fatal()
	}
}

func TestReferences_ResolveWith_Output(t *testing.T) {
	path := RefPathFromString("with_options")
	m := TestMessageWithOptions{Output: "x", Fqdn: "${with_options.output}"}
	var r References

	err := r.Collect(&m, path)
	if err != nil {
		t.Fatal(err)
	}

	err = r.ResolveWith(&m, path, ResolutionSkipOutputs)
	if err != nil {
		t.Fatal(err)
	}

	err = r.ResolveInlineRefs(&m, path)
	if err != nil {
		t.Fatal(err)
	}

	if m.Fqdn != "${with_options.output}" {
		t.Fatal(m)
	}

	err = r.ResolveWith(&m, path, ResolutionIncludeOutputs)
	if err != nil {
		t.Fatal(err)
	}

	err = r.ResolveInlineRefs(&m, path)
	if err != nil {
		t.Fatal(err)
	}

	if m.Fqdn != "x" {
		t.Fatal(m)
	}
}

func TestReferences_ExpandString(t *testing.T) {
	m := TestMessageWithOptions{Fqdn: "${repeated_field.a.name}${repeated_field.b.name}${repeated_field.c}${repeated_field.c.name}"}
	var r References

	err := r.Collect(&m, EmptyPath)
	if err != nil {
		t.Fatal(err)
	}

	if len(r.Unresolved) != 4 {
		t.Fatal()
	}

	w := TestMessageWithTypes{Name: "to", RepeatedField: []*TestGoodProto{
		&TestGoodProto{"a"},
		&TestGoodProto{"b"},
		&TestGoodProto{"c"},
	}}
	err = r.ResolveWith(&w, EmptyPath, ResolutionIncludeOutputs)
	if err != nil {
		t.Fatal(err)
	}
	if len(r.Unresolved) != 0 {
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
	var r References

	err := r.Collect(&m, EmptyPath)
	if err != nil || len(r.Unresolved) != 4 {
		t.Fatal()
	}

	w := TestMessageWithTypes{Name: "to", RepeatedField: []*TestGoodProto{
		&TestGoodProto{"a"},
		&TestGoodProto{"b"},
		&TestGoodProto{"c"},
	}}
	if err := r.ResolveWith(&w, EmptyPath, ResolutionIncludeOutputs); err != nil {
		t.Fatal(err)
	}

	err = r.ResolveInlineRefs(&m, EmptyPath)
	if err == nil {
		t.Fatal()
	}
	if !strings.Contains(err.Error(), "is not a string") || !strings.Contains(err.Error(), "fqdn") {
		t.Fatalf("unexpected error %#v", err)
	}
	m.Fqdn = "${repeated_field.a.name}${repeated_field.b.name}${repeated_field.c.name}"
	err = r.ResolveInlineRefs(&m, EmptyPath)
	if err != nil {
		t.Fatal(err)
	}
	if m.Fqdn != "abc" {
		t.Fatal(m.Fqdn)
	}
}
