// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package lab

import (
	"fmt"
	"go/lab/config"
	compute "google.golang.org/api/compute/v1"
	"net"
	"regexp"
	"strings"
)

// -----------------------------------------------------------------------------
// Validations
// -----------------------------------------------------------------------------

// Validate checks whether all required fields are present in the Config and
// that internal references are sound. Returns nil if there were no errors. If
// any inconsistencies are found, returns an error which describes what was
// missing.
//
// Limitation: this function only returns one error even if it could report
// more. Perhaps it would be more productive asset authors if it returned more.
//
// During the validation pass, the validators also populate the lookup maps in
// Config for named resources.
func (c *Config) Validate() error {
	// The order of the validator invocations is important since later
	// validators rely on the named resource maps being populated.

	if err := c.ValidateHostEnvironment(); err != nil {
		return err
	}

	if err := c.ValidateServiceAccounts(); err != nil {
		return err
	}

	if err := c.ValidateNetworks(); err != nil {
		return err
	}

	if err := c.ValidateSourceImages(); err != nil {
		return err
	}

	if err := c.ValidateExternalIps(); err != nil {
		return err
	}

	if err := c.ValidateInstanceTypes(); err != nil {
		return err
	}

	if err := c.ValidateInstances(); err != nil {
		return err
	}

	if err := c.ValidateWindowsDomains(); err != nil {
		return err
	}

	return nil
}

func (c *Config) ValidateHostEnvironment() error {
	bucket_m := matcher("gs://*/.*$")

	switch {
	case c.Project == "":
		return NewError("project ID is required")

	case !isIdentifier(c.GetProject()):
		return NewError("invalid project ID. Found: %s", c.GetProject())

	case c.DeploymentBucket == "":
		return NewError("deployment_bucket is required")

	case !bucket_m.MatchString(c.GetDeploymentBucket()):
		return NewError("invalid deployment bucket. Found %s", c.GetDeploymentBucket())

	case c.DomainKeyBucket == "":
		return NewError("domain_key_bucket is required")

	case !bucket_m.MatchString(c.GetDomainKeyBucket()):
		return NewError("invalid domain key bucket. Found %s", c.GetDomainKeyBucket())
	}
	return nil
}

func (c *Config) ValidateServiceAccount(s *config.ServiceAccount) error {
	if s.Id == "" {
		return NewError("missing ID for service account")
	}

	if !isIdentifier(s.Id) {
		return NewError("invalid ID for service account %s", s.Id)
	}

	return nil
}

func (c *Config) ValidateServiceAccounts() error {
	c.service_accounts = make(map[string]*config.ServiceAccount)
	for _, s := range c.ServiceAccount {
		err := c.ValidateServiceAccount(s)
		if err != nil {
			return err
		}

		if _, ok := c.service_accounts[s.Id]; ok {
			return NewError("duplicate service account Id : %s", s.Id)
		}

		c.service_accounts[s.Id] = s
	}

	return nil
}

func (c *Config) ValidateSubnetwork(s *config.Network_Subnetwork, n *config.Network, parent *net.IPNet) error {
	switch {
	case s.Name == "":
		return NewError("subnetwork name is required for network %s", n.Name)

	case !isIdentifier(s.Name):
		return NewError("invalid subnetwork name %s in network %s", s.Name, n.Name)

	case s.Region == "":
		return NewError("region is required for subnetwork %s in network %s", s.Name, n.Name)

	case !isIdentifier(s.Region):
		return NewError("invalid region name %s in subnetwork %s in network %s", s.Region, s.Name, n.Name)

	case s.Ipv4AddressRange == "":
		return NewError("ipv4_address_range is required for subnetwork %s in network %s", s.Name, n.Name)
	}

	ip, ipnet, err := net.ParseCIDR(s.Ipv4AddressRange)
	if err != nil {
		return NewError("invalid IPv4 address range %s for subnetwork %s in network %s", s.Ipv4AddressRange, s.Name, n.Name)
	}

	ones, bits := ipnet.Mask.Size()
	parent_ones, parent_bits := parent.Mask.Size()
	if !parent.Contains(ip) || ones <= parent_ones || bits != parent_bits {
		return NewError("IPv4 address range must be a proper subset of parent for subnetwork %s in network %s", s.Name, n.Name)
	}
	return nil
}

