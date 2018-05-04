// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"fmt"
	"github.com/pkg/errors"
	compute "google.golang.org/api/compute/v1"
	deploymentmanager "google.golang.org/api/deploymentmanager/v2beta"
	servicemanagement "google.golang.org/api/servicemanagement/v1"
	"reflect"
	"time"
)

const kWaitForOperationTimeout = time.Second * 2

type operationEvent struct {
	Message   string      `json:"m"`
	Error     error       `json:"err,omitempty"`
	Name      string      `json:"-"`
	Operation interface{} `json:"op,omitempty"`
}

func (o operationEvent) String() string {
	if o.Error != nil {
		return fmt.Sprintf("Op(%s): %s : %s", o.Name, o.Message, o.Error.Error())
	}
	return fmt.Sprintf("Op(%s): %s", o.Name, o.Message)
}

type operationRefresher func() (interface{}, error)

type isDone func(interface{}) bool

func joinGenericOperation(s *Session, op interface{}, refresher operationRefresher, isDone isDone) (o interface{}, err error) {
	v := reflect.ValueOf(op).Elem()
	opName := v.FieldByName("Name").String()

	retries_left := 5
	for !isDone(op) {
		time.Sleep(kWaitForOperationTimeout)

		op, err = refresher()
		if err != nil {
			if retries_left == 0 {
				s.Logger.Error(operationEvent{Name: opName, Message: "Timed out", Operation: op})
				return op, err
			}
		}
		v = reflect.ValueOf(op).Elem()
	}

	return op, nil
}

// JoinComputeOperation waits for the specific GCE compute operation to complete.
// These require periodic polling to make sure it's succeeded.
func JoinComputeOperation(s *Session, op *compute.Operation) (err error) {
	defer GcpLoggedServiceAction(s, ComputeServiceName, &err,
		"Compute %s", op.Name)()

	o, err := joinGenericOperation(s, op, func() (interface{}, error) {
		var cs *compute.Service
		cs, err = s.GetComputeService()
		if err != nil {
			return nil, err
		}

		return cs.ZoneOperations.Get(
			s.HostEnvironment.Project.Name,
			LastPathComponent(op.Zone),
			op.Name).Context(s.ctx).Do()
	}, func(op interface{}) bool {
		return op.(*compute.Operation).Status == "DONE"
	})

	if err != nil {
		return err
	}

	op = o.(*compute.Operation)
	if op.Error == nil || len(op.Error.Errors) == 0 {
		return nil
	}

	var el []error
	for _, e := range op.Error.Errors {
		err := errors.Errorf("Code: %s, Location: %s, Message: %s", e.Code, e.Location, e.Message)
		s.Logger.Debug(operationEvent{
			Name:  op.Name,
			Error: err})
		el = common.AppendErrorList(el, err)
	}
	return errors.Wrapf(common.WrapErrorList(el), "%s failed", op.Name)
}

func JoinDeploymentOperation(s *Session, op *deploymentmanager.Operation) (err error) {
	defer GcpLoggedServiceAction(s, DeploymentManagerServiceName, &err,
		"Deployment %s", op.Name)()

	o, err := joinGenericOperation(s, op, func() (interface{}, error) {
		dm, err := s.GetDeploymentManagerService()
		if err != nil {
			return nil, err
		}

		return dm.Operations.Get(s.HostEnvironment.Project.Name, op.Name).Context(s.ctx).Do()
	}, func(op interface{}) bool {
		return op.(*deploymentmanager.Operation).Status == "DONE"
	})

	if err != nil {
		return err
	}

	op = o.(*deploymentmanager.Operation)
	if op.Error == nil || len(op.Error.Errors) == 0 {
		return nil
	}

	var el []error
	for _, e := range op.Error.Errors {
		err := errors.Errorf("\n  Code: %s\n  Location: %s\n  Message: %s",
			e.Code, e.Location, e.Message)
		s.Logger.Debug(operationEvent{
			Name:  op.Name,
			Error: err})
		el = common.AppendErrorList(el, err)
	}
	return errors.Wrapf(common.WrapErrorList(el), "%s failed", op.Name)
}

func JoinServiceManagementOperation(s *Session, op *servicemanagement.Operation) (err error) {
	defer GcpLoggedServiceAction(s, ServiceManagementServiceName, &err,
		"Service Management %s", op.Name)()

	sm, err := servicemanagement.New(s.client)
	if err != nil {
		return err
	}
	o, err := joinGenericOperation(s, op, func() (interface{}, error) {
		return sm.Operations.Get(op.Name).Context(s.GetContext()).Do()
	}, func(op interface{}) bool {
		return op.(*servicemanagement.Operation).Done
	})

	if err != nil {
		return err
	}

	op = o.(*servicemanagement.Operation)
	if op.Error == nil {
		return nil
	}

	return errors.Errorf("%s failed. Code %d: %s", op.Name, op.Error.Code, op.Error.Message)
}
