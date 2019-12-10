// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// This module handles Environment Snapshots, a feature set to handle
// creating, restoring and managing CELab Environment Snapshots.
// (note: we only support saving/restoring instance disks)
//
// This is particularly useful in cases where manual steps are needed to finish
// configuring an environment after running "cel_ctl deploy".
//
// An EnvironmentSnapshot is a collection of Compute images used to restore
// instances to a previously saved state. CELab identifies those environment
// snapshots with a special prefix ("celab-{name}") and instances within that
// environment snapshot with a "--" delimiter (can't be used in a snapshot name).
//
// To create an environment snapshot, we create a Compute images for every
// instance in a given project. To restore these instances later, we'll also
// need to reassign their original internalIp. This data is stored alongside
// the Compute image in the immutable Description field.
//
// To restore the environment snapshot XYZ, we look for all Compute images
// prefixed with "cel-XYZ--" and compare instance names against those in
// the expected asset schema. Assuming those match, we inject the images
// in the deployment definition schema and proceed with a standard deployment.

package cel

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"path"
	"regexp"
	"strings"

	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/gcp"
	"chromium.googlesource.com/enterprise/cel/go/gcp/deploy"
	computepb "chromium.googlesource.com/enterprise/cel/go/schema/gcp/compute"
	runtimeconfigpb "chromium.googlesource.com/enterprise/cel/go/schema/gcp/runtimeconfig"
	"github.com/golang/protobuf/proto"
	compute "google.golang.org/api/compute/v1"
)

var (
	ErrSnapshotNotFound = errors.New("can't find environment snapshot")
)

type InstanceMetadata struct {
	InstanceName string
	NetworkIP    string
}

type InstanceSnapshot struct {
	Image    string
	Metadata *InstanceMetadata
}

type EnvironmentSnapshot struct {
	Name string

	// Map of all instances in an environment keyed by InstanceName.
	Instances map[string]InstanceSnapshot

	// The date that the environment snapshot was created in RFC3339 text format.
	CreationTimestamp string
}

type instanceOperation func(ctx common.Context, service *compute.InstancesService, project string, zone string, instance *compute.Instance) (*compute.Operation, error)

// Stops Compute Instances in the host project in `session`.
func GetRunningInstances(ctx common.Context, session *gcp.Session) (runningInstances []*compute.Instance, err error) {
	defer common.LoggedAction(ctx, &err, "GetRunningInstances")()

	computeService, err := session.GetComputeService()
	if err != nil {
		return nil, err
	}

	service := compute.NewInstancesService(computeService)
	project := session.GetProject()
	zone := session.HostEnvironment.GetProject().GetZone()

	instances, err := service.List(project, zone).Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	// Find all RUNNING instances.
	runningInstances = make([]*compute.Instance, 0)
	for _, instance := range instances.Items {
		if instance.Status == "RUNNING" {
			runningInstances = append(runningInstances, instance)
		}
	}

	return runningInstances, nil
}

// Stops all Compute Instances passed as argument.
func StopInstances(ctx common.Context, session *gcp.Session, instances []*compute.Instance) (err error) {
	defer common.LoggedAction(ctx, &err, "StopInstances")()

	stopInstance := func(ctx common.Context, service *compute.InstancesService, project string, zone string, instance *compute.Instance) (*compute.Operation, error) {
		session.Logger.Debug(common.MakeStringer("Stopping instance %s", instance.Name))
		return service.Stop(project, zone, instance.Name).Context(ctx).Do()
	}

	return DoOperationOnInstances(ctx, session, instances, stopInstance)
}

// Starts all Compute Instances passed as argument.
func StartInstances(ctx common.Context, session *gcp.Session, instances []*compute.Instance) (err error) {
	defer common.LoggedAction(ctx, &err, "StartInstances")()

	startInstance := func(ctx common.Context, service *compute.InstancesService, project string, zone string, instance *compute.Instance) (*compute.Operation, error) {
		session.Logger.Debug(common.MakeStringer("Starting instance %s", instance.Name))
		return service.Start(project, zone, instance.Name).Context(ctx).Do()
	}

	return DoOperationOnInstances(ctx, session, instances, startInstance)
}

