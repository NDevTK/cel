// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"strings"
	"testing"

	"chromium.googlesource.com/enterprise/cel/go/asset"
	"chromium.googlesource.com/enterprise/cel/go/gcp/compute"
)

func TestDeploymentManifest_Ref(t *testing.T) {
	m := &asset.Machine{}

	d := newDeploymentManifest()
	n := d.Ref(m)
	if !strings.HasPrefix(n, "machine-") {
		t.Errorf("resource name doesn't have correct prefix. Got \"%s\". Want \"machine-\"", n)
	}

	nn := d.Ref(m)
	if nn != n {
		t.Errorf("duplicate id issued for same asset")
	}

	m = &asset.Machine{}
	nn = d.Ref(m)
	if nn == n {
		t.Errorf("same id issued for different asset")
	}
}

func TestDeploymentManifest_Emit(t *testing.T) {
}

func TestDeploymentManifest_GetYaml(t *testing.T) {
	d := newDeploymentManifest()
	d.Emit(nil, &compute.Instance{
		Name:        "foo",
		Description: "description of foo instance",
		Metadata: &compute.Metadata{
			Items: []*compute.Metadata_Items{
				{
					Key:   "this is the key",
					Value: "this is the value",
				},
			},
		},
	})

	y, err := d.GetYaml()
	if err != nil {
		t.Error("unexpected error", err)
	}

	// This test relies on the fact that objects are emitted with sorted field
	// names.
	expected := `resources:
- name: instance-1
  properties:
    description: description of foo instance
    metadata:
      items:
      - key: this is the key
        value: this is the value
    name: foo
  type: compute.beta.instance
`
	if string(y) != expected {
		exp := strings.Split(expected, "\n")
		got := strings.Split(string(y), "\n")

		var diff []string

		// Not a diff.
		for i := 0; i < len(exp); i++ {
			if i >= len(got) {
				diff = append(diff, "- "+exp[i])
				continue
			}

			if got[i] != exp[i] {
				diff = append(diff, "- "+exp[i], "+ "+got[i])
			} else {
				diff = append(diff, "  "+exp[i])
			}
		}

		for i := len(exp); i < len(got); i++ {
			diff = append(diff, "+ "+got[i])
		}

		t.Errorf("Unexpected result. Diff:\n%s", strings.Join(diff, "\n"))
	}
}
