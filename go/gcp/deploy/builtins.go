// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"path"
)

const builtinsFileName string = "/deployment/gcp-builtins.host.textpb"

// GetBuiltinHostEnvironment returns the name and data for the build time
// embedded builtin host environment resources.
//
// The name can be used to infer the encoding. The schema is always
// HostEnvironment.
func GetBuiltinHostEnvironment() (name string, data []byte) {
	return path.Base(builtinsFileName), _escFSMustByte(false, builtinsFileName)
}
