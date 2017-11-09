// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

import (
	"chromium.googlesource.com/enterprise/cel/go/asset"
	"chromium.googlesource.com/enterprise/cel/go/common"
	"context"
)

type GcpHost struct {
}

func (g *GcpHost) ProvisionMachine(ctx context.Context, p common.RefPath, m *asset.Machine) error {
	return nil
}
