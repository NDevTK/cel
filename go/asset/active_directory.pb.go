// Code generated by protoc-gen-go.
// source: schema/asset/active_directory.proto
// DO NOT EDIT!

/*
Package asset is a generated protocol buffer package.

It is generated from these files:
	schema/asset/active_directory.proto
	schema/asset/cert.proto
	schema/asset/dns.proto
	schema/asset/iis.proto
	schema/asset/network.proto
	schema/asset/common.proto

It has these top-level messages:
	ActiveDirectoryDomain
	ActiveDirectoryDomainController
	WindowsContainer
	WindowsGroup
	GroupReference
	WindowsUser
	NetworkInterface
	WindowsMachine
	Certificate
	CertificatePool
	DNSZone
	DNSRecord
	Server
	Bindings
	Site
	Application
	Network
	NetworkPeer
	Address
	AddressRange
	FixedAddress
	FileReference
*/
package asset

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Active Directory functional level. A.k.a. Domain Mode. See
// https://docs.microsoft.com/en-us/windows-server/identity/ad-ds/active-directory-functional-levels
// for more details on the specific features that are available at each
// functional level.
type ActiveDirectoryDomain_Mode int32

const (
	// Use the default. The default functional level depends on the host OS and
	// on the other AD DS servers in the domain or forest.
	ActiveDirectoryDomain_DEFAULT   ActiveDirectoryDomain_Mode = 0
	ActiveDirectoryDomain_Win2003   ActiveDirectoryDomain_Mode = 2
	ActiveDirectoryDomain_Win2008   ActiveDirectoryDomain_Mode = 3
	ActiveDirectoryDomain_Win2008R2 ActiveDirectoryDomain_Mode = 4
	ActiveDirectoryDomain_Win2012   ActiveDirectoryDomain_Mode = 5
	ActiveDirectoryDomain_Win2012R2 ActiveDirectoryDomain_Mode = 6
	ActiveDirectoryDomain_Win2016   ActiveDirectoryDomain_Mode = 7
)

var ActiveDirectoryDomain_Mode_name = map[int32]string{
	0: "DEFAULT",
	2: "Win2003",
	3: "Win2008",
	4: "Win2008R2",
	5: "Win2012",
	6: "Win2012R2",
	7: "Win2016",
}
var ActiveDirectoryDomain_Mode_value = map[string]int32{
	"DEFAULT":   0,
	"Win2003":   2,
	"Win2008":   3,
	"Win2008R2": 4,
	"Win2012":   5,
	"Win2012R2": 6,
	"Win2016":   7,
}

func (x ActiveDirectoryDomain_Mode) String() string {
	return proto.EnumName(ActiveDirectoryDomain_Mode_name, int32(x))
}
func (ActiveDirectoryDomain_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

// Domain type
type ActiveDirectoryDomain_Type int32

const (
	// Child of an existing domain. (This is the default)
	ActiveDirectoryDomain_CHILD ActiveDirectoryDomain_Type = 0
	// Root of a new domain tree.
	ActiveDirectoryDomain_TREE ActiveDirectoryDomain_Type = 1
)

var ActiveDirectoryDomain_Type_name = map[int32]string{
	0: "CHILD",
	1: "TREE",
}
var ActiveDirectoryDomain_Type_value = map[string]int32{
	"CHILD": 0,
	"TREE":  1,
}

func (x ActiveDirectoryDomain_Type) String() string {
	return proto.EnumName(ActiveDirectoryDomain_Type_name, int32(x))
}
func (ActiveDirectoryDomain_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 1}
}

// Describes an Active Directory domain or forest.
type ActiveDirectoryDomain struct {
	// FQDN of ActiveDirectory domain.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Parent domain name. Only specify this if this domain is going to be a
	// child domain.
	ParentName string                     `protobuf:"bytes,2,opt,name=parent_name,json=parentName" json:"parent_name,omitempty"`
	DomainMode ActiveDirectoryDomain_Mode `protobuf:"varint,3,opt,name=domain_mode,json=domainMode,enum=asset.ActiveDirectoryDomain_Mode" json:"domain_mode,omitempty"`
	// NetBIOS name. Required if |name| is longer than 15 characters.
	NetbiosName string                     `protobuf:"bytes,4,opt,name=netbios_name,json=netbiosName" json:"netbios_name,omitempty"`
	Type        ActiveDirectoryDomain_Type `protobuf:"varint,5,opt,name=type,enum=asset.ActiveDirectoryDomain_Type" json:"type,omitempty"`
	// Forest functional level (only applicable when creating a new forest.
	ForestMode ActiveDirectoryDomain_Mode `protobuf:"varint,6,opt,name=forest_mode,json=forestMode,enum=asset.ActiveDirectoryDomain_Mode" json:"forest_mode,omitempty"`
}

