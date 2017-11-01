// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	pd "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/pkg/errors"
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
	c.ResolvedPath = filepath.Clean(filepath.Join(base_path, c.Path))
	return nil
}

func GetPathResolver(base_path string) WalkProtoFunc {
	return func(av reflect.Value, p *pd.FieldDescriptorProto) error {
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
