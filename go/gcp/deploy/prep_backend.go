// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"text/template"

	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/gcp"
	commonpb "chromium.googlesource.com/enterprise/cel/go/schema/common"
	cloudkmspb "chromium.googlesource.com/enterprise/cel/go/schema/gcp/cloudkms"
	cloudkms "google.golang.org/api/cloudkms/v1"
	"google.golang.org/api/cloudresourcemanager/v1"
	deploymentmanager "google.golang.org/api/deploymentmanager/v2beta"
	servicemanagement "google.golang.org/api/servicemanagement/v1"
	serviceusage "google.golang.org/api/serviceusage/v1"
	adminpb "google.golang.org/genproto/googleapis/iam/admin/v1"
)

// enableRequiredAPIs checks and enables the list of GCP services and APIs that
// are needed by the CEL toolchain. The authoritative list of services is
// requiredGcpServices.
func enableRequiredAPIs(ctx common.Context, s *gcp.Session) (err error) {
	defer gcp.GcpLoggedServiceAction(s, gcp.ServiceManagementServiceName, &err,
		"Enabling required services for GCP project \"%s\"", s.GetProject())()

	sm, err := servicemanagement.New(s.GetHttpClient())
	if err != nil {
		return err
	}

	su, err := serviceusage.New(s.GetHttpClient())
	if err != nil {
		return err
	}

	required := make(map[string]bool)
	for _, api := range gcp.RequiredGcpServices {
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

			req := &serviceusage.EnableServiceRequest{}
			svc := fmt.Sprintf("projects/%s/services/%s", s.GetProject(), a)
			op, err := su.Services.Enable(svc, req).Context(ctx).Do()
			if err != nil {
				return err
			}
			return gcp.JoinOperation(s, op, fmt.Sprintf("enabling service %s", a))
		})
	}

	return j.Join()
}

