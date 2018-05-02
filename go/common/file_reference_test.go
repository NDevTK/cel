// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"strings"
	"testing"
)

func TestFileReference_Validate(t *testing.T) {
	p := RefPath{}
	t.Run("Empty", func(t *testing.T) {
		v := &FileReference{}
		if err := ValidateProto(v, p); err == nil {
			t.Fail()
		}
	})

	t.Run("Absolute", func(t *testing.T) {
		v := &FileReference{Source: "/foo/bar"}
		if err := ValidateProto(v, p); err == nil || !strings.Contains(err.Error(), "cannot be absolute") {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("Backslash", func(t *testing.T) {
		v := &FileReference{Source: "foo\\bar"}
		if err := ValidateProto(v, p); err == nil || !strings.Contains(err.Error(), "backslash") {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("ParentRefPrefix", func(t *testing.T) {
		v := &FileReference{Source: "../bar"}
		if err := ValidateProto(v, p); err == nil || !strings.Contains(err.Error(), "parent path reference") {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("ParentRefSuffix", func(t *testing.T) {
		v := &FileReference{Source: "foo/bar/.."}
		if err := ValidateProto(v, p); err == nil || !strings.Contains(err.Error(), "parent path reference") {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("ParentRefInfix", func(t *testing.T) {
		v := &FileReference{Source: "foo/../bar"}
		if err := ValidateProto(v, p); err == nil || !strings.Contains(err.Error(), "parent path reference") {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("OutputParam", func(t *testing.T) {
		v := &FileReference{Source: "foo/bar", FullPath: "foo/bar"}
		if err := ValidateProto(v, p); err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("Valid", func(t *testing.T) {
		v := &FileReference{Source: "foo/bar"}
		if err := ValidateProto(v, p); err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("NoSource", func(t *testing.T) {
		v := &FileReference{}
		if err := ValidateProto(v, p); err == nil || !strings.Contains(err.Error(), "'source' is required") {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("NoSourceWithObjectReference", func(t *testing.T) {
		v := &FileReference{ObjectReference: "foo"}
		if err := ValidateProto(v, p); err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}
	})
}

func TestFileReference_Resolve(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		v := &FileReference{Source: "foo/bar"}
		err := v.ResolveRelativePath("")
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}

		if v.FullPath != "foo/bar" {
			t.Fatalf("unexpected resolved path: %s", v.FullPath)
		}
	})

	t.Run("Rel", func(t *testing.T) {
		v := &FileReference{Source: "foo/bar"}
		err := v.ResolveRelativePath("a/b")
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}

		if v.FullPath != "a/b/foo/bar" {
			t.Fatalf("unexpected resolved path: %s", v.FullPath)
		}
	})

	t.Run("RelSlash", func(t *testing.T) {
		v := &FileReference{Source: "foo/bar"}
		err := v.ResolveRelativePath("a/b/")
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}

		if v.FullPath != "a/b/foo/bar" {
			t.Fatalf("unexpected resolved path: %s", v.FullPath)
		}
	})

	t.Run("Abs", func(t *testing.T) {
		v := &FileReference{Source: "foo/bar"}
		err := v.ResolveRelativePath("/a/b/c")
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}

		if v.FullPath != "/a/b/c/foo/bar" {
			t.Fatalf("unexpected resolved path: %s", v.FullPath)
		}
	})

	t.Run("AbsSlash", func(t *testing.T) {
		v := &FileReference{Source: "foo/bar"}
		err := v.ResolveRelativePath("/a/b/c/")
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}

		if v.FullPath != "/a/b/c/foo/bar" {
			t.Fatalf("unexpected resolved path: %s", v.FullPath)
		}
	})
}

func TestFileReference_Resolver(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		v := &TestFileRefProto{}
		v.Ref = &FileReference{Source: "foo/bar"}
		err := WalkProtoMessage(v, RefPath{}, GetPathResolver("/a/b/c"))
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}

		if v.Ref.FullPath != "/a/b/c/foo/bar" {
			t.Fatalf("bad resolved path : %s", v.Ref.FullPath)
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
		v := &TestFileRefProto{Ref: &FileReference{Source: ""}}
		err := WalkProtoMessage(v, RefPath{}, GetPathResolver("/a/b/c"))
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}
	})
}
