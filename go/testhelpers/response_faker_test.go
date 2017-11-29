// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package testhelpers

import (
	"encoding/json"
	u "net/url"
	"strings"
	"testing"
)

func TestResponseFaker_Basic(t *testing.T) {
	url := "https://example.com/foo/bar"
	f := NewResponseFaker(t)
	exp := f.Expect(
		RestRequest{Url: url},
		RestResponse{BodyObject: `
		{
		  "A": "a",
		  "B": "b"
		}`})
	client := f.NewClient()

	r, err := client.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	if r.Header.Get("content-type") != "application/json" {
		t.Fatal(r.Header.Get("content-type"))
	}

	if exp.HitCount != 1 {
		t.Fatal(exp)
	}
	r.Body.Close()
}

func TestResponseFaker_Values(t *testing.T) {
	url := "https://example.com/foo/bar"

	v1 := make(u.Values)
	v1.Add("alt", "json")

	v2 := make(u.Values)
	v2.Add("alt", "text")

	f := NewResponseFaker(t)
	f.Expect(
		RestRequest{Url: url, Values: v1},
		RestResponse{BodyObject: `{ "r": 1 }`})
	f.Expect(
		RestRequest{Url: url, Values: v2},
		RestResponse{BodyObject: `{ "r": 2 }`})

	client := f.NewClient()

	t.Run("json", func(t *testing.T) {
		r, err := client.Get(url + "?alt=json")
		if err != nil {
			t.Fatal(err)
		}
		var v map[string]int
		defer r.Body.Close()
		err = json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatal(err)
		}
		if v["r"] != 1 {
			t.Fatal(v)
		}
	})

	t.Run("text", func(t *testing.T) {
		r, err := client.Get(url + "?alt=text")
		if err != nil {
			t.Fatal(err)
		}
		var v map[string]int
		defer r.Body.Close()
		err = json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatal(err)
		}
		if v["r"] != 2 {
			t.Fatal(v)
		}
	})
}

func TestResponseFaker_Body(t *testing.T) {
	url := "https://example.com/foo/bar"
	f := NewResponseFaker(t)
	f.Expect(
		RestRequest{
			Url:        url,
			Method:     "POST",
			BodyObject: `{ "A": "a" }`},
		RestResponse{
			BodyObject: `{ "request_number": 1 }`})
	f.Expect(
		RestRequest{
			Url:        url,
			Method:     "POST",
			BodyObject: `{ "A": "b" }`},
		RestResponse{
			BodyObject: `{ "request_number": 2 }`})

	client := f.NewClient()

	t.Run("first request", func(t *testing.T) {
		r, err := client.Post(
			url, "application/json",
			strings.NewReader(`{"A":"a", "B":"b"}`))
		if err != nil {
			t.Fatal(err)
		}
		defer r.Body.Close()

		var v map[string]int
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&v)
		if err != nil {
			t.Fatal(err)
		}
		if v["request_number"] != 1 {
			t.Fatal(v)
		}
	})

	t.Run("second request", func(t *testing.T) {
		r, err := client.Post(
			url, "application/json",
			strings.NewReader(`{"A":"b", "B":"b"}`))
		if err != nil {
			t.Fatal(err)
		}
		defer r.Body.Close()

		var v map[string]int
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&v)
		if err != nil {
			t.Fatal(err)
		}
		if v["request_number"] != 2 {
			t.Fatal(v)
		}
	})
}
