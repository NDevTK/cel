// Code generated by protoc-gen-go.
// source: go/common/testdata/testmsgs.proto
// DO NOT EDIT!

package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TestHasBadField struct {
	Name  string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Field *TestBadProto `protobuf:"bytes,2,opt,name=field" json:"field,omitempty"`
}

func (m *TestHasBadField) Reset()                    { *m = TestHasBadField{} }
func (m *TestHasBadField) String() string            { return proto.CompactTextString(m) }
func (*TestHasBadField) ProtoMessage()               {}
func (*TestHasBadField) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *TestHasBadField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TestHasBadField) GetField() *TestBadProto {
	if m != nil {
		return m.Field
	}
	return nil
}

type TestHasGoodField struct {
	Name  string         `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Field *TestGoodProto `protobuf:"bytes,2,opt,name=field" json:"field,omitempty"`
}

func (m *TestHasGoodField) Reset()                    { *m = TestHasGoodField{} }
func (m *TestHasGoodField) String() string            { return proto.CompactTextString(m) }
func (*TestHasGoodField) ProtoMessage()               {}
func (*TestHasGoodField) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *TestHasGoodField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TestHasGoodField) GetField() *TestGoodProto {
	if m != nil {
		return m.Field
	}
	return nil
}

type TestHasBadSlice struct {
	Name  string          `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Field []*TestBadProto `protobuf:"bytes,2,rep,name=field" json:"field,omitempty"`
}

func (m *TestHasBadSlice) Reset()                    { *m = TestHasBadSlice{} }
func (m *TestHasBadSlice) String() string            { return proto.CompactTextString(m) }
func (*TestHasBadSlice) ProtoMessage()               {}
func (*TestHasBadSlice) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *TestHasBadSlice) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TestHasBadSlice) GetField() []*TestBadProto {
	if m != nil {
		return m.Field
	}
	return nil
}

type TestHasGoodSlice struct {
	Name  string           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Field []*TestGoodProto `protobuf:"bytes,2,rep,name=field" json:"field,omitempty"`
}

func (m *TestHasGoodSlice) Reset()                    { *m = TestHasGoodSlice{} }
func (m *TestHasGoodSlice) String() string            { return proto.CompactTextString(m) }
func (*TestHasGoodSlice) ProtoMessage()               {}
func (*TestHasGoodSlice) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *TestHasGoodSlice) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TestHasGoodSlice) GetField() []*TestGoodProto {
	if m != nil {
		return m.Field
	}
	return nil
}

type TestGoodOneOf struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Types that are valid to be assigned to Opt:
	//	*TestGoodOneOf_Field
	Opt isTestGoodOneOf_Opt `protobuf_oneof:"opt"`
}

func (m *TestGoodOneOf) Reset()                    { *m = TestGoodOneOf{} }
func (m *TestGoodOneOf) String() string            { return proto.CompactTextString(m) }
func (*TestGoodOneOf) ProtoMessage()               {}
func (*TestGoodOneOf) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

type isTestGoodOneOf_Opt interface {
	isTestGoodOneOf_Opt()
}

type TestGoodOneOf_Field struct {
	Field *TestGoodProto `protobuf:"bytes,2,opt,name=field,oneof"`
}

func (*TestGoodOneOf_Field) isTestGoodOneOf_Opt() {}

func (m *TestGoodOneOf) GetOpt() isTestGoodOneOf_Opt {
	if m != nil {
		return m.Opt
	}
	return nil
}

func (m *TestGoodOneOf) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TestGoodOneOf) GetField() *TestGoodProto {
	if x, ok := m.GetOpt().(*TestGoodOneOf_Field); ok {
		return x.Field
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TestGoodOneOf) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TestGoodOneOf_OneofMarshaler, _TestGoodOneOf_OneofUnmarshaler, _TestGoodOneOf_OneofSizer, []interface{}{
		(*TestGoodOneOf_Field)(nil),
	}
}

func _TestGoodOneOf_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TestGoodOneOf)
	// opt
	switch x := m.Opt.(type) {
	case *TestGoodOneOf_Field:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Field); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TestGoodOneOf.Opt has unexpected type %T", x)
	}
	return nil
}

func _TestGoodOneOf_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TestGoodOneOf)
	switch tag {
	case 2: // opt.field
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TestGoodProto)
		err := b.DecodeMessage(msg)
		m.Opt = &TestGoodOneOf_Field{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TestGoodOneOf_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TestGoodOneOf)
	// opt
	switch x := m.Opt.(type) {
	case *TestGoodOneOf_Field:
		s := proto.Size(x.Field)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type TestBadOneOf struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Types that are valid to be assigned to Opt:
	//	*TestBadOneOf_Field
	Opt isTestBadOneOf_Opt `protobuf_oneof:"opt"`
}

