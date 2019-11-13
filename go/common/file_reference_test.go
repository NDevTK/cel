// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	commonpb "chromium.googlesource.com/enterprise/cel/go/schema/common"
	"path/filepath"
	"strings"
	"testing"
)

func TestFileReference_Validate(t *testing.T) {
	p := RefPath{}
	t.Run("Empty", func(t *testing.T) {
		v := &commonpb.FileReference{}
		if err := ValidateProto(v, p); err == nil {
			t.Fail()
		}
	})

	t.Run("Absolute", func(t *testing.T) {
		v := &commonpb.FileReference{Source: "/foo/bar"}
		if err := ValidateProto(v, p); err == nil || !strings.Contains(err.Error(), "cannot be absolute") {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("Backslash", func(t *testing.T) {
		v := &commonpb.FileReference{Source: "foo\\bar"}
		if err := ValidateProto(v, p); err == nil || !strings.Contains(err.Error(), "backslash") {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("ParentRefPrefix", func(t *testing.T) {
		v := &commonpb.FileReference{Source: "../bar"}
		if err := ValidateProto(v, p); err == nil || !strings.Contains(err.Error(), "parent path reference") {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("ParentRefSuffix", func(t *testing.T) {
		v := &commonpb.FileReference{Source: "foo/bar/.."}
		if err := ValidateProto(v, p); err == nil || !strings.Contains(err.Error(), "parent path reference") {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("ParentRefInfix", func(t *testing.T) {
		v := &commonpb.FileReference{Source: "foo/../bar"}
		if err := ValidateProto(v, p); err == nil || !strings.Contains(err.Error(), "parent path reference") {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("OutputParam", func(t *testing.T) {
		v := &commonpb.FileReference{Source: "foo/bar", FullPath: "foo/bar"}
		if err := ValidateProto(v, p); err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("Valid", func(t *testing.T) {
		v := &commonpb.FileReference{Source: "foo/bar"}
		if err := ValidateProto(v, p); err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("NoSource", func(t *testing.T) {
		v := &commonpb.FileReference{}
		if err := ValidateProto(v, p); err == nil || !strings.Contains(err.Error(), "'source' is required") {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("NoSourceWithObjectReference", func(t *testing.T) {
		v := &commonpb.FileReference{ObjectReference: "foo"}
		if err := ValidateProto(v, p); err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}
	})
}

func TestFileReference_Resolve(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		v := &commonpb.FileReference{Source: "foo/bar"}
		err := ResolveRelativePath(v, "")
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}

		if v.FullPath != "foo/bar" {
			t.Fatalf("unexpected resolved path: %s", v.FullPath)
		}
	})

	t.Run("Rel", func(t *testing.T) {
		v := &commonpb.FileReference{Source: "foo/bar"}
		err := ResolveRelativePath(v, "a/b")
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}

		if v.FullPath != "a/b/foo/bar" {
			t.Fatalf("unexpected resolved path: %s", v.FullPath)
		}
	})

	t.Run("RelSlash", func(t *testing.T) {
		v := &commonpb.FileReference{Source: "foo/bar"}
		err := ResolveRelativePath(v, "a/b/")
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}

		if v.FullPath != "a/b/foo/bar" {
			t.Fatalf("unexpected resolved path: %s", v.FullPath)
		}
	})

	t.Run("Abs", func(t *testing.T) {
		v := &commonpb.FileReference{Source: "foo/bar"}
		err := ResolveRelativePath(v, "/a/b/c")
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}

		if v.FullPath != "/a/b/c/foo/bar" {
			t.Fatalf("unexpected resolved path: %s", v.FullPath)
		}
	})

	t.Run("AbsSlash", func(t *testing.T) {
		v := &commonpb.FileReference{Source: "foo/bar"}
		err := ResolveRelativePath(v, "/a/b/c/")
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
		v := &commonpb.TestFileRefProto{}
		v.Ref = &commonpb.FileReference{Source: "foo/bar"}
		err := WalkProtoMessage(v, RefPath{}, GetPathResolver("/a/b/c"))
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}

		if v.Ref.FullPath != "/a/b/c/foo/bar" {
			t.Fatalf("bad resolved path : %s", v.Ref.FullPath)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		v := &commonpb.TestFileRefProto{}
		err := WalkProtoMessage(v, RefPath{}, GetPathResolver("/a/b/c"))
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}
	})

	t.Run("EmptyPath", func(t *testing.T) {
		v := &commonpb.TestFileRefProto{Ref: &commonpb.FileReference{Source: ""}}
		err := WalkProtoMessage(v, RefPath{}, GetPathResolver("/a/b/c"))
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}
	})
}

func TestReference_Store_unnamedBlob(t *testing.T) {
	o := &fakeObjectStore{}
	ctx := &fakeContext{
		objectStore: o,
	}

	v := &commonpb.FileReference{Source: filepath.Join("testdata", "tiny.bin")}
	err := ResolveRelativePath(v, ".")
	if err != nil {
		t.Fatal("unexpected error: ", err)
	}
	err = Store(ctx, v)
	if err != nil {
		t.Fatal("unexpected error: ", err)
	}
	if len(o.Log) != 1 {
		t.Fatal("wrong number of calls recorded")
	}
	if o.Log[0] != "PutObject:abcd" {
		t.Fatal("unexpected log value: ", o.Log[0])
	}
}

func TestFileReference_StoreFile(t *testing.T) {
	o := &fakeObjectStore{}
	ctx := &fakeContext{
		objectStore: o,
	}

	v := &commonpb.FileReference{}
	err := StoreFile(ctx, v, []byte("abcd"))
	if err != nil {
		t.Fatal("unexpected error: ", err)
	}
	if len(o.Log) != 1 {
		t.Fatal("wrong number of calls recorded")
	}
	if o.Log[0] != "PutObject:abcd" {
		t.Fatal("unexpected log value: ", o.Log[0])
	}
}

func TestFileReference_Store_withTarget(t *testing.T) {
	o := &fakeObjectStore{}
	ctx := &fakeContext{
		objectStore: o,
	}

	v := &commonpb.FileReference{
		Source:     filepath.Join("testdata", "tiny.bin"),
		TargetPath: "foo/targetfn.png",
	}
	err := ResolveRelativePath(v, ".")
	if err != nil {
		t.Fatal("unexpected error: ", err)
	}
	err = Store(ctx, v)
	if err != nil {
		t.Fatal("unexpected error: ", err)
	}
	if len(o.Log) != 1 {
		t.Fatal("wrong number of calls recorded")
	}
	if o.Log[0] != "PutNamedObject:targetfn.png:abcd" {
		t.Fatal("unexpected log value: ", o.Log[0])
	}
}
