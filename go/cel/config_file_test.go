// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cel

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"testing"
)

func TestLabConfiguration_MergeAssets_1(t *testing.T) {
	var l LabConfiguration

	err := l.MergeAssets("../../examples/schema/ad/one-domain.asset.textpb")
	if err != nil {
		t.Fatalf("failed to read asset example: %#v", err)
	}

	if len(l.AssetSourceFiles) == 0 {
		t.Fatal("expected one source file, but found none")
	}

	if len(l.AssetManifest.AdDomain) != 1 {
		t.Fatalf("expected to find one domain, but found %d", len(l.AssetManifest.AdDomain))
	}

	if l.AssetManifest.AdDomain[0].Name != "foo.example" {
		t.Fatalf("unexpected domain: %#v", l.AssetManifest.AdDomain[0])
	}

	err = l.MergeAssets("../../examples/schema/ad/one-domain.asset.textpb")
	if err == nil {
		t.Fatalf("duplicate file was not detected")
	}

	if _, ok := err.(*ConfigurationError); !ok {
		t.Fatal("unexpected error type")
	}
}

func TestLabConfiguration_MergeAssets_2(t *testing.T) {
	var l LabConfiguration

	err := l.MergeAssets("../../examples/schema/ad/one-domain.asset.textpb")
	if err != nil {
		t.Fatalf("failed to read asset example: %#v", err)
	}

	err = l.MergeAssets("../../examples/schema/ad/two-domains.asset.textpb")
	if err != nil {
		t.Fatalf("failed to read asset example: %#v", err)
	}

	if len(l.AssetSourceFiles) != 2 {
		t.Fatalf("unexpected number of source files. Expected 2, found %d", len(l.AssetSourceFiles))
	}

	if len(l.AssetManifest.GetAdDomain()) != 3 {
		t.Fatalf("unexpected number of domains. Expected 3, found %d", len(l.AssetManifest.GetAdDomain()))
	}
}

func TestLabConfiguration_MergeHost(t *testing.T) {
	var l LabConfiguration

	err := l.MergeHost("../../examples/schema/ad/one-domain.host.textpb")
	if err != nil {
		t.Fatalf("failed to read host example: %#v", err)
	}

	if len(l.HostSourceFiles) == 0 {
		t.Fatal("expected one source file, but found none")
	}

	if l.HostEnvironment.Project == nil {
		t.Fatalf("project entry failed to load")
	}

	if l.HostEnvironment.Project.Name != "my-test-gcp-project" {
		t.Fatalf("project name incorrect")
	}

	err = l.MergeHost("../../examples/schema/ad/one-domain.host.textpb")
	if err == nil {
		t.Fatalf("duplicate file was not detected")
	}

	if _, ok := err.(*ConfigurationError); !ok {
		t.Fatalf("unexpected error type")
	}
}

func TestLabConfiguration_Merge(t *testing.T) {
	var l LabConfiguration

	err := l.Merge("../../examples/schema/ad/one-domain.host.textpb")
	if err != nil {
		t.Fatalf("failed to read host example: %#v", err)
	}

	if len(l.HostSourceFiles) == 0 {
		t.Fatal("expected one source file, but found none")
	}

	if l.HostEnvironment.Project == nil {
		t.Fatalf("project entry failed to load")
	}

	if l.HostEnvironment.Project.Name != "my-test-gcp-project" {
		t.Fatalf("project name incorrect")
	}

	err = l.Merge("../../examples/schema/ad/one-domain.host.textpb")
	if err == nil {
		t.Fatalf("duplicate file was not detected")
	}

	if _, ok := err.(*ConfigurationError); !ok {
		t.Fatalf("unexpected error type")
	}

	err = l.Merge("../../examples/schema/ad/one-domain.asset.textpb")
	if err != nil {
		t.Fatalf("failed to read asset example: %#v", err)
	}

	err = l.Merge("../../examples/schema/ad/two-domains.asset.textpb")
	if err != nil {
		t.Fatalf("failed to read asset example: %#v", err)
	}

	if len(l.AssetSourceFiles) != 2 {
		t.Fatalf("unexpected number of source files. Expected 2, found %d", len(l.AssetSourceFiles))
	}

	if len(l.AssetManifest.GetAdDomain()) != 3 {
		t.Fatalf("unexpected number of domains. Expected 3, found %d", len(l.AssetManifest.GetAdDomain()))
	}
}

func TestLabConfiguration_Validate(t *testing.T) {
	var l LabConfiguration

	errl := common.AppendErrorList(nil,
		l.Merge("../../examples/schema/ad/one-domain.host.textpb"),
		l.Merge("../../examples/schema/ad/one-domain.asset.textpb"))
	if errl != nil {
		t.Fatal(errl)
	}

	err := l.Validate()
	if err != nil {
		t.Fatal(err)
	}
}

func TestLabConfiguration_ValidateBuiltins(t *testing.T) {
	var l LabConfiguration

	err := l.Merge("../../schema/host/builtins.textpb")
	if err != nil {
		t.Fatal(err)
	}

	err = l.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
