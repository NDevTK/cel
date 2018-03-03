// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

// WaitFor waits for |n| number of |error| objects from |c|. If any of them are
// non-nil, then it returns a non-nil error which has an aggregate error string
// which includes all the things that went wrong.
//
// If |c| is closed, then the function returns the errors that were observed so
// far.
//
// Expected usage is as follows:
//
//     v := make(chan error)
//     for _, x := range things {
//       go func(xx *XType) {
//          v <- HandleXThing(xx)
//       }(x)
//     }
//     return WaitFor(v, len(things))
//
// Basically, this makes it more convenience to gather the results of a number
// of operations that are sequenced in parallel via goroutines.
func WaitFor(ch <-chan error, n int) (err error) {
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

	return WrapErrorList(l)
}

type JobWaiterSource interface{}

// JobWaiter makes it easy to wait on an unspecified number of events jobs each
// of which is expected to asynchronously return an error value.
//
// For example:
//
//     func DoAllTheThings(n int) error {
//       w := NewJobWaiter(nil)
//       for i := 0; i < n; i += 1 {
//         go func(v int, e chan<- error) {
//            e <- DoSomething(v);
//         }(i, w.New())
//       }
//       return w.Join()
//     }
//
//     func DoSomething() error {
//       // Real lengthy operation.
//     }
//
type JobWaiter struct {
	errc   chan error
	count  int
	source JobWaiterSource
}

// NewJobWaiter constructs a new JobWaiter object.
//
// Optionally, you can specify a source object if one is needed for accounting
// purposes. The source is opaque to JobWaiter but can be queried via the
// Source() method.
func NewJobWaiter(m JobWaiterSource) *JobWaiter {
	return &JobWaiter{errc: make(chan error), count: 0, source: m}
}

// New adds a new job to a JobWaiter.
//
// The new task is considered complete when an error object is sent to the
// returned channel.
func (j *JobWaiter) New() chan<- error {
	j.count += 1
	return j.errc
}

func (j *JobWaiter) Source() JobWaiterSource {
	return j.source
}

// Join waits for all jobs to complete and returns an aggregate error message.
//
// New() should not be called while Join() is running. However, once Join()
// returned, the JobWaiter object can be reused by calling New().
func (j *JobWaiter) Join() error {
	defer func() {
		j.count = 0
	}()

	return WaitFor(j.errc, j.count)
}
