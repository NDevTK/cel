// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"fmt"
	"testing"
)

func TestResolve(t *testing.T) {
	A := &Assets{}

	a_asset := &testAsset{BaseNamedAsset{kTestNamespace, "a", nil}, false}
	if err := A.Add(a_asset); err != nil {
		t.Error(err)
	}

	// b depends on a
	b_asset := &testAsset{BaseNamedAsset{kTestNamespace, "b", []Asset{a_asset}}, false}
	if err := A.Add(b_asset); err != nil {
		t.Error(err)
	}

	roots, err := PrepareToResolve(A)
	if err != nil {
		t.Fatal(err)
	}

	if len(roots) != 1 {
		t.Errorf("incorrect number of roots. Expected %d, actual %d", 1, len(roots))
	}

	if A.GetNodeForAsset(b_asset).UnresolvedDependencies != 1 {
		t.Errorf("incorrect number of unresolved dependencies. Expected: %d, Actual: %d",
			1, A.GetNodeForAsset(b_asset).UnresolvedDependencies)
	}

	if len(A.GetNodeForAsset(a_asset).Dependents) != 1 {
		t.Error("incorrect dependents array")
	}

	if A.GetNodeForAsset(a_asset).Dependents[0] != A.GetNodeForAsset(b_asset) {
		t.Error("dependency incorrect")
	}

	if !A.GetNodeForAsset(a_asset).Ready() {
		t.Error("'a' node is not ready to be resolved despite not having any dependencies")
	}

	if A.GetNodeForAsset(b_asset).Ready() {
		t.Errorf("asset b is ready to be resolved despite depending on unresolved asset a")
	}

	if len(A.All) != 2 {
		t.Error("a.All incorrect")
	}

	if err := ResolveAssets(A); err != nil {
		t.Error(err)
	}

	if !a_asset.Resolved {
		t.Error("'a' not resolved")
	}

	if !b_asset.Resolved {
		t.Error("'b' not resolved")
	}
}

func TestResolve_Big(t *testing.T) {
	A := &Assets{}

	const kChainLength = 1000

	assets := make([]Asset, kChainLength)
	for i := 0; i < kChainLength; i += 1 {
		deps := []Asset{}
		if i > 0 {
			deps = []Asset{assets[i-1]}
		}
		assets[i] = &testAsset{BaseNamedAsset{kTestNamespace, fmt.Sprintf("asset-%d", i), deps}, false}
		A.Add(assets[i])
	}

	if err := ResolveAssets(A); err != nil {
		t.Error(err)
	}

	for i := 0; i < kChainLength; i += 1 {
		if !A.GetNodeForAsset(assets[i]).Processed {
			t.Fatalf("asset %s not resolved", assets[i].Id())
		}
	}
}

func TestResolve_Triangle(t *testing.T) {
	A := &Assets{}

	const kChainLength = 1000

	assets := make([]Asset, kChainLength)
	for i := 0; i < kChainLength; i += 1 {
		depends_on := make([]Asset, i)
		if i > 0 {
			copy(depends_on, assets[:i])
		}
		assets[i] = &testAsset{BaseNamedAsset{kTestNamespace, fmt.Sprintf("asset-%d", i), depends_on}, false}
		A.Add(assets[i])
	}

	_, err := PrepareToResolve(A)
	if err != nil {
		t.Fatal(err)
	}

	for i := 1; i < kChainLength; i += 1 {
		if A.GetNodeForAsset(assets[i]).UnresolvedDependencies != int32(i) {
			t.Fatalf("dependency count incorrect.")
		}
	}

	if err = ResolveAssets(A); err != nil {
		t.Error(err)
	}

	for i := 0; i < kChainLength; i += 1 {
		if !A.GetNodeForAsset(assets[i]).Processed {
			t.Fatalf("asset %s not resolved", assets[i].Id())
		}
	}
}

func TestAssetCycles(t *testing.T) {
	A := &Assets{}

	const kAssetCount = 10
	assets := make([]*testAsset, kAssetCount)
	for i := 0; i < kAssetCount; i += 1 {
		assets[i] = &testAsset{BaseNamedAsset{"foo", fmt.Sprintf("asset-%d", i), nil}, false}
		A.Add(assets[i])
	}

	t.Run("no dependents", func(t *testing.T) {
		roots, err := PrepareToResolve(A)
		if err != nil {
			t.Fatalf("a purely disjointed graph cannot have cycles")
		}

		if len(roots) != len(assets) {
			t.Fatalf("a disjointed graph should consider all nodes as roots")
		}
	})

	t.Run("cycle or order 2", func(t *testing.T) {
		assets[0].depends_on = []Asset{assets[1]}
		assets[1].depends_on = []Asset{assets[0]}

		_, err := PrepareToResolve(A)
		if err == nil {
			t.Fatalf("cycle of size 2 not detected")
		}
	})

	t.Run("cycle of order 4", func(t *testing.T) {
		assets[0].depends_on = []Asset{assets[1]}
		assets[1].depends_on = []Asset{assets[2]}
		assets[2].depends_on = []Asset{assets[3]}
		assets[3].depends_on = []Asset{assets[4]}
		assets[5].depends_on = []Asset{assets[4]}
		_, err := PrepareToResolve(A)
		if err != nil {
			t.Fatalf("graph with no cycles detected as a cycle: %s", err.Error())
		}

		assets[3].depends_on = []Asset{assets[0]}
		_, err = PrepareToResolve(A)
		if err == nil {
			t.Fatalf("cycle of size 4 not detected")
		}
	})

	t.Run("triangle", func(t *testing.T) {
		for i := 0; i < kAssetCount; i += 1 {
			if i == 0 {
				assets[i].depends_on = nil
				continue
			}
			deps := make([]Asset, i)
			// need to copy manually since deps[] and assets[] are different types.
			for j := 0; j < i; j += 1 {
				deps[j] = assets[j]
			}
			assets[i].depends_on = deps
		}

		_, err := PrepareToResolve(A)
		if err != nil {
			t.Fatalf("cycle found although there is none: %s", err.Error())
		}
	})
}
