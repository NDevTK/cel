// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cel

import (
	"chromium.googlesource.com/enterprise/cel/go/asset"
	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/host"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"io/ioutil"
	"path/filepath"
)

var (
	AssetRootPath = common.RefPath{"asset"}
	HostRootPath  = common.RefPath{"host"}
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

	// Paths of source files that were merged into AssetManifest. The mapping
	// is from the absolute path to the data that was loaded from it.
	AssetSourceFiles map[string]*asset.AssetManifest `json:"-"`

	// Paths of source files that were merged into HostEnvironment. The mapping
	// is from the absolute path to the data that was loaded from it.
	HostSourceFiles map[string]*host.HostEnvironment `json:"-"`

	// Contains cross references between assets and hosts as well as
	// constructed assets.
	References common.References `json:"-"`

	// Collection of all assets that need to be resolved in this configuration.
	Assets common.Assets `json:"-"`

	// sealed is a weak signal that no more source files should be added to
	// this configuration. It's set automatically after a Validate() call.
	sealed bool
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

	if l.AssetSourceFiles == nil {
		l.AssetSourceFiles = make(map[string]*asset.AssetManifest)
	}

	if _, ok := l.AssetSourceFiles[filename]; ok {
		return NewConfigurationError(filename, true, nil)
	}

	var a asset.AssetManifest
	err = unmarshallTextProtoFromFile(filename, AssetRootPath, &a)
	if err != nil {
		return NewConfigurationError(filename, false, err)
	}

	proto.Merge(&l.AssetManifest, &a)
	l.AssetSourceFiles[filename] = &a

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

	if l.HostSourceFiles == nil {
		l.HostSourceFiles = make(map[string]*host.HostEnvironment)
	}

	if _, ok := l.HostSourceFiles[filename]; ok {
		return NewConfigurationError(filename, true, nil)
	}

	var h host.HostEnvironment
	err = unmarshallTextProtoFromFile(filename, HostRootPath, &h)
	if err != nil {
		return err
	}

	proto.Merge(&l.HostEnvironment, &h)
	l.HostSourceFiles[filename] = &h
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

	l.References.AddSource(&l.AssetManifest, AssetRootPath)
	l.References.AddSource(&l.HostEnvironment, HostRootPath)

	err_list := []error{
		l.References.Resolve(common.ResolutionSkipOutputs),
		common.InvokeValidate(&l.AssetManifest, AssetRootPath),
		common.InvokeValidate(&l.HostEnvironment, HostRootPath),
	}

	l.References.Unresolved.Visit(func(p common.RefPath, i interface{}) bool {
		u := i.(*common.UnresolvedReference)
		// Output fields are expected to be unresolved at this point.
		if u.IsOutput {
			return true
		}
		err_list = append(err_list, errors.Errorf("unresolved reference to \"%s\" from \"%s\"",
			u.To.String(), u.From.String()))
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