func (m *TestBadOneOf) Reset()                    { *m = TestBadOneOf{} }
func (m *TestBadOneOf) String() string            { return proto.CompactTextString(m) }
func (*TestBadOneOf) ProtoMessage()               {}
func (*TestBadOneOf) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

type isTestBadOneOf_Opt interface {
	isTestBadOneOf_Opt()
}

type TestBadOneOf_Field struct {
	Field *TestBadProto `protobuf:"bytes,2,opt,name=field,oneof"`
}

func (*TestBadOneOf_Field) isTestBadOneOf_Opt() {}

func (m *TestBadOneOf) GetOpt() isTestBadOneOf_Opt {
	if m != nil {
		return m.Opt
	}
	return nil
}

func (m *TestBadOneOf) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TestBadOneOf) GetField() *TestBadProto {
	if x, ok := m.GetOpt().(*TestBadOneOf_Field); ok {
		return x.Field
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TestBadOneOf) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TestBadOneOf_OneofMarshaler, _TestBadOneOf_OneofUnmarshaler, _TestBadOneOf_OneofSizer, []interface{}{
		(*TestBadOneOf_Field)(nil),
	}
}

func _TestBadOneOf_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TestBadOneOf)
	// opt
	switch x := m.Opt.(type) {
	case *TestBadOneOf_Field:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Field); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TestBadOneOf.Opt has unexpected type %T", x)
	}
	return nil
}

func _TestBadOneOf_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TestBadOneOf)
	switch tag {
	case 2: // opt.field
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TestBadProto)
		err := b.DecodeMessage(msg)
		m.Opt = &TestBadOneOf_Field{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TestBadOneOf_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TestBadOneOf)
	// opt
	switch x := m.Opt.(type) {
	case *TestBadOneOf_Field:
		s := proto.Size(x.Field)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type TestBadValidateArgs struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *TestBadValidateArgs) Reset()                    { *m = TestBadValidateArgs{} }
func (m *TestBadValidateArgs) String() string            { return proto.CompactTextString(m) }
func (*TestBadValidateArgs) ProtoMessage()               {}
func (*TestBadValidateArgs) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *TestBadValidateArgs) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TestBadReturnType struct {
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *TestBadReturnType) Reset()                    { *m = TestBadReturnType{} }
func (m *TestBadReturnType) String() string            { return proto.CompactTextString(m) }
func (*TestBadReturnType) ProtoMessage()               {}
func (*TestBadReturnType) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *TestBadReturnType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TestBadProto struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *TestBadProto) Reset()                    { *m = TestBadProto{} }
func (m *TestBadProto) String() string            { return proto.CompactTextString(m) }
func (*TestBadProto) ProtoMessage()               {}
func (*TestBadProto) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *TestBadProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TestGoodProto struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *TestGoodProto) Reset()                    { *m = TestGoodProto{} }
func (m *TestGoodProto) String() string            { return proto.CompactTextString(m) }
func (*TestGoodProto) ProtoMessage()               {}
func (*TestGoodProto) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *TestGoodProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TestMessageWithOptions struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Key         string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Label       string `protobuf:"bytes,3,opt,name=label" json:"label,omitempty"`
	OptionalKey string `protobuf:"bytes,4,opt,name=optional_key,json=optionalKey" json:"optional_key,omitempty"`
	Fqdn        string `protobuf:"bytes,5,opt,name=fqdn" json:"fqdn,omitempty"`
	Reqd        string `protobuf:"bytes,6,opt,name=reqd" json:"reqd,omitempty"`
}

func (m *TestMessageWithOptions) Reset()                    { *m = TestMessageWithOptions{} }
func (m *TestMessageWithOptions) String() string            { return proto.CompactTextString(m) }
func (*TestMessageWithOptions) ProtoMessage()               {}
func (*TestMessageWithOptions) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func (m *TestMessageWithOptions) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TestMessageWithOptions) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *TestMessageWithOptions) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *TestMessageWithOptions) GetOptionalKey() string {
	if m != nil {
		return m.OptionalKey
	}
	return ""
}

func (m *TestMessageWithOptions) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *TestMessageWithOptions) GetReqd() string {
	if m != nil {
		return m.Reqd
	}
	return ""
}

type TestBadMessageWithOptions struct {
	Name int32 `protobuf:"varint,1,opt,name=name" json:"name,omitempty"`
}

