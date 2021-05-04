// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

import (
	hostpb "chromium.googlesource.com/enterprise/cel/go/schema/host"
	th "chromium.googlesource.com/enterprise/cel/go/testhelpers"
	"context"
	"testing"
)

func TestQueryGceState(t *testing.T) {
	f := th.NewResponseFaker(t)

	f.Expect(
		th.RestRequest{Url: "https://compute.googleapis.com/compute/v1/projects/test-project/zones"},
		th.RestResponse{BodyPath: "./testdata/compute_zoneList.json"})

	f.Expect(
		th.RestRequest{Url: "https://compute.googleapis.com/compute/v1/projects/test-project/aggregated/instances"},
		th.RestResponse{BodyPath: "./testdata/compute_instanceAggregatedList.json"})

	f.Expect(
		th.RestRequest{Url: "https://compute.googleapis.com/compute/v1/projects/test-project/global/firewalls"},
		th.RestResponse{BodyPath: "./testdata/compute_firewallList.json"})

	f.Expect(
		th.RestRequest{Url: "https://compute.googleapis.com/compute/v1/projects/test-project/global/networks"},
		th.RestResponse{BodyPath: "./testdata/compute_networkList.json"})

	f.Expect(
		th.RestRequest{Url: "https://compute.googleapis.com/compute/v1/projects/test-project/aggregated/addresses"},
		th.RestResponse{BodyPath: "./testdata/compute_addressAggregatedList.json"})

	f.Expect(
		th.RestRequest{Url: "https://compute.googleapis.com/compute/v1/projects/test-project/aggregated/subnetworks"},
		th.RestResponse{BodyPath: "./testdata/compute_subnetworkAggregatedList.json"})

	f.Expect(
		th.RestRequest{Url: "https://iam.googleapis.com/v1/projects/test-project/serviceAccounts"},
		th.RestResponse{BodyObject: `{}`})

	env := hostpb.HostEnvironment{
		Project: &hostpb.Project{Name: "test-project", Zone: "test-zone"},
	}

	_, err := QueryCloudState(context.Background(), f.NewClient(), &env)
	if err != nil {
		t.Fatal(err)
	}
}
