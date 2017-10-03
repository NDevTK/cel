// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cel

// Convenince function for generating errors.
func NewError(format string, v ...interface{}) error {
	return fmt.Errorf(format, v...)
}
