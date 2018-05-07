// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"chromium.googlesource.com/enterprise/cel/go/cel"
	"chromium.googlesource.com/enterprise/cel/go/common"

	// The following must be imported here for side-effects.
	_ "chromium.googlesource.com/enterprise/cel/go/asset/deploy"
	// _ "chromium.googlesource.com/enterprise/cel/go/common/deploy"
	gcpDeploy "chromium.googlesource.com/enterprise/cel/go/gcp/deploy"
)

// Deploy is the starting point for a deployment. All the required parameters
// should already have been configured in the DeployerSession.
//
// It invokes the remainder of the workflow in a pre-determined order.
//
// TODO(asanka): Here's what we need to do here:
//
// * Make sure that we are calling checkNamespaceIsReady() correctly. In
//   particular, ensure that the checked resources are already resolved by the
//   time we call checkNamespaceIsReady().
//
// * Add tests for verifying that the object store invocations are correct, at
//   least as far as we can see.
//
// * Add BackendObjectGenerator resolver type that will emit a valid deployment
//   manifest entry for each resource that we support. Need to figure out the
//   intermediate format for this. It's likely going to be a
//   map[string]interface{} for starters.
//
// * Ensure that the resulting deployment manifest is correct.
//
// * Start expanding the cel_agent binary to support a new on-host
//   configuration session.
//
// * The on-host session should support its own workflow where it can invoke
//   its own set of resolvers. These resolvers should be distinct from the
//   deployment resolvers.
//
// * One such resolver should be a common one that listens and waits for a
//   dependent foreign asset to be resolved. Once its resolved, the resolver
//   can query any runtime properties to fulfil the object.
//
// * Have a few examples of where in the resolution process we invoke
//   PowerShell.
//
// * Have a few basic examples that handle restarts correctly.
//
// * Think about how we should handle failures in the lab.
//
// * Think about how we should handle login sessions. Can cel_ctl construct and
//   invoke RDP sessions to the VMs? What about non-Windows platforms?
func Deploy(d *Session) (err error) {
	defer common.LoggedAction(d.GetContext(), &err, "Deployment %s", d.config.Lab.GenerationId)()

	// Right off the bat, lab and host.log_settings must be complete.
	err = checkNamespaceIsReady(d.GetConfiguration().GetNamespace(),
		[]common.RefPath{
			common.RefPathMust("host.log_settings"),
			common.RefPathMust("lab.generation_id")})
	if err != nil {
		return err
	}

	// Allow resolvers to add additional dependencies before doing anything
	// else. Note that this is the last opportunity to add dependencies before
	// we prune the graph.
	err = InvokeAdditionalDependencyResolvers(d)
	if err != nil {
		return err
	}

	// Drop any assets that aren't connected by now.
	err = Prune(d)
	if err != nil {
		return err
	}

	err = InvokeImmediateResolvers(d)
	if err != nil {
		return err
	}

	// host.project, and host.storage should be complete by now.
	err = checkNamespaceIsReady(d.GetConfiguration().GetNamespace(),
		[]common.RefPath{common.RefPathMust("host.project"), common.RefPathMust("host.storage")})
	if err != nil {
		return err
	}

	err = DeleteObsoleteDeployments(d)
	if err != nil {
		return err
	}

	err = StopAllVMs(d)
	if err != nil {
		return err
	}

	err = PrepBackend(d)
	if err != nil {
		return err
	}

	err = checkNamespaceIsReady(d.GetConfiguration().GetNamespace(),
		[]common.RefPath{common.RefPathMust("host.resources")})
	if err != nil {
		return err
	}

	err = InvokeGeneratedContentResolvers(d)
	if err != nil {
		return err
	}

	err = InvokeIndexedObjectResolvers(d)
	if err != nil {
		return err
	}

	err = InvokePreCompleteManifestResolvers(d)
	if err != nil {
		return err
	}

	err = checkNamespaceIsReady(d.GetConfiguration().GetNamespace(),
		[]common.RefPath{common.RefPathMust("asset"), common.RefPathMust("host")})
	if err != nil {
		return err
	}

	err = VerifyCompletedAssetManifest(d)
	if err != nil {
		return err
	}

	err = UpdateProjectMetadata(d)
	if err != nil {
		return err
	}

	err = GenerateDeploymentManifest(d)
	if err != nil {
		return err
	}

	return InvokeDeploymentManager(d)
}

// InvokeAdditionalDependencyResolvers step adds any explicit dependences that were
// not explicitly set in the input asset manifest. This step is needed so that
// required assets and host environment elements don't disappear during the
// pruning phase.
func InvokeAdditionalDependencyResolvers(d *Session) (err error) {
	defer common.LoggedAction(d.GetContext(), &err, "ResolveAdditionalDependencies")()

	return common.ApplyResolvers(d.ctx, d.config.GetNamespace(), common.AdditionalDependencyResolverKind)
}

