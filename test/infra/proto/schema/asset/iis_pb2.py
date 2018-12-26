# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema/asset/iis.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from schema.common import file_reference_pb2 as schema_dot_common_dot_file__reference__pb2
from schema.asset import network_pb2 as schema_dot_asset_dot_network__pb2
from schema.common import validation_pb2 as schema_dot_common_dot_validation__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='schema/asset/iis.proto',
  package='asset',
  syntax='proto3',
  serialized_options=_b('Z1chromium.googlesource.com/enterprise/cel/go/asset'),
  serialized_pb=_b('\n\x16schema/asset/iis.proto\x12\x05\x61sset\x1a\"schema/common/file_reference.proto\x1a\x1aschema/asset/network.proto\x1a\x1eschema/common/validation.proto\"O\n\tIISServer\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x34\n\x0fwindows_machine\x18\x02 \x01(\tB\x1b\x82\xb5\x18\x17\x12\x15\x61sset.windows_machine\"\x91\x01\n\x0bIISBindings\x12\x10\n\x08hostname\x18\x01 \x01(\t\x12!\n\x08protocol\x18\x02 \x01(\x0e\x32\x0f.asset.Protocol\x12\x0c\n\x04port\x18\x03 \x01(\r\x12.\n\x0b\x63\x65rtificate\x18\x04 \x01(\tB\x19\x82\xb5\x18\x15\x12\x11\x61sset.certificate\x18\x01\x12\x0f\n\x07use_sni\x18\x05 \x01(\x08\"\xb9\x01\n\x07IISSite\x12\x0c\n\x04name\x18\x01 \x01(\t\x12*\n\niis_server\x18\x02 \x01(\tB\x16\x82\xb5\x18\x12\x12\x10\x61sset.iis_server\x12$\n\x08\x62indings\x18\x03 \x01(\x0b\x32\x12.asset.IISBindings\x12\'\n\x08\x63ontents\x18\x04 \x01(\x0b\x32\x15.common.FileReference\x12%\n\tauth_type\x18\x05 \x01(\x0e\x32\x12.asset.IISAuthType\"\xba\x01\n\x0eIISApplication\x12\x0c\n\x04name\x18\x01 \x01(\t\x12&\n\x08iis_site\x18\x02 \x01(\tB\x14\x82\xb5\x18\x10\x12\x0e\x61sset.iis_site\x12\'\n\x08\x63ontents\x18\x03 \x01(\x0b\x32\x15.common.FileReference\x12.\n\x0fweb_config_file\x18\x04 \x01(\x0b\x32\x15.common.FileReference\x12\x19\n\x11web_config_string\x18\x05 \x01(\t*I\n\x0bIISAuthType\x12\x08\n\x04NONE\x10\x00\x12\x08\n\x04NTLM\x10\x01\x12\x0c\n\x08KERBEROS\x10\x02\x12\x18\n\x14KERBEROS_NEGOTIABLE2\x10\x03\x42\x33Z1chromium.googlesource.com/enterprise/cel/go/assetb\x06proto3')
  ,
  dependencies=[schema_dot_common_dot_file__reference__pb2.DESCRIPTOR,schema_dot_asset_dot_network__pb2.DESCRIPTOR,schema_dot_common_dot_validation__pb2.DESCRIPTOR,])

_IISAUTHTYPE = _descriptor.EnumDescriptor(
  name='IISAuthType',
  full_name='asset.IISAuthType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NONE', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NTLM', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='KERBEROS', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='KERBEROS_NEGOTIABLE2', index=3, number=3,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=735,
  serialized_end=808,
)
_sym_db.RegisterEnumDescriptor(_IISAUTHTYPE)

IISAuthType = enum_type_wrapper.EnumTypeWrapper(_IISAUTHTYPE)
NONE = 0
NTLM = 1
KERBEROS = 2
KERBEROS_NEGOTIABLE2 = 3



_IISSERVER = _descriptor.Descriptor(
  name='IISServer',
  full_name='asset.IISServer',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='asset.IISServer.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='windows_machine', full_name='asset.IISServer.windows_machine', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\027\022\025asset.windows_machine'), file=DESCRIPTOR),
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
  serialized_start=129,
  serialized_end=208,
)