func (c *Config) ValidateNetwork(n *config.Network) error {
	if !isIdentifier(n.Name) {
		return NewError("invalid network name: %s", n.Name)
	}

	if n.Ipv4AddressRange == "" {
		return NewError("IPv4 address range is required for network %s", n.Name)
	}

	ip, ipnet, err := net.ParseCIDR(n.Ipv4AddressRange)
	if err != nil {
		return NewError("failed to parse address range %s for network %s : %v", n.Ipv4AddressRange, n.Name, err)
	}

	ones, bits := ipnet.Mask.Size()
	if bits != 32 {
		return NewError("incorrect address range type for network %s. Should be an IPv4 range.", n.Name)
	}

	if ones <= 9 {
		return NewError("incorrect address range for network %s. Should be a subset of 10.0.0.0/9", n.Name)
	}

	if !(&net.IPNet{net.IPv4(10, 0, 0, 0), net.IPv4Mask(255, 128, 0, 0)}).Contains(ip) {
		return NewError("incorrect address range for network %s. Should be a subset of 10.0.0.0/9", n.Name)
	}

	for _, f := range n.FirewallRule {
		if !isIdentifier(f.Name) {
			return NewError("invalid firewall name:%s in %s", f.Name, n.Name)
		}

		if len(f.Allowed) == 0 {
			return NewError("no traffic allowed from firewall rule %s in %s", f.Name, n.Name)
		}

		if len(f.SourceTag) == 0 {
			return NewError("no source tags defined for firewall rule %s in %s", f.Name, n.Name)
		}

		if len(f.TargetTag) == 0 {
			return NewError("no target tags defined for firewall rule %s in %s", f.Name, n.Name)
		}
	}

	for _, s := range n.Subnetwork {
		if err := c.ValidateSubnetwork(s, n, ipnet); err != nil {
			return err
		}
	}

	if len(n.FirewallRule) == 0 {
		return NewError("no firewall rules specified for network %s", n.Name)
	}

	return nil
}

func (c *Config) ValidateNetworks() error {
	c.networks = make(map[string]*config.Network)

	for _, v := range c.GetNetwork() {
		err := c.ValidateNetwork(v)
		if err != nil {
			return err
		}

		if _, ok := c.networks[v.Name]; ok {
			return NewError("duplicate network name: %s", v.Name)
		}

		c.networks[v.Name] = v
	}

	if len(c.networks) == 0 {
		return NewError("no network configurations")
	}

	for _, v := range c.networks {
		if len(v.Peer) == 0 {
			continue
		}

		for _, peer := range v.Peer {
			if _, ok := c.networks[peer]; !ok {
				return NewError("network %s lists unknown network %s as a peer", v.Name, peer)
			}
		}
	}

	return nil
}

func (c *Config) ValidateSourceImage(i *config.SourceImage) error {
	switch {
	case !isIdentifier(i.Name):
		return NewError("invalid source image name %s", i.Name)

	case i.GetLatest() != nil:
		switch {
		case !isIdentifier(i.GetLatest().Family):
			return NewError("invalid family name for source image %s", i.Name)

		case !isIdentifier(i.GetLatest().Project):
			return NewError("invalid project name for source image %s", i.Name)
		}

	case i.GetFixed() != nil:
		if !matcher("projects/*/global/images/*").MatchString(i.GetFixed().ImageUrl) {
			return NewError("invalid url for source image %s: ", i.Name, i.GetFixed().ImageUrl)
		}

	case i.GetFixed() == nil && i.GetLatest() == nil:
		return NewError("no source details for source image %s. Either 'fixed' or 'latest' must be specified",
			i.Name)
	}

	return nil
}

func (c *Config) ValidateSourceImages() error {
	c.images = make(map[string]*config.SourceImage)
	for _, i := range c.SourceImage {
		err := c.ValidateSourceImage(i)
		if err != nil {
			return err
		}

		if _, ok := c.images[i.Name]; ok {
			return NewError("duplicate source image name :%s", i.Name)
		}
		c.images[i.Name] = i
	}

	if len(c.images) == 0 {
		return NewError("no source images found")
	}
	return nil
}

func (c *Config) ValidateExternalIP(i *config.ExternalIP) error {
	switch {
	case !isIdentifier(i.Name):
		return NewError("invalid external IP address name %s", i.Name)

	case i.Address == "":
		return NewError("external IP address specification %s missing address", i.Name)
	}

	ip := net.ParseIP(i.Address)
	if ip == nil {
		return NewError("invalid IP address %s for %s", i.Address, i.Name)
	}

	return nil
}

func (c *Config) ValidateExternalIps() error {
	c.static_ips = make(map[string]*config.ExternalIP)
	for _, i := range c.StaticIp {
		err := c.ValidateExternalIP(i)
		if err != nil {
			return err
		}

		if _, ok := c.static_ips[i.Name]; ok {
			return NewError("duplicate external ip name : %s", i.Name)
		}
		c.static_ips[i.Name] = i
	}
	return nil
}

func (c *Config) ValidateCryptoKey(name string, k *config.Key) error {
	switch {
	case k.Keyring == "":
		return NewError("missing keyring name for key in %s", name)

	case k.Cryptokey == "":
		return NewError("missing cryptokey name for key in %s", name)

	case !isIdentifier(k.Keyring):
		return NewError("invalid keyring name for key in %s", name)

	case !isIdentifier(k.Cryptokey):
		return NewError("invalid cryptokey name for key in %s", name)
	}

	return nil
}

