// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"github.com/golang/protobuf/proto"
	"reflect"
)

// ResolverKind represents one kind of resolver. Each kind has a specific
// purpose and is invoked at a specific time during deployment. The resolver is
// applied to all messages that match its associated type.
type ResolverKind reflect.Type

// resolverFunc is an internal type representing the function signature of a
// generic resolver method. All resolver methods are expected to take the same
// set of parameters and return the same type of return value.
type resolverFunc func(ctx Context, m proto.Message) error

// AdditionalDependencyResolver is a type of resolver that adds new
// dependencies to an asset.
//
// While some dependencies can be expressed in the asset and host environment
// schema, others may be easier to calculate. Still others may need to be
// calcuated. In such cases introduce a resolver of this type. During
// execution, ctx.PublishDependency() can be used to add new dependencies.
//
// This resolver kind is invoked before doing any other work that depends on
// inter-resource dependencies. In particular, these resolvers are invoked
// prior to pruning any resources that are not connected to the resource being
// deployed.
type AdditionalDependencyResolver interface {
	ResolveAdditionalDependencies(ctx Context, m proto.Message) error
}

// ImmediateResolver is a type of resolver that resolves an asset during
// deployment time.
//
// "Resolve" in this context refers to the act of ensuring the existence of the
// asset and querying its properties as required by the OUTPUT fields of the
// message. It is expected that once the ImmediateResolver runs, all OUTPUT
// fields of the message are available.
type ImmediateResolver interface {
	ResolveImmediate(ctx Context, m proto.Message) error
}

// GeneratedContentResolver generates content.
//
// Examples of resources that benefit from this type of resolver include assets
// where private keys or random numbers need to be generated and persisted at
// deployment time. E.g. for user passwords.
type GeneratedContentResolver interface {
	ResolveGeneratedContent(ctx Context, m proto.Message) error
}

// IndexedObjectResolver uploads indexed objects.
type IndexedObjectResolver interface {
	ResolveIndexedObjects(ctx Context, m proto.Message) error
}

// PreCompleteManifestResolver is invoked prior to constructing the completed
// asset manifest. This type of resolver is used to cleanup assets prior to
// being included in the final manifest.
type PreCompleteManifestResolver interface {
	ResolvePreManifestCompletion(ctx Context, m proto.Message) error
}

// ConstructedAssetResolver is invoked when resolving constructed assets. These
// are assets that need to be constructed within the hosting environment.
type ConstructedAssetResolver interface {
	ResolveConstructedAssets(ctx Context, m proto.Message) error
}

// OnHostResolver is a resolver type that's invoked exclusively on the VM
// hosting the corresponding asset.
//
// TODO(asanka): Document the semantics of this resolver once finalized.
// Dependency ordering for on-host assets are different from deployment time.
type OnHostResolver interface {
	ResolveOnHost(ctx Context, m proto.Message) error
}

// Keep sorted:

var AdditionalDependencyResolverKind ResolverKind = reflect.TypeOf((*AdditionalDependencyResolver)(nil)).Elem()
var ConstructedAssetResolverKind = reflect.TypeOf((*ConstructedAssetResolver)(nil)).Elem()
var GeneratedContentResolverKind ResolverKind = reflect.TypeOf((*GeneratedContentResolver)(nil)).Elem()
var ImmediateResolverKind ResolverKind = reflect.TypeOf((*ImmediateResolver)(nil)).Elem()
var IndexedObjectResolverKind ResolverKind = reflect.TypeOf((*IndexedObjectResolver)(nil)).Elem()
var OnHostResolverKind ResolverKind = reflect.TypeOf((*OnHostResolver)(nil)).Elem()
var PreCompleteManifestResolverKind = reflect.TypeOf((*PreCompleteManifestResolver)(nil)).Elem()

// Keep sorted:

var allResolverTypes = map[string]ResolverKind{
	"ResolveAdditionalDependencies": AdditionalDependencyResolverKind,
	"ResolveConstructedAssets":      ConstructedAssetResolverKind,
	"ResolveGeneratedContent":       GeneratedContentResolverKind,
	"ResolveImmediate":              ImmediateResolverKind,
	"ResolveIndexedObjects":         IndexedObjectResolverKind,
	"ResolveOnHost":                 OnHostResolverKind,
	"ResolvePreManifestCompletion":  PreCompleteManifestResolverKind,
}
