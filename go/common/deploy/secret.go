// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
)

type secret struct{}

func (secret) ResolvePreManifestCompletion(ctx common.Context, s *common.Secret) error {
	s.Hardcoded = ""
	s.Final = nil
	return nil
}

func init() {
	common.RegisterResolverClass(secret{})
}
