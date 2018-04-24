// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/host"
	"context"
	"github.com/pkg/errors"
	compute "google.golang.org/api/compute/v1"
	iam "google.golang.org/api/iam/v1"
	"net/http"
	"strings"
)

// A reasonably high limit such that a generic list of things isn't expected to
// exceed this.
const kMaxGenericListCount = 1000

// Maximum number of instances that we expect to see in a lab.
const kMaxInstanceListCount = kMaxGenericListCount

// Maximum number of service accounts we are going to support. It's just an
// arbitrarily large number that's big enough for anyone.
const kMaxServiceAccountCount = 100

// Maximum number of subnetworks for the entire projects. Also just an
// arbitrarily large number that should be big enough.
const kMaxSubnetworkListCount = 128

// Maximum number of reserved static IP addresses. Just setting this to the
// maximum number of instances since there can't be more reserved addresses
// than there are instances.
const kMaxAddressCount = kMaxGenericListCount

type NetworkAndRegion struct {
	Network string
	Region  string
}

type NetworkAndFirewall struct {
	Network  string
	Firewall string
}

// CloudState contains the last known state of a single project's GCP assets.
//
// For each asset type, there's a FooChanged member and a SetFooChanged()
// method. The method should be called if some consumer knows that the asset
// changed. This way a subsequent FetchStale() call can refresh just the pieces
// that need to be refreshed.
//
// If one needs to refresh the entire cache, the easiest method is to call
// QueryGceState() and construct a new cache.
//
// Annotated for easy serialization.
type CloudState struct {
	// GCP Project metadata.
	Project *compute.Project

	// All static external IP addresses reserved by the project.
	Addresses map[string]*compute.Address `json:"addresses"`

	// All service accounts for the project by unique ID.
	ServiceAccounts map[string]*iam.ServiceAccount `json:"serviceAccounts"`

	// Instances by instance label
	Instances map[string]*compute.Instance `json:"instances"`

	// Network by network label
	Networks map[string]*compute.Network `json:"networks"`

	// Subnetwork by network label, and region
	Subnetworks map[NetworkAndRegion]*compute.Subnetwork `json:"subnetworks"`

	// Firewalls by network label and firewall label.
	Firewalls map[NetworkAndFirewall]*compute.Firewall `json:"firewall"`

	// Images by image name. Only the Images that were requested are going to
	// be synchronized. The name of the image used here is the name used as the
	// image name used in |HostEnvironment|.
	Images map[string]*compute.Image `json:"image"`

	// Zones by zone label.
	Zones map[string]*compute.Zone `json:"zone"`

	// HostEnvironment corresponding to this CloudState instance. The project
	// and some of the cached metadata are specific to the host environment.
	// Should not be modified.
	HostEnvironment *host.HostEnvironment `json:"-"`
}

func (g *CloudState) FetchServiceAccounts(ctx context.Context, client *http.Client) (err error) {
	defer common.Action(&err, "querying service accounts")

	service, err := iam.New(client)
	if err != nil {
		return
	}

	ServiceAccounts := make(map[string]*iam.ServiceAccount)
	next_page_token := ""

	for {
		call := service.Projects.ServiceAccounts.List(ProjectResource(g.HostEnvironment.Project.Name))
		call.PageSize(kMaxServiceAccountCount).Context(ctx)
		if next_page_token != "" {
			call.PageToken(next_page_token)
		}
		l, err := call.Do()

		if IsNotFoundError(err) {
			return nil
		}

		if err != nil {
			return err
		}

		for _, a := range l.Accounts {
			name := a.Email
			idx := strings.IndexRune(a.Email, '@')
			if idx > 0 {
				name = a.Email[:idx]
			}
			ServiceAccounts[name] = a
		}

		if l.NextPageToken == "" {
			break
		}

		next_page_token = l.NextPageToken
	}

	g.ServiceAccounts = ServiceAccounts
	return nil
}

func (g *CloudState) FetchAddresses(ctx context.Context, service *compute.Service) (err error) {
	defer common.Action(&err, "querying addresses")
	g.Addresses = make(map[string]*compute.Address)
	aal, err := service.Addresses.AggregatedList(g.HostEnvironment.Project.Name).Context(ctx).MaxResults(kMaxAddressCount).Do()

	if IsNotFoundError(err) {
		return nil
	}

	if err != nil {
		return
	}

	for _, al := range aal.Items {
		for _, a := range al.Addresses {
			g.Addresses[a.Name] = a
		}
	}
	return nil
}

func (g *CloudState) FetchInstances(ctx context.Context, service *compute.Service) (err error) {
	defer common.Action(&err, "querying instances")
	g.Instances = make(map[string]*compute.Instance)
	il, err := service.Instances.AggregatedList(g.HostEnvironment.Project.Name).
		Context(ctx).MaxResults(kMaxInstanceListCount).Do()

	if IsNotFoundError(err) {
		return nil
	}

	if err != nil {
		return
	}

	for _, s := range il.Items {
		for _, i := range s.Instances {
			g.Instances[LastPathComponent(i.Name)] = i
		}
	}
	return nil
}

