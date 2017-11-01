// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cel

import (
	"chromium.googlesource.com/enterprise/cel/go/asset"
	"chromium.googlesource.com/enterprise/cel/go/host"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"io/ioutil"
	"path/filepath"
)

// LabConfiguration combines assets and host configuration from multiple source
// files. This is the recommended method for consuming asset and host
// manifests.
type LabConfiguration struct {
	// The combined asset manifest. Note that unfortunately using the asset
	// manifest in this manner loses valuable information about which source
	// files introduced the asset.
	AssetManifest asset.AssetManifest

	// The combined host environment.
	HostEnvironment host.HostEnvironment

	// Paths of source files that were merged into AssetManifest. The mapping
	// is from the absolute path to the data that was loaded from it.
	AssetSourceFiles map[string]*asset.AssetManifest

	// Paths of source files that were merged into HostEnvironment. The mapping
	// is from the absolute path to the data that was loaded from it.
	HostSourceFiles map[string]*host.HostEnvironment
}

// MergeAssets merges a text format AssetManifest into this LabConfiguration.
// Any errors will be propagated via |error| and will be annotated with the
// name of the configuration file that generated the error.
//
// See //schema/asset/asset_manifest.proto
func (l *LabConfiguration) MergeAssets(filename string) error {
	filename, err := filepath.Abs(filename)
	if err != nil {
		return errors.Wrapf(err, "can't determine absolute path for %s", filename)
	}

	if l.AssetSourceFiles == nil {
		l.AssetSourceFiles = make(map[string]*asset.AssetManifest)
	}

	if _, ok := l.AssetSourceFiles[filename]; ok {
		return NewConfigurationError(filename, true, nil)
	}

	var a asset.AssetManifest
	err = unmarshallTextProtoFromFile(filename, &a)
	if err != nil {
		return NewConfigurationError(filename, false, err)
	}

	proto.Merge(&l.AssetManifest, &a)
	l.AssetSourceFiles[filename] = &a

	return nil
}

// MergeHost merges a text format HostEnvironment into this LabConfiguration.
// Any errors will be propagated via |error| and will be annotated with the
// name of the configuration file that generated the error.
//
// See //schema/host/host_environment.proto
func (l *LabConfiguration) MergeHost(filename string) error {
	filename, err := filepath.Abs(filename)
	if err != nil {
		return errors.Wrapf(err, "can't determine absolute path for %s", filename)
	}

	if l.HostSourceFiles == nil {
		l.HostSourceFiles = make(map[string]*host.HostEnvironment)
	}

	if _, ok := l.HostSourceFiles[filename]; ok {
		return NewConfigurationError(filename, true, nil)
	}

	var h host.HostEnvironment
	err = unmarshallTextProtoFromFile(filename, &h)
	if err != nil {
		return err
	}

	proto.Merge(&l.HostEnvironment, &h)
	l.HostSourceFiles[filename] = &h
	return nil
}

func (l *LabConfiguration) Validate() error {
	return nil
}

// unmarshallTextProtoFromFile does what its name implies. The file is assumed
// to contain text format protobuf data that can be marshalled into |m|.
func unmarshallTextProtoFromFile(filename string, m proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.Wrapf(err, "while reading text format protobuf file %s", filename)
	}
	return proto.UnmarshalText(string(data), m)
}
