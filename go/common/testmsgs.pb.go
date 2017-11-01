// Code generated by protoc-gen-go.
// source: go/common/testdata/testmsgs.proto
// DO NOT EDIT!

package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

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
	Name           string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Key            string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Label          string `protobuf:"bytes,3,opt,name=label" json:"label,omitempty"`
	OptionalKey    string `protobuf:"bytes,4,opt,name=optional_key,json=optionalKey" json:"optional_key,omitempty"`
	Fqdn           string `protobuf:"bytes,5,opt,name=fqdn" json:"fqdn,omitempty"`
	Reqd           string `protobuf:"bytes,6,opt,name=reqd" json:"reqd,omitempty"`
	Output         string `protobuf:"bytes,7,opt,name=output" json:"output,omitempty"`
	OptionalString string `protobuf:"bytes,8,opt,name=optional_string,json=optionalString" json:"optional_string,omitempty"`
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

func (m *TestMessageWithOptions) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *TestMessageWithOptions) GetOptionalString() string {
	if m != nil {
		return m.OptionalString
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

type TestMessageWithTypes struct {
	Name          string           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	BoolValue     bool             `protobuf:"varint,2,opt,name=bool_value,json=boolValue" json:"bool_value,omitempty"`
	IntValue      int32            `protobuf:"varint,3,opt,name=int_value,json=intValue" json:"int_value,omitempty"`
	Field         *TestGoodProto   `protobuf:"bytes,4,opt,name=field" json:"field,omitempty"`
	RepeatedField []*TestGoodProto `protobuf:"bytes,5,rep,name=repeated_field,json=repeatedField" json:"repeated_field,omitempty"`
	// Types that are valid to be assigned to Optional:
	//	*TestMessageWithTypes_OptionalField
	Optional  isTestMessageWithTypes_Optional `protobuf_oneof:"optional"`
	MapField  map[string]*TestGoodProto       `protobuf:"bytes,7,rep,name=map_field,json=mapField" json:"map_field,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	MapString map[string]string               `protobuf:"bytes,8,rep,name=map_string,json=mapString" json:"map_string,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *TestMessageWithTypes) Reset()                    { *m = TestMessageWithTypes{} }
func (m *TestMessageWithTypes) String() string            { return proto.CompactTextString(m) }
func (*TestMessageWithTypes) ProtoMessage()               {}
func (*TestMessageWithTypes) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{13} }

type isTestMessageWithTypes_Optional interface {
	isTestMessageWithTypes_Optional()
}

type TestMessageWithTypes_OptionalField struct {
	OptionalField *TestGoodProto `protobuf:"bytes,6,opt,name=optional_field,json=optionalField,oneof"`
}

func (*TestMessageWithTypes_OptionalField) isTestMessageWithTypes_Optional() {}

func (m *TestMessageWithTypes) GetOptional() isTestMessageWithTypes_Optional {
	if m != nil {
		return m.Optional
	}
	return nil
}

func (m *TestMessageWithTypes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TestMessageWithTypes) GetBoolValue() bool {
	if m != nil {
		return m.BoolValue
	}
	return false
}

func (m *TestMessageWithTypes) GetIntValue() int32 {
	if m != nil {
		return m.IntValue
	}
	return 0
}

func (m *TestMessageWithTypes) GetField() *TestGoodProto {
	if m != nil {
		return m.Field
	}
	return nil
}

func (m *TestMessageWithTypes) GetRepeatedField() []*TestGoodProto {
	if m != nil {
		return m.RepeatedField
	}
	return nil
}

func (m *TestMessageWithTypes) GetOptionalField() *TestGoodProto {
	if x, ok := m.GetOptional().(*TestMessageWithTypes_OptionalField); ok {
		return x.OptionalField
	}
	return nil
}

func (m *TestMessageWithTypes) GetMapField() map[string]*TestGoodProto {
	if m != nil {
		return m.MapField
	}
	return nil
}

func (m *TestMessageWithTypes) GetMapString() map[string]string {
	if m != nil {
		return m.MapString
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TestMessageWithTypes) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TestMessageWithTypes_OneofMarshaler, _TestMessageWithTypes_OneofUnmarshaler, _TestMessageWithTypes_OneofSizer, []interface{}{
		(*TestMessageWithTypes_OptionalField)(nil),
	}
}

func _TestMessageWithTypes_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TestMessageWithTypes)
	// optional
	switch x := m.Optional.(type) {
	case *TestMessageWithTypes_OptionalField:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OptionalField); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TestMessageWithTypes.Optional has unexpected type %T", x)
	}
	return nil
}

