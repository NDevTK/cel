// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

import (
	"chromium.googlesource.com/enterprise/cel/go/cel"
	"chromium.googlesource.com/enterprise/cel/go/host"
	"golang.org/x/net/context"
	compute "google.golang.org/api/compute/v1"
	iam "google.golang.org/api/iam/v1"
	"net/http"
	"strings"
	"sync/atomic"
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
	// Project described by this cached state.
	Project string `json:"project"`

	// The host environment defines which resources need monitoring. Should not
	// be modified.
	HostEnvironment *host.HostEnvironment `json:"host_environment"`

	// All static external IP addresses reserved by the project.
	Addresses map[string]*compute.Address `json:"addresses"`

	// All service accounts for the project by unique ID.
	ServiceAccounts map[string]*iam.ServiceAccount `json:"service_accounts"`

	// Instances by instance label
	Instances map[string]*compute.Instance `json:"instances"`

	// Network by network label
	Networks map[string]*compute.Network `json:"networks"`

	// Subnetwork by network label, and region
	Subnetworks map[string]map[string]*compute.Subnetwork `json:"subnetworks"`

	// Firewalls by network label and firewall label.
	Firewalls map[string]map[string]*compute.Firewall `json:"firewall"`

	// Images by image name. Only the Images were requested are going to be
	// synchronized. The name of the image used here is the name used as the
	// image name used in |HostEnvironment|.
	Images map[string]*compute.Image `json:"image"`

	// Zones by zone label.
	Zones map[string]*compute.Zone `json:"zone"`
}

