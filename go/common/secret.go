// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"chromium.googlesource.com/enterprise/cel/go/schema"
	commonpb "chromium.googlesource.com/enterprise/cel/go/schema/common"
)

func validateSecret(s *commonpb.Secret) error { return nil }

func init() {
	schema.RegisterValidateFunction(validateSecret)
}
