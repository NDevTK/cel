// Code generated by protoc-gen-go.
// source: schema/asset/network.proto
// DO NOT EDIT!

package asset

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// List of web protocols. Not necessarily exhaustive. We just need a convenient
// enum so that we don't need to defined it everywhere. Values should be self
// explanatory.
type Protocol int32

const (
	Protocol_UNKNOWN Protocol = 0
	Protocol_HTTP    Protocol = 1
	Protocol_HTTPS   Protocol = 2
)

var Protocol_name = map[int32]string{
	0: "UNKNOWN",
	1: "HTTP",
	2: "HTTPS",
}
var Protocol_value = map[string]int32{
	"UNKNOWN": 0,
	"HTTP":    1,
	"HTTPS":   2,
}

func (x Protocol) String() string {
	return proto.EnumName(Protocol_name, int32(x))
}
func (Protocol) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type Network struct {
	// Name of the network.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The address range assigned to the network. If left unspecified, an address
	// range will be determined when deploying the assets. This is the preferred
	// option unless an explicit address range is required.
	//
	// Two networks in the same asset manifest can't have overlapping address
	// ranges even if they aren't peers.
	AddressRange *AddressRange `protobuf:"bytes,2,opt,name=address_range,json=addressRange" json:"address_range,omitempty"`
}

func (m *Network) Reset()                    { *m = Network{} }
func (m *Network) String() string            { return proto.CompactTextString(m) }
func (*Network) ProtoMessage()               {}
func (*Network) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Network) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Network) GetAddressRange() *AddressRange {
	if m != nil {
		return m.AddressRange
	}
	return nil
}

// NetworkPeer describes a peering group. All networks that are a member of a
// peering group can route traffic across each other.
type NetworkPeer struct {
	// List of networks that should form a full mesh. Individual networks are
	// isolated by default and are only able to talk to each other if:
	//   * They are a part of a peering group, or
	//   * They are connected via a VPN gateway, or
	//   * They are connected via a [virtual] router.
	//
	// A single nework can participate in multiple disjoint peering groups,
	// however peering is not transitive. I.e. If {A,B} is a peering group, and
	// {B,C} is a peering group, traffic from A still can't route to C.
	Network []string `protobuf:"bytes,1,rep,name=network" json:"network,omitempty"`
}

func (m *NetworkPeer) Reset()                    { *m = NetworkPeer{} }
func (m *NetworkPeer) String() string            { return proto.CompactTextString(m) }
func (*NetworkPeer) ProtoMessage()               {}
func (*NetworkPeer) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *NetworkPeer) GetNetwork() []string {
	if m != nil {
		return m.Network
	}
	return nil
}

