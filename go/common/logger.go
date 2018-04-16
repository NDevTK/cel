// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"fmt"
)

// Logger is an abstract logging facility. The log entries should be considered
// relatively heavyweight, and can carry large payloads.
//
// Each method corresponds to a severity level, and argument is the payload. It
// can be used as follows:
//
//     var l Logger
//     l.Error(p)
//
// The argument |p| should be JSON serializable into something sensible.
type Logger interface {
	// Debug entries are usually just displayed or logged to a file.
	Debug(v fmt.Stringer)

	// LogInfo logs a INFO event with |v| as the payload. See Logger for details.
	Info(v fmt.Stringer)

	// LogWarning logs a WARNING event with |v| as the payload. See Logger for details.
	Warning(v fmt.Stringer)

	// LogError logs a ERROR event with |v| as the payload. See Logger for details.
	Error(v fmt.Stringer)
}

// sstringer is a string Stringer. Used by MakeStringer() to generate a
// stringer from a string.
type sstringer struct {
	m string
}

func (t *sstringer) String() string {
	return t.m
}

// MakeStringer constucts a stringer from a string. Useful for ad-hoc log
// entries. In general, most heavy-weight tasks should be made available as
// JSON serializable structures. This way the resulting logs will be machine
// readable.
func MakeStringer(format string, v ...interface{}) fmt.Stringer {
	return &sstringer{m: fmt.Sprintf(format, v...)}
}
