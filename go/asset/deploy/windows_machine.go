// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"chromium.googlesource.com/enterprise/cel/go/asset"
	"chromium.googlesource.com/enterprise/cel/go/common"
)

type windowsMachine struct{}

func (*windowsMachine) ResolvePreManifestCompletion(ctx common.Context, m *asset.WindowsMachine) error {
	features := m.GetWindowsFeature()
	if features == nil {
		features = []string{}
	}

	// TODO(asanka): The list of features should be augmented based on which
	// services are hosted on this VM.
	return ctx.Publish(m, "all_features", features)
}

func init() {
	common.RegisterResolverClass(&windowsMachine{})
}
