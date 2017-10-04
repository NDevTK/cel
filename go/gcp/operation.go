// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

import (
	"cloud.google.com/go/logging"
	compute "google.golang.org/api/compute/v1"
	"time"
)

const kWaitForOperationTimeout = time.Second * 2

type operationEvent struct {
	Message   string             `json:"m"`
	Error     string             `json:"err,omitempty"`
	Operation *compute.Operation `json:"op,omitempty"`
}

func (e operationEvent) Entry(s logging.Severity) logging.Entry {
	return logging.Entry{Severity: s, Payload: e}
}

// WaitForOperation waits for the specific GCE compute operation to complete.
// These require periodic polling to make sure it's succeeded.
func WaitForOperation(s *Session, op *compute.Operation) (err error) {
	defer Action(&err, "fetching status of operation %s", op.Name)

	retries_left := 5
	for op.Status != "DONE" {
		time.Sleep(kWaitForOperationTimeout)
		op, err = s.GetComputeService().ZoneOperations.Get(
			s.Config.Project, LastPathComponent(op.Zone), op.Name).Context(s.Context).Do()
		if err != nil {
			if retries_left == 0 {
				s.LogError(operationEvent{
					Message:   "failed to fetch status of operation",
					Error:     err.Error(),
					Operation: op})
				return
			}
			retries_left -= 1
		}
	}

	s.LogInfo(operationEvent{
		Message:   "operation completed",
		Operation: op})
	return nil
}
