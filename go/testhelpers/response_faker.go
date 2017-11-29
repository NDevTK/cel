// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package testhelpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync/atomic"
	"testing"
)

// ReaderToReadCloser takes a Reader and adds a Close() method so that the
// reader can be used where a ReadCloser is required. The Close() method does
// nothing.
type ReaderToReadCloser struct {
	Reader io.Reader
}

func (r *ReaderToReadCloser) Close() error               { return nil }
func (r *ReaderToReadCloser) Read(p []byte) (int, error) { return r.Reader.Read(p) }

type urlAndMethod struct {
	Url    string
	Method string
}

type RestRequest struct {
	Url        string
	Values     url.Values
	Method     string
	Headers    http.Header
	BodyObject interface{}
}

type RestResponse struct {
	Code       int
	Status     string
	Headers    http.Header
	BodyObject interface{}
	BodyPath   string
}

type Expectation struct {
	Active   bool
	Request  RestRequest
	Response RestResponse
	HitCount int64
}

var kNotHandledError = errors.New("request not handled")

func (e *Expectation) Handle(req *http.Request) (resp *http.Response, err error) {
	if !e.Active {
		return nil, kNotHandledError
	}

	for k, want := range e.Request.Headers {
		if len(want) == 0 {
			continue
		}

		has, ok := req.Header[k]
		if !ok {
			return nil, kNotHandledError
		}
		has_map := make(map[string]bool)
		for _, h := range has {
			has_map[h] = true
		}
		for _, w := range want {
			// Value mismatch
			if _, ok := has_map[w]; !ok {
				return nil, kNotHandledError
			}
		}
	}

	if e.Request.Values != nil {
		if subst, _ := IsJsonSubset(e.Request.Values, req.URL.Query()); !subst {
			return nil, kNotHandledError
		}
	}

	if e.Request.BodyObject != nil {
		if req.GetBody == nil {
			return nil, kNotHandledError
		}

		body, err := req.GetBody()
		if err != nil {
			return nil, err
		}

		bytes, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, err
		}

		if subst, _ := IsJsonSubset(e.Request.BodyObject, bytes); !subst {
			return nil, kNotHandledError
		}
	}

	resp = &http.Response{
		Header:        e.Response.Headers,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Status:        e.Response.Status,
		StatusCode:    e.Response.Code,
		ContentLength: -1,
	}

	resp.Body, err = e.getResponseBody()
	if err != nil {
		return nil, err
	}

	if resp.Header == nil {
		resp.Header = make(http.Header)
	}

	if resp.Header.Get("Content-Type") == "" {
		resp.Header.Add("Content-Type", "application/json")
	}

	atomic.AddInt64(&e.HitCount, 1)
	return
}

func (e *Expectation) getResponseBody() (io.ReadCloser, error) {
	if e.Response.BodyObject != nil {
		njson, err := GetNormalizedJson(e.Response.BodyObject)
		if err != nil {
			return nil, err
		}

		body, err := json.Marshal(njson)
		if err != nil {
			return nil, err
		}
		return &ReaderToReadCloser{bytes.NewBuffer(body)}, nil
	}

	if e.Response.BodyPath != "" {
		body, err := ioutil.ReadFile(e.Response.BodyPath)
		if err != nil {
			return nil, err
		}
		return &ReaderToReadCloser{bytes.NewReader(body)}, nil
	}

	return nil, nil
}

func (e *Expectation) Deactivate() {
	e.Active = false
}

func (e *Expectation) Activate() {
	e.Active = true
}

type ResponseFaker struct {
	t            *testing.T
	expectations map[urlAndMethod][]*Expectation
}

func (f *ResponseFaker) Expect(req RestRequest, resp RestResponse) *Expectation {
	if req.Method == "" {
		req.Method = "GET"
	}

	if resp.Code == 0 {
		resp.Code = 200
	}

	if resp.Status == "" {
		resp.Status = fmt.Sprintf("%d Status %d", resp.Code, resp.Code)
	}

	if u, e := url.Parse(req.Url); e != nil || u.RawPath != "" {
		panic("url for a request expectation should be a valid URL without a query. " +
			"Use the Value property of the RestRequest object to specify query parameters.")
	}

	exp := &Expectation{Active: true, Request: req, Response: resp}
	if f.expectations == nil {
		f.expectations = make(map[urlAndMethod][]*Expectation)
	}
	m := urlAndMethod{Url: req.Url, Method: req.Method}
	f.expectations[m] = append(f.expectations[m], exp)
	return exp
}

func (f *ResponseFaker) NewClient() *http.Client {
	return &http.Client{Transport: f}
}

func (f *ResponseFaker) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	base_url := *req.URL
	base_url.RawQuery = ""
	m := urlAndMethod{Url: base_url.String(), Method: req.Method}
	exps, _ := f.expectations[m]

	for _, exp := range exps {
		resp, err = exp.Handle(req)

		if err == kNotHandledError {
			continue
		}

		return
	}

	dumped, _ := httputil.DumpRequestOut(req, true)
	f.t.Errorf("request for url '%s' with method '%s' was not handled\nRequest details:\n\n%s",
		req.URL.String(), req.Method, dumped)
	return notFoundError(), nil
}

func NewResponseFaker(t *testing.T) *ResponseFaker {
	return &ResponseFaker{t: t}
}

func notFoundError() *http.Response {
	return &http.Response{
		Status:        "404 Not found",
		StatusCode:    404,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		ContentLength: -1,
		Body:          &ReaderToReadCloser{strings.NewReader("not found\r\n")},
	}
}

func internalError(err error) *http.Response {
	body := err.Error()
	return &http.Response{
		Status:        "500 Internal",
		StatusCode:    500,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		ContentLength: -1,
		Body:          &ReaderToReadCloser{strings.NewReader(body)},
	}
}
