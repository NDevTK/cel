// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

import (
	"bytes"
	"fmt"
	"text/template"

	"chromium.googlesource.com/enterprise/cel/go/common"
	cloudkmspb "chromium.googlesource.com/enterprise/cel/go/gcp/cloudkms"
	cloudkms "google.golang.org/api/cloudkms/v1"
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

// CloudKMS KeyRing ID
const kmsKeyRingId string = "cel"

// CloudKMS CryptoKey ID.
const kmsCryptoKeyId string = "cel-manifest-key"

// All of our KMS resources are global.
// TODO(asanka): Consider restricting to the project's zone.
const kmsLocationId = "global"

// IAM role on the KMS crypto key to assign the instance service account.
const kmsRoleIdForServiceAccount = "roles/cloudkms.cryptoKeyDecrypter"

// Label key name for storing subresource integrity tokens.
const integrityTokenKey string = "integrity"

// enableRequiredAPIs checks and enables the list of GCP services and APIs that
// are needed by the CEL toolchain. The authoritative list of services is
// requiredGcpServices.
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

// constructBaseDeploymentManifest constructs the GCP deployment manifest for
// the set of base resources that must be present in every GCP project hosting
// a CEL lab.
func constructBaseDeploymentManifest(s *Session) ([]byte, error) {
	// Resource name for the deployment manifest *template*. The template itself
	// will be evaluated as a text/template with the a single parameter of type
	// celBaseTemplateParams.
	const deploymentManifestTemplate string = "/deployment/cel-base.yaml"

	// celBaseTemplateParams describes the inputs against which the
	// DeploymentManifestTemplate is evaluated. Note that the parameter types are
	// chosen to be strings so that the template evaluation doesn't have to do any
	// magical formatting.
	type celBaseTemplateParams struct {
		// GCP project ID. E.g. my-cel-project
		ProjectId string

		// Service account ID. E.g. cel-instance-service (not an email address)
		ServiceAccountId string

		// Full or partial service account resource name. E.g.
		// /projects/my-cel-project/serviceAccounts/...
		ServiceAccountName string
	}

	ttext, err := _escFSString(false, deploymentManifestTemplate)
	if err != nil {
		return nil, err
	}

	tmpl := template.Must(template.New("cel-base").Parse(ttext))

	var b bytes.Buffer
	saName := ServiceAccountResource(s.GetProject(), ServiceAccountEmail(s.GetProject(), ServiceAccountId))
	err = tmpl.Execute(&b, celBaseTemplateParams{
		ProjectId:          s.GetProject(),
		ServiceAccountId:   ServiceAccountId,
		ServiceAccountName: saName,
	})

	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
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

func deployKmsKey(ctx common.Context, s *Session) (err error) {
	defer GcpLoggedServiceAction(s, CloudKmsServiceName, &err, "Deploying KMS keys")()

	kms, err := s.GetCloudKmsService()
	if err != nil {
		return err
	}

	locName := KmsLocationResource(s.GetProject(), kmsLocationId)

	// First make sure our keyring exists.
	krName := locName + "/keyRings/" + kmsKeyRingId
	kr, err := kms.Projects.Locations.KeyRings.Get(krName).Context(ctx).Do()
	if IsNotFoundError(err) {
		kr, err = kms.Projects.Locations.KeyRings.
			Create(locName, &cloudkms.KeyRing{}).
			KeyRingId(kmsKeyRingId).
			Context(ctx).Do()
	}
	if err != nil {
		return err
	}

	// Then make sure our cryptokey exists.
	ckName := kr.Name + "/cryptoKeys/" + kmsCryptoKeyId
	ck, err := kms.Projects.Locations.KeyRings.CryptoKeys.Get(ckName).Context(ctx).Do()
	if IsNotFoundError(err) {
		ck, err = kms.Projects.Locations.KeyRings.CryptoKeys.
			Create(krName, &cloudkms.CryptoKey{
				Purpose: "ENCRYPT_DECRYPT",
			}).
			CryptoKeyId(kmsCryptoKeyId).
			Context(ctx).
			Do()
	}
	if err != nil {
		return err
	}

	// Verify ACLs
	pol, err := kms.Projects.Locations.KeyRings.CryptoKeys.GetIamPolicy(ck.Name).Context(ctx).Do()
	if err != nil {
		return err
	}

	var binding *cloudkms.Binding
	for _, b := range pol.Bindings {
		if b.Role == kmsRoleIdForServiceAccount {
			binding = b
			break
		}
	}
	if binding == nil {
		binding = &cloudkms.Binding{Role: kmsRoleIdForServiceAccount}
		pol.Bindings = append(pol.Bindings, binding)
	}

	aclLine := ServiceAccountAclBinding(s.HostEnvironment.Resources.ServiceAccount.Email)
	found := false
	for _, m := range binding.Members {
		if m == aclLine {
			found = true
			break
		}
	}

	if !found {
		binding.Members = append(binding.Members, aclLine)
		_, err = kms.Projects.Locations.KeyRings.CryptoKeys.
			SetIamPolicy(ck.Name, &cloudkms.SetIamPolicyRequest{Policy: pol, UpdateMask: "bindings"}).
			Context(ctx).Do()
	}

	var ckPb *cloudkmspb.CryptoKey
	err = common.HomomorphicCopy(&ck, &ckPb)
	if err != nil {
		return err
	}

	return ctx.Publish(s.HostEnvironment.Resources, "crypto_key", ckPb)
}

// deployBaseAssets ensures that the base set of assets required by the CEL
// toolchain are present in the target GCP project, deploying them if
// necessary.
func deployBaseAssets(ctx common.Context, s *Session) (err error) {
	defer GcpLoggedServiceAction(s, DeploymentManagerServiceName, &err,
		"Launching base assets for lab in GCP project \"%s\"", s.GetProject())()

	dm, err := s.GetDeploymentManagerService()
	if err != nil {
		return err
	}

	manifest, err := constructBaseDeploymentManifest(s)
	if err != nil {
		return err
	}

	reDeploy := true

	dep, err := dm.Deployments.Get(s.GetProject(), BaseAssetDeploymentName).Context(ctx).Do()
	if err == nil && (dep.Operation == nil || dep.Operation.Status == "DONE") {
		for _, l := range dep.Labels {
			if l.Key != integrityTokenKey {
				continue
			}

			if err := common.ValidateIntegrity(manifest, l.Value); err == nil {
				reDeploy = false
				break
			}
		}
	}

	if reDeploy {
		op, err := dm.Deployments.Delete(s.GetProject(), BaseAssetDeploymentName).
			Context(ctx).DeletePolicy("DELETE").Do()
		if err == nil {
			err = JoinDeploymentOperation(s, op)

			// This step can fail if there was no deployment named "cel-base" to
			// begin with. If so, we are going to ignore the error.
			if err != nil && !IsNotFoundError(err) {
				return err
			}
		}

		token := common.IntegrityLabel(manifest)
		dep = &deploymentmanager.Deployment{
			Description: "Chrome Enterprise Lab : Base environment deployment",
			Name:        BaseAssetDeploymentName,
			Target: &deploymentmanager.TargetConfiguration{
				Config: &deploymentmanager.ConfigFile{
					Content: string(manifest),
				},
			},
			Labels: []*deploymentmanager.DeploymentLabelEntry{
				{Key: integrityTokenKey, Value: token},
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
	}

	err = publishServiceAccount(ctx, s)
	if err != nil {
		return err
	}

	return deployKmsKey(ctx, s)
}

// uploadNamedResource fetches a named resource and uploads it to the
// ObjectStore.
//
// embeddedResource is the string identifying the embedded resources. E.g.
// /windows/instance-startup.ps1.
//
// fieldName is the name of the field in HostEnvironment.Resources where the
// resulting FileReference should be published.
func uploadNamedResource(ctx common.Context, s *Session, embeddedResource, fieldName string) (err error) {
	defer common.LoggedAction(ctx, &err, "uploading %s", embeddedResource)()

	data := _escFSMustByte(false, embeddedResource)
	fr := &common.FileReference{}
	err = fr.StoreFile(ctx, data)
	if err != nil {
		return err
	}

	return ctx.Publish(s.HostEnvironment.Resources, fieldName, fr)
}

// uploadStartupDependencies uploads the assets that are used during VM
// instance startup.
func uploadStartupDependencies(ctx common.Context, s *Session) error {
	err := uploadNamedResource(ctx, s, "/windows/instance-startup.ps1", "win_startup")
	if err != nil {
		return err
	}

	return uploadNamedResource(ctx, s, "/windows/gen/windows_amd64/cel_agent.exe", "win_agent_x64")
}

// PrepBackend prepares the backend for hosting a lab. The resources deployed
// here are not specific to inputs, but are based solely on the version of the
// toolchain.
func PrepBackend(ctx common.Context, s *Session) error {
	err := enableRequiredAPIs(ctx, s)
	if err != nil {
		return err
	}

	err = deployBaseAssets(ctx, s)
	if err != nil {
		return err
	}

	err = uploadStartupDependencies(ctx, s)
	if err != nil {
		return err
	}

	return nil
}
