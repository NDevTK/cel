// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"golang.org/x/sys/windows"
)

const (
	ShuttingDownIndex uintptr = uintptr(0x2000)
)

var user32 *windows.LazyDLL = windows.NewLazySystemDLL("user32.dll")

// We try to control when we reboot with DSC `RebootNodeIfNeeded = $false`, but this doesn't
// apply for configurations that require more than one reboot (e.g. setting up RemoteDesktop).
// In those scenarios, our LCM picks the config back up on boot and can restart while we're
// processing another script and trigger errors.
// IsRestarting is used to forgive/ignore failures that happen when Windows is shutting down.
func IsRestarting(logger common.Logger) bool {
	r1, r2, err := user32.NewProc("GetSystemMetrics").Call(ShuttingDownIndex)
	logger.Info(common.MakeStringer("GetSystemMetrics(SM_SHUTTINGDOWN): r1=%v, r2=%v, err=%v", r1, r2, err))
	return r1 != 0
}
