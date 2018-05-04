// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

import (
	"bytes"
	"fmt"
	"text/template"

	"chromium.googlesource.com/enterprise/cel/go/common"
	deploymentmanager "google.golang.org/api/deploymentmanager/v2beta"
	servicemanagement "google.golang.org/api/servicemanagement/v1"
	adminpb "google.golang.org/genproto/googleapis/iam/admin/v1"
)

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
	defer GcpLoggedServiceAction(s, ServiceManagementServiceName, &err,
		"Enabling required services for GCP project \"%s\"", s.GetProject())()

	sm, err := servicemanagement.New(s.client)
	if err != nil {
		return err
	}

	required := make(map[string]bool)
	for _, api := range RequiredGcpServices {
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

func publishServiceAccount(ctx common.Context, s *Session) (err error) {
	iamClient, err := s.GetIamClient()
	if err != nil {
		return err
	}

	saName := ServiceAccountResource(s.GetProject(), ServiceAccountEmail(s.GetProject(), ServiceAccountId))
	sa, err := iamClient.GetServiceAccount(ctx, &adminpb.GetServiceAccountRequest{Name: saName})
	if err != nil {
		return err
	}

	return ctx.Publish(s.HostEnvironment.Resources, "service_account", sa)
}

func launchBaseAssets(ctx common.Context, s *Session) (err error) {
	defer GcpLoggedServiceAction(s, DeploymentManagerServiceName, &err,
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

	return publishServiceAccount(ctx, s)
}

func uploadNamedResource(ctx common.Context, s *Session, name string) (fr *common.FileReference, err error) {
	defer common.LoggedAction(ctx, &err, "uploading %s", name)()
	data := _escFSMustByte(false, name)
	fr = &common.FileReference{}
	err = fr.StoreFile(ctx, data)
	if err != nil {
		return nil, err
	}
	return fr, nil
}

func uploadStartupDependencies(ctx common.Context, s *Session) error {
	fr, err := uploadNamedResource(ctx, s, "/windows/instance-startup.ps1")
	if err != nil {
		return err
	}

	err = ctx.Publish(s.HostEnvironment.Resources, "win_startup", fr)
	if err != nil {
		return err
	}

	fr, err = uploadNamedResource(ctx, s, "/windows/gen/windows_amd64/cel_agent.exe")
	if err != nil {
		return err
	}

	err = ctx.Publish(s.HostEnvironment.Resources, "win_agent_x64", fr)
	if err != nil {
		return err
	}
	return nil
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

	err = uploadStartupDependencies(ctx, s)
	if err != nil {
		return err
	}

	return nil
}
