// Code generated by protoc-gen-go.
// source: schema/asset/network.proto
// DO NOT EDIT!

package asset

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "chromium.googlesource.com/enterprise/cel/go/common"

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
	Ip string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *Address) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

// AddressRange is an IPv4 or IPv6 CIDR range.
type AddressRange struct {
	Cidr string `protobuf:"bytes,1,opt,name=cidr" json:"cidr,omitempty"`
}

func (m *AddressRange) Reset()                    { *m = AddressRange{} }
func (m *AddressRange) String() string            { return proto.CompactTextString(m) }
func (*AddressRange) ProtoMessage()               {}
func (*AddressRange) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *AddressRange) GetCidr() string {
	if m != nil {
		return m.Cidr
	}
	return ""
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
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xcf, 0x8f, 0x9a, 0x50,
	0x10, 0xc7, 0x85, 0xaa, 0xc8, 0xa0, 0xd6, 0x3e, 0x93, 0x96, 0xd8, 0x8b, 0xe5, 0x64, 0x4c, 0x0b,
	0xa9, 0x5e, 0x7a, 0xf0, 0x52, 0x0e, 0x8d, 0x49, 0x13, 0x4a, 0xa8, 0x8d, 0x49, 0x2f, 0x86, 0x3e,
	0x26, 0x48, 0x0a, 0x0c, 0x79, 0x0f, 0xb3, 0xbb, 0x57, 0x4f, 0xfb, 0x0f, 0xf9, 0xff, 0x6d, 0xe4,
	0x47, 0xd6, 0xdd, 0xdb, 0xcc, 0x7c, 0x3f, 0xf3, 0x65, 0x66, 0x78, 0x30, 0x93, 0xfc, 0x88, 0x59,
	0xe8, 0x84, 0x52, 0x62, 0xe9, 0xe4, 0x58, 0xde, 0x91, 0xf8, 0x6f, 0x17, 0x82, 0x4a, 0x62, 0xbd,
	0xaa, 0x38, 0xfb, 0xd8, 0x20, 0x9c, 0xb2, 0x8c, 0x72, 0x87, 0x8a, 0x32, 0xa1, 0x5c, 0xd6, 0x8c,
	0xb5, 0x07, 0xcd, 0xab, 0x9b, 0x18, 0x83, 0x6e, 0x1e, 0x66, 0x68, 0x2a, 0x73, 0x65, 0xa1, 0x07,
	0x55, 0xcc, 0xbe, 0xc1, 0x28, 0x8c, 0x22, 0x81, 0x52, 0x1e, 0x44, 0x98, 0xc7, 0x68, 0xaa, 0x73,
	0x65, 0x61, 0xac, 0xa6, 0x76, 0x65, 0x6d, 0x7f, 0xaf, 0xb5, 0xe0, 0x2a, 0x05, 0xc3, 0xf0, 0x26,
	0xb3, 0x36, 0x60, 0x34, 0xc6, 0x3e, 0xa2, 0x60, 0x5f, 0x40, 0x6b, 0x86, 0x33, 0x95, 0xf9, 0x9b,
	0x85, 0xee, 0x4e, 0xcf, 0x17, 0xf3, 0x2d, 0x1b, 0xd5, 0x3e, 0x8d, 0x14, 0xb4, 0x8c, 0xf5, 0x09,
	0xb4, 0xc6, 0x9b, 0xbd, 0x07, 0x35, 0x29, 0xea, 0xa1, 0xdc, 0xfe, 0xf9, 0x62, 0xaa, 0x03, 0x25,
	0x50, 0x93, 0xc2, 0x5a, 0xc2, 0xf0, 0xf6, 0xf3, 0x6c, 0x06, 0x5d, 0x9e, 0x44, 0xe2, 0x15, 0x59,
	0xd5, 0xac, 0x47, 0x05, 0x86, 0x3f, 0x92, 0x7b, 0x8c, 0x5a, 0xd3, 0x25, 0x68, 0xcd, 0xb4, 0x15,
	0x6f, 0xac, 0xc6, 0x2f, 0x37, 0xda, 0x76, 0x82, 0x16, 0x60, 0x1b, 0x68, 0x37, 0x3b, 0x14, 0x44,
	0x69, 0x75, 0x02, 0xdd, 0xfd, 0x70, 0xbe, 0x98, 0x53, 0xf6, 0xee, 0x48, 0xb2, 0xb4, 0x6f, 0xe5,
	0x6d, 0x27, 0x30, 0x9a, 0xdc, 0x27, 0x4a, 0xdd, 0xf1, 0x73, 0x77, 0xf9, 0x50, 0xe0, 0xf2, 0x33,
	0x0c, 0xfc, 0xeb, 0xe5, 0x39, 0xa5, 0xcc, 0x00, 0xed, 0x8f, 0xf7, 0xd3, 0xfb, 0xb5, 0xf7, 0x26,
	0x1d, 0x36, 0x80, 0xee, 0x76, 0xb7, 0xf3, 0x27, 0x0a, 0xd3, 0xa1, 0x77, 0x8d, 0x7e, 0x4f, 0x54,
	0x77, 0xfd, 0xf7, 0x2b, 0x3f, 0x0a, 0xca, 0x92, 0x53, 0x66, 0xc7, 0x44, 0x71, 0x8a, 0x92, 0x4e,
	0x82, 0xa3, 0xcd, 0x29, 0x73, 0x30, 0x2f, 0x51, 0x14, 0x22, 0x91, 0xe8, 0x70, 0x4c, 0x9d, 0x98,
	0xea, 0x57, 0xf0, 0xaf, 0x5f, 0xfd, 0xda, 0xf5, 0x53, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x7d,
	0xda, 0x1f, 0x1c, 0x02, 0x00, 0x00,
}
