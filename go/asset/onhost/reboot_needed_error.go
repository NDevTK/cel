// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

import "github.com/pkg/errors"

// The error returned by on-host resolvers to indicate that reboot is needed.
var ErrRebootNeeded = errors.New("reboot needed")