func (g *CloudState) FetchNetworks(ctx context.Context, service *compute.Service) (err error) {
	defer common.Action(&err, "querying networks and subnetworks")

	g.Networks = make(map[string]*compute.Network)
	g.Subnetworks = make(map[NetworkAndRegion]*compute.Subnetwork)

	nl, err := service.Networks.List(g.HostEnvironment.Project.Name).Context(ctx).Do()

	if IsNotFoundError(err) {
		return nil
	}

	if err != nil {
		return
	}

	for _, n := range nl.Items {
		g.Networks[LastPathComponent(n.Name)] = n
	}

	// And subnetworks:
	snl, err := service.Subnetworks.AggregatedList(g.HostEnvironment.Project.Name).Context(ctx).
		MaxResults(kMaxSubnetworkListCount).Do()

	if IsNotFoundError(err) {
		return nil
	}

	if err != nil {
		return
	}

	for _, p := range snl.Items {
		for _, s := range p.Subnetworks {
			k := NetworkAndRegion{
				Network: LastPathComponent(s.Network),
				Region:  LastPathComponent(s.Region)}

			g.Subnetworks[k] = s
		}
	}
	return nil
}

func (g *CloudState) FetchFirewalls(ctx context.Context, service *compute.Service) (err error) {
	defer common.Action(&err, "querying firewall rules")

	g.Firewalls = make(map[NetworkAndFirewall]*compute.Firewall)

	fl, err := service.Firewalls.List(g.HostEnvironment.Project.Name).Context(ctx).Do()

	if IsNotFoundError(err) {
		return nil
	}

	if err != nil {
		return
	}

	for _, f := range fl.Items {
		k := NetworkAndFirewall{
			Network:  LastPathComponent(f.Network),
			Firewall: LastPathComponent(f.Name)}
		g.Firewalls[k] = f
	}
	return nil
}

func (g *CloudState) FetchImages(ctx context.Context, service *compute.Service) (err error) {
	defer common.Action(&err, "querying source images")
	g.Images = make(map[string]*compute.Image)
	for _, i := range g.HostEnvironment.Image {
		l := i.GetLatest()
		if l == nil {
			continue
		}

		var f *compute.Image
		f, err = service.Images.GetFromFamily(l.Project, l.Family).Context(ctx).
			Fields("selfLink", "status").Do()
		if err != nil {
			return
		}

		if f.Status != "READY" {
			return errors.Errorf("image is not ready for use for project %s, family %s. Status is %s",
				l.Project, l.Family, f.Status)
		}

		i.Url = f.SelfLink

		g.Images[i.Name] = f
	}
	return nil
}

func (g *CloudState) FetchZones(ctx context.Context, service *compute.Service) (err error) {
	defer common.Action(&err, "querying zones")
	zl, err := service.Zones.List(g.HostEnvironment.Project.Name).Context(ctx).Do()
	if err != nil {
		return
	}

	g.Zones = make(map[string]*compute.Zone)
	for _, z := range zl.Items {
		g.Zones[LastPathComponent(z.Name)] = z
	}
	return nil
}

// FetchAll fetches metadata from GCE. Attempts
// to parallelize the work as much as possible.
func (g *CloudState) FetchAll(ctx context.Context, client *http.Client) (err error) {
	defer common.Action(&err, "querying cloud state")

	service, err := compute.New(client)
	if err != nil {
		return
	}

	j := common.NewTasks(nil)
	j.Go(func() error { return g.FetchServiceAccounts(ctx, client) })
	j.Go(func() error { return g.FetchAddresses(ctx, service) })
	j.Go(func() error { return g.FetchInstances(ctx, service) })
	j.Go(func() error { return g.FetchNetworks(ctx, service) })
	j.Go(func() error { return g.FetchFirewalls(ctx, service) })
	j.Go(func() error { return g.FetchImages(ctx, service) })
	j.Go(func() error { return g.FetchZones(ctx, service) })
	return j.Join()
}

// QueryCloudState queries GCE assets and returns a CloudState object containing the cached state.
//
// |project| is the name of the GCE project to cache, and |images| is a map of
// GCE images that should be queried. There are too many images for this
// function to query all possible images. Hence the need for a narrowed down
// list.
//
// The |images| map should be a mapping from an arbitrary string to a
// config.SourceImage object. The same key will be used as the mapping key to
// strong the corresopnding compute.Image record.
func QueryCloudState(ctx context.Context, client *http.Client, env *host.HostEnvironment) (*CloudState, error) {

	g := CloudState{HostEnvironment: env}

	err := g.FetchAll(ctx, client)
	if err != nil {
		return nil, err
	}
	return &g, nil
}
