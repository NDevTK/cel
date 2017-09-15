// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package lab

import (
	"cloud.google.com/go/logging"
	"fmt"
	"go/lab/config"
	compute "google.golang.org/api/compute/v1"
	"strings"
)

const kInstanceNamespace = "instances"
const kGceInstanceNamespace = "gceInstances"
const kGceInstanceRunNamespace = "runningGceInstances"

// instanceLogEntry is-a LogEntrySource for GCE VM instance operations. It
// encodes an instance name in addition to a message and an error field.
type instanceLogEntry struct {
	Message  string `json:"m"`
	Instance string `json:"instance,omitempty"`
	Error    string `json:"err,omitemtpy"`
}

func (i instanceLogEntry) Entry(s logging.Severity) logging.Entry {
	return logging.Entry{Severity: s, Payload: i}
}

type gceInstance struct {
	BaseNamedAsset
	Project         *Project
	Instance        *config.Instance
	Networks        map[string]*Network
	Addresses       map[string]*Address
	ServiceAccount  *ServiceAccount
	ComputeInstance *compute.Instance
}

// ComputeInstanceTemplate constructs a GCE compute.Instance using a Config and
// a canonicalized config.Instance.
func (i *gceInstance) ComputeInstanceTemplate(c *Config) (*compute.Instance, error) {
	co := i.Instance.CreateOptions
	auto_restart := co.AutomaticRestart

	scopes := append([]string{}, co.ServiceAccount.Scope...)
	for i, v := range scopes {
		if strings.Contains(v, "/") {
			continue
		}
		scopes[i] = fmt.Sprintf("https://www.googleapis.com/auth/%s", v)
	}
	service_account := &compute.ServiceAccount{
		Email:  i.ServiceAccount.IamServiceAccount.Email,
		Scopes: scopes}

	var interfaces []*compute.NetworkInterface
	for _, n := range co.Interface {
		network := i.Networks[n.Network]
		if network == nil {
			return nil, NewError("network %s was not found in the dependency tree")
		}
		if network.ComputeNetwork == nil {
			panic("dependency not resolved")
		}

		var access_configs []*compute.AccessConfig
		if n.ExternallyVisible {
			var ext_ip string
			if n.ExternalIpName != "" {
				ext_ip = i.Addresses[n.ExternalIpName].ComputeAddress.Address
			}
			access_configs = []*compute.AccessConfig{
				&compute.AccessConfig{
					Type:  "ONE_TO_ONE_NAT",
					NatIP: ext_ip}}
		}
		interfaces = append(interfaces, &compute.NetworkInterface{
			Network:       network.ComputeNetwork.SelfLink,
			NetworkIP:     n.InternalIp,
			AccessConfigs: access_configs})
	}

	return &compute.Instance{
		Name:        i.Instance.Name,
		Zone:        co.Zone,
		Description: i.Instance.Description,

		Disks: []*compute.AttachedDisk{{
			AutoDelete: true,
			Boot:       true,
			Interface:  "SCSI",
			Mode:       "READ_WRITE",
			Type:       "PERSISTENT",
			InitializeParams: &compute.AttachedDiskInitializeParams{
				SourceImage: c.Cloud.Images[co.Image].SelfLink}}},

		MachineType: MachineTypeResource(c.Project, co.Zone, co.MachineType),

		Metadata: metadataFromKVMap(co.Metadata),

		NetworkInterfaces: interfaces,

		Scheduling: &compute.Scheduling{
			AutomaticRestart:  &auto_restart,
			OnHostMaintenance: co.OnHostMaintenance.String(),
			Preemptible:       co.Preemptible},

		Tags: &compute.Tags{Items: co.Tag},

		ServiceAccounts: []*compute.ServiceAccount{service_account}}, nil
}

