// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"chromium.googlesource.com/enterprise/cel/go/common"
	"github.com/ghodss/yaml"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

type Resource struct {
	Name       string        `json:"name"`
	Type       string        `json:"type"`
	Properties proto.Message `json:"properties"`
}

type resourceList struct {
	Resources []Resource `json:"resources"`
}

type DeploymentManifest struct {
	refs   map[proto.Message]string
	schema resourceList
	counts map[string]int
}

func (d *DeploymentManifest) Ref(m proto.Message) string {
	if n, ok := d.refs[m]; ok {
		return n
	}

	t := reflect.ValueOf(m)
	var base string
	var named bool
	if nm, ok := m.(common.NamedProto); ok {
		base = nm.GetName()
		named = true
	} else {
		base = strings.ToLower(t.Elem().Type().Name())
	}

	var name string
	if c, ok := d.counts[base]; ok {
		c += 1
		d.counts[base] = c
		name = fmt.Sprintf("%s-%d", base, c)
	} else {
		d.counts[base] = 1
		if named {
			name = fmt.Sprintf("%s", base)
		} else {
			name = fmt.Sprintf("%s-1", base)
		}
	}
	d.refs[m] = name
	return name
}

func (d *DeploymentManifest) Emit(m proto.Message, v proto.Message) error {
	tn, ok := getResourceType(v)
	if !ok {
		ty := reflect.TypeOf(v)
		return errors.Errorf("resource type not found for type (%s)%s", ty.PkgPath(), ty.Name())
	}

	var name string
	if m != nil {
		name = d.Ref(m)
	} else {
		name = d.Ref(v)
	}

	r := Resource{
		Name:       name,
		Type:       tn,
		Properties: v,
	}

	d.schema.Resources = append(d.schema.Resources, r)
	return nil
}

func (d *DeploymentManifest) GetResources() []Resource {
	return d.schema.Resources
}

func (d *DeploymentManifest) GetYaml() ([]byte, error) {
	return yaml.Marshal(&d.schema)
}

var once sync.Once
var manifestSingleon *DeploymentManifest

func newDeploymentManifest() *DeploymentManifest {
	return &DeploymentManifest{
		refs:   make(map[proto.Message]string),
		counts: make(map[string]int),
	}
}

func GetDeploymentManifest() *DeploymentManifest {
	once.Do(func() {
		manifestSingleon = newDeploymentManifest()
	})
	return manifestSingleon
}