func (c *Config) ValidateInstanceServiceAccount(name string,
	sa *config.InstanceServiceAccount) error {

	switch {
	case sa.Id == "":
		return NewError("no id specified for service account in %s", name)

	case len(sa.Scope) == 0:
		return NewError("no scopes specified for service account in %s", name)
	}

	if _, ok := c.service_accounts[sa.Id]; !ok {
		return NewError("invalid service account name for %s: %s", name, sa.Id)
	}

	return nil
}

func (c *Config) ValidateInstanceNetwork(name string, i *config.NetworkInterface) error {
	switch {
	case i.Network == "":
		return NewError("missing network name for instance %s", name)

	case !i.ExternallyVisible && i.ExternalIpName != "":
		return NewError("external static IP address requested for instance %s with no external visibility", name)

	case i.ExternalIpName != "":
		if _, found := c.static_ips[i.ExternalIpName]; !found {
			return NewError("external static IP address %s for instance %s is not defined in host environment",
				i.ExternalIpName, name)
		}

	case i.InternalIp != "":
		if net.ParseIP(i.InternalIp) == nil {
			return NewError("internal static IP %s for instance %s is invalid", i.InternalIp, name)
		}
	}

	if _, found := c.networks[i.Network]; !found {
		return NewError("network name %s specified for instance %s is invalid", i.Network, name)
	}
	return nil
}

func (c *Config) ValidateInstanceCreateOptions(name string, co *config.InstanceCreateOptions, is_instance_type bool) error {

	if is_instance_type {
		// The following fields are ignored for an instance type's create
		// options. If specified, it's likely a configuration error.
		switch {
		case co.Metadata != nil:
			return NewError("metadata specified for an instance type, but will be ignored")

		case co.Tag != nil:
			return NewError("tag specified for an instance type, but will be ignored")

		case co.OnHostMaintenance == config.InstanceCreateOptions_TERMINATE:
			return NewError("on_host_maintenance specified for an instance type, but will be ignored")
		}

		for _, i := range co.Interface {
			switch {
			case i.ExternallyVisible:
				return NewError("interface.externally_visible has no effect when specified for an instance type")

			case i.InternalIp != "":
				return NewError("interface.internal_ip specified for an instance type, but will be ignored")

			case i.ExternalIpName != "":
				return NewError("interface.external_ip_name specified for an instance type, but will be ignored")
			}
		}
	} else {
		// Make sure all the required fields are there. Note that we are
		// verifying a canonicalized config.InstanceCreateOptions here. So we
		// don't need to check against the instance type's create options. All
		// required fields should've been populated as a result of the
		// canonicalization logic.
		switch {
		case co.Zone == "":
			return NewError("missing zone for %s", name)

		case co.Image == "":
			return NewError("missing disk image for %s", name)

		case co.MachineType == "":
			return NewError("missing machine type for %s", name)

		case co.Interface == nil:
			return NewError("missing network interface for %s", name)

		case co.ServiceAccount == nil:
			return NewError("no service accounts configured for %s", name)

		case co.Cryptokey == nil:
			return NewError("no crypto_key for %s", name)
		}
	}

	_, valid_image := c.images[co.Image]

	// The following fields, if present, must be valid.
	switch {
	case co.Zone != "" && !isIdentifier(co.Zone):
		return NewError("invalid zone for %s: %s", name, co.Zone)

	case co.Image != "" && !valid_image:
		return NewError("invalid disk image for %s: %s", name, co.Image)

	case co.MachineType != "" && !isIdentifier(co.MachineType):
		return NewError("invalid machine type for %s: %s", name, co.MachineType)
	}

	if co.Cryptokey != nil {
		if err := c.ValidateCryptoKey(name, co.Cryptokey); err != nil {
			return err
		}
	}

	if co.ServiceAccount != nil {
		if err := c.ValidateInstanceServiceAccount(name, co.ServiceAccount); err != nil {
			return err
		}
	}

	for _, i := range co.Interface {
		if err := c.ValidateInstanceNetwork(name, i); err != nil {
			return err
		}
	}

	return nil
}

func (c *Config) ValidateInstance(i *config.Instance) error {
	switch {
	case i.Name == "":
		return NewError("instance name is required")

	case i.Role == "":
		return NewError("role is required for instance name %s", i.Name)

	case i.Description == "":
		return NewError("description not specified for instance %s", i.Name)

	case !isIdentifier(i.Name):
		return NewError("invalid instance name. Found %s", i.Name)

	case !isIdentifier(i.Role):
		return NewError("invalid role for instance %s: %s", i.Name, i.Role)
	}

	if _, ok := c.instance_types[i.Type]; !ok && i.Type != "" {
		return NewError("instance type %s for instance %s is unknown", i.Type, i.Name)
	}

	return c.ValidateInstanceCreateOptions(fmt.Sprintf("create options for instance %s", i.Name), i.CreateOptions, false)
}

