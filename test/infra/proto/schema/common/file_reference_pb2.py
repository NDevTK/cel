# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema/common/file_reference.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from schema.common import validation_pb2 as schema_dot_common_dot_validation__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='schema/common/file_reference.proto',
  package='common',
  syntax='proto3',
  serialized_options=_b('Z2chromium.googlesource.com/enterprise/cel/go/common'),
  serialized_pb=_b('\n\"schema/common/file_reference.proto\x12\x06\x63ommon\x1a\x1eschema/common/validation.proto\"\xf9\x01\n\rFileReference\x12\x0e\n\x06source\x18\x01 \x01(\t\x12\x13\n\x0btarget_path\x18\x02 \x01(\t\x12\x19\n\tfull_path\x18\x03 \x01(\tB\x06\x82\xb5\x18\x02\x08\x05\x12 \n\x10object_reference\x18\x04 \x01(\tB\x06\x82\xb5\x18\x02\x08\x05\x12\x19\n\tintegrity\x18\x05 \x01(\tB\x06\x82\xb5\x18\x02\x08\x05\x12\x39\n\rresolved_type\x18\x06 \x01(\x0e\x32\x1a.common.FileReference.TypeB\x06\x82\xb5\x18\x02\x08\x05\"0\n\x04Type\x12\r\n\tUNDEFINED\x10\x00\x12\x08\n\x04\x46ILE\x10\x01\x12\x0f\n\x0bZIP_ARCHIVE\x10\x02\x42\x34Z2chromium.googlesource.com/enterprise/cel/go/commonb\x06proto3')
  ,
  dependencies=[schema_dot_common_dot_validation__pb2.DESCRIPTOR,])



_FILEREFERENCE_TYPE = _descriptor.EnumDescriptor(
  name='Type',
  full_name='common.FileReference.Type',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNDEFINED', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FILE', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ZIP_ARCHIVE', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=280,
  serialized_end=328,
)
_sym_db.RegisterEnumDescriptor(_FILEREFERENCE_TYPE)


_FILEREFERENCE = _descriptor.Descriptor(
  name='FileReference',
  full_name='common.FileReference',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='source', full_name='common.FileReference.source', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='target_path', full_name='common.FileReference.target_path', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='full_path', full_name='common.FileReference.full_path', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\005'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='object_reference', full_name='common.FileReference.object_reference', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\005'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='integrity', full_name='common.FileReference.integrity', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\005'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='resolved_type', full_name='common.FileReference.resolved_type', index=5,
      number=6, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\005'), file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _FILEREFERENCE_TYPE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=79,
  serialized_end=328,
)

_FILEREFERENCE.fields_by_name['resolved_type'].enum_type = _FILEREFERENCE_TYPE
_FILEREFERENCE_TYPE.containing_type = _FILEREFERENCE
DESCRIPTOR.message_types_by_name['FileReference'] = _FILEREFERENCE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

FileReference = _reflection.GeneratedProtocolMessageType('FileReference', (_message.Message,), dict(
  DESCRIPTOR = _FILEREFERENCE,
  __module__ = 'schema.common.file_reference_pb2'
  # @@protoc_insertion_point(class_scope:common.FileReference)
  ))
_sym_db.RegisterMessage(FileReference)


DESCRIPTOR._options = None
_FILEREFERENCE.fields_by_name['full_path']._options = None
_FILEREFERENCE.fields_by_name['object_reference']._options = None
_FILEREFERENCE.fields_by_name['integrity']._options = None
_FILEREFERENCE.fields_by_name['resolved_type']._options = None
# @@protoc_insertion_point(module_scope)
