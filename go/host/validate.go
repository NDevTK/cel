// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package host

import (
	"strings"

	"github.com/pkg/errors"
)

func (*AddressPool) Validate() error     { return nil }
func (*HostEnvironment) Validate() error { return nil }
func (*Image_Family) Validate() error    { return nil }
func (*LogSettings) Validate() error     { return nil }
func (*Project) Validate() error         { return nil }
func (*RuntimeSupport) Validate() error  { return nil }
func (*Startup) Validate() error         { return nil }

func (i *Image) Validate() error {
	if i.GetFixed() == "" && i.GetLatest() == nil {
		return errors.New("either 'url' or 'latest' must be specified for an Image")
	}
	return nil
}

func (s *Storage) Validate() error {
	if s.Prefix == "" {
		return nil
	}

	if strings.HasSuffix(s.Prefix, "/") {
		return errors.Errorf("the GCS object name prefix must not end with a slash. Found name \"%s\"", s.Prefix)
	}
	return nil
}

func (t *MachineType) Validate() error {
	if t.GetBase() == nil {
		return errors.New("either the 'instance_properties' or 'instance_template' must be specified for a 'machine_type'")
	}

	switch b := t.GetBase().(type) {
	case *MachineType_InstanceTemplate:
		break

	case *MachineType_InstanceProperties:
		if b.InstanceProperties == nil {
			return errors.New("'instance_properties' cannot be empty")
		}

		if b.InstanceProperties.NetworkInterfaces != nil {
			return errors.Errorf("'instance_properties' cannot specify 'network_interfaces' in machine type %s", t.Name)
		}

		if len(b.InstanceProperties.Disks) == 0 {
			return errors.Errorf("at least one disk must be specified for machine type %s", t.Name)
		}

		for _, d := range b.InstanceProperties.Disks {
			if d.InitializeParams != nil && !d.AutoDelete {
				return errors.Errorf("when specifying 'initialize_params', you must also set 'auto_delete' to true in 'instance_properties' for machine type %s", t.Name)
			}
		}
	}
	return nil
}

func (n *NestedVM) Validate() error {
	if n.DiskSizeGb != 0 {
		if n.DiskSizeGb < 10 {
			return errors.New("'diskSizeGb' cannot be smaller than 10")
		}

		if n.DiskSizeGb > 500 {
			return errors.New("'diskSizeGb' cannot be bigger than 500")
		}
	}

	return nil
}
