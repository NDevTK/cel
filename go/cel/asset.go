// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cel

import ()

// Asset is the foundation for a deployment. It represents something that can
// be deployed, either directly by the deployer, or indirectly via scripts that
// are generated and then executed within a lab.
//
// All the properties exposed in this interface should be intrinsic and
// constant. I.e. they should always return the same values.
type Asset interface {
	// Namespace returns a string identifying a namespace within which the name
	// of the asset is unique.
	//
	// The string returned by Namespace() must satisfy IsRFC1035Label().
	Namespace() string

	// Id returns a string that identifies this asset within the namespace
	// specified by Namespace().
	//
	// The string returned by Id() must satisfy IsRFC1035Label().
	Id() string

	// FullName returns a name combining Id() and Namespace(). This can be used
	// as a convenience for logging etc. but is less efficient than using
	// Namespace() and Id() directly.
	//
	// TODO(asanka): Rename this to CombinedName(). FullName() is what we use
	// in assets to specify names that may be invalid or otherwise exceed the
	// limits set forth for regular names.
	FullName() string

	// DependsOn() returns a list of assets that must be successfully resolved
	// prior to resolving this asset. The ordering of the assets is not
	// significant.
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

	// PermanentAsset method is only used for identifying this interface and to
	// prevent a ResolvableAsset from being misidentified as a PermanentAsset
	// during a type assertion.
	IsPermanentAsset()
}

// BaseNamedAsset is a convenient struct which can be embedded in a struct to
// make it an Asset.
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
