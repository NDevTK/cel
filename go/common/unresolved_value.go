// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"fmt"
	"strings"
)

// UnresolvedValue represents a node in our namespace which hasn't been
// resolved yet. The type or reason for not being resolved is indicated by
// which of the detail fields is non-nil.
//
// UnresolvedValue is returned by References.VisitUnresolved().
type UnresolvedValue struct {
	Location RefPath // Location of the unresolved node.
	Value    UnresolvedValueType
}

func (u UnresolvedValue) String() string {
	return fmt.Sprintf("unresolved value at \"%s\"\n%s\n", u.Location, u.Value)
}

// UnresolvedValueType identifies the kind of unresolved value.
type UnresolvedValueType interface {
	fmt.Stringer
	sealed()
}

// UnresolvedValue_String indicates that the corresponding node is a string
// with external field references. Each node that's referred to by this node is
// listed in the `DependsOn` field.
type UnresolvedValue_String struct {
	DependsOn []RefPath
}

func (u UnresolvedValue_String) sealed() {}
func (u UnresolvedValue_String) String() string {
	return fmt.Sprintf("A string field has unresolved references to the following paths : %s",
		refListToString(u.DependsOn))
}

// UnresolvedValue_Output indicates that the corresponding field is an output
// field whose value has not been published yet. If any other fields refer to
// the output of this field, then those nodes are listed in `Dependents`.
type UnresolvedValue_Output struct {
	Dependents []RefPath
}

func (u UnresolvedValue_Output) sealed() {}
func (u UnresolvedValue_Output) String() string {
	reason := "An OUTPUT field is waiting to be assigned a value."
	if len(u.Dependents) != 0 {
		reason += fmt.Sprintf(" The following fields depend on this output : %s",
			refListToString(u.Dependents))
	}
	return reason
}

// UnresolvedValue_Placeholder indicates that the corresponding node was
// created because one or more other nodes referred to it. The node in question
// currently does not correspond to any known node in a grafted message. Nodes
// that refer to this node are listed in `Referrers`.
type UnresolvedValue_Placeholder struct {
	Referrers []RefPath
}

func (u UnresolvedValue_Placeholder) sealed() {}
func (u UnresolvedValue_Placeholder) String() string {
	return fmt.Sprintf(
		"The following paths have references to this field : %s\n"+
			"However, none of the objects known to the namespace map to this path.",
		refListToString(u.Referrers))
}

func refListToString(rl []RefPath) string {
	if len(rl) == 0 {
		return "(empty)"
	}
	var s []string
	for _, r := range rl {
		s = append(s, "\""+r.String()+"\"")
	}
	return strings.Join(s, ", ")
}