func (m *ActiveDirectoryDomain) Reset()                    { *m = ActiveDirectoryDomain{} }
func (m *ActiveDirectoryDomain) String() string            { return proto.CompactTextString(m) }
func (*ActiveDirectoryDomain) ProtoMessage()               {}
func (*ActiveDirectoryDomain) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ActiveDirectoryDomain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ActiveDirectoryDomain) GetParentName() string {
	if m != nil {
		return m.ParentName
	}
	return ""
}

func (m *ActiveDirectoryDomain) GetDomainMode() ActiveDirectoryDomain_Mode {
	if m != nil {
		return m.DomainMode
	}
	return ActiveDirectoryDomain_DEFAULT
}

func (m *ActiveDirectoryDomain) GetNetbiosName() string {
	if m != nil {
		return m.NetbiosName
	}
	return ""
}

func (m *ActiveDirectoryDomain) GetType() ActiveDirectoryDomain_Type {
	if m != nil {
		return m.Type
	}
	return ActiveDirectoryDomain_CHILD
}

func (m *ActiveDirectoryDomain) GetForestMode() ActiveDirectoryDomain_Mode {
	if m != nil {
		return m.ForestMode
	}
	return ActiveDirectoryDomain_DEFAULT
}

// Describes a single Active Directory Domain Controller.
type ActiveDirectoryDomainController struct {
	// Name of the domain. Must match a ActiveDirectoryDomain entry.
	Domain string `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
	// Machine hosting the ADDS. Must match a WindowsMachine entry.
	Machine string `protobuf:"bytes,2,opt,name=machine" json:"machine,omitempty"`
	// Whether or not to install a DNS server on this machine. The default is
	// almost always |true| unless the domain already exists and the existing
	// domain controller does not host a DNS server.
	//
	// Don't specify the option if you would like the domain controller to do the
	// default action. Or specify it to force one or the other.
	//
	// Types that are valid to be assigned to OptionalDns:
	//	*ActiveDirectoryDomainController_InstallDns
	OptionalDns isActiveDirectoryDomainController_OptionalDns `protobuf_oneof:"optional_dns"`
	// Assume DNS service is not available on the network. Only applicable when
	// installing DNS services. If this field is not set, or set to false, then
	// the installation can assume that the TCP/IP client settings of the host OS
	// specifies the DNS server to use.
	NoDnsOnNetwork bool `protobuf:"varint,4,opt,name=no_dns_on_network,json=noDnsOnNetwork" json:"no_dns_on_network,omitempty"`
	// This domain controller should not be a global catalog server. Default is
	// to run with global catalog for Win2012 or later.
	NoGlobalCatalog bool `protobuf:"varint,5,opt,name=no_global_catalog,json=noGlobalCatalog" json:"no_global_catalog,omitempty"`
	// If true, attempts to create a DNS delegation for the new DNS server. Only
	// applicable when installing a DNS server. E.g.: If the authoritative DNS
	// server for foo.example.com is using ActiveDirectory, and we are installing
	// the subordinate domain bar, then setting this value to true causes
	// foo.example.com to delegate the bar domain to the new DNS server.
	CreateDnsDelegation bool `protobuf:"varint,6,opt,name=create_dns_delegation,json=createDnsDelegation" json:"create_dns_delegation,omitempty"`
}

func (m *ActiveDirectoryDomainController) Reset()                    { *m = ActiveDirectoryDomainController{} }
func (m *ActiveDirectoryDomainController) String() string            { return proto.CompactTextString(m) }
func (*ActiveDirectoryDomainController) ProtoMessage()               {}
func (*ActiveDirectoryDomainController) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isActiveDirectoryDomainController_OptionalDns interface {
	isActiveDirectoryDomainController_OptionalDns()
}

type ActiveDirectoryDomainController_InstallDns struct {
	InstallDns bool `protobuf:"varint,3,opt,name=install_dns,json=installDns,oneof"`
}

func (*ActiveDirectoryDomainController_InstallDns) isActiveDirectoryDomainController_OptionalDns() {}

func (m *ActiveDirectoryDomainController) GetOptionalDns() isActiveDirectoryDomainController_OptionalDns {
	if m != nil {
		return m.OptionalDns
	}
	return nil
}

func (m *ActiveDirectoryDomainController) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *ActiveDirectoryDomainController) GetMachine() string {
	if m != nil {
		return m.Machine
	}
	return ""
}

func (m *ActiveDirectoryDomainController) GetInstallDns() bool {
	if x, ok := m.GetOptionalDns().(*ActiveDirectoryDomainController_InstallDns); ok {
		return x.InstallDns
	}
	return false
}

func (m *ActiveDirectoryDomainController) GetNoDnsOnNetwork() bool {
	if m != nil {
		return m.NoDnsOnNetwork
	}
	return false
}

func (m *ActiveDirectoryDomainController) GetNoGlobalCatalog() bool {
	if m != nil {
		return m.NoGlobalCatalog
	}
	return false
}

func (m *ActiveDirectoryDomainController) GetCreateDnsDelegation() bool {
	if m != nil {
		return m.CreateDnsDelegation
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ActiveDirectoryDomainController) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ActiveDirectoryDomainController_OneofMarshaler, _ActiveDirectoryDomainController_OneofUnmarshaler, _ActiveDirectoryDomainController_OneofSizer, []interface{}{
		(*ActiveDirectoryDomainController_InstallDns)(nil),
	}
}

func _ActiveDirectoryDomainController_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ActiveDirectoryDomainController)
	// optional_dns
	switch x := m.OptionalDns.(type) {
	case *ActiveDirectoryDomainController_InstallDns:
		t := uint64(0)
		if x.InstallDns {
			t = 1
		}
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("ActiveDirectoryDomainController.OptionalDns has unexpected type %T", x)
	}
	return nil
}

func _ActiveDirectoryDomainController_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ActiveDirectoryDomainController)
	switch tag {
	case 3: // optional_dns.install_dns
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.OptionalDns = &ActiveDirectoryDomainController_InstallDns{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _ActiveDirectoryDomainController_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ActiveDirectoryDomainController)
	// optional_dns
	switch x := m.OptionalDns.(type) {
	case *ActiveDirectoryDomainController_InstallDns:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Describes a container that a Windows asset can reside in.
//
// Resources like machines, users, and groups can be specified per domain, per
// machine, or per organizational unit. When specifying one of these asset
// types, use the WindowsContainer member to specify where to create the asset.
type WindowsContainer struct {
	// Types that are valid to be assigned to Container:
	//	*WindowsContainer_Domain
	//	*WindowsContainer_Machine
	//	*WindowsContainer_Ou
	Container isWindowsContainer_Container `protobuf_oneof:"container"`
}

func (m *WindowsContainer) Reset()                    { *m = WindowsContainer{} }
func (m *WindowsContainer) String() string            { return proto.CompactTextString(m) }
func (*WindowsContainer) ProtoMessage()               {}
func (*WindowsContainer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isWindowsContainer_Container interface {
	isWindowsContainer_Container()
}

type WindowsContainer_Domain struct {
	Domain string `protobuf:"bytes,1,opt,name=domain,oneof"`
}
type WindowsContainer_Machine struct {
	Machine string `protobuf:"bytes,2,opt,name=machine,oneof"`
}
type WindowsContainer_Ou struct {
	Ou string `protobuf:"bytes,3,opt,name=ou,oneof"`
}

func (*WindowsContainer_Domain) isWindowsContainer_Container()  {}
func (*WindowsContainer_Machine) isWindowsContainer_Container() {}
func (*WindowsContainer_Ou) isWindowsContainer_Container()      {}

func (m *WindowsContainer) GetContainer() isWindowsContainer_Container {
	if m != nil {
		return m.Container
	}
	return nil
}

func (m *WindowsContainer) GetDomain() string {
	if x, ok := m.GetContainer().(*WindowsContainer_Domain); ok {
		return x.Domain
	}
	return ""
}

func (m *WindowsContainer) GetMachine() string {
	if x, ok := m.GetContainer().(*WindowsContainer_Machine); ok {
		return x.Machine
	}
	return ""
}

func (m *WindowsContainer) GetOu() string {
	if x, ok := m.GetContainer().(*WindowsContainer_Ou); ok {
		return x.Ou
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*WindowsContainer) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _WindowsContainer_OneofMarshaler, _WindowsContainer_OneofUnmarshaler, _WindowsContainer_OneofSizer, []interface{}{
		(*WindowsContainer_Domain)(nil),
		(*WindowsContainer_Machine)(nil),
		(*WindowsContainer_Ou)(nil),
	}
}

func _WindowsContainer_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*WindowsContainer)
	// container
	switch x := m.Container.(type) {
	case *WindowsContainer_Domain:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Domain)
	case *WindowsContainer_Machine:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Machine)
	case *WindowsContainer_Ou:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Ou)
	case nil:
	default:
		return fmt.Errorf("WindowsContainer.Container has unexpected type %T", x)
	}
	return nil
}

func _WindowsContainer_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*WindowsContainer)
	switch tag {
	case 1: // container.domain
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Container = &WindowsContainer_Domain{x}
		return true, err
	case 2: // container.machine
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Container = &WindowsContainer_Machine{x}
		return true, err
	case 3: // container.ou
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Container = &WindowsContainer_Ou{x}
		return true, err
	default:
		return false, nil
	}
}

func _WindowsContainer_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*WindowsContainer)
	// container
	switch x := m.Container.(type) {
	case *WindowsContainer_Domain:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Domain)))
		n += len(x.Domain)
	case *WindowsContainer_Machine:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Machine)))
		n += len(x.Machine)
	case *WindowsContainer_Ou:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Ou)))
		n += len(x.Ou)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Descibes an Active Directory or Windows local group.
type WindowsGroup struct {
	// Name of the group. Exclude the domain name. The name alone is not
	// sufficient if this group corresponds to a Well Known group. Use the
	// |well_known_sid| field for that.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Human readable description of the group.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// Container for the group. A container must be specified for a WindowsGroup.
	Container *WindowsContainer `protobuf:"bytes,3,opt,name=container" json:"container,omitempty"`
	// Well-known security identifier. The string should be of the form S-* and
	// should correspond to a known SID as described in
	// https://support.microsoft.com/en-us/help/243330/well-known-security-identifiers-in-windows-operating-systems.
	//
	// Only specify this field if this group corresponds to a well known group.
	WellKnownSid string `protobuf:"bytes,4,opt,name=well_known_sid,json=wellKnownSid" json:"well_known_sid,omitempty"`
}

func (m *WindowsGroup) Reset()                    { *m = WindowsGroup{} }
func (m *WindowsGroup) String() string            { return proto.CompactTextString(m) }
func (*WindowsGroup) ProtoMessage()               {}
func (*WindowsGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *WindowsGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WindowsGroup) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *WindowsGroup) GetContainer() *WindowsContainer {
	if m != nil {
		return m.Container
	}
	return nil
}

func (m *WindowsGroup) GetWellKnownSid() string {
	if m != nil {
		return m.WellKnownSid
	}
	return ""
}

// A reference to a group. The combination of |name| and |container| must match
// one of the WindowsGroup entries.
type GroupReference struct {
	// The name of the group.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Location. Since GroupReference messages are typically specified as a field
	// of an object that already has a container, omiting this field results in
	// the GroupReference inheriting the parent object's container. Take for
	// example, the following WindowsUser definition:
	//
	//     windows_user {
	//       name: 'joe'
	//       container: { domain: 'foo.example' }
	//       member_of: { name: 'bar' }
	//     }
	//
	// This results in the user being a member of the group 'bar' in the
	// 'foo.example' AD domain because that's the enclosing container.  Note
	// however, that inheriting in this manner isn't always correct since it is
	// possible for users to be members of groups from other containers.
	Container *WindowsContainer `protobuf:"bytes,2,opt,name=container" json:"container,omitempty"`
}

func (m *GroupReference) Reset()                    { *m = GroupReference{} }
func (m *GroupReference) String() string            { return proto.CompactTextString(m) }
func (*GroupReference) ProtoMessage()               {}
func (*GroupReference) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GroupReference) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GroupReference) GetContainer() *WindowsContainer {
	if m != nil {
		return m.Container
	}
	return nil
}

// Describes a Active Directory or a Windows local user.
type WindowsUser struct {
	// Name of the user. Exclude the domain name.
	//
	// E.g.: joe
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Description. A.k.a. Full name.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// Container for the user. A container must be specified for a WindowsUser.
	Container *WindowsContainer `protobuf:"bytes,3,opt,name=container" json:"container,omitempty"`
	// Password. Can be left blank in which case each instantiation of the lab
	// will cause a new password to be generated.
	Password string `protobuf:"bytes,5,opt,name=password" json:"password,omitempty"`
	// List of groups that the user belongs to.
	MemberOf []*GroupReference `protobuf:"bytes,6,rep,name=member_of,json=memberOf" json:"member_of,omitempty"`
}

func (m *WindowsUser) Reset()                    { *m = WindowsUser{} }
func (m *WindowsUser) String() string            { return proto.CompactTextString(m) }
func (*WindowsUser) ProtoMessage()               {}
func (*WindowsUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *WindowsUser) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WindowsUser) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *WindowsUser) GetContainer() *WindowsContainer {
	if m != nil {
		return m.Container
	}
	return nil
}

func (m *WindowsUser) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *WindowsUser) GetMemberOf() []*GroupReference {
	if m != nil {
		return m.MemberOf
	}
	return nil
}

// NetworkInterface describes a single network interface on a machine.
type NetworkInterface struct {
	// Name of Network entry describing the network that this interface is
	// attached to.
	Network string `protobuf:"bytes,1,opt,name=network" json:"network,omitempty"`
	// Fixed address, if this interface is to have one. Leave undefined if the
	// interface should obatain an address automatically.
	FixedAddress *FixedAddress `protobuf:"bytes,2,opt,name=fixed_address,json=fixedAddress" json:"fixed_address,omitempty"`
}

func (m *NetworkInterface) Reset()                    { *m = NetworkInterface{} }
func (m *NetworkInterface) String() string            { return proto.CompactTextString(m) }
func (*NetworkInterface) ProtoMessage()               {}
func (*NetworkInterface) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *NetworkInterface) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *NetworkInterface) GetFixedAddress() *FixedAddress {
	if m != nil {
		return m.FixedAddress
	}
	return nil
}

// A Windows machine.
type WindowsMachine struct {
	// Name of the machine. This name will become the hostname for the machine,
	// both absolute and domain relative (if applicable). Hence must be globally
	// unique.
	//
	// For Windows machines, it's advisable to have *short* hostnames, ideally
	// shorter than 11 characters. This allows the name to do double duty as a
	// NetBios name as well as a DNS hostname.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Container. Only |domain| and |ou| values are acceptable. Currently
	// |machine| is not a valid option. If no container is specified, the machine
	// will be brought up as a standalone workstation or server depending on the
	// installed operating system.
	//
	// This field should be empty for a machine that's referenced in a
	// ActiveDirectoryDomainController entry.
	//
	// Specifying this field results in the machine being joined to the specified
	// domain and, if necessary, placed in the specified container.
	Container *WindowsContainer `protobuf:"bytes,3,opt,name=container" json:"container,omitempty"`
	// The name of a host.MachineType entry that describes the host machine.
	HostMachineType string `protobuf:"bytes,4,opt,name=host_machine_type,json=hostMachineType" json:"host_machine_type,omitempty"`
	// System locale. If left unspecified, the default is left unchanged. Use the
	// following PowerShell command to determine the list of available locales on
	// a Windows machine:
	//
	//     ``` ps1
	//     [System.Globalization.CultureInfo]::GetCultures([System.Globalization.CultureTypes]::AllCultures).name
	//     ```
	//
	// PS DSC Reference:
	// * https://github.com/PowerShell/SystemLocaleDsc
	Locale string `protobuf:"bytes,5,opt,name=locale" json:"locale,omitempty"`
	// Set the system timezone. If left unspecified, the default is left
	// unchanged. Use the following PowerShell command to determine the lsit of
	// available timezone identifiers on a Windows machine:
	//
	//     ``` ps1
	//     [System.TimeZoneInfo]::GetSystemTimeZones().Id
	//     ```
	//
	// PS DSC Reference:
	// * https://github.com/PowerShell/xTimeZone
	Timezone string `protobuf:"bytes,6,opt,name=timezone" json:"timezone,omitempty"`
	// Network interfaces. There can be more than one for multihomed machines.
	// There MUST be at least one of these.
	NetworkInterface []*NetworkInterface `protobuf:"bytes,7,rep,name=network_interface,json=networkInterface" json:"network_interface,omitempty"`
	// List of additional Windows features or roles to install. The values here
	// should be valid for the selected host machine type. You can use the
	// 'Get-WindowsFeature' PowerShell commandlet to retrieve a list of available
	// Windows features.
	//
	// E.g.: windows_feature: "Web-Server"
	//
	// Note: This method cannot be used to specify all sub-features. All
	// features that needs to be installed should be listed explicitly.
	//
	// Note: Addition of roles can cause features to be installed implicitly.
	// E.g. specifying a machine as the host for an IIS site will automatically
	// install the necessary web server roles. The |windows_feature| field should
	// be used for features that otherwise won't be installed as part of any such
	// role assignment.
	WindowsFeature []string `protobuf:"bytes,8,rep,name=windows_feature,json=windowsFeature" json:"windows_feature,omitempty"`
	// A configuration file. Specify this if you've run Windows Server Manager
	// and produced a configration file already.
	ConfigurationFile *FileReference `protobuf:"bytes,9,opt,name=configuration_file,json=configurationFile" json:"configuration_file,omitempty"`
}

