// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package testhelpers

import (
	"testing"
)

func Test_GetNormalizedJson_Object(t *testing.T) {
	v, err := GetNormalizedJson(struct {
		A string `json:"a"`
		B int    `json:"b"`
		C bool   `json:"c"`
	}{"a", 10, true})

	if err != nil {
		t.Fatal()
	}

	if v["a"].(string) != "a" {
		t.Fatal()
	}

	if v["b"].(float64) != 10 {
		t.Fatal()
	}

	if !v["c"].(bool) {
		t.Fatal()
	}
}

func Test_GetNormalizedJson_String(t *testing.T) {
	v, err := GetNormalizedJson(
		`{
		  "a": "foo",
		  "b": 10
		 }`)
	if err != nil {
		t.Fatal()
	}

	if v["a"].(string) != "foo" {
		t.Fatal()
	}
}

func Test_IsJsonSubset_Homogenous(t *testing.T) {
	type foo struct {
		A string
		B int
	}

	if subset, err := IsJsonSubset(foo{"a", 10}, foo{"a", 10}); !subset {
		t.Fatal(err)
	}

	if subset, err := IsJsonSubset(foo{"a", 10}, foo{"b", 10}); subset {
		t.Fatal(err)
	}
}

func Test_IsJsonSubset_Heterogenous(t *testing.T) {
	type foo struct {
		A string
		B int
	}

	if subset, err := IsJsonSubset(foo{"a", 10}, `
	{
	  "A": "a",
	  "B": 10
	}`); !subset {
		t.Fatal(err)
	}

	if subset, err := IsJsonSubset(foo{"a", 10}, `
	{
	  "A": "a",
	  "B": 10,
	  "C": true
	}`); !subset {
		t.Fatal(err)
	}

	if subset, err := IsJsonSubset(&foo{"a", 10}, `
	{
	  "A": "a",
	  "B": 10,
	  "C": true
	}`); !subset {
		t.Fatal(err)
	}

	if subset, _ := IsJsonSubset(foo{"a", 10}, `
	{
	  "A": "a",
	  "B": 9
	}`); subset {
		t.Fatal()
	}

	if subset, _ := IsJsonSubset(foo{"a", 10}, `
	{
	  "A": "b",
	  "B": 10,
	  "C": true
	}`); subset {
		t.Fatal()
	}
}

func Test_IsJsonSubset_Zero(t *testing.T) {
	type one struct {
		A string `json:"a,omitempty"`
		B int    `json:"b,omitempty"`
		C bool   `json:"c,omitempty"`
	}

	if subset, err := IsJsonSubset(one{}, `
  {
	"D": "d",
	"E": 1
  }`); !subset {
		t.Fatal(err)
	}
}

func Test_IsJsonSubset_Slice(t *testing.T) {
	type one struct {
		A string
		B []string
	}

	type two struct {
		A string
		B []one
	}

	if subset, err := IsJsonSubset(one{"a", []string{"x", "y"}}, `
	{
	  "A": "a",
	  "B": ["x","y"]
	}`); !subset {
		t.Fatal(err)
	}
}
