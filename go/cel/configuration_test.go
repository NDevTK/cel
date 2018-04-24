// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cel

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"github.com/pkg/errors"
	"strings"
	"testing"
)

func TestConfiguration_Merge_singleTextPb(t *testing.T) {
	var l Configuration

	err := l.Merge("../../examples/schema/ad/one-domain.asset.textpb")
	if err != nil {
		t.Fatalf("failed to read asset example: %s", err)
	}

	if len(l.assetSources) == 0 {
		t.Fatal("expected one source file, but found none")
	}

	if len(l.AssetManifest.AdDomain) != 1 {
		t.Fatalf("expected to find one domain, but found %d", len(l.AssetManifest.AdDomain))
	}

	if l.AssetManifest.AdDomain[0].Name != "foo.example" {
		t.Fatalf("unexpected domain: %#v", l.AssetManifest.AdDomain[0])
	}
}

func TestConfiguration_Merge_singleYaml(t *testing.T) {
	var l Configuration

	err := l.Merge("../../examples/schema/ad/one-domain.asset.yaml")
	if err != nil {
		t.Fatalf("failed to read asset example: %s", err)
	}

	if len(l.assetSources) == 0 {
		t.Fatal("expected one source file, but found none")
	}

	if len(l.AssetManifest.AdDomain) != 1 {
		t.Fatalf("expected to find one domain, but found %d", len(l.AssetManifest.AdDomain))
	}

	if l.AssetManifest.AdDomain[0].Name != "foo.example" {
		t.Fatalf("unexpected domain: %#v", l.AssetManifest.AdDomain[0])
	}

	if len(l.AssetManifest.WindowsMachine) != 2 {
		t.Fatal("repeated object not parsed properly")
	}
}

func TestConfiguration_Merge_duplicate(t *testing.T) {
	var l Configuration

	err := l.Merge("../../examples/schema/ad/one-domain.asset.textpb")
	if err != nil {
		t.Fatalf("failed to read asset example: %#v", err)
	}

	err = l.Merge("../../examples/schema/ad/one-domain.asset.textpb")
	if errors.Cause(err) != ConfigurationAlreadyLoadedError {
		t.Fatalf("duplicate file was not detected")
	}
}

func TestConfiguration_Merge_nonexistent(t *testing.T) {
	var l Configuration

	err := l.Merge("../../examples/schema/ad/one-domain.asset.textpbnotreally")
	if err == nil || !strings.Contains(errors.Cause(err).Error(), "no such file or directory") {
		t.Fail()
	}
}

func TestConfiguration_Merge_noschema(t *testing.T) {
	var l Configuration

	err := l.MergeContents("../../examples/schema/ad/one-domain.textpb", []byte("asdf"))
	if errors.Cause(err) != IncorrectFilenameFormatError {
		t.Fail()
	}
}

func TestConfiguration_Merge_unsupportedType(t *testing.T) {
	var l Configuration

	err := l.MergeContents("../../examples/schema/ad/one-domain.asset.abc", []byte("asdf"))
	if errors.Cause(err) != IncorrectFilenameFormatError {
		t.Fail()
	}
}

func TestConfiguration_Merge_cantParse(t *testing.T) {
	var l Configuration

	err := l.MergeContents("../../examples/schema/ad/one-domain.asset.textpb", []byte("asdf"))
	if err == nil || !strings.Contains(errors.Cause(err).Error(), "unknown field name \"asdf\"") {
		t.Fatal(err)
	}
}

func TestConfiguration_Merge_twoDomains(t *testing.T) {
	var l Configuration

	err := l.Merge("../../examples/schema/ad/one-domain.asset.textpb")
	if err != nil {
		t.Fatalf("failed to read asset example: %#v", err)
	}

	err = l.Merge("../../examples/schema/ad/two-domains.asset.textpb")
	if err != nil {
		t.Fatalf("failed to read asset example: %#v", err)
	}

	if len(l.assetSources) != 2 {
		t.Fatalf("unexpected number of source files. Expected 2, found %d", len(l.assetSources))
	}

	if len(l.AssetManifest.GetAdDomain()) != 3 {
		t.Fatalf("unexpected number of domains. Expected 3, found %d", len(l.AssetManifest.GetAdDomain()))
	}
}

func TestConfiguration_Merge_oneHost(t *testing.T) {
	var l Configuration

	err := l.Merge("../../examples/schema/ad/one-domain.host.textpb")
	if err != nil {
		t.Fatalf("failed to read host example: %#v", err)
	}

	if len(l.hostSources) == 0 {
		t.Fatal("expected one source file, but found none")
	}

	if l.HostEnvironment.Project == nil {
		t.Fatalf("project entry failed to load")
	}

	if l.HostEnvironment.Project.Name != "my-test-gcp-project" {
		t.Fatalf("project name incorrect")
	}

	err = l.Merge("../../examples/schema/ad/one-domain.host.textpb")
	if errors.Cause(err) != ConfigurationAlreadyLoadedError {
		t.Fatalf("duplicate file was not detected")
	}
}

func TestConfiguration_Merge_assetAndHost(t *testing.T) {
	var l Configuration

	err := l.Merge("../../examples/schema/ad/one-domain.host.textpb")
	if err != nil {
		t.Fatalf("failed to read host example: %#v", err)
	}

	if len(l.hostSources) == 0 {
		t.Fatal("expected one source file, but found none")
	}

	if l.HostEnvironment.Project == nil {
		t.Fatalf("project entry failed to load")
	}

	if l.HostEnvironment.Project.Name != "my-test-gcp-project" {
		t.Fatalf("project name incorrect")
	}

	err = l.Merge("../../examples/schema/ad/one-domain.asset.textpb")
	if err != nil {
		t.Fatalf("failed to read asset example: %#v", err)
	}

	err = l.Merge("../../examples/schema/ad/two-domains.asset.textpb")
	if err != nil {
		t.Fatalf("failed to read asset example: %#v", err)
	}

	if len(l.assetSources) != 2 {
		t.Fatalf("unexpected number of source files. Expected 2, found %d", len(l.assetSources))
	}

	if len(l.AssetManifest.GetAdDomain()) != 3 {
		t.Fatalf("unexpected number of domains. Expected 3, found %d", len(l.AssetManifest.GetAdDomain()))
	}
}

func TestConfiguration_Validate(t *testing.T) {
	var l Configuration

	errl := common.AppendErrorList(nil,
		l.Merge("../../examples/schema/ad/one-domain.host.textpb"),
		l.Merge("../../examples/schema/ad/one-domain.asset.textpb"),
		l.Merge("../../schema/gcp/builtins.host.textpb"))
	if errl != nil {
		t.Fatal(errl)
	}

	err := l.Validate()
	if err != nil {
		t.Fatal(err)
	}

	// Once sealed, Merge() no longer allows changes.
	err = l.Merge("../../examples/schema/ad/two-domains.asset.textpb")
	if errors.Cause(err) != ConfigurationSealedError {
		t.Fatal("merges allowed after validation", err)
	}
}

func TestConfiguration_ValidateBuiltins(t *testing.T) {
	var l Configuration

	err := l.Merge("../../schema/gcp/builtins.host.textpb")
	if err != nil {
		t.Fatal(err)
	}

	// This shouldn't be necessary, but by merging these two, we get a well
	// formed set of project settings. Otherwise validation will fail since
	// project and log parameters are required.
	err = l.Merge("../../examples/schema/ad/one-domain.host.textpb")
	if err != nil {
		t.Fatal(err)
	}

	err = l.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
