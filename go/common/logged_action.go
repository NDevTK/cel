// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"fmt"
	"github.com/pkg/errors"
)

// LoggedAction is similar to action, but indicates an action whose start and
// end are logged to a Logger instance.
//
// Usage is different from Action, and should include an additional function
// invocation as follows:
//
//     func DoSomething(l Logger, i int) (err error) {
//       defer LoggedAction(l, &err, "create object number %d", i)()
//       //                                                       ^^
//       //          Note the additional function invocation here !!
//       //          -----------------------------------------------
//       DoStuff()
//     }
//
// The biggest difference is that LoggedAction logs both the start and the end
// of the function (or at least the point at which it was invoked). It is the
// responsibility of the Logger to record timing information. LoggedAction by
// itself doesn't explicitly record timing.
//
//
// Vetting LoggedAction Calls
//
// The go/tools/vet_annotations tool vets invocations to all known LoggedAction
// derivatives. Please keep the derivative list in
// go/tools/vet_annotations/vet.go up-to-date.
func LoggedAction(l Logger, err *error, action string, v ...interface{}) func() {
	msg := fmt.Sprintf(action, v...)
	l.Info(event{Step: msg, State: actionBegin})
	return func() {
		// Inner function call to Action() can't call recover on behalf of the
		// call frame of our invoker.
		if r := recover(); r != nil {
			if rerr, ok := r.(error); ok {
				*err = errors.Wrapf(rerr, "panic")
			} else {
				*err = errors.Errorf("panic: %v", r)
			}
		}

		Action(err, action, v...)
		if *err != nil {
			l.Error(event{Step: msg, State: actionFailed, Error: *err})
		} else {
			l.Info(event{Step: msg, State: actionEnd})
		}
	}
}
