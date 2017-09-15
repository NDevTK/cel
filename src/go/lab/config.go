// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package lab

import (
	"github.com/golang/protobuf/proto"
	"go/lab/config"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
)

// Config is the entrypoint for everything that has to do with the lab.
//
// Construct one using LoadConfigFile(). Currently you need to call
// SyncGceState() before it's useful.
type Config struct {
	config.HostEnvironment
	config.Assets

	// The following fields are populated during Validate() for easy lookup by
	// resource name.
	domains          map[string]*config.WindowsDomain
	images           map[string]*config.SourceImage
	networks         map[string]*config.Network
	service_accounts map[string]*config.ServiceAccount
	instance_types   map[string]*config.InstanceType
	static_ips       map[string]*config.ExternalIP

	// effective_instances contains a map from instance name to *canonicalized* Instance
	// (see CanonicalizeInstance() below).  This is different from other
	// convenience maps where the mapping is to the raw object found in the
	// configuration file.
	effective_instances map[string]*config.Instance

	// Last known Google Cloud state. Updated each time SyncGceState() is
	// called.
	Cloud *CloudState `json:"cloud_state,omitempty"`
}

// loadTextProto reads the contents of the file at |path| as a text proto and
// unmarshals it into |o|.
func loadTextProto(path string, o proto.Message) error {
	var d []byte
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return proto.UnmarshalText(string(d), o)
}

// LoadConfigFile loads the configuration files specified by |hosting_config|
// and |asset_configs| and performs a set of validations on it.
//
// All configuration files are expected to bet in textpb format. The file at
// |hosting_config| must define a config.HostEnvironment object, while the
// files at |asset_configs| should describe config.Assets objects.
func LoadConfigFiles(hosting_config string, asset_configs []string) (c *Config, err error) {
	defer Action(&err, "loading hosting configuration at %s and assets at %v", hosting_config, asset_configs)

	c = &Config{}

	err = loadTextProto(hosting_config, &c.HostEnvironment)
	if err != nil {
		return
	}

	for _, p := range asset_configs {
		var assets_from_p config.Assets
		err = loadTextProto(p, &assets_from_p)
		if err != nil {
			return
		}

		proto.Merge(&c.Assets, &assets_from_p)
	}

	err = c.Validate()
	if err != nil {
		return
	}
	return
}

// SyncWithCloud queries the current GCE project and caches the current set of
// relevant GCE resource metadata. This includes:
//   * The list of VM instances.
//   * The list of networks and their subnetworks.
//   * The VM images that are referenced from the current configuration.
//
// Assuming that this metadata doesn't change during the course of deploying a
// lab configuration, these cached values can help speed things up.
func (c *Config) SyncWithCloud(ctx context.Context, client *http.Client) error {
	if c.Cloud != nil {
		return c.Cloud.FetchStale(ctx, client)
	}

	s, err := QueryCloudState(ctx, client, c.Project, c.images)
	if err != nil {
		return err
	}
	c.Cloud = s
	return nil
}

func (c *Config) EffectiveInstance(i *config.Instance) (ci *config.Instance) {
	var do *config.InstanceCreateOptions

	// If c.instance_types[] doesn't have a record for this instance type, it's
	// considered an error only if i.Type is not empty. This is detected during
	// the validation pass and is ignored here.
	if _, ok := c.instance_types[i.Type]; ok {
		do = c.instance_types[i.Type].CreateOptions
	} else {
		do = &config.InstanceCreateOptions{}
	}
	co := i.CreateOptions

	service_account := co.ServiceAccount
	if service_account == nil {
		service_account = do.ServiceAccount
	}

	crypto_key := co.Cryptokey
	if crypto_key == nil {
		crypto_key = do.Cryptokey
	}

	network := co.Interface
	if network == nil {
		network = do.Interface
	}

	return &config.Instance{
		Name:        i.Name,
		Description: i.Description,
		Role:        i.Role,
		Type:        i.Type,
		CreateOptions: &config.InstanceCreateOptions{
			Zone:              oneOfString(co.Zone, do.Zone),
			Image:             oneOfString(co.Image, do.Image),
			MachineType:       oneOfString(co.MachineType, do.MachineType),
			Metadata:          co.Metadata,
			Interface:         network,
			AutomaticRestart:  co.AutomaticRestart,
			OnHostMaintenance: co.OnHostMaintenance,
			Preemptible:       co.Preemptible || do.Preemptible,
			ServiceAccount:    service_account,
			Cryptokey:         crypto_key,
			Tag:               co.Tag}}
}
