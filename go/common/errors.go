// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"fmt"
)

// UnresolvedReferenceError is returned to indicate that a reference could no
// be resolved.
type UnresolvedReferenceError struct {
	To        RefPath
	From      RefPath
	InlineRef string
	Reason    string
}

func (u *UnresolvedReferenceError) Error() string {
	if len(u.From) != 0 && len(u.To) != 0 {
		return fmt.Sprintf("reference from \"%s\" to \"%s\" could not be resolved: ",
			u.From.String(), u.To.String()) + u.Reason
	}
	return fmt.Sprintf("reference to \"%s\" could not be resolved: ", u.To.String()) + u.Reason
}

// ServiceNotFoundError is returned when a required service was not found in a
// context.Context. This implies that the context was not properly constructed
// or the wrong context was passed in to a function.
type ServiceNotFoundError struct {
	Service string
}

func (s *ServiceNotFoundError) Error() string {
	return fmt.Sprintf("service \"%s\" not found", s.Service)
}
