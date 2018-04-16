// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"fmt"
	"github.com/pkg/errors"
)

// Action annotates an error by wrapping it with an enclosing error.
//
// During normal operation there can be several nested operations going on at
// the same time. Bubbling up an |error| object doesn't help identify which
// operations were involved. Hence we use ScopedAction objects to annotate the
// error messages.
//
// Say we start some operation like so:
//
//     func DoFoo() error {
//        err := bar()
//        if err != nil {
//           return err
//        }
//	      ...
//     }
//
//     func DoBar() error {
//        err := DoFoo()
//        if err != nil {
//           return err
//        }
//	      ...
//     }
//
// In this case, if bar() returns an error, it bubbles to the top. It may not
// be obvious what went wrong. The same code using ScopedActions could look
// like this:
//
//     func DoFoo() (err error) {
//        defer Action(&err, "doing foo")
//        err := bar()
//        if err != nil {
//           return
//        }
//	      ...
//     }
//
//     func DoBar() (err error) {
//        defer Action(&err, "doing bar")
//        err := DoFoo()
//        if err != nil {
//           return
//        }
//	      ...
//     }
//
// Now, the error returned by bar() is wrapped by both DoFoo()'s and DoBar()'s
// actions. Hence the err.Error() output looks like:
//
//    something went wrong with bar()
//    ... while doing foo
//    ... while doing bar
//
// It should be obvious that we can do something similar using stack traces.
// But this approach is a more usable and much less noisy when multiple
// external libraries are involved.
func Action(err *error, action string, v ...interface{}) {
	if r := recover(); r != nil {
		if rerr, ok := r.(error); ok {
			*err = rerr
		} else {
			*err = errors.Errorf("panic: %v", r)
		}
	}

	if *err != nil {
		*err = errors.WithMessage(*err, fmt.Sprintf(action, v...))
	}
}
