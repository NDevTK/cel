// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"archive/zip"
	"bytes"
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
	if c.Source == "" && c.ObjectReference == "" {
		return errors.New("'source' is required")
	}
	if c.Source != "" {
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
	}
	return nil
}

func (c *FileReference) ResolveRelativePath(basePath string) error {
	if c.Source == "" {
		errors.New("source is empty")
	}
	c.FullPath = filepath.Clean(filepath.Join(basePath, c.Source))
	return nil
}

func (c *FileReference) Store(ctx Context) (err error) {
	defer Action(&err, "storing FileReference with source \"%s\"", c.FullPath)

	if c.FullPath == "" {
		return errors.Errorf("path to \"%s\" is not resolved", c.Source)
	}

	fi, err := os.Stat(c.FullPath)
	if err != nil {
		return
	}

	if fi.IsDir() {
		return c.storeArchive(ctx)
	}
	return c.storeFile(ctx)
}

func (c *FileReference) StoreFile(ctx Context, contents []byte) (err error) {
	defer Action(&err, "storing blob of length %d", len(contents))

	c.ResolvedType = FileReference_FILE
	c.FullPath = "/?"
	return c.storeBlob(ctx, contents)
}

func GetPathResolver(basePath string) WalkProtoFunc {
	return func(av reflect.Value, p RefPath, fd *pd.FieldDescriptorProto) (bool, error) {
		if av.Kind() != reflect.Ptr || av.IsNil() || av.Elem().Kind() != reflect.Struct {
			return true, nil
		}

		fr, ok := av.Interface().(*FileReference)
		if !ok {
			return true, nil
		}

		return true, fr.ResolveRelativePath(basePath)
	}
}

func (c *FileReference) storeFile(ctx Context) (err error) {
	defer Action(&err, "storing file at \"%s\"", c.FullPath)

	contents, err := ioutil.ReadFile(c.FullPath)
	if err != nil {
		return
	}
	c.ResolvedType = FileReference_FILE
	return c.storeBlob(ctx, contents)
}

func (c *FileReference) storeBlob(ctx Context, contents []byte) (err error) {
	c.Integrity = IntegrityToken(contents)
	if c.TargetPath != "" {
		last := path.Base(c.TargetPath)
		c.ObjectReference, err = ctx.GetObjectStore().PutNamedObject(last, contents)
	} else {
		c.ObjectReference, err = ctx.GetObjectStore().PutObject(contents)
	}
	return
}

func (c *FileReference) storeArchive(ctx Context) (err error) {
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
	return c.storeBlob(ctx, contents)
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