func (c *Config) ValidateInstances() error {
	c.effective_instances = make(map[string]*config.Instance)

	for _, i := range c.Instance {
		ci := c.EffectiveInstance(i)
		err := c.ValidateInstance(ci)
		if err != nil {
			return err
		}

		if _, ok := c.effective_instances[i.Name]; ok {
			return NewError("duplicate instance name : %s", i.Name)
		}
		c.effective_instances[i.Name] = ci
	}

	if len(c.effective_instances) == 0 {
		return NewError("no instances defined")
	}
	return nil
}

func (c *Config) ValidateInstanceType(it *config.InstanceType) error {
	switch {
	case it.Name == "":
		return NewError("instance type name is required")

	case !isIdentifier(it.Name):
		return NewError("instance type name is invalid : %s", it.Name)
	}

	return c.ValidateInstanceCreateOptions(
		fmt.Sprintf("create options for instance type %s", it.Name), it.CreateOptions, true)
}

func (c *Config) ValidateInstanceTypes() error {
	c.instance_types = make(map[string]*config.InstanceType)
	for _, it := range c.InstanceType {
		if err := c.ValidateInstanceType(it); err != nil {
			return err
		}

		if _, ok := c.instance_types[it.Name]; ok {
			return NewError("duplicate instance type name : %s", it.Name)
		}
		c.instance_types[it.Name] = it
	}
	return nil
}

func (c *Config) ValidateWindowsDomain(d *config.WindowsDomain) error {
	if !isIdentifier(d.Name) {
		return NewError("invalid domain name %s", d.Name)
	}

	for _, a := range d.Account {
		if !isIdentifier(a.Name) {
			return NewError("invalid username %s in domain %s", a.Name, d.Name)
		}
	}
	return nil
}

func (c *Config) ValidateWindowsDomains() error {
	c.domains = make(map[string]*config.WindowsDomain)

	for _, d := range c.WindowsDomain {
		err := c.ValidateWindowsDomain(d)
		if err != nil {
			return err
		}

		if _, ok := c.domains[d.Name]; ok {
			return NewError("duplicate domain name %s", d.Name)
		}
		c.domains[d.Name] = d
	}

	return nil
}

// Most GCE identifiers should be valid RFC 1035 labels and must match
// [a-z]([-a-z0-9]*[a-z0-9])?.
//
// In addition to this, project names look like "<domain-name>:<label>." Some
// APIs enforce additional constraints -- e.g. length limits. Those are outside
// the scope for validation here.
//
// The regexp below only intends to be /kinda/ correct and is only here to
// catch cases where something egregiously wrong was entered into the
// configuration file.
//
// During development, we want to catch errors early as possible. Hence we are
// going to use a lenient regexp that'll catch errors like entering a full or
// partial URL into a field which should only contain a single identifier. But
// it won't catch subtle errors like using the wrong identifier.
const kIdentifierMatcher = "[a-z][-a-z0-9:.]*"

var idMatcher = matcher("*")

// isIdentifier returns true if |s| is probably a valid token. It returns no false
// negatives, but may return false positives.
func isIdentifier(s string) bool {
	return idMatcher.MatchString(s)
}

// matcher returns a RegExp object that matches a globbed pattern like
// "foo/*/bar" where each "*" represents an identifier.
func matcher(m string) *regexp.Regexp {
	// This is typical.
	if m == "*" {
		return regexp.MustCompile("^" + kIdentifierMatcher + "$")
	}

	m = strings.Replace(m, "/*/", "/"+kIdentifierMatcher+"/", -1)

	if strings.HasSuffix(m, "/*") {
		m = m[0:len(m)-1] + kIdentifierMatcher
	}

	return regexp.MustCompile("^" + m + "$")
}

// oneOfString returns the first string argument that is not empty.
func oneOfString(s ...string) string {
	for _, v := range s {
		if v != "" {
			return v
		}
	}
	return ""
}

// oneOfStringList returns the first []string argument that is not empty.
func oneOfStringList(s ...[]string) []string {
	for _, v := range s {
		if len(v) != 0 {
			return v
		}
	}
	return nil
}

// metadataFromKVMap constructs a compute.Metadata object from a
// map[string]string. The keys and values in the map become the Key and Value
// attributes of the MeatadataItem.
func metadataFromKVMap(m map[string]string) *compute.Metadata {
	items := []*compute.MetadataItems{}

	for k, v := range m {
		items = append(items, &compute.MetadataItems{
			Key:   k,
			Value: &v})
	}
	return &compute.Metadata{
		Items: items}
}
