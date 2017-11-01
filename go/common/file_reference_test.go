// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"strings"
	"testing"
)

func TestFileReference_Basic(t *testing.T) {
	p := RefPath{}
	t.Run("Empty", func(t *testing.T) {
		v := &FileReference{}
		if err := InvokeValidate(v, p); err == nil {
			t.Fail()
		}
	})

	t.Run("Absolute", func(t *testing.T) {
		v := &FileReference{Path: "/foo/bar"}
		if err := InvokeValidate(v, p); err == nil || !strings.Contains(err.Error(), "cannot be absolute") {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("Backslash", func(t *testing.T) {
		v := &FileReference{Path: "foo\\bar"}
		if err := InvokeValidate(v, p); err == nil || !strings.Contains(err.Error(), "backslash") {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("ParentRefPrefix", func(t *testing.T) {
		v := &FileReference{Path: "../bar"}
		if err := InvokeValidate(v, p); err == nil || !strings.Contains(err.Error(), "parent path reference") {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("ParentRefSuffix", func(t *testing.T) {
		v := &FileReference{Path: "foo/bar/.."}
		if err := InvokeValidate(v, p); err == nil || !strings.Contains(err.Error(), "parent path reference") {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("ParentRefInfix", func(t *testing.T) {
		v := &FileReference{Path: "foo/../bar"}
		if err := InvokeValidate(v, p); err == nil || !strings.Contains(err.Error(), "parent path reference") {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("OutputParam", func(t *testing.T) {
		v := &FileReference{ResolvedPath: "foo/bar"}
		if err := InvokeValidate(v, p); err == nil || !strings.Contains(err.Error(), "marked as output") {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("Valid", func(t *testing.T) {
		v := &FileReference{Path: "foo/bar"}
		if err := InvokeValidate(v, p); err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}
	})
}

func TestFileReference_Resolve(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		v := &FileReference{Path: "foo/bar"}
		err := v.ResolveRelativePath("")
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}

		if v.ResolvedPath != "foo/bar" {
			t.Fatalf("unexpected resolved path: %s", v.ResolvedPath)
		}
	})

	t.Run("Rel", func(t *testing.T) {
		v := &FileReference{Path: "foo/bar"}
		err := v.ResolveRelativePath("a/b")
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}

		if v.ResolvedPath != "a/b/foo/bar" {
			t.Fatalf("unexpected resolved path: %s", v.ResolvedPath)
		}
	})

	t.Run("RelSlash", func(t *testing.T) {
		v := &FileReference{Path: "foo/bar"}
		err := v.ResolveRelativePath("a/b/")
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}

		if v.ResolvedPath != "a/b/foo/bar" {
			t.Fatalf("unexpected resolved path: %s", v.ResolvedPath)
		}
	})

	t.Run("Abs", func(t *testing.T) {
		v := &FileReference{Path: "foo/bar"}
		err := v.ResolveRelativePath("/a/b/c")
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}

		if v.ResolvedPath != "/a/b/c/foo/bar" {
			t.Fatalf("unexpected resolved path: %s", v.ResolvedPath)
		}
	})

	t.Run("AbsSlash", func(t *testing.T) {
		v := &FileReference{Path: "foo/bar"}
		err := v.ResolveRelativePath("/a/b/c/")
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}

		if v.ResolvedPath != "/a/b/c/foo/bar" {
			t.Fatalf("unexpected resolved path: %s", v.ResolvedPath)
		}
	})
}

func TestFileReference_Resolver(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		v := &TestFileRefProto{}
		v.Ref = &FileReference{Path: "foo/bar"}
		err := WalkProtoMessage(v, RefPath{}, GetPathResolver("/a/b/c"))
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}

		if v.Ref.ResolvedPath != "/a/b/c/foo/bar" {
			t.Fatalf("bad resolved path : %s", v.Ref.ResolvedPath)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		v := &TestFileRefProto{}
		err := WalkProtoMessage(v, RefPath{}, GetPathResolver("/a/b/c"))
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("EmptyPath", func(t *testing.T) {
		v := &TestFileRefProto{Ref: &FileReference{Path: ""}}
		err := WalkProtoMessage(v, RefPath{}, GetPathResolver("/a/b/c"))
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}
	})
}
