// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package lab

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Environment variable that controls whether the RequestReplayer is in
// 'Record' mode or 'Replay' mode. This variable must exist and contains a
// non-empty value for the RequestReplayer to switch to 'Record' mode.
const kLabTestRecordEnv = "LAB_RECORD"

// ReaderToReadCloser takes a Reader and adds a Close() method so that the
// reader can be used where a ReadCloser is required. The Close() method does
// nothing.
type ReaderToReadCloser struct {
	Reader io.Reader
}

func (r *ReaderToReadCloser) Close() error               { return nil }
func (r *ReaderToReadCloser) Read(p []byte) (int, error) { return r.Reader.Read(p) }

// RequestReplayer is-a http.RoundTripper that can record and replay network
// requests.
//
// It has two modes of operation: replay and record. The normal mode of
// operation is 'Replay', but it will switch to 'Record' mode if the
// 'LAB_RECORD' environment variable is defined and has a non-empty value.
//
// Replay
// ------
// During the replay phase, it intercepts outgoing network requests. If a
// corresponding cached response is found, it sends the cached response to the
// caller. If no response is found, it responds with a synthesized HTTP 500
// response with a status of "not implemented".
//
// Record
// ------
// The record phase is similar to the Replay phase, but in the absence of a
// cached response, it sends the response to the network and records both the
// request and the response.
//
// File format
// -----------
// This class uses directories named 'requests' and 'responses' under the
// DataPath directory for storing the requests and the responses respectively.
// Both are stored in wire format and can be read and written to directly by
// http.Request and http.Response objects.
//
// The filename for both the request and the response is the hex encoded SHA1
// digest of the wire format of the request.
type RequestReplayer struct {
	DataPath string
}

func (t *RequestReplayer) requestDir(digest string) string {
	if t.DataPath == "" {
		panic("no TestDataPath specified for mock request handler")
	}
	return filepath.Join(t.DataPath, "requests", digest)
}

func (t *RequestReplayer) responseDir(digest string) string {
	if t.DataPath == "" {
		panic("no TestDataPath specified for mock request handler")
	}
	return filepath.Join(t.DataPath, "responses", digest)
}

func (t *RequestReplayer) recordExternalRequest(req *http.Request, rs, digest string) (err error) {
	if os.Getenv(kLabTestRecordEnv) == "" {
		return NewError("not implemented")
	}

	// Record and playback mode
	client, err := GetDefaultClient(req.Context())
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	for k, _ := range resp.Header {
		if strings.HasPrefix(strings.ToLower(k), "x-goog") {
			resp.Header.Del(k)
		}
	}

	err = ioutil.WriteFile(t.requestDir(digest), bytes.NewBufferString(rs).Bytes(), 0644)
	if err != nil {
		return
	}

	f, err := os.Create(t.responseDir(digest))
	if err != nil {
		return
	}
	defer f.Close()

	print(fmt.Sprintf("recorded request to: %s. Received status: %d\n", req.URL.String(), resp.StatusCode))
	return resp.Write(f)
}

func requestToString(req *http.Request) (string, string, error) {
	var b bytes.Buffer
	err := req.Write(&b)
	return string(b.Bytes()), fmt.Sprintf("%x", sha1.Sum(b.Bytes())), err
}

func (t *RequestReplayer) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	rs, digest, err := requestToString(req)
	if err != nil {
		return
	}

	if _, err = os.Stat(t.responseDir(digest)); err != nil {
		err = t.recordExternalRequest(req, rs, digest)
	}
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile(t.responseDir(digest), os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	return http.ReadResponse(bufio.NewReader(f), req)
}
