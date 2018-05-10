// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

// Must eats the error argument for cases where the caller definitely doesn't
// expect absolutely anything to go wrong. If the error is non-nil, it panics.
// Otherwise the remaining argument is funneled through.
func Must(i interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return i
}
