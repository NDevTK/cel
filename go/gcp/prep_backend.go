// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

import (
	"bytes"
	"chromium.googlesource.com/enterprise/cel/go/common"
	protoIam "chromium.googlesource.com/enterprise/cel/go/gcp/iam"
	"fmt"
	deploymentmanager "google.golang.org/api/deploymentmanager/v2beta"
	servicemanagement "google.golang.org/api/servicemanagement/v1"
	"text/template"
)

const (
	computeServiceName              = "compute.googleapis.com"
	deploymentManagerServiceName    = "deploymentmanager.googleapis.com"
	iamServiceName                  = "iam.googleapis.com"
	loggingServiceName              = "logging.googleapis.com"
	cloudKmsServiceName             = "cloudkms.googleapis.com"
	monitoringServiceName           = "monitoring.googleapis.com"
	serviceManagementServiceName    = "servicemanagement.googleapis.com"
	cloudResourceManagerServiceName = "cloudresourcemanager.googleapis.com"
)

// These are the APIs that are required by the CEL toolchain. During deployment
// these APIs will be automatically enabled on new projects.
var RequiredGcpApis = []string{
	cloudKmsServiceName,
	computeServiceName,
	deploymentManagerServiceName,
	iamServiceName,
	loggingServiceName,
	monitoringServiceName,
	cloudResourceManagerServiceName,
}

// The name of the GCP Deployment Manager Deployment that is used to create
// some of the base assets.
//
// Unfortunately, not all required asset types are currently supported by the
// Deployment Manager. So some of them have to be constructed "manually".
const BaseAssetDeploymentName string = "cel-base"

// The account ID for the CEL instance service account.
const ServiceAccountId string = "cel-instance-service"

type celBaseTemplateParams struct {
	ProjectId          string
	ServiceAccountId   string
	ServiceAccountName string
}

func enableRequiredAPIs(ctx common.Context, s *Session) (err error) {
	defer GcpLoggedServiceAction(s, serviceManagementServiceName, &err,
		"Enabling required services for GCP project \"%s\"", s.GetProject())()

	sm, err := servicemanagement.New(s.client)
	if err != nil {
		return err
	}

	required := make(map[string]bool)
	for _, api := range RequiredGcpApis {
		required[api] = true
	}

	// The consumer ID as understood by the services API.
	consumerId := fmt.Sprintf("project:%s", s.GetProject())

	// First try to list the enabled services and remove any APIs that are
	// already enabled.
	lr, err := sm.Services.List().ConsumerId(consumerId).Context(ctx).Do()
	if err == nil {
		for _, svc := range lr.Services {
			if _, ok := required[svc.ServiceName]; ok {
				s.Logger.Debug(common.MakeStringer("Service %s is already enabled", svc.ServiceName))
				delete(required, svc.ServiceName)
			}
		}
	}

	if len(required) == 0 {
		return nil
	}

	j := common.NewTasks(nil)

	// Parallelize API activations. These tend to take a while. It is also
	// possible that the API usage will fail immediately after enabling the
	// activation. So far the only remedy is to wait and retry.
	for api, _ := range required {
		a := api
		j.Go(func() error {
			s.Logger.Debug(common.MakeStringer("Enabling service %s", a))

			req := &servicemanagement.EnableServiceRequest{
				ConsumerId: fmt.Sprintf("project:%s", s.GetProject()),
			}
			op, err := sm.Services.Enable(a, req).Context(ctx).Do()
			if err != nil {
				return err
			}
			return JoinServiceManagementOperation(s, op)
		})
	}

	return j.Join()
}

func launchBaseAssets(ctx common.Context, s *Session) (err error) {
	defer GcpLoggedServiceAction(s, deploymentManagerServiceName, &err,
		"Launching base assets for lab in GCP project \"%s\"", s.GetProject())()

	dm, err := s.GetDeploymentManagerService()
	if err != nil {
		return err
	}

	op, err := dm.Deployments.Delete(s.GetProject(), BaseAssetDeploymentName).Context(ctx).DeletePolicy("DELETE").Do()
	if err == nil {
		err = JoinDeploymentOperation(s, op)

		// This step can fail if there was no deployment named "cel-base" to
		// begin with. If so, we are going to ignore the error.
		if err != nil && !IsNotFoundError(err) {
			return err
		}
	}

	ttext, err := _escFSString(false, "/deployment/cel-base.yaml")
	if err != nil {
		return err
	}

	tmpl := template.Must(template.New("cel-base.yaml").Parse(ttext))

	saName := ServiceAccountResource(s.GetProject(), ServiceAccountEmail(s.GetProject(), ServiceAccountId))
	var b bytes.Buffer
	err = tmpl.Execute(&b, celBaseTemplateParams{
		ProjectId:          s.GetProject(),
		ServiceAccountId:   ServiceAccountId,
		ServiceAccountName: saName,
	})
	if err != nil {
		return err
	}

	dep := &deploymentmanager.Deployment{
		Description: "Chrome Enterprise Lab : Base environment deployment",
		Name:        BaseAssetDeploymentName,
		Target: &deploymentmanager.TargetConfiguration{
			Config: &deploymentmanager.ConfigFile{
				Content: b.String(),
			},
		},
	}
	op, err = dm.Deployments.Insert(s.GetProject(), dep).Context(ctx).Do()
	if err != nil {
		return err
	}

	err = JoinDeploymentOperation(s, op)
	if err != nil {
		return err
	}

	dep, err = dm.Deployments.Get(s.GetProject(), BaseAssetDeploymentName).Context(ctx).Do()
	if err != nil {
		return err
	}

	iam, err := s.GetIamService()
	if err != nil {
		return err
	}

	sa, err := iam.Projects.ServiceAccounts.Get(saName).Context(s.GetContext()).Do()
	if err != nil {
		return err
	}

	var pIam *protoIam.ServiceAccount
	err = common.HomomorphicCopy(&sa, &pIam)
	if err != nil {
		return err
	}
	return ctx.Publish(s.HostEnvironment.Resources, "service_account", pIam)
}

func PrepBackend(ctx common.Context, s *Session) error {
	err := enableRequiredAPIs(ctx, s)
	if err != nil {
		return err
	}

	err = launchBaseAssets(ctx, s)
	if err != nil {
		return err
	}

	return nil
}
