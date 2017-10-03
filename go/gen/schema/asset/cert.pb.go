// Code generated by protoc-gen-go.
// source: schema/asset/cert.proto
// DO NOT EDIT!

/*
Package asset is a generated protocol buffer package.

It is generated from these files:
	schema/asset/cert.proto

It has these top-level messages:
	Certificate
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

type Certificate struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Certificate) Reset()                    { *m = Certificate{} }
func (m *Certificate) String() string            { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()               {}
func (*Certificate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Certificate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Certificate)(nil), "cert.Certificate")
}

func init() { proto.RegisterFile("schema/asset/cert.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 96 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x4f, 0x2c, 0x2e, 0x4e, 0x2d, 0xd1, 0x4f, 0x4e, 0x2d, 0x2a, 0xd1, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x14, 0xb9, 0xb8, 0x9d, 0x53, 0x8b, 0x4a, 0x32,
	0xd3, 0x32, 0x93, 0x13, 0x4b, 0x52, 0x85, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x27, 0xf6, 0x28, 0x56, 0xb0, 0xe6, 0x24, 0x36, 0xb0,
	0x46, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x30, 0xf1, 0x43, 0x0d, 0x53, 0x00, 0x00, 0x00,
}