func (m *TestBadMessageWithOptions) Reset()                    { *m = TestBadMessageWithOptions{} }
func (m *TestBadMessageWithOptions) String() string            { return proto.CompactTextString(m) }
func (*TestBadMessageWithOptions) ProtoMessage()               {}
func (*TestBadMessageWithOptions) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

func (m *TestBadMessageWithOptions) GetName() int32 {
	if m != nil {
		return m.Name
	}
	return 0
}

type TestFileRefProto struct {
	Ref *FileReference `protobuf:"bytes,1,opt,name=ref" json:"ref,omitempty"`
}

func (m *TestFileRefProto) Reset()                    { *m = TestFileRefProto{} }
func (m *TestFileRefProto) String() string            { return proto.CompactTextString(m) }
func (*TestFileRefProto) ProtoMessage()               {}
func (*TestFileRefProto) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{12} }

func (m *TestFileRefProto) GetRef() *FileReference {
	if m != nil {
		return m.Ref
	}
	return nil
}

func init() {
	proto.RegisterType((*TestHasBadField)(nil), "common.TestHasBadField")
	proto.RegisterType((*TestHasGoodField)(nil), "common.TestHasGoodField")
	proto.RegisterType((*TestHasBadSlice)(nil), "common.TestHasBadSlice")
	proto.RegisterType((*TestHasGoodSlice)(nil), "common.TestHasGoodSlice")
	proto.RegisterType((*TestGoodOneOf)(nil), "common.TestGoodOneOf")
	proto.RegisterType((*TestBadOneOf)(nil), "common.TestBadOneOf")
	proto.RegisterType((*TestBadValidateArgs)(nil), "common.TestBadValidateArgs")
	proto.RegisterType((*TestBadReturnType)(nil), "common.TestBadReturnType")
	proto.RegisterType((*TestBadProto)(nil), "common.TestBadProto")
	proto.RegisterType((*TestGoodProto)(nil), "common.TestGoodProto")
	proto.RegisterType((*TestMessageWithOptions)(nil), "common.TestMessageWithOptions")
	proto.RegisterType((*TestBadMessageWithOptions)(nil), "common.TestBadMessageWithOptions")
	proto.RegisterType((*TestFileRefProto)(nil), "common.TestFileRefProto")
}

