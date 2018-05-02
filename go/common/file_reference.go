// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"archive/zip"
	"bytes"
	"context"
	pd "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
)

func (c *FileReference) Validate() error {
	if len(c.Source) == 0 {
		return errors.New("'source' is required")
	}
	if filepath.IsAbs(c.Source) {
		return errors.New("'source' cannot be absolute")
	}
	if strings.Contains(c.Source, "\\") {
		return errors.New("'source' cannot contain backslashes")
	}
	components := strings.Split(c.Source, "/")
	for _, c := range components {
		if c == ".." {
			return errors.New("'source' contains parent path references")
		}
	}
	return nil
}

func (c *FileReference) ResolveRelativePath(base_path string) error {
	if c.Source == "" {
		errors.New("source is empty")
	}
	c.FullPath = filepath.Clean(filepath.Join(base_path, c.Source))
	return nil
}

func (c *FileReference) Store(ctx context.Context, o ObjectStore) (err error) {
	defer Action(&err, "storing FileReference with source \"%s\"", c.FullPath)

	if c.FullPath == "" {
		return errors.Errorf("path to \"%s\" is not resolved", c.Source)
	}

	fi, err := os.Stat(c.FullPath)
	if err != nil {
		return
	}

	if fi.IsDir() {
		return c.storeArchive(ctx, o)
	}
	return c.storeFile(ctx, o)
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

func (c *FileReference) storeFile(ctx context.Context, o ObjectStore) (err error) {
	defer Action(&err, "storing file at \"%s\"", c.FullPath)

	contents, err := ioutil.ReadFile(c.FullPath)
	if err != nil {
		return
	}
	c.ResolvedType = FileReference_FILE
	return c.storeBlob(ctx, o, contents)
}

func (c *FileReference) storeBlob(ctx context.Context, o ObjectStore, contents []byte) (err error) {
	c.Integrity = IntegrityToken(contents)
	c.ObjectReference, err = o.PutObject(contents)
	return
}

func (c *FileReference) storeArchive(ctx context.Context, o ObjectStore) (err error) {
	defer Action(&err, "storing directory at \"%s\"", c.FullPath)

	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)

	err = addDirectoryToZip(w, c.FullPath, "")
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}

	contents := buf.Bytes()
	c.ResolvedType = FileReference_ZIP_ARCHIVE
	return c.storeBlob(ctx, o, contents)
}

func addDirectoryToZip(w *zip.Writer, fullPath, base string) error {
	files, err := ioutil.ReadDir(fullPath)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() {
			err = addDirectoryToZip(w, filepath.Join(fullPath, f.Name()), path.Join(base, f.Name()))
			if err != nil {
				return err
			}
			continue
		}

		if !f.Mode().IsRegular() {
			continue
		}

		fw, err := w.Create(path.Join(base, f.Name()))
		if err != nil {
			return err
		}

		fr, err := os.Open(filepath.Join(fullPath, f.Name()))
		if err != nil {
			return err
		}

		_, err = io.Copy(fw, fr)
		fr.Close()

		if err != nil {
			return err
		}
	}
	return nil
}