func (m *WindowsMachine) Reset()                    { *m = WindowsMachine{} }
func (m *WindowsMachine) String() string            { return proto.CompactTextString(m) }
func (*WindowsMachine) ProtoMessage()               {}
func (*WindowsMachine) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *WindowsMachine) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WindowsMachine) GetContainer() *WindowsContainer {
	if m != nil {
		return m.Container
	}
	return nil
}

func (m *WindowsMachine) GetHostMachineType() string {
	if m != nil {
		return m.HostMachineType
	}
	return ""
}

func (m *WindowsMachine) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

func (m *WindowsMachine) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *WindowsMachine) GetNetworkInterface() []*NetworkInterface {
	if m != nil {
		return m.NetworkInterface
	}
	return nil
}

func (m *WindowsMachine) GetWindowsFeature() []string {
	if m != nil {
		return m.WindowsFeature
	}
	return nil
}

func (m *WindowsMachine) GetConfigurationFile() *FileReference {
	if m != nil {
		return m.ConfigurationFile
	}
	return nil
}

func init() {
	proto.RegisterType((*ActiveDirectoryDomain)(nil), "asset.ActiveDirectoryDomain")
	proto.RegisterType((*ActiveDirectoryDomainController)(nil), "asset.ActiveDirectoryDomainController")
	proto.RegisterType((*WindowsContainer)(nil), "asset.WindowsContainer")
	proto.RegisterType((*WindowsGroup)(nil), "asset.WindowsGroup")
	proto.RegisterType((*GroupReference)(nil), "asset.GroupReference")
	proto.RegisterType((*WindowsUser)(nil), "asset.WindowsUser")
	proto.RegisterType((*NetworkInterface)(nil), "asset.NetworkInterface")
	proto.RegisterType((*WindowsMachine)(nil), "asset.WindowsMachine")
	proto.RegisterEnum("asset.ActiveDirectoryDomain_Mode", ActiveDirectoryDomain_Mode_name, ActiveDirectoryDomain_Mode_value)
	proto.RegisterEnum("asset.ActiveDirectoryDomain_Type", ActiveDirectoryDomain_Type_name, ActiveDirectoryDomain_Type_value)
}

