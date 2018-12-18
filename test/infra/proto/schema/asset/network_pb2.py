# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema/asset/network.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from schema.common import validation_pb2 as schema_dot_common_dot_validation__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='schema/asset/network.proto',
  package='asset',
  syntax='proto3',
  serialized_options=_b('Z1chromium.googlesource.com/enterprise/cel/go/asset'),
  serialized_pb=_b('\n\x1aschema/asset/network.proto\x12\x05\x61sset\x1a\x1eschema/common/validation.proto\"C\n\x07Network\x12\x0c\n\x04name\x18\x01 \x01(\t\x12*\n\raddress_range\x18\x02 \x01(\x0b\x32\x13.asset.AddressRange\"3\n\x0bNetworkPeer\x12$\n\x07network\x18\x01 \x03(\tB\x13\x82\xb5\x18\x0f\x12\rasset.network\"\x1d\n\x07\x41\x64\x64ress\x12\x12\n\x02ip\x18\x01 \x01(\tB\x06\x82\xb5\x18\x02\x08\x01\"$\n\x0c\x41\x64\x64ressRange\x12\x14\n\x04\x63idr\x18\x01 \x01(\tB\x06\x82\xb5\x18\x02\x08\x01\"\xa4\x01\n\x0c\x46ixedAddress\x12!\n\x07\x61\x64\x64ress\x18\x01 \x01(\x0b\x32\x0e.asset.AddressH\x00\x12/\n\x0c\x61\x64\x64ress_pool\x18\x02 \x01(\tB\x17\x82\xb5\x18\x13\x12\x11host.address_poolH\x00\x12\x30\n\x10resolved_address\x18\x03 \x01(\x0b\x32\x0e.asset.AddressB\x06\x82\xb5\x18\x02\x08\x05\x42\x0e\n\x0c\x61\x64\x64ress_type*,\n\x08Protocol\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x08\n\x04HTTP\x10\x01\x12\t\n\x05HTTPS\x10\x02\x42\x33Z1chromium.googlesource.com/enterprise/cel/go/assetb\x06proto3')
  ,
  dependencies=[schema_dot_common_dot_validation__pb2.DESCRIPTOR,])

_PROTOCOL = _descriptor.EnumDescriptor(
  name='Protocol',
  full_name='asset.Protocol',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HTTP', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HTTPS', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=427,
  serialized_end=471,
)
_sym_db.RegisterEnumDescriptor(_PROTOCOL)

Protocol = enum_type_wrapper.EnumTypeWrapper(_PROTOCOL)
UNKNOWN = 0
HTTP = 1
HTTPS = 2



_NETWORK = _descriptor.Descriptor(
  name='Network',
  full_name='asset.Network',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='asset.Network.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='address_range', full_name='asset.Network.address_range', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=69,
  serialized_end=136,
)


_NETWORKPEER = _descriptor.Descriptor(
  name='NetworkPeer',
  full_name='asset.NetworkPeer',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='network', full_name='asset.NetworkPeer.network', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\017\022\rasset.network'), file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=138,
  serialized_end=189,
)


_ADDRESS = _descriptor.Descriptor(
  name='Address',
  full_name='asset.Address',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ip', full_name='asset.Address.ip', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\001'), file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=191,
  serialized_end=220,
)


_ADDRESSRANGE = _descriptor.Descriptor(
  name='AddressRange',
  full_name='asset.AddressRange',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cidr', full_name='asset.AddressRange.cidr', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\001'), file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=222,
  serialized_end=258,
)


_FIXEDADDRESS = _descriptor.Descriptor(
  name='FixedAddress',
  full_name='asset.FixedAddress',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='address', full_name='asset.FixedAddress.address', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='address_pool', full_name='asset.FixedAddress.address_pool', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\023\022\021host.address_pool'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='resolved_address', full_name='asset.FixedAddress.resolved_address', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\005'), file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='address_type', full_name='asset.FixedAddress.address_type',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=261,
  serialized_end=425,
)

_NETWORK.fields_by_name['address_range'].message_type = _ADDRESSRANGE
_FIXEDADDRESS.fields_by_name['address'].message_type = _ADDRESS
_FIXEDADDRESS.fields_by_name['resolved_address'].message_type = _ADDRESS
_FIXEDADDRESS.oneofs_by_name['address_type'].fields.append(
  _FIXEDADDRESS.fields_by_name['address'])
_FIXEDADDRESS.fields_by_name['address'].containing_oneof = _FIXEDADDRESS.oneofs_by_name['address_type']
_FIXEDADDRESS.oneofs_by_name['address_type'].fields.append(
  _FIXEDADDRESS.fields_by_name['address_pool'])
_FIXEDADDRESS.fields_by_name['address_pool'].containing_oneof = _FIXEDADDRESS.oneofs_by_name['address_type']
DESCRIPTOR.message_types_by_name['Network'] = _NETWORK
DESCRIPTOR.message_types_by_name['NetworkPeer'] = _NETWORKPEER
DESCRIPTOR.message_types_by_name['Address'] = _ADDRESS
DESCRIPTOR.message_types_by_name['AddressRange'] = _ADDRESSRANGE
DESCRIPTOR.message_types_by_name['FixedAddress'] = _FIXEDADDRESS
DESCRIPTOR.enum_types_by_name['Protocol'] = _PROTOCOL
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Network = _reflection.GeneratedProtocolMessageType('Network', (_message.Message,), dict(
  DESCRIPTOR = _NETWORK,
  __module__ = 'schema.asset.network_pb2'
  # @@protoc_insertion_point(class_scope:asset.Network)
  ))
_sym_db.RegisterMessage(Network)

NetworkPeer = _reflection.GeneratedProtocolMessageType('NetworkPeer', (_message.Message,), dict(
  DESCRIPTOR = _NETWORKPEER,
  __module__ = 'schema.asset.network_pb2'
  # @@protoc_insertion_point(class_scope:asset.NetworkPeer)
  ))
_sym_db.RegisterMessage(NetworkPeer)

Address = _reflection.GeneratedProtocolMessageType('Address', (_message.Message,), dict(
  DESCRIPTOR = _ADDRESS,
  __module__ = 'schema.asset.network_pb2'
  # @@protoc_insertion_point(class_scope:asset.Address)
  ))
_sym_db.RegisterMessage(Address)

AddressRange = _reflection.GeneratedProtocolMessageType('AddressRange', (_message.Message,), dict(
  DESCRIPTOR = _ADDRESSRANGE,
  __module__ = 'schema.asset.network_pb2'
  # @@protoc_insertion_point(class_scope:asset.AddressRange)
  ))
_sym_db.RegisterMessage(AddressRange)

FixedAddress = _reflection.GeneratedProtocolMessageType('FixedAddress', (_message.Message,), dict(
  DESCRIPTOR = _FIXEDADDRESS,
  __module__ = 'schema.asset.network_pb2'
  # @@protoc_insertion_point(class_scope:asset.FixedAddress)
  ))
_sym_db.RegisterMessage(FixedAddress)


DESCRIPTOR._options = None
_NETWORKPEER.fields_by_name['network']._options = None
_ADDRESS.fields_by_name['ip']._options = None
_ADDRESSRANGE.fields_by_name['cidr']._options = None
_FIXEDADDRESS.fields_by_name['address_pool']._options = None
_FIXEDADDRESS.fields_by_name['resolved_address']._options = None
# @@protoc_insertion_point(module_scope)
