// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

// Service names
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

// The name of the GCP Deployment Manager Deployment that is used to create
// some of the base assets.
//
// Unfortunately, not all required asset types are currently supported by the
// Deployment Manager. So some of them have to be constructed "manually".
const BaseAssetDeploymentName string = "cel-base"

// The account ID for the CEL instance service account.
const ServiceAccountId string = "cel-instance-service"

// CloudKMS KeyRing ID
const KmsKeyRingId string = "cel"

// CloudKMS CryptoKey ID.
const KmsCryptoKeyId string = "cel-manifest-key"

// All of our KMS resources are global.
// TODO(asanka): Consider restricting to the project's zone.
const KmsLocationId = "global"

// IAM role on the KMS crypto key to assign the instance service account.
const KmsRoleIdForServiceAccount = "roles/cloudkms.cryptoKeyDecrypter"

// Label key name for storing subresource integrity tokens.
const IntegrityTokenKey string = "integrity"

// Project scoped metadata key for cel-manifest GCS path.
const CelManifestMetadataKey string = "cel-manifest"

// Project scoped metadata key for cel-agent.
const CelAgentMetadataKey string = "cel-agent"

// Project scoped metadata key for cel-admin-log.
const CelAdminLogMetadataKey string = "cel-admin-log"
