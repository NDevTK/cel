# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema/asset/remote_desktop.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from schema.asset import machine_pb2 as schema_dot_asset_dot_machine__pb2
from schema.common import validation_pb2 as schema_dot_common_dot_validation__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='schema/asset/remote_desktop.proto',
  package='asset',
  syntax='proto3',
  serialized_options=_b('Z1chromium.googlesource.com/enterprise/cel/go/asset'),
  serialized_pb=_b('\n!schema/asset/remote_desktop.proto\x12\x05\x61sset\x1a\x1aschema/asset/machine.proto\x1a\x1eschema/common/validation.proto\"\x82\x01\n\x11RemoteDesktopHost\x12\x34\n\x0fwindows_machine\x18\x01 \x01(\tB\x1b\x82\xb5\x18\x17\x12\x15\x61sset.windows_machine\x12\x17\n\x0f\x63ollection_name\x18\x02 \x01(\t\x12\x1e\n\x16\x63ollection_description\x18\x03 \x01(\tB3Z1chromium.googlesource.com/enterprise/cel/go/assetb\x06proto3')
  ,
  dependencies=[schema_dot_asset_dot_machine__pb2.DESCRIPTOR,schema_dot_common_dot_validation__pb2.DESCRIPTOR,])




_REMOTEDESKTOPHOST = _descriptor.Descriptor(
  name='RemoteDesktopHost',
  full_name='asset.RemoteDesktopHost',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='windows_machine', full_name='asset.RemoteDesktopHost.windows_machine', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\027\022\025asset.windows_machine'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='collection_name', full_name='asset.RemoteDesktopHost.collection_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='collection_description', full_name='asset.RemoteDesktopHost.collection_description', index=2,
      number=3, type=9, cpp_type=9, label=1,
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
  serialized_start=105,
  serialized_end=235,
)

DESCRIPTOR.message_types_by_name['RemoteDesktopHost'] = _REMOTEDESKTOPHOST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

RemoteDesktopHost = _reflection.GeneratedProtocolMessageType('RemoteDesktopHost', (_message.Message,), dict(
  DESCRIPTOR = _REMOTEDESKTOPHOST,
  __module__ = 'schema.asset.remote_desktop_pb2'
  # @@protoc_insertion_point(class_scope:asset.RemoteDesktopHost)
  ))
_sym_db.RegisterMessage(RemoteDesktopHost)


DESCRIPTOR._options = None
_REMOTEDESKTOPHOST.fields_by_name['windows_machine']._options = None
# @@protoc_insertion_point(module_scope)
