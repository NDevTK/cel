// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
)

func IsRestarting(logger common.Logger) bool {
	logger.Warning(common.MakeStringer("IsRestarting: Not implemented on Linux. Returning false."))
	return false
}