func (g *CloudState) FetchServiceAccounts(ctx context.Context, client *http.Client) (err error) {
	defer cel.Action(&err, "querying service accounts")

	service, err := iam.New(client)
	if err != nil {
		return
	}

	ServiceAccounts := make(map[string]*iam.ServiceAccount)
	next_page_token := ""

	for {
		call := s.rvice.Projects.ServiceAccounts.List(ProjectResource(g.Project))
		call.PageSize(kMaxServiceAccountCount).Context(ctx)
		if next_page_token != "" {
			call.PageToken(next_page_token)
		}
		l, err := call.Do()

		if IsNotFoundError(err) {
			return nil
		}

		if err != nil {
			return
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
	defer Action(&err, "querying addresses")
	g.Addresses = make(map[string]*compute.Address)
	aal, err := service.Addresses.AggregatedList(g.Project).Context(ctx).MaxResults(kMaxAddressCount).Do()

	if IsNotFoundError(err) {
		g.addresses_uptodate = 1
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
	g.addresses_uptodate = 1
	return nil
}

func (g *CloudState) FetchInstances(ctx context.Context, service *compute.Service) (err error) {
	defer Action(&err, "querying instances")
	g.Instances = make(map[string]*compute.Instance)
	il, err := service.Instances.AggregatedList(g.Project).
		Context(ctx).MaxResults(kMaxInstanceListCount).Do()

	if IsNotFoundError(err) {
		g.instances_uptodate = 1
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
	g.instances_uptodate = 1
	return nil
}

func (g *CloudState) FetchNetworks(ctx context.Context, service *compute.Service) (err error) {
	defer Action(&err, "querying networks and subnetworks")

	g.Networks = make(map[string]*compute.Network)
	g.Subnetworks = make(map[string]map[string]*compute.Subnetwork)

	nl, err := service.Networks.List(g.Project).Context(ctx).Do()

	if IsNotFoundError(err) {
		g.networks_uptodate = 1
		return nil
	}

	if err != nil {
		return
	}

	for _, n := range nl.Items {
		g.Networks[LastPathComponent(n.Name)] = n
	}

	// And subnetworks:
	snl, err := service.Subnetworks.AggregatedList(g.Project).Context(ctx).
		MaxResults(kMaxSubnetworkListCount).Do()

	if IsNotFoundError(err) {
		g.networks_uptodate = 1
		return nil
	}

	if err != nil {
		return
	}

	for _, p := range snl.Items {
		for _, s := range p.Subnetworks {
			n := LastPathComponent(s.Network)
			r := LastPathComponent(s.Region)

			if _, ok := g.Subnetworks[n]; !ok {
				g.Subnetworks[n] = make(map[string]*compute.Subnetwork)
			}

			g.Subnetworks[n][r] = s
		}
	}
	g.networks_uptodate = 1
	return nil
}

func (g *CloudState) FetchFirewalls(ctx context.Context, service *compute.Service) (err error) {
	defer Action(&err, "querying firewall rules")

	g.Firewalls = make(map[string]map[string]*compute.Firewall)

	fl, err := service.Firewalls.List(g.Project).Context(ctx).Do()

	if IsNotFoundError(err) {
		g.firewalls_uptodate = 1
		return nil
	}

	if err != nil {
		return
	}

	for _, f := range fl.Items {
		n := LastPathComponent(f.Network)

		if _, ok := g.Firewalls[n]; !ok {
			g.Firewalls[n] = make(map[string]*compute.Firewall)
		}
		fn := LastPathComponent(f.Name)
		g.Firewalls[n][fn] = f
	}
	g.firewalls_uptodate = 1
	return nil
}

func (g *CloudState) FetchImages(ctx context.Context, service *compute.Service) (err error) {
	defer Action(&err, "querying source images")
	g.Images = make(map[string]*compute.Image)
	for k, i := range g.MonitoredImages {
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
			return NewError("image is not ready for use for project %s, family %s. Status is %s",
				l.Project, l.Family, f.Status)
		}

		g.Images[k] = f
	}
	g.images_uptodate = 1
	return nil
}

func (g *CloudState) FetchZones(ctx context.Context, service *compute.Service) (err error) {
	defer Action(&err, "querying zones")
	zl, err := service.Zones.List(g.Project).Context(ctx).Do()
	if err != nil {
		return
	}

	g.Zones = make(map[string]*compute.Zone)
	for _, z := range zl.Items {
		g.Zones[LastPathComponent(z.Name)] = z
	}
	g.zones_uptodate = 1
	return nil
}

// FetchStale fetches stale metadata from GCE. Staleness is identified by the
// *_changed flags that were set by called to Set*Changed() methods. Attempts
// to parallelize the work as much as possible.
func (g *CloudState) FetchStale(ctx context.Context, client *http.Client) (err error) {
	defer Action(&err, "querying cloud state")

	c := make(chan error)
	pending := 0

	if g.service_accounts_uptodate == 0 {
		pending += 1
		go func() {
			c <- g.FetchServiceAccounts(ctx, client)
		}()
	}

	service, err := compute.New(client)
	if err != nil {
		return
	}

	if g.addresses_uptodate == 0 {
		pending += 1
		go func() {
			c <- g.FetchAddresses(ctx, service)
		}()
	}

	if g.instances_uptodate == 0 {
		pending += 1
		go func() {
			c <- g.FetchInstances(ctx, service)
		}()
	}

	if g.networks_uptodate == 0 {
		pending += 1
		go func() {
			c <- g.FetchNetworks(ctx, service)
		}()
	}

	if g.firewalls_uptodate == 0 {
		pending += 1
		go func() {
			c <- g.FetchFirewalls(ctx, service)
		}()
	}

	if g.images_uptodate == 0 {
		pending += 1
		go func() {
			c <- g.FetchImages(ctx, service)
		}()
	}

	if g.zones_uptodate == 0 {
		pending += 1
		go func() {
			c <- g.FetchZones(ctx, service)
		}()
	}

	return WaitFor(c, pending)
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
func QueryCloudState(ctx context.Context, client *http.Client, project string, dependencies *host.Dependencies) (*CloudState, error) {

	g := CloudState{
		Project: project,

		MonitoredImages: images}

	err := g.FetchStale(ctx, client)
	if err != nil {
		return nil, err
	}
	return &g, nil
}
