// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cel

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

// Error returns a string describing the collection of errors contained in this
// error object.
func (w *wrappedErrors) Error() string {
	s := ""
	for _, e := range w.innerErrors {
		s += e.Error() + "\n"
	}
	return s
}