// Executes an operation on Compute Instances passed as argument.
func DoOperationOnInstances(ctx common.Context, session *gcp.Session, instances []*compute.Instance, doInstanceOperation instanceOperation) (err error) {
	computeService, err := session.GetComputeService()
	if err != nil {
		return err
	}

	service := compute.NewInstancesService(computeService)
	project := session.GetProject()
	zone := session.HostEnvironment.GetProject().GetZone()

	// Start all operations on all instances - we'll wait later.
	operations := make([]*compute.Operation, 0, len(instances))
	for _, instance := range instances {
		operation, err := doInstanceOperation(ctx, service, project, zone, instance)
		if err != nil {
			return err
		}

		operations = append(operations, operation)
	}

	// Wait for all instance operations to complete.
	session.Logger.Debug(common.MakeStringer("Waiting for %d operations to finish...", len(operations)))
	for _, operation := range operations {
		err := gcp.JoinOperation(session, operation, "Waiting for operation to finish")
		if err != nil {
			return err
		}
	}

	return nil
}

// Create snapshots of every Compute Instances in the host project in `session`.
func CreateSnapshots(ctx common.Context, session *gcp.Session, snapshotName string) (err error) {
	defer common.LoggedAction(ctx, &err, "CreateSnapshots")()

	computeService, err := session.GetComputeService()
	if err != nil {
		return err
	}

	service := compute.NewInstancesService(computeService)
	project := session.GetProject()
	zone := session.HostEnvironment.GetProject().GetZone()

	instances, err := service.List(project, zone).Context(ctx).Do()
	if err != nil {
		return err
	}

	// Start the operations for creating an image for every compute instance we have.
	imagesService := compute.NewImagesService(computeService)
	operations := make([]*compute.Operation, 0, len(instances.Items))
	for _, instance := range instances.Items {
		session.Logger.Debug(common.MakeStringer("Creating image for instance %s", instance.Name))

		// Images only support [a-z0-9-]. It's ok to modify instanceName if needed
		// because the one we'll use for Restore will come from InstanceMetadata.
		imageName := strings.ToLower(fmt.Sprintf("cel-%s--%s", snapshotName, instance.Name))

		// Save Snapshot Metadata alongside the image (NetworkIP for now)
		metadata := InstanceMetadata{InstanceName: instance.Name, NetworkIP: instance.NetworkInterfaces[0].NetworkIP}
		json, err := json.Marshal(metadata)
		if err != nil {
			return err
		}

		// Create the image
		insert, err := imagesService.Insert(project, &compute.Image{
			Name:        imageName,
			SourceDisk:  instance.Disks[0].Source,
			Description: string(json),
		}).Context(ctx).Do()
		if err != nil {
			return err
		}

		operations = append(operations, insert)
	}

	// Wait for all images creation operations to complete.
	session.Logger.Debug(common.MakeStringer("Waiting for %d images to be created...", len(operations)))
	maxRetries := 5
	for _, operation := range operations {
		retries := 0
		for {
			err := gcp.JoinOperation(session, operation, "Creating images")
			if err == nil {
				break
			}

			retries++
			if retries > maxRetries {
				return err
			}
		}
	}

	session.Logger.Debug(common.MakeStringer("All images created."))

	return nil
}

// Deploys a snapshot to the session host's celab environment.
// `snapshot` can be a snapshot that lives in a different Compute project.
// This function is eventually going to invoke the standard deployment flow,
// but only after injecting environment snapshot images/configuration.
func RestoreDeploymentConfigurationFromSnapshot(ctx common.Context, snapshot *EnvironmentSnapshot) (err error) {
	defer common.LoggedAction(ctx, &err, "RestoreDeploymentConfigurationFromSnapshot")()

	manifest := deploy.GetDeploymentManifest()

	resources := manifest.GetResources()
	for _, resource := range resources {
		if instance, ok := resource.Properties.(*computepb.Instance); ok {
			// Instances need to be modified to use the snapshot images and their original internal IP.
			instanceSnapshot, err := getInstanceInSnapshot(snapshot, instance.Name)
			if err != nil {
				return err
			}

			if len(instance.Disks) != 1 {
				return fmt.Errorf("We only support environment snapshots on single-disk machines.")
			}

			// Instances in the manifest share machine_type definitions, so we must create/edit a copy.
			diskProto := proto.Clone(instance.Disks[0])
			disk, ok := diskProto.(*computepb.AttachedDisk)
			if !ok {
				return fmt.Errorf("Failed to copy AttachedDisk for %s", instance.Name)
			}

			// NestedVMs don't use InitializeParams and have a distinct disk resource
			// where we'll inject the image (below).
			if disk.InitializeParams != nil {
				disk.InitializeParams.SourceImage = instanceSnapshot.Image
				instance.Disks = []*computepb.AttachedDisk{disk}
			}

			// Both normal instances and NestedVMs have own network interfaces, so we modify that directly.
			instance.NetworkInterfaces[0].NetworkIP = instanceSnapshot.Metadata.NetworkIP
		} else if disk, ok := resource.Properties.(*computepb.Disk); ok {
			// NestedVMs have separate disk entries where we must inject the image.
			parts := strings.Split(disk.Name, "-disk")
			if len(parts) != 2 || parts[1] != "" {
				// Skip the disk as it doesn't seem to be a nested VM disk.
				continue
			}

			instanceSnapshot, err := getInstanceInSnapshot(snapshot, parts[0])
			if err != nil {
				return err
			}

			disk.SourceImage = instanceSnapshot.Image
		} else if variable, ok := resource.Properties.(*runtimeconfigpb.Variable); ok {
			// All asset runtimeconfig variables should be set to "ready".
			variable.Text = "ready"
		}
	}

	return nil
}

