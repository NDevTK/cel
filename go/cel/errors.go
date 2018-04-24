// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cel

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"fmt"
	"strings"
)

type NotImplementedError struct {
	thing string
}

func (n *NotImplementedError) Thing() string {
	return n.thing
}

func (n *NotImplementedError) Error() string {
	return fmt.Sprintf("not implemented: %s", n.thing)
}

func NewNotImplementedError(thing string) error {
	return &NotImplementedError{thing: thing}
}

type NotReadyError struct {
	Root     common.RefPath
	messages []string
}

func (n *NotReadyError) Error() string {
	return fmt.Sprintf(`The subtree rooted at path "%s" is not fully resolved. The following issues were found:

%s
`, n.Root, strings.Join(n.messages, "\n"))
}

func NewNotReadyError(refs *common.Namespace, root common.RefPath) error {
	var s []string
	refs.VisitUnresolved(root, func(v common.UnresolvedValue) bool {
		s = append(s, v.String())
		return true
	})
	return &NotReadyError{Root: root, messages: s}
}
