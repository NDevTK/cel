// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

const (
	ComputeServiceName              = "compute.googleapis.com"
	DeploymentManagerServiceName    = "deploymentmanager.googleapis.com"
	IamServiceName                  = "iam.googleapis.com"
	LoggingServiceName              = "logging.googleapis.com"
	CloudKmsServiceName             = "cloudkms.googleapis.com"
	MonitoringServiceName           = "monitoring.googleapis.com"
	ServiceManagementServiceName    = "servicemanagement.googleapis.com"
	CloudResourceManagerServiceName = "cloudresourcemanager.googleapis.com"
)

// These are the APIs that are required by the CEL toolchain. During deployment
// these APIs will be automatically enabled on new projects.
var RequiredGcpServices = []string{
	CloudKmsServiceName,
	ComputeServiceName,
	DeploymentManagerServiceName,
	IamServiceName,
	LoggingServiceName,
	MonitoringServiceName,
	CloudResourceManagerServiceName,
}
