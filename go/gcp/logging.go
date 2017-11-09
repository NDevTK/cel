// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

import (
	"cloud.google.com/go/logging"
)

// LogEntrySource is an interface for retrieving a logging.Entry structure
// given a severity. It is to be used with Log*() methods of Config as follows:
//
// First define your LogEntrySource:
//
//   type myLogEntrySource struct {
//       Message string `json:"m"`
//       Error string `json:"err,omitempty"`
//       ...
//   }
//
// Then introduce an Entry() method that returns a logging.Entry:
//
//   func (e myLogEntrySource) Entry(s logging.Severity) logging.Entry {
//     return logging.Entry{
//       Severity: s,
//       Payload: e}
//   }
//
// Note that Entry.Payload needs to be serializable via encoding/json, hence
// the `json` annotations on the entry structure. Common JSON fields are:
//
//     * "m"   : For a descriptive message. E.g. : "failed to fetch foo"
//     * "err" : For a text serialized error mesage describing the error. This
//               can usually be sourced from error.Error().
//
// Then you can log an event like so:
//
//     c.LogInfo(myLogEntrySource{
//         Message: "failed to fetch foo",
//         Error: err.Error(),
//         ...
//     })
//
type LogEntrySource interface {
	Entry(s logging.Severity) logging.Entry
}