// Address is an IPv4 or IPv6 address.
type Address struct {
	// Types that are valid to be assigned to AddressType:
	//	*Address_Ipv4
	//	*Address_Ipv6
	AddressType isAddress_AddressType `protobuf_oneof:"address_type"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

type isAddress_AddressType interface {
	isAddress_AddressType()
}

type Address_Ipv4 struct {
	Ipv4 string `protobuf:"bytes,1,opt,name=ipv4,oneof"`
}
type Address_Ipv6 struct {
	Ipv6 string `protobuf:"bytes,2,opt,name=ipv6,oneof"`
}

func (*Address_Ipv4) isAddress_AddressType() {}
func (*Address_Ipv6) isAddress_AddressType() {}

func (m *Address) GetAddressType() isAddress_AddressType {
	if m != nil {
		return m.AddressType
	}
	return nil
}

func (m *Address) GetIpv4() string {
	if x, ok := m.GetAddressType().(*Address_Ipv4); ok {
		return x.Ipv4
	}
	return ""
}

func (m *Address) GetIpv6() string {
	if x, ok := m.GetAddressType().(*Address_Ipv6); ok {
		return x.Ipv6
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Address) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Address_OneofMarshaler, _Address_OneofUnmarshaler, _Address_OneofSizer, []interface{}{
		(*Address_Ipv4)(nil),
		(*Address_Ipv6)(nil),
	}
}

func _Address_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Address)
	// address_type
	switch x := m.AddressType.(type) {
	case *Address_Ipv4:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Ipv4)
	case *Address_Ipv6:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Ipv6)
	case nil:
	default:
		return fmt.Errorf("Address.AddressType has unexpected type %T", x)
	}
	return nil
}

func _Address_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Address)
	switch tag {
	case 1: // address_type.ipv4
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.AddressType = &Address_Ipv4{x}
		return true, err
	case 2: // address_type.ipv6
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.AddressType = &Address_Ipv6{x}
		return true, err
	default:
		return false, nil
	}
}

func _Address_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Address)
	// address_type
	switch x := m.AddressType.(type) {
	case *Address_Ipv4:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Ipv4)))
		n += len(x.Ipv4)
	case *Address_Ipv6:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Ipv6)))
		n += len(x.Ipv6)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// AddressRange is an IPv4 or IPv6 CIDR range.
type AddressRange struct {
	// Types that are valid to be assigned to AddressType:
	//	*AddressRange_Ipv4Range
	//	*AddressRange_Ipv6Range
	AddressType isAddressRange_AddressType `protobuf_oneof:"address_type"`
}

func (m *AddressRange) Reset()                    { *m = AddressRange{} }
func (m *AddressRange) String() string            { return proto.CompactTextString(m) }
func (*AddressRange) ProtoMessage()               {}
func (*AddressRange) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

type isAddressRange_AddressType interface {
	isAddressRange_AddressType()
}

type AddressRange_Ipv4Range struct {
	Ipv4Range string `protobuf:"bytes,1,opt,name=ipv4_range,json=ipv4Range,oneof"`
}
type AddressRange_Ipv6Range struct {
	Ipv6Range string `protobuf:"bytes,2,opt,name=ipv6_range,json=ipv6Range,oneof"`
}

func (*AddressRange_Ipv4Range) isAddressRange_AddressType() {}
func (*AddressRange_Ipv6Range) isAddressRange_AddressType() {}

func (m *AddressRange) GetAddressType() isAddressRange_AddressType {
	if m != nil {
		return m.AddressType
	}
	return nil
}

func (m *AddressRange) GetIpv4Range() string {
	if x, ok := m.GetAddressType().(*AddressRange_Ipv4Range); ok {
		return x.Ipv4Range
	}
	return ""
}

func (m *AddressRange) GetIpv6Range() string {
	if x, ok := m.GetAddressType().(*AddressRange_Ipv6Range); ok {
		return x.Ipv6Range
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AddressRange) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AddressRange_OneofMarshaler, _AddressRange_OneofUnmarshaler, _AddressRange_OneofSizer, []interface{}{
		(*AddressRange_Ipv4Range)(nil),
		(*AddressRange_Ipv6Range)(nil),
	}
}

func _AddressRange_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AddressRange)
	// address_type
	switch x := m.AddressType.(type) {
	case *AddressRange_Ipv4Range:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Ipv4Range)
	case *AddressRange_Ipv6Range:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Ipv6Range)
	case nil:
	default:
		return fmt.Errorf("AddressRange.AddressType has unexpected type %T", x)
	}
	return nil
}

func _AddressRange_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AddressRange)
	switch tag {
	case 1: // address_type.ipv4_range
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.AddressType = &AddressRange_Ipv4Range{x}
		return true, err
	case 2: // address_type.ipv6_range
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.AddressType = &AddressRange_Ipv6Range{x}
		return true, err
	default:
		return false, nil
	}
}

func _AddressRange_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AddressRange)
	// address_type
	switch x := m.AddressType.(type) {
	case *AddressRange_Ipv4Range:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Ipv4Range)))
		n += len(x.Ipv4Range)
	case *AddressRange_Ipv6Range:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Ipv6Range)))
		n += len(x.Ipv6Range)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// FixedAddress described an address that is determined either by the host
// environment or by the asset manifest.
type FixedAddress struct {
	// Types that are valid to be assigned to AddressType:
	//	*FixedAddress_Address
	//	*FixedAddress_AddressPool
	AddressType isFixedAddress_AddressType `protobuf_oneof:"address_type"`
}

func (m *FixedAddress) Reset()                    { *m = FixedAddress{} }
func (m *FixedAddress) String() string            { return proto.CompactTextString(m) }
func (*FixedAddress) ProtoMessage()               {}
func (*FixedAddress) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

type isFixedAddress_AddressType interface {
	isFixedAddress_AddressType()
}

type FixedAddress_Address struct {
	Address *Address `protobuf:"bytes,1,opt,name=address,oneof"`
}
type FixedAddress_AddressPool struct {
	AddressPool string `protobuf:"bytes,2,opt,name=address_pool,json=addressPool,oneof"`
}

func (*FixedAddress_Address) isFixedAddress_AddressType()     {}
func (*FixedAddress_AddressPool) isFixedAddress_AddressType() {}

func (m *FixedAddress) GetAddressType() isFixedAddress_AddressType {
	if m != nil {
		return m.AddressType
	}
	return nil
}

func (m *FixedAddress) GetAddress() *Address {
	if x, ok := m.GetAddressType().(*FixedAddress_Address); ok {
		return x.Address
	}
	return nil
}

func (m *FixedAddress) GetAddressPool() string {
	if x, ok := m.GetAddressType().(*FixedAddress_AddressPool); ok {
		return x.AddressPool
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*FixedAddress) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _FixedAddress_OneofMarshaler, _FixedAddress_OneofUnmarshaler, _FixedAddress_OneofSizer, []interface{}{
		(*FixedAddress_Address)(nil),
		(*FixedAddress_AddressPool)(nil),
	}
}

func _FixedAddress_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*FixedAddress)
	// address_type
	switch x := m.AddressType.(type) {
	case *FixedAddress_Address:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Address); err != nil {
			return err
		}
	case *FixedAddress_AddressPool:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.AddressPool)
	case nil:
	default:
		return fmt.Errorf("FixedAddress.AddressType has unexpected type %T", x)
	}
	return nil
}

func _FixedAddress_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*FixedAddress)
	switch tag {
	case 1: // address_type.address
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Address)
		err := b.DecodeMessage(msg)
		m.AddressType = &FixedAddress_Address{msg}
		return true, err
	case 2: // address_type.address_pool
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.AddressType = &FixedAddress_AddressPool{x}
		return true, err
	default:
		return false, nil
	}
}

func _FixedAddress_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*FixedAddress)
	// address_type
	switch x := m.AddressType.(type) {
	case *FixedAddress_Address:
		s := proto.Size(x.Address)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *FixedAddress_AddressPool:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.AddressPool)))
		n += len(x.AddressPool)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Network)(nil), "asset.Network")
	proto.RegisterType((*NetworkPeer)(nil), "asset.NetworkPeer")
	proto.RegisterType((*Address)(nil), "asset.Address")
	proto.RegisterType((*AddressRange)(nil), "asset.AddressRange")
	proto.RegisterType((*FixedAddress)(nil), "asset.FixedAddress")
	proto.RegisterEnum("asset.Protocol", Protocol_name, Protocol_value)
}

func init() { proto.RegisterFile("schema/asset/network.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0xdf, 0x4b, 0x83, 0x60,
	0x14, 0x9d, 0x6b, 0xcb, 0x79, 0x5d, 0x63, 0x7c, 0xf5, 0x20, 0xbd, 0x34, 0xec, 0xa1, 0x31, 0x42,
	0x69, 0x0b, 0xe9, 0xb5, 0x41, 0x31, 0x08, 0x4c, 0x6c, 0x31, 0xe8, 0x65, 0x99, 0xbb, 0x38, 0x49,
	0xbd, 0xf2, 0xe9, 0xfa, 0xf1, 0xdf, 0x87, 0x9f, 0x5a, 0x56, 0x6f, 0xf7, 0x9e, 0x7b, 0xee, 0x39,
	0x47, 0xbf, 0x0b, 0xc7, 0x99, 0xbf, 0xc5, 0xd8, 0x33, 0xbd, 0x2c, 0xc3, 0xdc, 0x4c, 0x30, 0x7f,
	0x27, 0xfe, 0x6a, 0xa4, 0x9c, 0x72, 0x62, 0x5d, 0x01, 0xea, 0x2b, 0x90, 0xed, 0x12, 0x67, 0x0c,
	0x3a, 0x89, 0x17, 0xa3, 0x26, 0x8d, 0xa4, 0xb1, 0xe2, 0x8a, 0x9a, 0x5d, 0xc1, 0x81, 0xb7, 0xd9,
	0x70, 0xcc, 0xb2, 0x35, 0xf7, 0x92, 0x00, 0xb5, 0xf6, 0x48, 0x1a, 0xab, 0xd3, 0x43, 0x43, 0x6c,
	0x1b, 0xd7, 0xe5, 0xcc, 0x2d, 0x46, 0x6e, 0xdf, 0x6b, 0x74, 0xfa, 0x19, 0xa8, 0x95, 0xb0, 0x83,
	0xc8, 0x99, 0x06, 0x72, 0xe5, 0xaf, 0x49, 0xa3, 0xbd, 0xb1, 0xe2, 0xd6, 0xad, 0x7e, 0x03, 0x72,
	0x25, 0xc3, 0x8e, 0xa0, 0x13, 0xa6, 0x6f, 0x97, 0x65, 0x82, 0x45, 0xcb, 0x15, 0x5d, 0x85, 0x5a,
	0xc2, 0xba, 0x46, 0xad, 0xf9, 0x00, 0x6a, 0xbf, 0x75, 0xfe, 0x99, 0xa2, 0xfe, 0x0c, 0xfd, 0x66,
	0x1a, 0x76, 0x02, 0x50, 0x6c, 0x57, 0xb1, 0x6b, 0x45, 0xa5, 0xc0, 0x9a, 0x04, 0xab, 0xf1, 0x5d,
	0x35, 0xc1, 0x12, 0x84, 0x7f, 0x0e, 0x04, 0xfd, 0xdb, 0xf0, 0x03, 0x37, 0x75, 0xda, 0x09, 0xc8,
	0xd5, 0x5c, 0xc8, 0xab, 0xd3, 0xc1, 0xef, 0xbf, 0xb2, 0x68, 0xb9, 0x35, 0x81, 0x9d, 0xfe, 0x68,
	0xa5, 0x44, 0xd1, 0xb7, 0x9d, 0x5a, 0xa1, 0x0e, 0x51, 0xf4, 0xd7, 0x70, 0x72, 0x0e, 0x3d, 0xa7,
	0x78, 0x2b, 0x9f, 0x22, 0xa6, 0x82, 0xfc, 0x68, 0xdf, 0xd9, 0xf7, 0x2b, 0x7b, 0xd8, 0x62, 0x3d,
	0xe8, 0x2c, 0x96, 0x4b, 0x67, 0x28, 0x31, 0x05, 0xba, 0x45, 0xf5, 0x30, 0x6c, 0xcf, 0x67, 0x4f,
	0x17, 0xfe, 0x96, 0x53, 0x1c, 0xee, 0x62, 0x23, 0x20, 0x0a, 0x22, 0xcc, 0x68, 0xc7, 0x7d, 0x34,
	0x7c, 0x8a, 0x4d, 0x4c, 0x72, 0xe4, 0x29, 0x0f, 0x33, 0x34, 0x7d, 0x8c, 0xcc, 0x80, 0xca, 0x9b,
	0x78, 0xd9, 0x17, 0xc7, 0x30, 0xfb, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x62, 0x6d, 0x98, 0x2a,
	0x02, 0x00, 0x00,
}