func init() { proto.RegisterFile("go/common/testdata/testmsgs.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xd7, 0xfc, 0xa9, 0x8a, 0x3b, 0x34, 0xc8, 0x00, 0x85, 0xc2, 0x61, 0x04, 0xa1, 0x0d,
	0xca, 0x1a, 0x69, 0x20, 0x71, 0xe0, 0x44, 0x0e, 0x63, 0x12, 0x42, 0x05, 0x6f, 0x62, 0x82, 0x0b,
	0xf2, 0x92, 0x37, 0xa9, 0x35, 0x27, 0xce, 0x6c, 0xf7, 0xd0, 0x6b, 0xbf, 0x53, 0xbf, 0x0a, 0x9f,
	0x07, 0x39, 0x4e, 0xbb, 0x14, 0xd2, 0x31, 0x6e, 0xd6, 0xfb, 0x3c, 0xfe, 0xf5, 0x79, 0xdf, 0xbe,
	0x31, 0x7a, 0x96, 0xf1, 0x30, 0xe6, 0x79, 0xce, 0x8b, 0x50, 0x81, 0x54, 0x09, 0x51, 0xa4, 0x3a,
	0xe4, 0x32, 0x93, 0xa3, 0x52, 0x70, 0xc5, 0xbd, 0xae, 0xd1, 0x07, 0x4f, 0x64, 0x3c, 0x81, 0x9c,
	0x2c, 0xed, 0x29, 0x65, 0x20, 0x20, 0x35, 0xa6, 0x3f, 0x45, 0x5e, 0x2a, 0xca, 0x8b, 0x9a, 0x10,
	0x7c, 0x45, 0x3b, 0x67, 0x20, 0xd5, 0x09, 0x91, 0x11, 0x49, 0x8e, 0x29, 0xb0, 0xc4, 0xf3, 0x90,
	0x53, 0x90, 0x1c, 0xfc, 0xce, 0x5e, 0xe7, 0xe0, 0x0e, 0xae, 0xce, 0xde, 0x2b, 0xe4, 0xa6, 0x5a,
	0xf4, 0xad, 0xbd, 0xce, 0x41, 0xff, 0xe8, 0xc1, 0xc8, 0xc0, 0x46, 0xfa, 0x6e, 0x44, 0x92, 0x2f,
	0x9a, 0x85, 0x8d, 0x25, 0x38, 0x45, 0xf7, 0x6a, 0xe4, 0x47, 0xce, 0x6f, 0x60, 0x0e, 0xd7, 0x99,
	0x0f, 0x9b, 0x4c, 0x7d, 0x73, 0x0d, 0xba, 0x96, 0xf3, 0x94, 0xd1, 0x18, 0xfe, 0x95, 0xd3, 0xfe,
	0xbf, 0x9c, 0x9b, 0x99, 0xc3, 0x75, 0xe6, 0xcd, 0x39, 0xbf, 0xa3, 0xbb, 0xcb, 0xfa, 0xb8, 0x80,
	0x71, 0xda, 0x4a, 0x3c, 0xbc, 0x4d, 0xe7, 0x27, 0x5b, 0x35, 0x33, 0x72, 0x91, 0xcd, 0x4b, 0x15,
	0x9c, 0xa3, 0xed, 0xba, 0x8d, 0xcd, 0xe4, 0xd7, 0xb7, 0xf8, 0x9f, 0xfe, 0x02, 0xbf, 0x44, 0xbb,
	0xb5, 0xfe, 0x8d, 0x30, 0x9a, 0x10, 0x05, 0x1f, 0x44, 0x26, 0xdb, 0xf8, 0xc1, 0x3e, 0xba, 0x5f,
	0x5b, 0x31, 0xa8, 0xa9, 0x28, 0xce, 0x66, 0xe5, 0xf5, 0xd0, 0xac, 0x86, 0x31, 0x58, 0x85, 0xad,
	0x7e, 0xb3, 0x15, 0xf6, 0xfc, 0x7a, 0x56, 0x9b, 0x4d, 0xbf, 0x3a, 0xe8, 0x91, 0x76, 0x7d, 0x06,
	0x29, 0x49, 0x06, 0xe7, 0x54, 0x4d, 0xc6, 0x66, 0x83, 0x5b, 0x07, 0xf0, 0x02, 0xd9, 0x97, 0x30,
	0x33, 0x51, 0xa2, 0xdd, 0xf9, 0xc2, 0xdf, 0xe9, 0x59, 0x5e, 0x3f, 0xe5, 0x02, 0x68, 0x56, 0x1c,
	0x5e, 0xc2, 0x0c, 0x6b, 0xdd, 0x7b, 0x8a, 0x5c, 0x46, 0x2e, 0x80, 0xf9, 0x76, 0x65, 0xec, 0xce,
	0x17, 0xbe, 0xd5, 0x73, 0xb0, 0x29, 0x7a, 0x43, 0xb4, 0x6d, 0xbe, 0x12, 0xc2, 0x7e, 0x6a, 0x9a,
	0x53, 0x99, 0x7a, 0xf3, 0x85, 0xef, 0xf4, 0x2c, 0xbf, 0x83, 0xfb, 0x4b, 0xf5, 0x13, 0xcc, 0xbc,
	0x01, 0x72, 0xd2, 0xab, 0xa4, 0xf0, 0xdd, 0x06, 0xc9, 0xc5, 0x55, 0x4d, 0x6b, 0x02, 0xae, 0x12,
	0xbf, 0xdb, 0xd0, 0xb6, 0x70, 0x55, 0x0b, 0xde, 0xa1, 0xc7, 0xf5, 0x84, 0x5a, 0x5a, 0x1b, 0x34,
	0x5a, 0x73, 0x57, 0xf1, 0xcc, 0x44, 0xde, 0x9b, 0xbd, 0x3d, 0xa6, 0x0c, 0x30, 0xa4, 0x66, 0x72,
	0xfb, 0xc8, 0x16, 0x90, 0x56, 0xf6, 0xc6, 0x3e, 0xd5, 0x16, 0x10, 0x50, 0xc4, 0x80, 0xb5, 0x23,
	0x7a, 0xfb, 0xe3, 0x28, 0x9e, 0x08, 0x9e, 0xd3, 0x69, 0x3e, 0xca, 0x38, 0xcf, 0x18, 0x48, 0x3e,
	0x15, 0x31, 0xe8, 0x2b, 0x21, 0x14, 0x0a, 0x44, 0x29, 0xa8, 0x84, 0x30, 0x06, 0x16, 0xae, 0xde,
	0x9f, 0x8b, 0x6e, 0xf5, 0x58, 0xbc, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x16, 0xa7, 0x1d, 0xad,
	0x93, 0x04, 0x00, 0x00,
}