func (i *gceInstance) Resolve(s *Session) (err error) {
	defer Action(&err, "ensuring instance %s exists", i.id)
	if i.ComputeInstance = s.Config.Cloud.Instances[i.id]; i.ComputeInstance != nil {
		// Instance already exists. We could verify that at least some of the
		// properties match here.
		return
	}

	defer Action(&err, "creating new instance %s", i.id)
	i.ComputeInstance, err = i.ComputeInstanceTemplate(s.Config)
	if err != nil {
		return
	}

	s.LogInfo(instanceLogEntry{
		Message:  "creating instance",
		Instance: i.id})

	op, err := s.GetComputeService().Instances.Insert(
		i.Project.id, i.ComputeInstance.Zone, i.ComputeInstance).Context(s.Context).Do()
	if err != nil {
		s.LogError(instanceLogEntry{
			Message:  "couldn't insert new instance",
			Error:    err.Error(),
			Instance: i.id})
		return
	}

	err = WaitForOperation(s, op)
	if err != nil {
		return
	}

	i.ComputeInstance, err = s.GetComputeService().Instances.Get(
		i.Project.id, i.ComputeInstance.Zone, i.id).Context(s.Context).Do()
	return
}

func (i *gceInstance) Check(s *Session) (error, bool) {
	panic("not implemented")
}

func (i *gceInstance) Purge(s *Session) error {
	panic("not implemented")
}

type gceInstanceRun struct {
	BaseNamedAsset
	Instance *gceInstance
}

func (r *gceInstanceRun) Resolve(s *Session) (err error) {
	if r.Instance.ComputeInstance.Status == "RUNNING" {
		// We are done.
		return
	}

	defer Action(&err, "starting instance %s", r.Instance.id)
	op, err := s.GetComputeService().Instances.Start(
		r.Instance.Project.id, r.Instance.Instance.CreateOptions.Zone, r.Instance.id).
		Context(s.Context).Do()
	if err != nil {
		return
	}

	return WaitForOperation(s, op)
}

func (r *gceInstanceRun) Check(s *Session) (error, bool) {
	panic("not implemented")
}

func (r *gceInstanceRun) Purge(s *Session) error {
	panic("not implemented")
}

type Instance struct {
	BaseNamedAsset
	Instance        *gceInstance
	ComputeInstance *compute.Instance
}

func (i *Instance) Resolve(s *Session) error {
	i.ComputeInstance = i.Instance.ComputeInstance
	return nil
}

func (i *Instance) Check(s *Session) (error, bool) {
	panic("not implemented")
}

func (i *Instance) Purge(s *Session) error {
	panic("not implemented")
}

func LookupInstance(A *Assets, id string) *Instance {
	if a := A.Get(kInstanceNamespace, id); a != nil {
		return a.(*Instance)
	}
	panic("instance not found")
}

func ConstructInstanceAssets(A *Assets, c *Config) (err error) {
	p := LookupProject(A, c.Project)

	// c.instances contains canonicalized instances as opposed to c.Instance
	// which is a slice of raw Instance objects are specified in the
	// configuration file.
	for _, ci := range c.effective_instances {
		sa := LookupServiceAccount(A, ci.CreateOptions.ServiceAccount.Id)

		deps := []Asset{p, sa}
		networks := make(map[string]*Network)
		addresses := make(map[string]*Address)

		for _, ni := range ci.CreateOptions.Interface {
			networks[ni.Network] = LookupNetwork(A, ni.Network)
			if networks[ni.Network] == nil {
				return nil
			}

			deps = append(deps, networks[ni.Network])
			if !ni.ExternallyVisible || ni.ExternalIpName == "" {
				continue
			}

			named_ip := c.static_ips[ni.ExternalIpName]

			addresses[ni.ExternalIpName] = NewAddress(A, p, ci.CreateOptions.Zone, named_ip)
			deps = append(deps, addresses[ni.ExternalIpName])
		}

		i := gceInstance{BaseNamedAsset{kGceInstanceNamespace, ci.Name, deps}, p, ci, networks, addresses, sa, nil}
		A.Add(&i)

		ra := gceInstanceRun{BaseNamedAsset{kGceInstanceRunNamespace, i.id, []Asset{&i}}, &i}
		A.Add(&ra)

		I := Instance{BaseNamedAsset{kInstanceNamespace, ci.Name, []Asset{&i, &ra}}, &i, nil}
		A.Add(&I)
	}
	return nil
}
