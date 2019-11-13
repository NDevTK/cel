// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package asset

import (
	"net"
	"strings"

	"chromium.googlesource.com/enterprise/cel/go/schema"
	assetpb "chromium.googlesource.com/enterprise/cel/go/schema/asset"
	"github.com/pkg/errors"
)

// Fields which are properly annotated and have no additional validation
// requirements can have a trivial validator. The derived validators that are
// based on proto annotations will do all the work.

// See //go/schema/validator_registry.go and schema guidelines
// https://chromium.googlesource.com/enterprise/cel/+/master/docs/schema-guidelines.md
// for details on how to author new schema and how to add a validate method.

var validateFunctions = []interface{}{

	func(*assetpb.ActiveDirectoryDomainController) error { return nil },
	func(*assetpb.ActiveDirectoryGroupPolicy) error { return nil },
	func(*assetpb.ActiveDirectoryGroupPolicyLink) error { return nil },
	func(*assetpb.ActiveDirectoryOrganizationalUnit) error { return nil },
	func(*assetpb.ActiveDirectoryRegistryPolicy) error { return nil },
	func(*assetpb.ActiveDirectoryRegistryPrefPolicy) error { return nil },
	func(*assetpb.Certificate) error { return nil },
	func(*assetpb.GroupReference) error { return nil },
	func(*assetpb.IISApplication) error { return nil },
	func(*assetpb.IISBindings) error { return nil },
	func(*assetpb.IISServer) error { return nil },
	func(*assetpb.Machine) error { return nil },
	func(*assetpb.Network) error { return nil },
	func(*assetpb.NetworkInterface) error { return nil },
	func(*assetpb.UserReference) error { return nil },
	func(*assetpb.WindowsGroup) error { return nil },

	func(u *assetpb.WindowsUser) error {
		// Strings starting with a hyphen will be interpreted as flags by powershell
		if strings.HasPrefix(u.Description, "-") {
			return errors.Errorf("description '%s' cannot start with a '-'", u.Description)
		}
		return nil
	},

	// Validate that all assets are coherent with one another.
	func(a *assetpb.AssetManifest) error {
		apb := (*assetpb.AssetManifest)(a)
		for _, remoteDesktopHost := range a.RemoteDesktopHost {
			if err := validateRemoteDesktopHostWithAssetManifest(remoteDesktopHost, apb); err != nil {
				return err
			}
		}

		for _, iisSite := range a.IisSite {
			if err := validateIISSiteWithAssetManifest(iisSite, apb); err != nil {
				return err
			}
		}

		return nil
	},

	func(u *assetpb.UserOrGroupReference) error {
		if u.Entity == nil {
			return errors.New("either 'user' or 'group' is required.")
		}
		return nil
	},

	func(n *assetpb.NetworkPeer) error {
		if len(n.Network) < 2 {
			return errors.New("a NetworkPeer declaration must specify at least two networks")
		}
		return nil
	},

	func(a *assetpb.Address) error {
		ip := net.ParseIP(a.Ip)
		if ip == nil {
			return errors.Errorf("failed to parse address %s", a.Ip)
		}

		return nil
	},

	func(a *assetpb.AddressRange) error {
		_, _, err := net.ParseCIDR(a.Cidr)
		if err != nil {
			return errors.Wrapf(err, "failed to parse address range %s", a.Cidr)
		}

		return nil
	},

	func(f *assetpb.FixedAddress) error {
		fpb := (*assetpb.FixedAddress)(f)
		if fpb.GetAddress() == nil && fpb.GetAddressPool() == "" {
			return errors.New("either the 'address' or 'address_pool' field must be specified for a 'FixedAddress'")
		}
		return nil
	},

	func(d *assetpb.DNSZone) error {
		if len(d.Record) == 0 {
			return errors.New("at least one DNSRecord must be specified for a DNSZone")
		}
		return nil
	},

	func(d *assetpb.DNSRecord) error { return nil },

	func(a *assetpb.ActiveDirectoryDomain) error {
		if len(a.Name) > 15 && a.NetbiosName == "" {
			return errors.New("'netbios_name' is required if 'name' is longer than 15 characters")
		}

		// Strings starting with a hyphen will be interpreted as flags by powershell
		if strings.HasPrefix(a.NetbiosName, "-") {
			return errors.Errorf("netbios_name '%s' cannot start with a '-'", a.NetbiosName)
		}

		return nil
	},

	func(w *assetpb.WindowsContainer) error {
		wpb := (*assetpb.WindowsContainer)(w)
		if wpb.GetAdOrganizationalUnit() == "" && wpb.GetAdDomain() == "" && wpb.GetWindowsMachine() == "" {
			return errors.New("one of 'ad_organizational_unit', 'ad_domain', or 'windows_machine' must be specified for a WindowsContainer")
		}
		return nil
	},

	func(w *assetpb.WindowsMachine) error {
		switch {
		case len(w.NetworkInterface) == 0:
			return errors.Errorf("at least one 'network_interface' is required for WindowsMachine named %s", w.Name)
		case w.Container != nil && w.Container.GetWindowsMachine() != "":
			return errors.Errorf("'container' cannot specify 'windows_machine' for WindowsMachine named %s", w.Name)
		case len(w.Name) > 15:
			return errors.Errorf("WindowsMachine name %s is too long. Its length is %d, while the max length allowed is 15.",
				w.Name, len(w.Name))
		}

		return nil
	},

	func(c *assetpb.CertificatePool) error {
		if len(c.IncludeNamed) == 0 && len(c.IncludeFile) == 0 {
			return errors.Errorf("at least one 'include_named' or 'include_file' is required for CertificatePool %s", c.Name)
		}

		for _, n := range c.IncludeNamed {
			if n == "" {
				return errors.Errorf("'include_named' can't contain empty names in CertificatePool named %s", c.Name)
			}
		}

		return nil
	},

	func(r *assetpb.RegistryKey) error {
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
	},

	func(v *assetpb.RegistryValue) error {
		return nil
	},

	func(v *assetpb.RegistryValue_MultiString) error {
		for _, s := range v.Value {
			if len(s) == 0 {
				return errors.New("registry MultiString value cannot be empty")
			}

			if strings.ContainsRune(s, 0) {
				return errors.New("registry MultiString value cannot contain \\0")
			}
		}
		return nil
	},

	func(rd *assetpb.RemoteDesktopHost) error {
		// Strings starting with a hyphen will be interpreted as flags by powershell
		if strings.HasPrefix(rd.CollectionName, "-") {
			return errors.Errorf("collection_name '%s' cannot start with a '-'", rd.CollectionName)
		} else if strings.HasPrefix(rd.CollectionDescription, "-") {
			return errors.Errorf("collection_description '%s' cannot start with a '-'", rd.CollectionDescription)
		}

		if len(rd.CollectionName) > 15 {
			return errors.New("'collection_name' can't be longer than 15 characters")
		}

		return nil
	},

	func(s *assetpb.IISSite) error {
		if s.Bindings.Protocol == assetpb.Protocol_HTTPS {
			return errors.Errorf("HTTPS is not yet supported for IIS Sites")
		}
		return nil
	},
}

