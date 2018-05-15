// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

import (
	"encoding/json"
	"time"

	"chromium.googlesource.com/enterprise/cel/go/common"
	"github.com/pkg/errors"
	compute "google.golang.org/api/compute/v1"
	deploymentmanager "google.golang.org/api/deploymentmanager/v2beta"
	servicemanagement "google.golang.org/api/servicemanagement/v1"
)

// Amount of time to wait for next round of polling for an operation.
const pollDuration = time.Second * 2

// Maximum number of times an operation poll will be retried before giving up.
// The amount of time before an operation times out would be (maxRetries *
// pollDuration) + (overhead for operation state polling).
const maxRetries = 30

var (
	ErrOperationTimedOut    = errors.New("operation timed out")
	ErrUnknownOperationType = errors.New("unknown operation type")
)

func JoinOperation(s *Session, oi interface{}, desc string) error {
	switch o := oi.(type) {
	case *compute.Operation:
		return joinComputeOperation(s, o, desc)

	case *deploymentmanager.Operation:
		return joinDeploymentOperation(s, o, desc)

	case *servicemanagement.Operation:
		return joinServiceManagementOperation(s, o, desc)
	}

	return ErrUnknownOperationType
}

// JoinComputeOperation waits for the specific GCE compute operation to complete.
// These require periodic polling to make sure it's succeeded.
func joinComputeOperation(s *Session, op *compute.Operation, desc string) (err error) {
	defer GcpLoggedServiceAction(s, ComputeServiceName, &err, "%s", desc)()

	if op.Description != "" {
		s.Logger.Debug(common.MakeStringer("Description %s", op.Description))
	}

	cs, err := s.GetComputeService()
	if err != nil {
		return err
	}

	for retriesLeft := maxRetries; op.Status != "DONE" && retriesLeft > 0; retriesLeft-- {
		s.Logger.Debug(common.MakeStringer("Operation %s %s", op.Status, op.StatusMessage))

		time.Sleep(pollDuration)

		var nop *compute.Operation

		switch {
		case op.Zone != "":
			nop, err = cs.ZoneOperations.Get(s.GetProject(), LastPathComponent(op.Zone), op.Name).
				Context(s.GetContext()).
				Do()

		case op.Region != "":
			nop, err = cs.RegionOperations.Get(s.GetProject(), LastPathComponent(op.Region), op.Name).
				Context(s.GetContext()).
				Do()

		default:
			nop, err = cs.GlobalOperations.Get(s.GetProject(), op.Name).
				Context(s.GetContext()).
				Do()
		}

		if err == nil {
			op = nop
			continue
		}

		// It is possible for this lookup to fail if we poll immediately after
		// operation start. The error in this case would be a 404.
		if IsNotFoundError(err) {
			s.Logger.Debug(common.MakeStringer("Operation not found: %s", err.Error()))
			s.Logger.Debug(common.MakeStringer(
				"Failed getting project=\"%s\",zone=\"%s\",region=\"%s\",op=\"%s\"",
				s.GetProject(), op.Zone, op.Region, op.Name))
			continue
		}

		return err
	}

	if op.Status == "DONE" {
		return nil
	}

	if op.Error == nil || len(op.Error.Errors) == 0 {
		return ErrOperationTimedOut
	}

	var el []error
	for _, e := range op.Error.Errors {
		err := errors.Errorf("Code: %s, Location: %s, Message: %s", e.Code, e.Location, e.Message)
		s.Logger.Debug(common.MakeStringer("%s", err.Error()))
		el = common.AppendErrorList(el, err)
	}
	return errors.Wrapf(common.WrapErrorList(el), "%s failed", op.Name)
}

func joinDeploymentOperation(s *Session, op *deploymentmanager.Operation, desc string) (err error) {
	defer GcpLoggedServiceAction(s, DeploymentManagerServiceName, &err, "%s", desc)()

	if op.Description != "" {
		s.Logger.Debug(common.MakeStringer("Description %s", op.Description))
	}

	ds, err := s.GetDeploymentManagerService()
	if err != nil {
		return err
	}

	for retriesLeft := maxRetries; op.Status != "DONE" && retriesLeft > 0; retriesLeft-- {
		s.Logger.Debug(common.MakeStringer("Status: (%s) %s", op.Status, op.StatusMessage))

		time.Sleep(pollDuration)

		nop, err := ds.Operations.Get(s.GetProject(), op.Name).Context(s.GetContext()).Do()
		if err == nil {
			op = nop
			continue
		}

		// It is possible for this lookup to fail if we poll immediately after
		// operation start. The error in this case would be a 404.
		if IsNotFoundError(err) {
			continue
		}

		return err
	}

	if op.Status == "DONE" {
		return nil
	}

	if op.Error == nil || len(op.Error.Errors) == 0 {
		return ErrOperationTimedOut
	}

	var el []error
	for _, e := range op.Error.Errors {
		err := errors.Errorf("Code: %s, Location: %s, Message: %s", e.Code, e.Location, e.Message)
		s.Logger.Debug(common.MakeStringer("%s", err.Error()))
		el = common.AppendErrorList(el, err)
	}
	return errors.Wrapf(common.WrapErrorList(el), "%s failed", op.Name)
}

func joinServiceManagementOperation(s *Session, op *servicemanagement.Operation, desc string) (err error) {
	defer GcpLoggedServiceAction(s, ServiceManagementServiceName, &err, "%s", desc)()

	sm, err := servicemanagement.New(s.client)
	if err != nil {
		return err
	}

	for retriesLeft := maxRetries; !op.Done; retriesLeft-- {
		if metadata, err := json.Marshal(op.Metadata); err == nil {
			s.Logger.Debug(common.MakeStringer("Metadata: %s", string(metadata)))
		}

		time.Sleep(pollDuration)

		nop, err := sm.Operations.Get(op.Name).Context(s.GetContext()).Do()
		if err == nil {
			op = nop
			continue
		}

		if IsNotFoundError(err) {
			continue
		}

		return err
	}

	if op.Done {
		return nil
	}

	if op.Error == nil {
		return ErrOperationTimedOut
	}

	if m, err := json.Marshal(op.Error.Details); err == nil {
		return errors.Errorf("operation failed: %s. Details: %s", op.Error.Message, string(m))
	} else {
		return errors.Errorf("operation failed (%s, %s)", op.Error.Code, op.Error.Message)
	}
}
