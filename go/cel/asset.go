// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cel

import ()

// Asset is the foundation for a deployment.
type Asset interface {
	Namespace() string
	Id() string
	FullName() string
	DependsOn() []Asset
}

type Checker interface {
	Check() error
}

type Resolver interface {
	Resolve() error
	Purge() error
}

type ScriptGenerator interface {
	GenerateScript() error
}

type ResolvableAsset interface {
	Asset
	Checker
	Resolver
}

type ScriptAsset interface {
	Asset
	ScriptGenerator
}

type PermanentAsset interface {
	Asset
	Checker

	// PermanentAsset method is only used for identifying this interface.
	IsPermanentAsset()
}

type BaseNamedAsset struct {
	namespace  string
	id         string
	depends_on []Asset
}

func (s *BaseNamedAsset) Namespace() string {
	return s.namespace
}

func (s *BaseNamedAsset) Id() string {
	return s.id
}

func (s *BaseNamedAsset) FullName() string {
	return s.namespace + ":" + s.id
}

func (s *BaseNamedAsset) DependsOn() []Asset {
	return s.depends_on
}

type assetLogEntry struct {
	Namespace string `json:"namespace"`
	Id        string `json:"id"`
	Error     string `json:"error_string"`
}

func LogAssetError(l Logger, a Asset, err error) {
	if l == nil {
		return
	}
	l.LogError(assetLogEntry{
		Namespace: a.Namespace(),
		Id:        a.Id(),
		Error:     err.Error()})
}