// Prune removes unneeded resources from the namespace.
func Prune(d *Session) (err error) {
	defer common.LoggedAction(d.GetContext(), &err, "Prune")()

	// TODO(asanka): The operator should be able to override the "asset" entry
	// if they just want to deploy a subset of the assets from a manifest.
	var anchors = []common.RefPath{
		common.RefPathMust("asset"),        // Unless overridden, "asset" is the thing we are deploying.
		common.RefPathMust("lab"),          // "lab" contains global metadata.
		common.RefPathMust("host.project"), // All assets implicitly depend on "project", "storage", and "log_settings".
		common.RefPathMust("host.storage"),
		common.RefPathMust("host.resources"),
		common.RefPathMust("host.log_settings"),
	}

	return d.config.GetNamespace().Prune(anchors)
}

// InvokeImmediateResolvers is the discovery phase where the GCP logic performs
// lookups for project metadata.
func InvokeImmediateResolvers(d *Session) (err error) {
	defer common.LoggedAction(d.GetContext(), &err, "InvokeImmediateResolvers")()

	return common.ApplyResolvers(d.ctx, d.config.GetNamespace(), common.ImmediateResolverKind)
}

// InvokeGeneratedContentResolvers step generates assets that must be generated as
// deployemnt time. These are assets like generated passwords and certificates.
func InvokeGeneratedContentResolvers(d *Session) (err error) {
	defer common.LoggedAction(d.GetContext(), &err, "ResolveGeneratedContent")()

	return common.ApplyResolvers(d.ctx, d.config.GetNamespace(), common.GeneratedContentResolverKind)
}

// DeleteObsoleteDeployments removes any deployments that are not needed. The
// CEL toolchain assumes that all the deployments that exist within the project
// are under its control. If any deployments are not recognized, the toolchain
// will remove them.
func DeleteObsoleteDeployments(d *Session) (err error) {
	defer common.LoggedAction(d.GetContext(), &err, "DeleteObsoleteDeployments")()

	if true {
		return nil
	}
	return cel.NewNotImplementedError("DeleteObsoleteDeployments")
}

// StopAllVMs stops all running instances. This step ensures that even if we
// reuse instance, they will still pick up the correct up-to-date configuration
// that is about to be applied to the project. Configuration is typically only
// reqd during instance startup.
func StopAllVMs(d *Session) (err error) {
	defer common.LoggedAction(d.GetContext(), &err, "StopAllVMs")()

	if true {
		return nil
	}
	return cel.NewNotImplementedError("StopAllVMs")
}

// PrepBackend prepares the base set of services on the hosting environment.
//
// In this phase, the toolchain ensures that there are service accounts and KMS
// keys as required by the deployment process. It also enables the services and
// APIs on the target project that are required by the CEL toolchain.
func PrepBackend(d *Session) (err error) {
	defer common.LoggedAction(d.GetContext(), &err, "PrepBackend")()

	return gcpDeploy.PrepBackend(d.ctx, d.backend)
}

// InvokeIndexedObjectResolvers uploads blobs of data to the object store in order to
// make them available to lab VMs. Currently these are FileReference and Secret
// objects.
func InvokeIndexedObjectResolvers(d *Session) (err error) {
	defer common.LoggedAction(d.GetContext(), &err, "ResolveIndexedObjects")()

	return common.ApplyResolvers(d.ctx, d.config.GetNamespace(), common.IndexedObjectResolverKind)
}

// InvokePreCompleteManifestResolvers step cleans up assets prior to them being
// included in the final completed asset manifest.
func InvokePreCompleteManifestResolvers(d *Session) (err error) {
	defer common.LoggedAction(d.GetContext(), &err, "ResolvePreManifestCompletion")()

	return common.ApplyResolvers(d.ctx, d.config.GetNamespace(), common.PreCompleteManifestResolverKind)
}

// VerifyCompletedAssetManifest ensures that all OUTPUT fields in the namespace
// have values. It then serializes and uploads the asset manifest to the
// project's object storage.
func VerifyCompletedAssetManifest(d *Session) (err error) {
	defer common.LoggedAction(d.GetContext(), &err, "VerifyCompletedAssetManifest")()

	manifest, err := d.config.GenerateCompletedManifest()
	if err != nil {
		return
	}

	return d.config.Manifest.StoreFile(d.GetContext(), manifest)
}

// UpdateProjectMetadata sets the project scoped metadata.
func UpdateProjectMetadata(d *Session) (err error) {
	defer common.LoggedAction(d.GetContext(), &err, "UpdateProjectMetadata")()

	return gcpDeploy.UpdateProjectMetadata(d.ctx, d.backend, &d.config.Manifest)
}

// GenerateDeploymentManifest emits the deployment manifest for lab assets.
func GenerateDeploymentManifest(d *Session) error {
	return cel.NewNotImplementedError("GenerateDeploymentManifest")
}

// InvokeDeploymentManager uploads and creates a new deployment based on the
// manifest that was generated in the prior steps.
func InvokeDeploymentManager(d *Session) error {
	return cel.NewNotImplementedError("InvokeDeploymentManager")
}

func checkNamespaceIsReady(r *common.Namespace, ns []common.RefPath) error {
	for _, p := range ns {
		if !r.Ready(p) {
			return cel.NewNotReadyError(r, p)
		}
	}
	return nil
}
