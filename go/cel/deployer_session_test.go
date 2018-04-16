// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cel

import (
	"testing"
)

func TestCreateGenerationId(t *testing.T) {
	g, err := createGenerationId()
	if err != nil {
		t.Fatal(err)
	}

	if len(g) != 32 {
		t.Errorf("Incorrect length for generation ID: \"%s\"", g)
	}

	// for giggles, let's try this a few times. It'll fail 1 in 2^18 times.
	h, err := createGenerationId()
	if err != nil {
		t.Fatal(err)
	}

	if g == h {
		t.Error("generation ID conflict. What's wrong with our source of randomness?")
	}
}
