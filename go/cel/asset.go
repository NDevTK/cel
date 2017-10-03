// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cel

import (
	"cloud.google.com/go/logging"
)

type Asset interface {
	Namespace() string
	Id() string
	FullName() string
	DependsOn() []Asset
}

type Checker interface {
	Check(s *Session) error
}

type ResolvableAsset interface {
	Asset
	Checker

	Resolve(s *Session) error
	Purge(s *Session) error
}

type ScriptAsset interface {
	Asset

	GenerateScript(s *Session) error
}

type PermanentAsset interface {
	Asset
	Checker

	PermanentAsset()
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
	Namespace string
	Id        string
	Error     string
}

func (d assetLogEntry) Entry(s logging.Severity) logging.Entry {
	return logging.Entry{
		Severity: s,
		Payload:  d}
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