func _TestMessageWithTypes_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TestMessageWithTypes)
	switch tag {
	case 6: // optional.optional_field
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TestGoodProto)
		err := b.DecodeMessage(msg)
		m.Optional = &TestMessageWithTypes_OptionalField{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TestMessageWithTypes_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TestMessageWithTypes)
	// optional
	switch x := m.Optional.(type) {
	case *TestMessageWithTypes_OptionalField:
		s := proto.Size(x.OptionalField)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

var E_X = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         51000,
	Name:          "common.x",
	Tag:           "varint,51000,opt,name=x",
	Filename:      "go/common/testdata/testmsgs.proto",
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
	proto.RegisterType((*TestMessageWithTypes)(nil), "common.TestMessageWithTypes")
	proto.RegisterExtension(E_X)
}

func init() { proto.RegisterFile("go/common/testdata/testmsgs.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 772 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0xcd, 0x6e, 0xeb, 0x44,
	0x14, 0xc7, 0x71, 0x12, 0xe7, 0x26, 0xe7, 0xde, 0xf4, 0x16, 0x53, 0x90, 0x9b, 0x52, 0x68, 0xc3,
	0xa2, 0xa5, 0xa1, 0x8e, 0x28, 0x48, 0xa0, 0x52, 0x21, 0x11, 0x41, 0x5b, 0x81, 0xaa, 0xc2, 0xb4,
	0x6a, 0x05, 0x9b, 0x68, 0x62, 0x1f, 0x27, 0x16, 0xb6, 0xc7, 0x1d, 0x4f, 0xa0, 0xd9, 0xf6, 0x01,
	0x78, 0x9b, 0x4a, 0x2c, 0x59, 0xb0, 0xe2, 0x75, 0x78, 0x01, 0x34, 0x33, 0x76, 0xea, 0x14, 0x27,
	0xed, 0xdd, 0xd9, 0xe7, 0xfc, 0xcf, 0x6f, 0x8e, 0xcf, 0x87, 0x07, 0xb6, 0x47, 0xac, 0xe7, 0xb2,
	0x28, 0x62, 0x71, 0x4f, 0x60, 0x2a, 0x3c, 0x2a, 0xa8, 0x7a, 0x88, 0xd2, 0x51, 0xea, 0x24, 0x9c,
	0x09, 0x66, 0xd5, 0xb5, 0xbf, 0xbd, 0x35, 0x62, 0x6c, 0x14, 0x62, 0x4f, 0x59, 0x87, 0x13, 0xbf,
	0xe7, 0x61, 0xea, 0xf2, 0x20, 0x11, 0x8c, 0x6b, 0x65, 0x7b, 0x23, 0x75, 0xc7, 0x18, 0xd1, 0x1c,
	0xe8, 0x07, 0x21, 0x72, 0xf4, 0xcb, 0x9d, 0x2c, 0x11, 0x01, 0x8b, 0xb3, 0x33, 0x3a, 0x3f, 0xc1,
	0xeb, 0x4b, 0x4c, 0xc5, 0x29, 0x4d, 0xfb, 0xd4, 0x3b, 0x0e, 0x30, 0xf4, 0x2c, 0x0b, 0x6a, 0x31,
	0x8d, 0xd0, 0x36, 0xb6, 0x8c, 0xdd, 0x26, 0x51, 0xcf, 0xd6, 0x1e, 0x98, 0xbe, 0x74, 0xda, 0x95,
	0x2d, 0x63, 0xf7, 0xe5, 0xc1, 0x9a, 0xa3, 0x61, 0x8e, 0x8c, 0xed, 0x53, 0xef, 0x47, 0xc9, 0x22,
	0x5a, 0xd2, 0xb9, 0x80, 0xd5, 0x0c, 0x79, 0xc2, 0xd8, 0x12, 0x66, 0x77, 0x9e, 0xf9, 0x6e, 0x91,
	0x29, 0x23, 0xe7, 0xa0, 0x73, 0x79, 0x5e, 0x84, 0x81, 0x8b, 0x4f, 0xe5, 0x59, 0x7d, 0xb3, 0x3c,
	0x17, 0x33, 0xbb, 0xf3, 0xcc, 0xe5, 0x79, 0xfe, 0x0c, 0xad, 0xdc, 0x7e, 0x1e, 0xe3, 0xb9, 0x5f,
	0x4a, 0xdc, 0x7f, 0xce, 0x97, 0x9f, 0xbe, 0x95, 0x31, 0xfb, 0x26, 0x54, 0x59, 0x22, 0x3a, 0xd7,
	0xf0, 0x2a, 0xfb, 0x8c, 0xc5, 0xe4, 0x4f, 0x9e, 0xd1, 0xa7, 0xff, 0x81, 0x3f, 0x86, 0x77, 0x32,
	0xff, 0x15, 0x0d, 0x03, 0x8f, 0x0a, 0xfc, 0x86, 0x8f, 0xd2, 0x32, 0x7e, 0x67, 0x07, 0xde, 0xce,
	0xa4, 0x04, 0xc5, 0x84, 0xc7, 0x97, 0xd3, 0xe4, 0xa1, 0x68, 0x95, 0x82, 0xb0, 0x33, 0x4b, 0x56,
	0x9d, 0x59, 0x0a, 0xfb, 0xe8, 0xa1, 0x56, 0x8b, 0x45, 0xff, 0x54, 0xe0, 0x3d, 0xa9, 0x3a, 0xc3,
	0x34, 0xa5, 0x23, 0xbc, 0x0e, 0xc4, 0xf8, 0x5c, 0x4f, 0x70, 0x69, 0x01, 0x3e, 0x85, 0xea, 0xaf,
	0x38, 0xd5, 0xa9, 0xf4, 0x3f, 0xbc, 0xbb, 0xb7, 0x37, 0xac, 0xf5, 0xdf, 0x03, 0x31, 0x1e, 0x88,
	0x69, 0x82, 0xa9, 0xc3, 0x31, 0x41, 0x2a, 0xd0, 0x1b, 0xa8, 0x02, 0x10, 0xa9, 0xb5, 0xde, 0x07,
	0x33, 0xa4, 0x43, 0x0c, 0xed, 0xaa, 0x0a, 0xaa, 0xdf, 0xdd, 0xdb, 0x95, 0x46, 0x95, 0x68, 0xa3,
	0xf5, 0x2d, 0xbc, 0xd2, 0x1b, 0x43, 0xc3, 0x81, 0x24, 0xd7, 0x94, 0x68, 0xfb, 0xee, 0xde, 0xde,
	0x5c, 0x42, 0xb6, 0x0d, 0xf2, 0x32, 0x0f, 0xfb, 0x01, 0xa7, 0x56, 0x1b, 0x6a, 0xfe, 0x8d, 0x17,
	0xdb, 0x66, 0xe1, 0x88, 0x1a, 0x51, 0x36, 0xe9, 0xe3, 0x78, 0xe3, 0xd9, 0xf5, 0x82, 0xcf, 0x20,
	0xca, 0x66, 0x7d, 0x00, 0x75, 0x36, 0x11, 0xc9, 0x44, 0xd8, 0x2f, 0x0a, 0xde, 0x0a, 0xc9, 0xac,
	0xd6, 0x3e, 0xbc, 0x9e, 0x65, 0x97, 0x0a, 0x1e, 0xc4, 0x23, 0xbb, 0xa1, 0x84, 0xb5, 0xbf, 0xfe,
	0xb5, 0x0d, 0xb2, 0x92, 0x3b, 0x2f, 0x94, 0xaf, 0xf3, 0x05, 0xac, 0x67, 0x5d, 0x29, 0x29, 0x67,
	0xbb, 0x50, 0x4e, 0x73, 0x56, 0x06, 0xdd, 0x85, 0xaf, 0xf4, 0xae, 0x1c, 0x07, 0x21, 0x12, 0xf4,
	0x75, 0xb7, 0x76, 0xa0, 0xca, 0xd1, 0x57, 0xf2, 0xc2, 0x0c, 0x67, 0x12, 0xe4, 0x18, 0xbb, 0x48,
	0xa4, 0xa2, 0xf3, 0x77, 0x0d, 0xd6, 0x1e, 0xb5, 0x50, 0xce, 0x4d, 0x79, 0x03, 0x37, 0x01, 0x86,
	0x8c, 0x85, 0x83, 0xdf, 0x68, 0x38, 0xd1, 0x23, 0xd5, 0x20, 0x4d, 0x69, 0xb9, 0x92, 0x06, 0x6b,
	0x03, 0x9a, 0x41, 0x2c, 0x32, 0xaf, 0x6c, 0x98, 0x49, 0x1a, 0x41, 0x2c, 0xb4, 0x73, 0xb6, 0xa9,
	0xb5, 0xa7, 0xff, 0x28, 0xd6, 0x11, 0xac, 0x3c, 0xea, 0x99, 0xb9, 0x6c, 0xbf, 0x5b, 0xb9, 0x58,
	0xff, 0xd0, 0xbe, 0x86, 0x59, 0x6d, 0xb3, 0xe8, 0xfa, 0xf2, 0x5d, 0x6e, 0xe5, 0x72, 0x1d, 0x7f,
	0x02, 0xcd, 0x88, 0x26, 0x59, 0xe8, 0x0b, 0x75, 0xf0, 0x5e, 0x31, 0xf4, 0x71, 0xad, 0x9c, 0x33,
	0x9a, 0xa8, 0xc8, 0xef, 0x62, 0xc1, 0xa7, 0xa4, 0x11, 0x65, 0xaf, 0xd6, 0xf7, 0x00, 0x12, 0x34,
	0x6b, 0xbe, 0x24, 0x75, 0x9f, 0x22, 0xe9, 0x71, 0xd0, 0x28, 0x99, 0x87, 0x7e, 0x6f, 0x13, 0x68,
	0xcd, 0x1d, 0x63, 0xad, 0xea, 0x6d, 0xd2, 0xfd, 0x51, 0xcb, 0xd2, 0x05, 0xf3, 0xa1, 0x33, 0x8b,
	0x4b, 0xac, 0x34, 0x87, 0x95, 0x2f, 0x8d, 0xf6, 0x11, 0xac, 0xcc, 0x1f, 0x58, 0x02, 0x5d, 0x2b,
	0x42, 0x9b, 0x85, 0xe8, 0x3e, 0x40, 0x23, 0xaf, 0xdb, 0xe1, 0x3e, 0x18, 0xb7, 0xd6, 0xa6, 0xa3,
	0x2f, 0x43, 0x27, 0xbf, 0x0c, 0x1d, 0x95, 0x6e, 0x36, 0xc3, 0xf6, 0x9f, 0x7f, 0x54, 0xd5, 0xc4,
	0x18, 0xb7, 0xfd, 0xcf, 0x7f, 0x39, 0x70, 0xc7, 0x9c, 0x45, 0xc1, 0x24, 0xca, 0x62, 0x52, 0x36,
	0xe1, 0x2e, 0xca, 0x8c, 0x7b, 0x18, 0x0b, 0xe4, 0x09, 0x0f, 0x52, 0xec, 0xb9, 0x18, 0xf6, 0x66,
	0x77, 0xf1, 0xb0, 0xae, 0xb8, 0x9f, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x26, 0xfe, 0xd0,
	0x9f, 0x07, 0x00, 0x00,
}
