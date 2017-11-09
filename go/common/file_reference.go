// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"chromium.googlesource.com/enterprise/cel/go/meta"
	"context"
	"github.com/golang/protobuf/proto"
	pd "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

func (c *FileReference) Validate() error {
	if len(c.Path) == 0 {
		return errors.New("'path' is required")
	}
	if filepath.IsAbs(c.Path) {
		return errors.New("'path' cannot be absolute")
	}
	if strings.Contains(c.Path, "\\") {
		return errors.New("'path' cannot contain backslashes")
	}
	components := strings.Split(c.Path, "/")
	for _, c := range components {
		if c == ".." {
			return errors.New("'path' contains parent path references")
		}
	}
	return nil
}

func (c *FileReference) ResolveRelativePath(base_path string) error {
	if c.Path == "" {
		errors.New("path is empty")
	}
	c.ResolvedLocalPath = filepath.Clean(filepath.Join(base_path, c.Path))
	return nil
}

func (c *FileReference) Resolve(ctx context.Context, p RefPath) (err error) {
	Action(&err, "resolving FileReference to \"%s\" (local path \"%s\")", c.Path, c.ResolvedLocalPath)
	o, ok := ObjectStoreServiceFromContext(ctx)
	if !ok {
		return &ServiceNotFoundError{Service: "ObjectStore"}
	}

	if c.ResolvedLocalPath == "" {
		return errors.Errorf("path to \"%s\" is not resolved", c.Path)
	}

	var (
		ref  string
		size int64
	)
	j := NewJobWaiter()
	storePath(ctx, o, c.ResolvedLocalPath, &ref, &size, j.Collect())
	err = j.Join()
	if err == nil {
		c.ResolvedObject = ref
		// TODO(asanka): Publish this reference.
	}
	return
}

func (c *FileReference) Purge(ctx context.Context, p RefPath) error {
	return nil
}

func GetPathResolver(base_path string) WalkProtoFunc {
	return func(av reflect.Value, p RefPath, fd *pd.FieldDescriptorProto) error {
		if av.Kind() != reflect.Ptr || av.IsNil() || av.Elem().Kind() != reflect.Struct {
			return nil
		}

		fr, ok := av.Interface().(*FileReference)
		if !ok {
			return nil
		}

		return fr.ResolveRelativePath(base_path)
	}
}

func storePath(ctx context.Context, o ObjectStoreService, path string, ref *string, size *int64, result chan<- error) {
	fi, err := os.Stat(path)
	if err != nil {
		result <- err
		return
	}

	if fi.IsDir() {
		storeDir(ctx, o, path, ref, size, result)
	} else {
		storeFile(ctx, o, path, ref, size, result)
	}
}

func storeFile(ctx context.Context, o ObjectStoreService, path string, ref *string, size *int64, result chan<- error) {
	var err error
	defer AsyncAction(&err, result, "storing file at \"%s\"", path)

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	j := NewJobWaiter()
	o.PutObject(ctx, contents, ref, j.Collect())
	err = j.Join()
	if err != nil {
		return
	}
	*size = int64(len(contents))
}

func storeDir(ctx context.Context, o ObjectStoreService, path string, ref *string, size *int64, result chan<- error) {
	var err error
	defer AsyncAction(&err, result, "storing directory at \"%s\"", path)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}

	var tree meta.Tree
	tree.File = make([]*meta.FileReference, len(files))

	j := NewJobWaiter()
	for i, f := range files {
		fr := &meta.FileReference{
			Name: f.Name(),
			Type: meta.FileReference_FILE}
		new_path := filepath.Join(path, f.Name())
		if f.IsDir() {
			fr.Type = meta.FileReference_DIRECTORY
			storeDir(ctx, o, new_path, &fr.Reference, &fr.Size, j.Collect())
		} else {
			storeFile(ctx, o, new_path, &fr.Reference, &fr.Size, j.Collect())
		}
		tree.File[i] = fr
	}
	err = j.Join()
	if err != nil {
		return
	}

	tree.Canonicalize()
	encoded, err := proto.Marshal(&tree)
	if err != nil {
		return
	}

	*size = int64(len(encoded))

	o.PutObject(ctx, encoded, ref, j.Collect())
	err = j.Join()
}
