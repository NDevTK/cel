// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"github.com/golang/protobuf/proto"
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

type JobWaiter struct {
	errc   chan error
	count  int
	source proto.Message
}

func NewJobWaiter(m proto.Message) *JobWaiter {
	return &JobWaiter{errc: make(chan error), count: 0, source: m}
}

func (j *JobWaiter) New() chan<- error {
	j.count += 1
	return j.errc
}

func (j *JobWaiter) Source() proto.Message {
	return j.source
}

func (j *JobWaiter) Join() error {
	defer func() {
		j.count = 0
	}()

	return WaitFor(j.errc, j.count)
}
