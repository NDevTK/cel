// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"encoding/json"

	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/gcp"
	commonpb "chromium.googlesource.com/enterprise/cel/go/schema/common"
	gcpPb "chromium.googlesource.com/enterprise/cel/go/schema/gcp"
	compute "google.golang.org/api/compute/v1"
)

func UpdateProjectMetadata(ctx common.Context, s *gcp.Session, manifest *commonpb.FileReference) (err error) {
	cs, err := s.GetComputeService()
	if err != nil {
		return err
	}

	var m *compute.Metadata
	if s.HostEnvironment.GetProject().GetProject().GetCommonInstanceMetadata() == nil {
		m = &compute.Metadata{}
	} else {
		err = common.HomomorphicCopy(&s.HostEnvironment.Project.Project.CommonInstanceMetadata, &m)
		if err != nil {
			return err
		}
	}

	agentText, err := computeAgentMetadata(s)
	if err != nil {
		return err
	}

	bucket := s.HostEnvironment.Storage.Bucket

	modified := false
	modified = setMetadata(m, gcp.CelManifestMetadataKey,
		gcp.AbsoluteReference(bucket, manifest.ObjectReference)) || modified
	modified = setMetadata(m, gcp.CelAgentMetadataKey, string(agentText)) || modified
	modified = setMetadata(m, gcp.CelAdminLogMetadataKey,
		s.HostEnvironment.LogSettings.AdminLog) || modified

	if modified {
		o, err := cs.Projects.SetCommonInstanceMetadata(s.GetProject(), m).Context(ctx).Do()
		if err != nil {
			return err
		}

		return gcp.JoinOperation(s, o, "Updating project metadata")
	}

	return nil
}

func computeAgentMetadata(s *gcp.Session) ([]byte, error) {
	md := &gcpPb.CelAgentMetadata{}
	md.WinAgentX64 = &gcpPb.CelAgentMetadata_GCSObject{}
	md.WinAgentX64.AbsPath = gcp.AbsoluteReference(
		s.HostEnvironment.Storage.Bucket,
		s.HostEnvironment.Resources.Startup.WinAgentX64.ObjectReference)
	md.WinAgentX64.Integrity = s.HostEnvironment.Resources.Startup.WinAgentX64.Integrity

	md.WinUiAgentX64 = &gcpPb.CelAgentMetadata_GCSObject{}
	md.WinUiAgentX64.AbsPath = gcp.AbsoluteReference(
		s.HostEnvironment.Storage.Bucket,
		s.HostEnvironment.Resources.Startup.WinUiAgentX64.ObjectReference)
	md.WinUiAgentX64.Integrity = s.HostEnvironment.Resources.Startup.WinUiAgentX64.Integrity

	md.LinuxAgentX64 = &gcpPb.CelAgentMetadata_GCSObject{}
	md.LinuxAgentX64.AbsPath = gcp.AbsoluteReference(
		s.HostEnvironment.Storage.Bucket,
		s.HostEnvironment.Resources.Startup.LinuxAgentX64.ObjectReference)
	md.LinuxAgentX64.Integrity = s.HostEnvironment.Resources.Startup.LinuxAgentX64.Integrity

	return json.Marshal(md)
}

func setMetadata(m *compute.Metadata, key, value string) bool {
	for _, i := range m.Items {
		if i.Key == key {
			if *i.Value == value {
				return false
			}
			i.Value = &value
			return true
		}
	}

	m.Items = append(m.Items, &compute.MetadataItems{
		Key:   key,
		Value: &value,
	})
	return true
}
