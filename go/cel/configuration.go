// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cel

import (
	"chromium.googlesource.com/enterprise/cel/go/asset"
	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/host"
	"chromium.googlesource.com/enterprise/cel/go/lab"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"io/ioutil"
	"path/filepath"
)

var (
	AssetRootPath = common.RefPathFromComponents("asset")
	HostRootPath  = common.RefPathFromComponents("host")
	LabRootPath   = common.RefPathFromComponents("lab")
)

// Configuration combines assets and host configuration from multiple source
// files. This is the recommended method for consuming asset and host
// manifests.
type Configuration struct {
	// The combined asset manifest. Note that unfortunately using the asset
	// manifest in this manner loses valuable information about which source
	// files introduced the asset.
	AssetManifest asset.AssetManifest `json:"asset"`

	// The combined host environment.
	HostEnvironment host.HostEnvironment `json:"host"`

	// Lab. Lab? Lab!
	Lab lab.Lab `json:"lab"`

	// Runtime resources.
	Resources host.RuntimeSupport `json:"-"`

	// Paths of source files that were merged into AssetManifest. The mapping
	// is from the absolute path to the data that was loaded from it.
	assetSources map[string]*asset.AssetManifest

	// Paths of source files that were merged into HostEnvironment. The mapping
	// is from the absolute path to the data that was loaded from it.
	hostSources map[string]*host.HostEnvironment

	// Contains a cross referenced view of the combined asset manifest and host
	// environment.
	references common.Namespace

	// sealed is a weak signal that no more source files should be added to
	// this configuration. It's set automatically after a Validate() call.
	sealed bool
}

// GetNamespace returns the loaded and validated namespace for this Configuration.
func (l *Configuration) GetNamespace() *common.Namespace {
	return &l.references
}

// MergeAssets merges a text format AssetManifest into this Configuration.
// Any errors will be propagated via |error| and will be annotated with the
// name of the configuration file that generated the error.
//
// See /schema/asset/asset_manifest.proto
func (l *Configuration) MergeAssets(filename string) error {
	// If you see this error, it means that an attempt was made to load another
	// configuration file after it was sealed. Sealing happens when Validate()
	// is called.
	if l.sealed {
		return NewConfigurationError(filename, false, errors.Errorf("configuration is already sealed"))
	}

	filename, err := filepath.Abs(filename)
	if err != nil {
		return errors.Wrapf(err, "can't determine absolute path for %s", filename)
	}

	if l.assetSources == nil {
		l.assetSources = make(map[string]*asset.AssetManifest)
	}

	if _, ok := l.assetSources[filename]; ok {
		return NewConfigurationError(filename, true, nil)
	}

	var a asset.AssetManifest
	err = unmarshallTextProtoFromFile(filename, AssetRootPath, &a)
	if err != nil {
		return NewConfigurationError(filename, false, err)
	}

	proto.Merge(&l.AssetManifest, &a)
	l.assetSources[filename] = &a

	return nil
}

// MergeHost merges a text format HostEnvironment into this Configuration.
// Any errors will be propagated via |error| and will be annotated with the
// name of the configuration file that generated the error.
//
// See /schema/host/host_environment.proto
func (l *Configuration) MergeHost(filename string) error {
	// If you see this error, it means that an attempt was made to load another
	// configuration file after it was sealed. Sealing happens when Validate()
	// is called.
	if l.sealed {
		return NewConfigurationError(filename, false, errors.Errorf("configuration is already sealed"))
	}

	filename, err := filepath.Abs(filename)
	if err != nil {
		return errors.Wrapf(err, "can't determine absolute path for %s", filename)
	}

	if l.hostSources == nil {
		l.hostSources = make(map[string]*host.HostEnvironment)
	}

	if _, ok := l.hostSources[filename]; ok {
		return NewConfigurationError(filename, true, nil)
	}

	var h host.HostEnvironment
	err = unmarshallTextProtoFromFile(filename, HostRootPath, &h)
	if err != nil {
		return err
	}

	proto.Merge(&l.HostEnvironment, &h)
	l.hostSources[filename] = &h
	return nil
}

// Merge attempts to parse |filename| as an asset manifest, and failing that, a
// host environment.
func (l *Configuration) Merge(filename string) error {
	if l.sealed {
		return NewConfigurationError(filename, false, errors.Errorf("configuration is already sealed"))
	}

	if err := l.MergeAssets(filename); err == nil {
		return nil
	}

	return l.MergeHost(filename)
}

// Validate ensures that the cross references between assets and the host
// environment are sound.
func (l *Configuration) Validate() error {
	// Validate() is idempotent and should be albe to be called multiple times.
	// Hence not checking whether sealed is already true.
	l.sealed = true

	l.HostEnvironment.Resources = &l.Resources
	l.references.Graft(&l.AssetManifest, AssetRootPath)
	l.references.Graft(&l.HostEnvironment, HostRootPath)
	l.references.Graft(&l.Lab, LabRootPath)

	err_list := []error{
		common.ValidateProto(&l.AssetManifest, AssetRootPath),
		common.ValidateProto(&l.HostEnvironment, HostRootPath),
	}

	l.references.VisitUnresolved(common.EmptyPath, func(v common.UnresolvedValue) bool {
		if _, ok := v.Value.(common.UnresolvedValue_Placeholder); ok {
			err_list = append(err_list, errors.New(v.String()))
		}
		return true
	})

	return common.WrapErrorList(common.AppendErrorList(nil, err_list...))
}

// unmarshallTextProtoFromFile does what its name implies. The file is assumed
// to contain text format protobuf data that can be marshalled into |m|.
func unmarshallTextProtoFromFile(filename string, p common.RefPath, m proto.Message) (err error) {
	defer common.Action(&err, "reading text format protobuf file \"%s\"", filename)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	err = proto.UnmarshalText(string(data), m)
	if err != nil {
		return
	}

	// Paths should be resolved pretty much as soon as the file is loaded.
	return common.WalkProtoMessage(m, p, common.GetPathResolver(filepath.Dir(filename)))
}
