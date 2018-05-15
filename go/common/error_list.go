// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

// ErrorUnpacker can unpack an enclosed list of errors.
type ErrorUnpacker interface {
	// UnpackErrors returns the list of errors that are contained herein.
	UnpackErrors() []error
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
	w.innerErrors = make([]error, len(list))
	copy(w.innerErrors, list)
	return &w
}

// wrappedErrors aggregates a list of errors. Commonly used when returning the
// status of a list of parallelized tasks. See WaitFor().
type wrappedErrors struct {
	innerErrors []error
}

// UnpackErrors returns the list of wrapped errors.
func (w *wrappedErrors) UnpackErrors() []error {
	return w.innerErrors
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

// AppendErrorList appends one or more errors to a list of errors.
//
// As a special case, if any of the new errors happen to be a wrapped error
// returned from a prior call to WrapErrorList, that list is unpacked and used
// to extend the list of errors. Thus AppendErrorList can be used to keep a
// large collection of errors flattened rather than becoming tree shaped.
func AppendErrorList(list []error, values ...error) []error {
	for _, v := range values {
		if v == nil {
			continue
		}
		switch vt := v.(type) {
		case *wrappedErrors:
			list = append(list, vt.innerErrors...)

		default:
			list = append(list, vt)
		}
	}
	return list
}

// UnpackErrorList takes as input a possibly wrapped error and returns the
// underlying error list.
func UnpackErrorList(err error) []error {
	if u, ok := err.(ErrorUnpacker); ok {
		return u.UnpackErrors()
	}
	return []error{err}
}