func init() { proto.RegisterFile("schema/asset/active_directory.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 875 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x5d, 0x6e, 0xdb, 0x46,
	0x10, 0xb6, 0x7e, 0x2c, 0x4b, 0x43, 0x45, 0xa1, 0x37, 0x75, 0xaa, 0xaa, 0x0f, 0x91, 0xd9, 0x02,
	0x75, 0xf3, 0x20, 0xc5, 0x34, 0x52, 0xe4, 0xd5, 0xb6, 0xec, 0x38, 0x68, 0x7e, 0x00, 0xd6, 0x41,
	0x80, 0xf6, 0x81, 0x58, 0x93, 0x43, 0x79, 0x11, 0x72, 0x47, 0xd8, 0xa5, 0xaa, 0xba, 0xb7, 0xe8,
	0x11, 0x7a, 0x85, 0x1e, 0xa3, 0x97, 0xe9, 0x15, 0x0a, 0xee, 0xae, 0x24, 0xcb, 0x30, 0x50, 0x3f,
	0xf5, 0x6d, 0x67, 0xbe, 0xe1, 0xcc, 0x37, 0xdf, 0xcc, 0x2e, 0xe1, 0x1b, 0x9d, 0x5c, 0x63, 0xc1,
	0xc7, 0x5c, 0x6b, 0x2c, 0xc7, 0x3c, 0x29, 0xc5, 0xaf, 0x18, 0xa7, 0x42, 0x61, 0x52, 0x92, 0xba,
	0x19, 0xcd, 0x14, 0x95, 0xc4, 0xb6, 0x0d, 0x3a, 0xf8, 0x6a, 0x23, 0x36, 0xa1, 0xa2, 0x20, 0x69,
	0x23, 0x06, 0x83, 0x0d, 0x48, 0x62, 0xb9, 0x20, 0xf5, 0xd9, 0x62, 0xc1, 0x5f, 0x0d, 0xd8, 0x3b,
	0x36, 0x89, 0x27, 0xcb, 0xbc, 0x13, 0x2a, 0xb8, 0x90, 0x8c, 0x41, 0x53, 0xf2, 0x02, 0xfb, 0xb5,
	0x61, 0xed, 0xa0, 0x13, 0x99, 0x33, 0x7b, 0x06, 0xde, 0x8c, 0x2b, 0x94, 0x65, 0x6c, 0xa0, 0xba,
	0x81, 0xc0, 0xba, 0xde, 0x57, 0x01, 0x27, 0xe0, 0xa5, 0xe6, 0xf3, 0xb8, 0xa0, 0x14, 0xfb, 0x8d,
	0x61, 0xed, 0xa0, 0x17, 0xee, 0x8f, 0x4c, 0xe5, 0xd1, 0xbd, 0x75, 0x46, 0xef, 0x28, 0xc5, 0x08,
	0xec, 0x57, 0xd5, 0x99, 0xed, 0x43, 0x57, 0x62, 0x79, 0x25, 0x48, 0xdb, 0x2a, 0x4d, 0x53, 0xc5,
	0x73, 0x3e, 0x53, 0xe6, 0x25, 0x34, 0xcb, 0x9b, 0x19, 0xf6, 0xb7, 0x1f, 0x90, 0xff, 0xf2, 0x66,
	0x86, 0x91, 0x09, 0xaf, 0xd8, 0x65, 0xa4, 0x50, 0x97, 0x96, 0x5d, 0xeb, 0xc1, 0xec, 0xec, 0x57,
	0xd5, 0x39, 0x40, 0x68, 0x1a, 0x96, 0x1e, 0xec, 0x4c, 0xce, 0xce, 0x8f, 0x3f, 0xbe, 0xbd, 0xf4,
	0xb7, 0x2a, 0xe3, 0x93, 0x90, 0xe1, 0x8b, 0x17, 0x47, 0x7e, 0x7d, 0x6d, 0xbc, 0xf2, 0x1b, 0xec,
	0x11, 0x74, 0x9c, 0x11, 0x85, 0x7e, 0x73, 0x85, 0x1d, 0x86, 0xfe, 0xf6, 0x0a, 0x3b, 0x0c, 0xa3,
	0xd0, 0x6f, 0xad, 0xb1, 0x1f, 0xfc, 0x9d, 0xe0, 0x6b, 0x68, 0x56, 0xc4, 0x59, 0x07, 0xb6, 0x4f,
	0x2f, 0xde, 0xbc, 0x9d, 0xf8, 0x5b, 0xac, 0x0d, 0xcd, 0xcb, 0xe8, 0xec, 0xcc, 0xaf, 0x05, 0x7f,
	0xd4, 0xe1, 0xd9, 0xbd, 0x74, 0x4f, 0x49, 0x96, 0x8a, 0xf2, 0x1c, 0x15, 0x7b, 0x0a, 0x2d, 0xab,
	0xa9, 0x1b, 0xa0, 0xb3, 0x58, 0x1f, 0x76, 0x0a, 0x9e, 0x5c, 0x0b, 0xb9, 0x1c, 0xdf, 0xd2, 0x64,
	0xfb, 0xe0, 0x09, 0xa9, 0x4b, 0x9e, 0xe7, 0x71, 0x2a, 0xb5, 0x99, 0x5d, 0xfb, 0x62, 0x2b, 0x02,
	0xe7, 0x9c, 0x48, 0xcd, 0xbe, 0x87, 0x5d, 0x49, 0x15, 0x1a, 0x93, 0x8c, 0xdd, 0x22, 0x99, 0xf9,
	0xb4, 0xa3, 0x9e, 0xa4, 0x89, 0xd4, 0x1f, 0xe4, 0x7b, 0xeb, 0x65, 0xcf, 0x4d, 0xe8, 0x34, 0xa7,
	0x2b, 0x9e, 0xc7, 0x09, 0x2f, 0x79, 0x4e, 0x53, 0x33, 0xaf, 0x76, 0xf4, 0x58, 0xd2, 0x6b, 0xe3,
	0x3f, 0xb5, 0x6e, 0x16, 0xc2, 0x5e, 0xa2, 0x90, 0x97, 0x68, 0x52, 0xa7, 0x98, 0xe3, 0x94, 0x97,
	0x82, 0xa4, 0x99, 0x50, 0x3b, 0x7a, 0x62, 0xc1, 0x89, 0xd4, 0x93, 0x15, 0x74, 0xd2, 0x83, 0x2e,
	0xcd, 0xaa, 0x13, 0x37, 0x74, 0x83, 0x29, 0xf8, 0x9f, 0x84, 0x4c, 0x69, 0xa1, 0x2b, 0x11, 0xb8,
	0x90, 0xa8, 0x58, 0x7f, 0x53, 0x83, 0x8b, 0xad, 0x95, 0x0a, 0x83, 0x3b, 0x2a, 0x5c, 0x6c, 0xad,
	0x75, 0xf0, 0xa1, 0x4e, 0x73, 0xd3, 0x7e, 0xe5, 0xae, 0xd3, 0xfc, 0xc4, 0x83, 0x4e, 0xb2, 0x4c,
	0x1a, 0xfc, 0x59, 0x83, 0xae, 0xab, 0xf4, 0x5a, 0xd1, 0x7c, 0x76, 0xef, 0x45, 0x19, 0x82, 0x97,
	0xa2, 0x4e, 0x94, 0x30, 0x14, 0x9d, 0xd2, 0xb7, 0x5d, 0xec, 0xe5, 0xad, 0x9c, 0xa6, 0x98, 0x17,
	0x7e, 0xe9, 0x36, 0xf1, 0x6e, 0x1f, 0xd1, 0x3a, 0x92, 0x7d, 0x0b, 0xbd, 0x05, 0xe6, 0x79, 0xfc,
	0x59, 0xd2, 0x42, 0xc6, 0x5a, 0xa4, 0xee, 0x7a, 0x74, 0x2b, 0xef, 0x8f, 0x95, 0xf3, 0x27, 0x91,
	0x06, 0xbf, 0x40, 0xcf, 0x70, 0x8b, 0x30, 0x43, 0x85, 0x32, 0xc1, 0x7b, 0x49, 0x6e, 0x50, 0xa8,
	0x3f, 0x94, 0x42, 0xf0, 0x77, 0x0d, 0x3c, 0x87, 0x7f, 0xd4, 0xa8, 0xfe, 0xdf, 0xfe, 0x07, 0xd0,
	0x9e, 0x71, 0xad, 0x17, 0xa4, 0x52, 0xb3, 0x4d, 0x9d, 0x68, 0x65, 0xb3, 0x10, 0x3a, 0x05, 0x16,
	0x57, 0xa8, 0x62, 0xca, 0xfa, 0xad, 0x61, 0xe3, 0xc0, 0x0b, 0xf7, 0x5c, 0xca, 0x4d, 0x35, 0xa2,
	0xb6, 0x8d, 0xfb, 0x90, 0x05, 0x19, 0xf8, 0x6e, 0x63, 0xdf, 0xc8, 0x12, 0x55, 0xc6, 0x13, 0xac,
	0xae, 0xc8, 0x72, 0xb7, 0x6d, 0x4f, 0x4b, 0x93, 0xbd, 0x82, 0x47, 0x99, 0xf8, 0x0d, 0xd3, 0x98,
	0xa7, 0xa9, 0x42, 0xad, 0x9d, 0x6a, 0x4f, 0x5c, 0x95, 0xf3, 0x0a, 0x3b, 0xb6, 0x50, 0xd4, 0xcd,
	0x6e, 0x59, 0xc1, 0x3f, 0x75, 0xe8, 0xb9, 0xbe, 0xde, 0xb9, 0x3d, 0xfb, 0xcf, 0x91, 0x3c, 0x5c,
	0x95, 0xe7, 0xb0, 0x7b, 0x4d, 0xd5, 0xb3, 0x66, 0x53, 0xc7, 0xe6, 0x71, 0xb4, 0x8b, 0xf1, 0xb8,
	0x02, 0x5c, 0x49, 0xf3, 0xa2, 0x3c, 0x85, 0x56, 0x4e, 0x09, 0xcf, 0xd1, 0xe9, 0xe7, 0xac, 0x4a,
	0xd9, 0x52, 0x14, 0xf8, 0x3b, 0x49, 0xfb, 0x32, 0x76, 0xa2, 0x95, 0xcd, 0x26, 0xb0, 0xeb, 0x24,
	0x88, 0xc5, 0x52, 0xa6, 0xfe, 0x8e, 0x51, 0x78, 0x49, 0xef, 0xae, 0x8a, 0x91, 0x2f, 0xef, 0xea,
	0xfa, 0x1d, 0x3c, 0x5e, 0xd8, 0x26, 0xe2, 0x0c, 0x79, 0x39, 0x57, 0xd8, 0x6f, 0x0f, 0x1b, 0x07,
	0x9d, 0xa8, 0xe7, 0xdc, 0xe7, 0xd6, 0xcb, 0x4e, 0x81, 0x25, 0x24, 0x33, 0x31, 0x9d, 0x2b, 0x73,
	0xd9, 0xe3, 0x4c, 0xe4, 0xd8, 0xef, 0x18, 0x39, 0xbe, 0x58, 0x69, 0x9d, 0xe3, 0x7a, 0xa0, 0xbb,
	0x1b, 0xf1, 0x15, 0x76, 0x72, 0xf4, 0xf3, 0x61, 0x72, 0xad, 0xa8, 0x10, 0xf3, 0x62, 0x34, 0x25,
	0x9a, 0xe6, 0xa8, 0x69, 0xae, 0x12, 0x1c, 0x25, 0x54, 0x8c, 0xb1, 0x22, 0x35, 0x53, 0x42, 0xe3,
	0x38, 0xc1, 0x7c, 0x3c, 0x25, 0xfb, 0x73, 0xbc, 0x6a, 0x99, 0xbf, 0xe2, 0xd1, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xde, 0xd2, 0x5c, 0x88, 0x7a, 0x07, 0x00, 0x00,
}
