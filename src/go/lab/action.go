// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package lab

import (
	"fmt"
	"runtime/debug"
)

// Action wraps an error with an action scope as described by the |action|
// string and its arguments.
//
// The expected usage is as follows:
//
// func DoSomething(o string, args ...interface{}) (err error) {
//   defer Action(&err, "doing something with %s", o)
//   ...
// }
//
// This way the returned error will be automatically wrapped in a ScopedAction.
// The err.Error() string will look like:
//
//   something went wrong with bar()
//   ... while doing something with foo
//
// Schedule the Action() call as the first |defer| in a function. This will
// ensure that the error code is wrapped if non-nil regardless of other defered
// operations.
func Action(err *error, action string, v ...interface{}) error {
	*err = ScopedAction{action: fmt.Sprintf(action, v...)}.Wrap(*err)
	if *err != nil {
		print((*err).Error())
	}
	return *err
}

// ScopedAction contains information for generating better error messages.
//
// During normal operation there can be several nested operations going on at
// the same time. Bubbling up an |error| object doesn't help identify which
// operations were involved. Hence we use ScopedAction objects to annotate the
// error messages.
//
// Say we start some operation like so:
//
//    func DoFoo() error {
//       err := bar()
//       if err != nil {
//          return err
//       }
//	     ...
//    }
//
//    func DoBar() error {
//       err := DoFoo()
//       if err != nil {
//          return err
//       }
//	     ...
//    }
//
// In this case, if bar() returns an error, it bubbles to the top. It may not
// be obvious what went wrong. The same code using ScopedActions could look
// like this:
//
//    func DoFoo() (err error) {
//       defer Action(&err, "doing foo")
//       err := bar()
//       if err != nil {
//          return
//       }
//	     ...
//    }
//
//    func DoBar() (err error) {
//       defer Action(&err, "doing bar")
//       err := DoFoo()
//       if err != nil {
//          return
//       }
//	     ...
//    }
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
type ScopedAction struct {
	action string
}

func (c ScopedAction) actionLine() string {
	return "\n...  while " + c.action
}

// Wrap takes an error and wraps the message with the current action's
// description.
func (c ScopedAction) Wrap(e error) error {
	if e == nil {
		return nil
	}
	return &wrappedError{
		innerError: e,
		s:          e.Error() + c.actionLine(),
		stack:      debug.Stack()}
}

// WaitFor waits for |n| number of |error| objects from |c|. If any of them are
// non-nil, then it returns a non-nil error which has an aggregate error string
// which includes all the things that went wrong.
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
	var w wrappedErrors
	for i := 0; i < n; i += 1 {
		e := <-ch
		if e != nil {
			w.innerErrors = append(w.innerErrors, e)
		}
	}

	if len(w.innerErrors) > 0 {
		return &w
	}
	return nil
}

// newAction constructs a new ScopedAction with a formatted message. The
// |action| string and the variadic arguments are as per fmt.Sprintf().
func newAction(action string, v ...interface{}) ScopedAction {
	return ScopedAction{action: fmt.Sprintf(action, v...)}
}

// Convenince function for generating errors.
func NewError(format string, v ...interface{}) error {
	return fmt.Errorf(format, v...)
}

// UnwrapError returns the innner error for |err| assuming |err| was previously
// wrapped via a ScopedAction.
func UnwrapError(err error) error {
	if err.(*wrappedError) != nil {
		return err.(*wrappedError).innerError
	}
	return err
}

// wrappedErrors aggregates a list of errors. Commonly used when returning the
// status of a list of parallelized tasks. See WaitFor().
type wrappedErrors struct {
	innerErrors []error
}

// Wraps a list of errors in a single error. Useful for aggregating a list of
// errors.
func WrapErrorList(list []error) error {
	if len(list) == 0 {
		return nil
	}

	if len(list) == 1 {
		return list[0]
	}

	w := wrappedErrors{}
	copy(w.innerErrors, list)
	return &w
}

// Error returns a string describing the collection of errors contained in this
// error object.
func (w *wrappedErrors) Error() string {
	s := ""
	for _, e := range w.innerErrors {
		s += e.Error() + "\n"
	}
	return s
}

// wrappedError is the type of error returned by ScopedAction. It's not
// publicly instantiated. The recommended method of unwrapping the result of a
// ScopedAction is to use UnwrapError().
type wrappedError struct {
	innerError error
	s          string
	stack      []byte
}

// Error returns a string describing this error including a stack trace.
func (w *wrappedError) Error() string {
	return w.s + "\n" + string(w.stack)
}

// WrapErrorWithMessage can be used to add a message to another error.
// Convenient when code needs to annotate a returned error with more user
// friendly context.
func WrapErrorWithMessage(e error, format string, v ...interface{}) error {
	if e == nil {
		return nil
	}

	return &wrappedError{
		innerError: e,
		s:          fmt.Sprintf(format, v...)}
}
