// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"github.com/pkg/errors"
)

// Tasks makes it easy to wait on an unspecified number of tasks each of which
// is expected to asynchronously return an error value.
type Tasks struct {
	// Error channel. Used for collecting errors from goroutines.
	errc chan error

	// Number of outstanding jobs.
	count int

	// Source. Opaque to Jobs. Only used for accounting by external parties.
	source interface{}
}

// NewTasks constructs a new Tasks object.
//
// Optionally, you can specify a source object if one is needed for accounting
// purposes. The source is opaque to JobWaiter but can be queried via the
// Source() method.
func NewTasks(source interface{}) *Tasks {
	return &Tasks{errc: make(chan error), count: 0, source: source}
}

// add adds a new task to a Tasks object.
//
// The caller MUST send exactly one error to the returned channel. Otherwise a
// subsequent Join() will hang or terminal too early depending on whether no
// errors were sent or more than one error was sent. It is the callers
// responsibility to deal with panic()/recover() as appropriate.
//
// As soon as the error is sent, the job is considered done.
//
// Note that it is possible for the channel to block until someone calls
// Join().
//
// Whenever possible, use Go() instead of add() when working with Jobs.
func (t *Tasks) add() chan<- error {
	t.count += 1
	return t.errc
}

// Go invokes a new goroutine that executes `worker` and collects its return
// value.
//
// The job is considered done when the function returns. Panics are captured
// and result in an error.
func (t *Tasks) Go(worker func() error) {
	go func(e chan<- error) {
		defer func() {
			if r := recover(); r != nil {
				if er, ok := r.(error); ok {
					e <- er
				} else {
					e <- errors.Errorf("panic: %#v", r)
				}
			}
		}()
		e <- worker()
	}(t.add())
}

// Source returns the source metadata that was passed in when the Jobs object
// was created.
func (t *Tasks) Source() interface{} {
	return t.source
}

// Join waits for all jobs to complete and returns an aggregate error message.
//
// New() should not be called while Join() is running. However, once Join()
// returned, the JobWaiter object can be reused by calling New() or Go().
func (t *Tasks) Join() error {
	defer func() {
		t.count = 0
	}()

	return WaitFor(t.errc, t.count)
}
