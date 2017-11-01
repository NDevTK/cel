// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cel

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
)

// WaitFor waits for |n| number of |error| objects from |c|. If any of them are
// non-nil, then it returns a non-nil error which has an aggregate error string
// which includes all the things that went wrong.
//
// If |c| is closed, then the function returns the errors that were observed so
// far.
//
// Expected usage is as follows:
//
//    v := make(chan error)
//    for _, x := range things {
//      go func(xx *XType) {
//         v <- HandleXThing(xx)
//      }(x)
//    }
//    return WaitFor(v, len(things))
//
// Basically, this makes it more convenience to gather the results of a number
// of operations that are sequenced in parallel via goroutines.
func WaitFor(ch chan error, n int) (err error) {
	var l []error
	for i := 0; i < n; i += 1 {
		e, ok := <-ch
		if e != nil {
			l = append(l, e)
		}

		if !ok {
			break
		}
	}

	return common.WrapErrorList(l)
}
