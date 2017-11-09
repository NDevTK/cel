// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

// Logger is an abstract logging facility. The log entries should be considered
// relatively heavyweight, and can carry large payloads.
//
// Each method corresponds to a severity level, and argument is the payload. It
// can be used as follows:
//
//     var l Logger
//     l.LogError(p)
//
// The argument |p| should be JSON serializable into something sensible.
type Logger interface {
	// LogDebug logs a DEBUG event with |v| as the payload. See Logger for details.
	LogDebug(v interface{})

	// LogInfo logs a INFO event with |v| as the payload. See Logger for details.
	LogInfo(v interface{})

	// LogWarning logs a WARNING event with |v| as the payload. See Logger for details.
	LogWarning(v interface{})

	// LogError logs a ERROR event with |v| as the payload. See Logger for details.
	LogError(v interface{})
}
