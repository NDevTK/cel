// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"reflect"

	"chromium.googlesource.com/enterprise/cel/go/gcp/compute"
	"chromium.googlesource.com/enterprise/cel/go/gcp/onhost"
	"github.com/golang/protobuf/proto"
)

var typeToResourceName = map[reflect.Type]string{
	reflect.TypeOf(&compute.Network{}):                    "compute.beta.network",
	reflect.TypeOf(&compute.Instance{}):                   "compute.beta.instance",
	reflect.TypeOf(&onhost.RuntimeConfigConfigVariable{}): "runtimeconfig.v1beta1.variable",
	reflect.TypeOf(&onhost.RuntimeConfigConfig{}):         "runtimeconfig.v1beta1.config",
}

func getResourceType(m proto.Message) (string, bool) {
	v, ok := typeToResourceName[reflect.TypeOf(m)]
	return v, ok
}
