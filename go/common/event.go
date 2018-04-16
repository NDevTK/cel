// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"fmt"
)

type workflowState string

const actionBegin workflowState = "BEGIN"
const actionFailed workflowState = "FAIL"
const actionEnd workflowState = "END"

// event represents a log event. It serializes with a common and hopefully
// consistent format. This is used by LoggedAction.
type event struct {
	Step    string        `json:"m,omitempty"`
	State   workflowState `json:"s,omitempty"`
	Error   error         `json:"err,omitempty"`
	Details interface{}   `json:"details,omitempty"`
}

func (w event) String() string {
	if w.Error != nil {
		return fmt.Sprintf("[%5s] %s", w.State, w.Error.Error())
	}
	return fmt.Sprintf("[%5s] %s", w.State, w.Step)
}
