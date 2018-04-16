// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"testing"
)

type largeA struct {
	A int
	B string
	C *smallA
	D []*smallA
	E smallA
	F int
}

type largeB struct {
	A int
	B string
	C *smallB
	D []*smallB
	E smallB
}

type smallA struct {
	A int
}

type smallB struct {
	A int
}

func TestHomomorphicCopy_small(t *testing.T) {
	var sa smallA
	sa.A = 1

	var sb smallB

	err := HomomorphicCopy(&sa, &sb)
	if err != nil {
		t.Error(err)
	}

	if sb.A != 1 {
		t.Error("failed to copy field")
	}

	var psa *smallA = &sa
	var psb *smallB
	err = HomomorphicCopy(&psa, &psb)
	if err != nil {
		t.Error(err)
	}

	if psb.A != 1 {
		t.Error("failed to copy field")
	}
}

func TestHomomorphicCopy_large(t *testing.T) {
	la := largeA{A: 1, B: "foo", C: &smallA{A: 2}, D: []*smallA{{A: 1}, {A: 2}}, E: smallA{A: 4}}
	var lb largeB

	err := HomomorphicCopy(&la, &lb)
	if err != nil {
		t.Error(err)
	}

	if lb.A != 1 || lb.B != "foo" || lb.C == nil || lb.C.A != 2 || len(lb.D) != 2 || lb.D[0].A != 1 || lb.D[1].A != 2 || lb.E.A != 4 {
		t.Errorf("failed to copy: %#v", lb)
	}
}
