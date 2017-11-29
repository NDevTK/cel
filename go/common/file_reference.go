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

func (c *FileReference) Resolve(ctx context.Context, resolver ResolverService) (err error) {
	defer Action(&err, "resolving FileReference to \"%s\" (local path \"%s\")", c.Path, c.ResolvedLocalPath)
	o, err := ObjectStoreServiceFromContext(ctx)
	if err != nil {
		return
	}

	if c.ResolvedLocalPath == "" {
		return errors.Errorf("path to \"%s\" is not resolved", c.Path)
	}

	fi, err := os.Stat(c.ResolvedLocalPath)
	if err != nil {
		return
	}

	if fi.IsDir() {
		c.ResolvedObject, _, err = c.storeDir(ctx, o, c.ResolvedLocalPath)
	} else {
		c.ResolvedObject, _, err = c.storeFile(ctx, o, c.ResolvedLocalPath)
	}
	return
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

func (c *FileReference) storeFile(ctx context.Context, o ObjectStoreService, path string) (ref string, size int64, err error) {
	defer Action(&err, "storing file at \"%s\"", path)

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	size = int64(len(contents))

	j := NewJobWaiter(c)
	o.PutObject(ctx, contents, &ref, j.New())
	err = j.Join()
	return
}

func (c *FileReference) storeDir(ctx context.Context, o ObjectStoreService, path string) (ref string, size int64, err error) {
	defer Action(&err, "storing directory at \"%s\"", path)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}

	var tree meta.Tree
	tree.File = make([]*meta.FileReference, len(files))

	j := NewJobWaiter(c)
	for i, f := range files {
		fr := &meta.FileReference{
			Name: f.Name(),
			Type: meta.FileReference_FILE}
		new_path := filepath.Join(path, f.Name())
		if f.IsDir() {
			fr.Type = meta.FileReference_DIRECTORY
			go func(result chan<- error) {
				var err error
				fr.Reference, fr.Size, err = c.storeDir(ctx, o, new_path)
				result <- err
			}(j.New())
		} else {
			go func(result chan<- error) {
				var err error
				fr.Reference, fr.Size, err = c.storeFile(ctx, o, new_path)
				result <- err
			}(j.New())
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
	size = int64(len(encoded))

	o.PutObject(ctx, encoded, &ref, j.New())
	err = j.Join()
	return
}