func getInstanceInSnapshot(snapshot *EnvironmentSnapshot, instanceName string) (*InstanceSnapshot, error) {
	instanceSnapshot, ok := snapshot.Instances[instanceName]
	if !ok {
		return nil, fmt.Errorf("Missing image for compute instance %s", instanceName)
	} else if instanceSnapshot.Metadata == nil {
		return nil, fmt.Errorf("Missing metadata for compute instance %s", instanceName)
	}

	return &instanceSnapshot, nil
}

// Get a list of EnvironmentSnapshots in a given project.
func GetAllEnvironmentSnapshots(ctx context.Context, client *http.Client, project string) ([]*EnvironmentSnapshot, error) {
	computeService, err := compute.New(client)
	if err != nil {
		return nil, err
	}

	imagesService := compute.NewImagesService(computeService)

	images, err := imagesService.List(project).OrderBy("name").Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	// Find image names that fit cel's format. That's how we identify them.
	re := regexp.MustCompile("cel-(.*?)--(.*)")
	var snapshots []*EnvironmentSnapshot
	var snapshot *EnvironmentSnapshot
	for _, image := range images.Items {
		parts := re.FindStringSubmatch(image.Name)
		if len(parts) != 3 {
			continue
		}

		snapshotName := parts[1]
		instanceName := parts[2]

		if snapshot == nil || snapshot.Name != snapshotName {
			snapshot = &EnvironmentSnapshot{
				Name:              snapshotName,
				Instances:         make(map[string]InstanceSnapshot),
				CreationTimestamp: image.CreationTimestamp,
			}
			snapshots = append(snapshots, snapshot)
		}

		pathToImage := fmt.Sprintf("projects/%s/global/images/%s", project, image.Name)
		instance := InstanceSnapshot{Image: pathToImage}

		// Try to parse Metadata but don't abort if it fails because we still want to be able to list/delete malformed snapshots.
		var metadata InstanceMetadata
		err := json.Unmarshal([]byte(image.Description), &metadata)
		if err == nil {
			instanceName = metadata.InstanceName
			instance.Metadata = &metadata
		}

		snapshot.Instances[instanceName] = instance
	}

	return snapshots, nil
}

// Find a specific project EnvironmentSnapshot by name.
func FindEnvironmentSnapshot(ctx context.Context, client *http.Client, project string, snapshotName string) (*EnvironmentSnapshot, error) {
	snapshots, err := GetAllEnvironmentSnapshots(ctx, client, project)
	if err != nil {
		return nil, err
	}

	for _, snapshot := range snapshots {
		if snapshot.Name == snapshotName {
			return snapshot, nil
		}
	}

	return nil, ErrSnapshotNotFound
}

// Deletes all images in an environment snapshot.
func DeleteEnvironmentSnapshot(ctx context.Context, client *http.Client, project string, snapshot *EnvironmentSnapshot) error {
	computeService, err := compute.New(client)
	if err != nil {
		return err
	}

	imagesService := compute.NewImagesService(computeService)

	for _, instance := range snapshot.Instances {
		_, err := imagesService.Delete(project, path.Base(instance.Image)).Context(ctx).Do()
		if err != nil {
			return err
		}
	}

	return nil
}
