// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

// base struct for windows instance implementations
type windowsInstanceBase struct {
	// the windows version
	winVersion windowsVersion
}

func (instance *windowsInstanceBase) SetWindowsVersion(ver windowsVersion) {
	instance.winVersion = ver
}
func (instance *windowsInstanceBase) GetWindowsVersion() windowsVersion {
	return instance.winVersion
}