// constructBaseDeploymentManifest constructs the GCP deployment manifest for
// the set of base resources that must be present in every GCP project hosting
// a CEL lab.
func constructBaseDeploymentManifest(s *gcp.Session) ([]byte, error) {
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
	saName := gcp.ServiceAccountResource(s.GetProject(), gcp.ServiceAccountEmail(s.GetProject(), gcp.ServiceAccountId))
	err = tmpl.Execute(&b, celBaseTemplateParams{
		ProjectId:          s.GetProject(),
		ServiceAccountId:   gcp.ServiceAccountId,
		ServiceAccountName: saName,
	})

	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func publishServiceAccount(ctx common.Context, s *gcp.Session) (err error) {
	iamClient, err := s.GetIamClient()
	if err != nil {
		return err
	}

	saName := gcp.ServiceAccountResource(s.GetProject(), gcp.ServiceAccountEmail(s.GetProject(), gcp.ServiceAccountId))
	sa, err := iamClient.GetServiceAccount(ctx, &adminpb.GetServiceAccountRequest{Name: saName})
	if err != nil {
		return err
	}

	return ctx.Publish(s.HostEnvironment.Resources, "service_account", sa)
}

func deployKmsKey(ctx common.Context, s *gcp.Session) (err error) {
	defer gcp.GcpLoggedServiceAction(s, gcp.CloudKmsServiceName, &err, "Deploying KMS keys")()

	kms, err := s.GetCloudKmsService()
	if err != nil {
		return err
	}

	locName := gcp.KmsLocationResource(s.GetProject(), gcp.KmsLocationId)

	// First make sure our keyring exists.
	krName := locName + "/keyRings/" + gcp.KmsKeyRingId
	kr, err := kms.Projects.Locations.KeyRings.Get(krName).Context(ctx).Do()
	if gcp.IsNotFoundError(err) {
		kr, err = kms.Projects.Locations.KeyRings.
			Create(locName, &cloudkms.KeyRing{}).
			KeyRingId(gcp.KmsKeyRingId).
			Context(ctx).Do()
	}
	if err != nil {
		return err
	}

	// Then make sure our cryptokey exists.
	ckName := kr.Name + "/cryptoKeys/" + gcp.KmsCryptoKeyId
	ck, err := kms.Projects.Locations.KeyRings.CryptoKeys.Get(ckName).Context(ctx).Do()
	if gcp.IsNotFoundError(err) {
		ck, err = kms.Projects.Locations.KeyRings.CryptoKeys.
			Create(krName, &cloudkms.CryptoKey{
				Purpose: "ENCRYPT_DECRYPT",
			}).
			CryptoKeyId(gcp.KmsCryptoKeyId).
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
		if b.Role == gcp.KmsRoleIdForServiceAccount {
			binding = b
			break
		}
	}
	if binding == nil {
		binding = &cloudkms.Binding{Role: gcp.KmsRoleIdForServiceAccount}
		pol.Bindings = append(pol.Bindings, binding)
	}

	aclLine := gcp.ServiceAccountAclBinding(s.HostEnvironment.Resources.ServiceAccount.Email)
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
func deployBaseAssets(ctx common.Context, s *gcp.Session) (err error) {
	defer gcp.GcpLoggedServiceAction(s, gcp.DeploymentManagerServiceName, &err,
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

	dep, err := dm.Deployments.Get(s.GetProject(), gcp.BaseAssetDeploymentName).Context(ctx).Do()
	if err == nil && (dep.Operation == nil || dep.Operation.Status == "DONE") {
		for _, l := range dep.Labels {
			if l.Key != gcp.IntegrityTokenKey {
				continue
			}

			if err := common.ValidateIntegrity(manifest, l.Value); err == nil {
				reDeploy = false
				break
			}
		}
	}

	if reDeploy {
		op, err := dm.Deployments.Delete(s.GetProject(), gcp.BaseAssetDeploymentName).
			Context(ctx).DeletePolicy("DELETE").Do()
		if err == nil {
			err = gcp.JoinOperation(s, op, "removing existing stale base asset deployment")

			// This step can fail if there was no deployment named "cel-base" to
			// begin with. If so, we are going to ignore the error.
			if err != nil && !gcp.IsNotFoundError(err) {
				return err
			}
		}

		token := common.IntegrityLabel(manifest)
		dep = &deploymentmanager.Deployment{
			Description: "Chrome Enterprise Lab : Base environment deployment",
			Name:        gcp.BaseAssetDeploymentName,
			Target: &deploymentmanager.TargetConfiguration{
				Config: &deploymentmanager.ConfigFile{
					Content: string(manifest),
				},
			},
			Labels: []*deploymentmanager.DeploymentLabelEntry{
				{Key: gcp.IntegrityTokenKey, Value: token},
			},
		}
		op, err = dm.Deployments.Insert(s.GetProject(), dep).Context(ctx).Do()
		if err != nil {
			return err
		}

		err = gcp.JoinOperation(s, op, "running base asset deployment")
		if err != nil {
			return err
		}

		// make the service account an edit so that it can access the start-up script.
		saEmail := gcp.ServiceAccountEmail(s.GetProject(), gcp.ServiceAccountId)
		err = AddServiceAccountBinding(ctx, s, saEmail, "roles/editor")
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

// Adds a binding for a service account to the project for `s`.
func AddServiceAccountBinding(ctx common.Context, s *gcp.Session, saEmail string, role string) (err error) {
	s.Logger.Debug(common.MakeStringer("Set IAM policy of service account '%s'", saEmail))

	cloudResourceManagerService, err := s.GetCloudResourceManagerService()
	if err != nil {
		return err
	}

	policy, err := cloudResourceManagerService.Projects.GetIamPolicy(s.GetProject(),
		&cloudresourcemanager.GetIamPolicyRequest{}).Context(ctx).Do()
	if err != nil {
		return err
	}

	policy, err = cloudResourceManagerService.Projects.SetIamPolicy(
		s.GetProject(),
		&cloudresourcemanager.SetIamPolicyRequest{
			Policy: &cloudresourcemanager.Policy{
				Bindings: append(policy.Bindings,
					&cloudresourcemanager.Binding{
						Role: role,
						Members: []string{
							fmt.Sprintf("serviceAccount:%s", saEmail),
						},
					}),
			},
		},
	).Context(ctx).Do()

	return err
}

// uploadLocalResource fetches a local resource and uploads it to the
// ObjectStore.
//
// localResource is the path to the local resource, relative to the current exe
// e.g. ./resource/cel_agent.exe.
//
// fieldName is the name of the field in HostEnvironment.Resources where the
// resulting FileReference should be published.
func uploadLocalResource(ctx common.Context, s *gcp.Session, localResource, fieldName string) (err error) {
	defer common.LoggedAction(ctx, &err, "uploading %s", localResource)()

	var baseDir string
	if baseDir, err = currentBinDir(); err != nil {
		return err
	}

	var data []byte
	if data, err = ioutil.ReadFile(path.Join(baseDir, localResource)); err != nil {
		return err
	}

	fr := &commonpb.FileReference{
		TargetPath: path.Join("/cel", path.Base(localResource)),
	}
	if err = common.StoreFile(ctx, fr, data); err != nil {
		return err
	}

	return ctx.Publish(s.HostEnvironment.Resources.Startup, fieldName, fr)
}

// uploadNamedResource fetches a named resource and uploads it to the
// ObjectStore.
//
// embeddedResource is the string identifying the embedded resources. E.g.
// /windows/instance-startup.ps1.
//
// fieldName is the name of the field in HostEnvironment.Resources where the
// resulting FileReference should be published.
func uploadNamedResource(ctx common.Context, s *gcp.Session, embeddedResource, fieldName string) (err error) {
	defer common.LoggedAction(ctx, &err, "uploading %s", embeddedResource)()

	data := _escFSMustByte(false, embeddedResource)

	fr := &commonpb.FileReference{
		TargetPath: path.Join("/cel", path.Base(embeddedResource)),
	}
	err = common.StoreFile(ctx, fr, data)
	if err != nil {
		return err
	}

	return ctx.Publish(s.HostEnvironment.Resources.Startup, fieldName, fr)
}

func currentBinDir() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}

// uploadStartupDependencies uploads the assets that are used during VM
// instance startup.
func uploadStartupDependencies(ctx common.Context, s *gcp.Session) error {
	err := uploadNamedResource(ctx, s, "/windows/instance-startup.ps1", "win_startup")
	if err != nil {
		return err
	}

	err = uploadLocalResource(ctx, s, "./resources/cel_agent.exe", "win_agent_x64")
	if err != nil {
		return err
	}

	err = uploadLocalResource(ctx, s, "./resources/cel_ui_agent.exe", "win_ui_agent_x64")
	if err != nil {
		return err
	}

	err = uploadNamedResource(ctx, s, "/linux/instance-startup.py", "linux_startup")
	if err != nil {
		return err
	}

	return uploadLocalResource(ctx, s, "./resources/cel_agent", "linux_agent_x64")
}

// PrepBackend prepares the backend for hosting a lab. The resources deployed
// here are not specific to inputs, but are based solely on the version of the
// toolchain.
func PrepBackend(ctx common.Context, s *gcp.Session) error {
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
