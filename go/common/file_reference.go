// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"archive/zip"
	"bytes"
	"chromium.googlesource.com/enterprise/cel/go/schema"
	commonpb "chromium.googlesource.com/enterprise/cel/go/schema/common"
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

func validateFileReference(c *commonpb.FileReference) error {
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

func ResolveRelativePath(c *commonpb.FileReference, basePath string) error {
	if c.Source == "" {
		errors.New("source is empty")
	}
	c.FullPath = filepath.Clean(filepath.Join(basePath, c.Source))
	return nil
}

func Store(ctx Context, c *commonpb.FileReference) (err error) {
	defer Action(&err, "storing FileReference with source \"%s\"", c.FullPath)

	if c.FullPath == "" {
		return errors.Errorf("path to \"%s\" is not resolved", c.Source)
	}

	fi, err := os.Stat(c.FullPath)
	if err != nil {
		return
	}

	if fi.IsDir() {
		return storeArchive(ctx, c)
	}
	return storeFile(ctx, c)
}

func StoreFile(ctx Context, c *commonpb.FileReference, contents []byte) (err error) {
	defer Action(&err, "storing blob of length %d", len(contents))

	c.ResolvedType = commonpb.FileReference_FILE
	c.FullPath = "/?"
	return storeBlob(ctx, c, contents)
}

func GetPathResolver(basePath string) WalkProtoFunc {
	return func(av reflect.Value, p RefPath, fd *pd.FieldDescriptorProto) (bool, error) {
		if av.Kind() != reflect.Ptr || av.IsNil() || av.Elem().Kind() != reflect.Struct {
			return true, nil
		}

		fr, ok := av.Interface().(*commonpb.FileReference)
		if !ok {
			return true, nil
		}

		return true, ResolveRelativePath(fr, basePath)
	}
}

func storeFile(ctx Context, c *commonpb.FileReference) (err error) {
	defer Action(&err, "storing file at \"%s\"", c.FullPath)

	contents, err := ioutil.ReadFile(c.FullPath)
	if err != nil {
		return
	}
	c.ResolvedType = commonpb.FileReference_FILE
	return storeBlob(ctx, c, contents)
}

func storeBlob(ctx Context, c *commonpb.FileReference, contents []byte) (err error) {
	c.Integrity = IntegrityToken(contents)
	if c.TargetPath != "" {
		last := path.Base(c.TargetPath)
		c.ObjectReference, err = ctx.GetObjectStore().PutNamedObject(last, contents)
	} else {
		c.ObjectReference, err = ctx.GetObjectStore().PutObject(contents)
	}
	return
}

func storeArchive(ctx Context, c *commonpb.FileReference) (err error) {
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
	c.ResolvedType = commonpb.FileReference_ZIP_ARCHIVE
	return storeBlob(ctx, c, contents)
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

func init() {
	schema.RegisterValidateFunction(validateFileReference)
}
