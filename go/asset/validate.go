// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package asset

import (
	"net"
	"strings"

	"github.com/pkg/errors"
)

// Fields which are properly annotated and have no additional validation
// requirements can have a trivial validator. The derived validators that are
// based on proto annotations will do all the work.

// See the Validator interface and schema guidelines
// https://chromium.googlesource.com/enterprise/cel/+/master/docs/schema-guidelines.md
// for details on how to author new schema and how to add a Validate() method.

func (*ActiveDirectoryDomainController) Validate() error   { return nil }
func (*ActiveDirectoryGroupPolicy) Validate() error        { return nil }
func (*ActiveDirectoryGroupPolicyLink) Validate() error    { return nil }
func (*ActiveDirectoryOrganizationalUnit) Validate() error { return nil }
func (*ActiveDirectoryRegistryPolicy) Validate() error     { return nil }
func (*ActiveDirectoryRegistryPrefPolicy) Validate() error { return nil }
func (*AssetManifest) Validate() error                     { return nil }
func (*Certificate) Validate() error                       { return nil }
func (*GroupReference) Validate() error                    { return nil }
func (*IISApplication) Validate() error                    { return nil }
func (*IISBindings) Validate() error                       { return nil }
func (*IISServer) Validate() error                         { return nil }
func (*IISSite) Validate() error                           { return nil }
func (*Machine) Validate() error                           { return nil }
func (*Network) Validate() error                           { return nil }
func (*NetworkInterface) Validate() error                  { return nil }
func (*UserReference) Validate() error                     { return nil }
func (*WindowsGroup) Validate() error                      { return nil }

func (u *WindowsUser) Validate() error {
	// Strings starting with a hyphen will be interpreted as flags by powershell
	if strings.HasPrefix(u.Description, "-") {
		return errors.Errorf("description '%s' cannot start with a '-'", u.Description)
	}
	return nil
}

func (u *UserOrGroupReference) Validate() error {
	if u.Entity == nil {
		return errors.New("either 'user' or 'group' is required.")
	}
	return nil
}

func (n *NetworkPeer) Validate() error {
	if len(n.Network) < 2 {
		return errors.New("a NetworkPeer declaration must specify at least two networks")
	}
	return nil
}

func (a *Address) Validate() error {
	ip := net.ParseIP(a.Ip)
	if ip == nil {
		return errors.Errorf("failed to parse address %s", a.Ip)
	}

	return nil
}

func (a *AddressRange) Validate() error {
	_, _, err := net.ParseCIDR(a.Cidr)
	if err != nil {
		return errors.Wrapf(err, "failed to parse address range %s", a.Cidr)
	}

	return nil
}

func (f *FixedAddress) Validate() error {
	if f.GetAddress() == nil && f.GetAddressPool() == "" {
		return errors.New("either the 'address' or 'address_pool' field must be specified for a 'FixedAddress'")
	}
	return nil
}

func (d *DNSZone) Validate() error {
	if len(d.Record) == 0 {
		return errors.New("at least one DNSRecord must be specified for a DNSZone")
	}
	return nil
}

func (d *DNSRecord) Validate() error { return nil }

func (a *ActiveDirectoryDomain) Validate() error {
	if len(a.Name) > 15 && a.NetbiosName == "" {
		return errors.New("'netbios_name' is required if 'name' is longer than 15 characters")
	}

	// Strings starting with a hyphen will be interpreted as flags by powershell
	if strings.HasPrefix(a.NetbiosName, "-") {
		return errors.Errorf("netbios_name '%s' cannot start with a '-'", a.NetbiosName)
	}

	return nil
}

func (w *WindowsContainer) Validate() error {
	if w.GetAdOrganizationalUnit() == "" && w.GetAdDomain() == "" && w.GetWindowsMachine() == "" {
		return errors.New("one of 'ad_organizational_unit', 'ad_domain', or 'windows_machine' must be specified for a WindowsContainer")
	}
	return nil
}

func (w *WindowsMachine) Validate() error {
	switch {
	case len(w.NetworkInterface) == 0:
		return errors.Errorf("at least one 'network_interface' is required for WindowsMachine named %s", w.Name)
	case w.Container != nil && w.Container.GetWindowsMachine() != "":
		return errors.Errorf("'container' cannot specify 'windows_machine' for WindowsMachine named %s", w.Name)
	}

	return nil
}

func (c *CertificatePool) Validate() error {
	if len(c.IncludeNamed) == 0 && len(c.IncludeFile) == 0 {
		return errors.Errorf("at least one 'include_named' or 'include_file' is required for CertificatePool %s", c.Name)
	}

	for _, n := range c.IncludeNamed {
		if n == "" {
			return errors.Errorf("'include_named' can't contain empty names in CertificatePool named %s", c.Name)
		}
	}

	return nil
}

var validRegistryHives = [...]string{
	"HKEY_CLASSES_ROOT", "HKCR",
	"HKEY_CURRENT_USER", "HKCU",
	"HKEY_LOCAL_MACHINE", "HKLM",
	"HKEY_USERS", "HKU",
	"HKEY_CURRENT_CONFIG", "HKCC"}

const registryKeyPathSeparator = "/"

func (r *RegistryKey) Validate() error {
	if !strings.Contains(r.Path, registryKeyPathSeparator) {
		return errors.Errorf("registry key path is invalid: \"%s\"", r.Path)
	}
	hive := r.Path[:strings.Index(r.Path, registryKeyPathSeparator)]
	found := false
	for _, h := range validRegistryHives {
		if h == hive {
			found = true
			break
		}
	}
	if !found {
		return errors.Errorf("registry hive is \"%s\" which doesn't match allowed list: %v", hive, validRegistryHives)
	}

	if strings.ContainsRune(r.Path, '/') {
		return errors.Errorf("found forward slashes in registry path \"%s\". These are not allowed due to the possibility of inconsistent behavior", r.Path)
	}

	return nil
}

func (v *RegistryValue) Validate() error {
	return nil
}

func (v *RegistryValue_MultiString) Validate() error {
	for _, s := range v.Value {
		if len(s) == 0 {
			return errors.New("registry MultiString value cannot be empty")
		}

		if strings.ContainsRune(s, 0) {
			return errors.New("registry MultiString value cannot contain \\0")
		}
	}
	return nil
}

func (rd *RemoteDesktopHost) Validate() error {
	// Strings starting with a hyphen will be interpreted as flags by powershell
	if strings.HasPrefix(rd.CollectionName, "-") {
		return errors.Errorf("collection_name '%s' cannot start with a '-'", rd.CollectionName)
	} else if strings.HasPrefix(rd.CollectionDescription, "-") {
		return errors.Errorf("collection_description '%s' cannot start with a '-'", rd.CollectionDescription)
	}

	return nil
}
