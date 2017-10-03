// Code generated by protoc-gen-go.
// source: schema/asset/dns.proto
// DO NOT EDIT!

/*
Package asset is a generated protocol buffer package.

It is generated from these files:
	schema/asset/dns.proto

It has these top-level messages:
	Zone
	Record
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

// dns.Zone describes a DNS zone. Conceptually it encompasses the same
// information included in a DNS zone file as described in
// https://en.wikipedia.org/wiki/Zone_file
type Zone struct {
	Origin string    `protobuf:"bytes,1,opt,name=origin" json:"origin,omitempty"`
	Record []*Record `protobuf:"bytes,2,rep,name=record" json:"record,omitempty"`
}

func (m *Zone) Reset()                    { *m = Zone{} }
func (m *Zone) String() string            { return proto.CompactTextString(m) }
func (*Zone) ProtoMessage()               {}
func (*Zone) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Zone) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Zone) GetRecord() []*Record {
	if m != nil {
		return m.Record
	}
	return nil
}

// dns.Record describes a single DNS record in a Zone.
type Record struct {
	// Name on record.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// TTL in seconds.
	Ttl int32 `protobuf:"varint,2,opt,name=ttl" json:"ttl,omitempty"`
	// Must be omitted or is always IN.
	RecordClass string `protobuf:"bytes,3,opt,name=record_class,json=recordClass" json:"record_class,omitempty"`
	// Type of record. E.g. A, AAAA, NS, MX, ...
	RecordType string `protobuf:"bytes,4,opt,name=record_type,json=recordType" json:"record_type,omitempty"`
	// Priority value.
	Priority int32 `protobuf:"varint,5,opt,name=priority" json:"priority,omitempty"`
	// Answer section.
	Answer string `protobuf:"bytes,6,opt,name=answer" json:"answer,omitempty"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Record) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Record) GetTtl() int32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *Record) GetRecordClass() string {
	if m != nil {
		return m.RecordClass
	}
	return ""
}

func (m *Record) GetRecordType() string {
	if m != nil {
		return m.RecordType
	}
	return ""
}

func (m *Record) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Record) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

func init() {
	proto.RegisterType((*Zone)(nil), "dns.Zone")
	proto.RegisterType((*Record)(nil), "dns.Record")
}

func init() { proto.RegisterFile("schema/asset/dns.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0xd9, 0xee, 0x6e, 0xd4, 0x59, 0x0f, 0x32, 0x87, 0x12, 0xbc, 0xb8, 0xd6, 0xcb, 0x9e,
	0x5a, 0xd0, 0x37, 0xb0, 0x6f, 0x10, 0x3c, 0xf5, 0x52, 0xe2, 0xee, 0xa0, 0x81, 0x36, 0x09, 0x99,
	0x80, 0xec, 0x0b, 0xf9, 0x9c, 0x92, 0x3f, 0x78, 0x9b, 0xdf, 0x2f, 0x93, 0x0f, 0xe6, 0x83, 0x2d,
	0xcf, 0xdf, 0x74, 0xd5, 0x07, 0xcd, 0x4c, 0xf1, 0xb0, 0x58, 0xde, 0xfb, 0xe0, 0xa2, 0xc3, 0x76,
	0xb1, 0xbc, 0x3b, 0x42, 0x77, 0x72, 0x96, 0x70, 0x0b, 0xc2, 0x05, 0xf3, 0x65, 0xac, 0x6c, 0xc6,
	0x66, 0xba, 0x53, 0x95, 0xf0, 0x05, 0x44, 0xa0, 0xd9, 0x85, 0x45, 0x6e, 0xc6, 0x76, 0x1a, 0x5e,
	0x87, 0x7d, 0x0a, 0x50, 0x59, 0xa9, 0xfa, 0xb4, 0xfb, 0x6d, 0x40, 0x14, 0x85, 0x08, 0x9d, 0xd5,
	0x57, 0xaa, 0x29, 0x79, 0xc6, 0x07, 0x68, 0x63, 0xbc, 0xc8, 0xcd, 0xd8, 0x4c, 0xbd, 0x4a, 0x23,
	0x3e, 0xc3, 0x7d, 0xf9, 0x7a, 0x9e, 0x2f, 0x9a, 0x59, 0xb6, 0x79, 0x7b, 0x28, 0xee, 0x98, 0x14,
	0x3e, 0x41, 0xc5, 0x73, 0x5c, 0x3d, 0xc9, 0x2e, 0x6f, 0x40, 0x51, 0x1f, 0xab, 0x27, 0x7c, 0x84,
	0x5b, 0x1f, 0x8c, 0x0b, 0x26, 0xae, 0xb2, 0xcf, 0xd1, 0xff, 0x9c, 0xae, 0xd1, 0x96, 0x7f, 0x28,
	0x48, 0x51, 0xae, 0x29, 0xf4, 0x7e, 0x73, 0xea, 0x73, 0x0b, 0x9f, 0x22, 0x57, 0xf0, 0xf6, 0x17,
	0x00, 0x00, 0xff, 0xff, 0xe8, 0xe2, 0x97, 0x67, 0x1c, 0x01, 0x00, 0x00,
}