var validRegistryHives = [...]string{
	"HKEY_CLASSES_ROOT", "HKCR",
	"HKEY_CURRENT_USER", "HKCU",
	"HKEY_LOCAL_MACHINE", "HKLM",
	"HKEY_USERS", "HKU",
	"HKEY_CURRENT_CONFIG", "HKCC"}

const registryKeyPathSeparator = "\\"

func validateRemoteDesktopHostWithAssetManifest(rd *assetpb.RemoteDesktopHost, a *assetpb.AssetManifest) error {
	// RDS is not supported outside of a domain
	windows_machine, err := FindWindowsMachine(a, rd.WindowsMachine)
	if err != nil {
		return err
	}

	_, err = FindActiveDirectoryDomainFor(a, windows_machine)
	if err != nil {
		return errors.Errorf("RDS is not supported on machine '%s' (not in a domain)", rd.WindowsMachine)
	}

	return nil
}

func validateIISSiteWithAssetManifest(s *assetpb.IISSite, a *assetpb.AssetManifest) error {
	// Kerberos is not supported outside of a domain
	if s.AuthType == assetpb.IISAuthType_KERBEROS || s.AuthType == assetpb.IISAuthType_KERBEROS_NEGOTIABLE2 {
		iis_server, err := FindIISServer(a, s.IisServer)
		if err != nil {
			return err
		}

		windows_machine, err := FindWindowsMachine(a, iis_server.WindowsMachine)
		if err != nil {
			return err
		}

		_, err = FindActiveDirectoryDomainFor(a, windows_machine)
		if err != nil {
			return errors.Errorf("Kerberos is unsupported for site %s on machine '%s' (not in a domain)", s.Name, windows_machine.Name)
		}
	}

	return nil
}

func init() {
	schema.RegisterAllValidateFunctions(validateFunctions)
}