_IISBINDINGS = _descriptor.Descriptor(
  name='IISBindings',
  full_name='asset.IISBindings',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='hostname', full_name='asset.IISBindings.hostname', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='protocol', full_name='asset.IISBindings.protocol', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='port', full_name='asset.IISBindings.port', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='certificate', full_name='asset.IISBindings.certificate', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\025\022\021asset.certificate\030\001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='use_sni', full_name='asset.IISBindings.use_sni', index=4,
      number=5, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
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
  serialized_start=211,
  serialized_end=356,
)


_IISSITE = _descriptor.Descriptor(
  name='IISSite',
  full_name='asset.IISSite',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='asset.IISSite.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='iis_server', full_name='asset.IISSite.iis_server', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\022\022\020asset.iis_server'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='bindings', full_name='asset.IISSite.bindings', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='contents', full_name='asset.IISSite.contents', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='auth_type', full_name='asset.IISSite.auth_type', index=4,
      number=5, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=359,
  serialized_end=544,
)


_IISAPPLICATION = _descriptor.Descriptor(
  name='IISApplication',
  full_name='asset.IISApplication',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='asset.IISApplication.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='iis_site', full_name='asset.IISApplication.iis_site', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\020\022\016asset.iis_site'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='contents', full_name='asset.IISApplication.contents', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='web_config_file', full_name='asset.IISApplication.web_config_file', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='web_config_string', full_name='asset.IISApplication.web_config_string', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=547,
  serialized_end=733,
)

_IISBINDINGS.fields_by_name['protocol'].enum_type = schema_dot_asset_dot_network__pb2._PROTOCOL
_IISSITE.fields_by_name['bindings'].message_type = _IISBINDINGS
_IISSITE.fields_by_name['contents'].message_type = schema_dot_common_dot_file__reference__pb2._FILEREFERENCE
_IISSITE.fields_by_name['auth_type'].enum_type = _IISAUTHTYPE
_IISAPPLICATION.fields_by_name['contents'].message_type = schema_dot_common_dot_file__reference__pb2._FILEREFERENCE
_IISAPPLICATION.fields_by_name['web_config_file'].message_type = schema_dot_common_dot_file__reference__pb2._FILEREFERENCE
DESCRIPTOR.message_types_by_name['IISServer'] = _IISSERVER
DESCRIPTOR.message_types_by_name['IISBindings'] = _IISBINDINGS
DESCRIPTOR.message_types_by_name['IISSite'] = _IISSITE
DESCRIPTOR.message_types_by_name['IISApplication'] = _IISAPPLICATION
DESCRIPTOR.enum_types_by_name['IISAuthType'] = _IISAUTHTYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

IISServer = _reflection.GeneratedProtocolMessageType('IISServer', (_message.Message,), dict(
  DESCRIPTOR = _IISSERVER,
  __module__ = 'schema.asset.iis_pb2'
  # @@protoc_insertion_point(class_scope:asset.IISServer)
  ))
_sym_db.RegisterMessage(IISServer)

IISBindings = _reflection.GeneratedProtocolMessageType('IISBindings', (_message.Message,), dict(
  DESCRIPTOR = _IISBINDINGS,
  __module__ = 'schema.asset.iis_pb2'
  # @@protoc_insertion_point(class_scope:asset.IISBindings)
  ))
_sym_db.RegisterMessage(IISBindings)

IISSite = _reflection.GeneratedProtocolMessageType('IISSite', (_message.Message,), dict(
  DESCRIPTOR = _IISSITE,
  __module__ = 'schema.asset.iis_pb2'
  # @@protoc_insertion_point(class_scope:asset.IISSite)
  ))
_sym_db.RegisterMessage(IISSite)

IISApplication = _reflection.GeneratedProtocolMessageType('IISApplication', (_message.Message,), dict(
  DESCRIPTOR = _IISAPPLICATION,
  __module__ = 'schema.asset.iis_pb2'
  # @@protoc_insertion_point(class_scope:asset.IISApplication)
  ))
_sym_db.RegisterMessage(IISApplication)


DESCRIPTOR._options = None
_IISSERVER.fields_by_name['windows_machine']._options = None
_IISBINDINGS.fields_by_name['certificate']._options = None
_IISSITE.fields_by_name['iis_server']._options = None
_IISAPPLICATION.fields_by_name['iis_site']._options = None
# @@protoc_insertion_point(module_scope)