# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema/asset/cert.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from schema.common import file_reference_pb2 as schema_dot_common_dot_file__reference__pb2
from schema.common import validation_pb2 as schema_dot_common_dot_validation__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='schema/asset/cert.proto',
  package='asset',
  syntax='proto3',
  serialized_options=_b('Z1chromium.googlesource.com/enterprise/cel/go/asset'),
  serialized_pb=_b('\n\x17schema/asset/cert.proto\x12\x05\x61sset\x1a\"schema/common/file_reference.proto\x1a\x1eschema/common/validation.proto\"z\n\x0b\x43\x65rtificate\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x31\n\ncertficate\x18\x02 \x01(\x0b\x32\x15.common.FileReferenceB\x06\x82\xb5\x18\x02\x08\x01\x12*\n\x0bprivate_key\x18\x03 \x01(\x0b\x32\x15.common.FileReference\"|\n\x0f\x43\x65rtificatePool\x12\x0c\n\x04name\x18\x01 \x01(\t\x12.\n\rinclude_named\x18\x02 \x03(\tB\x17\x82\xb5\x18\x13\x12\x11\x61sset.certificate\x12+\n\x0cinclude_file\x18\x03 \x03(\x0b\x32\x15.common.FileReferenceB3Z1chromium.googlesource.com/enterprise/cel/go/assetb\x06proto3')
  ,
  dependencies=[schema_dot_common_dot_file__reference__pb2.DESCRIPTOR,schema_dot_common_dot_validation__pb2.DESCRIPTOR,])




_CERTIFICATE = _descriptor.Descriptor(
  name='Certificate',
  full_name='asset.Certificate',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='asset.Certificate.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='certficate', full_name='asset.Certificate.certficate', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='private_key', full_name='asset.Certificate.private_key', index=2,
      number=3, type=11, cpp_type=10, label=1,
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
  serialized_start=102,
  serialized_end=224,
)


_CERTIFICATEPOOL = _descriptor.Descriptor(
  name='CertificatePool',
  full_name='asset.CertificatePool',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='asset.CertificatePool.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='include_named', full_name='asset.CertificatePool.include_named', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\023\022\021asset.certificate'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='include_file', full_name='asset.CertificatePool.include_file', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=226,
  serialized_end=350,
)

_CERTIFICATE.fields_by_name['certficate'].message_type = schema_dot_common_dot_file__reference__pb2._FILEREFERENCE
_CERTIFICATE.fields_by_name['private_key'].message_type = schema_dot_common_dot_file__reference__pb2._FILEREFERENCE
_CERTIFICATEPOOL.fields_by_name['include_file'].message_type = schema_dot_common_dot_file__reference__pb2._FILEREFERENCE
DESCRIPTOR.message_types_by_name['Certificate'] = _CERTIFICATE
DESCRIPTOR.message_types_by_name['CertificatePool'] = _CERTIFICATEPOOL
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Certificate = _reflection.GeneratedProtocolMessageType('Certificate', (_message.Message,), dict(
  DESCRIPTOR = _CERTIFICATE,
  __module__ = 'schema.asset.cert_pb2'
  # @@protoc_insertion_point(class_scope:asset.Certificate)
  ))
_sym_db.RegisterMessage(Certificate)

CertificatePool = _reflection.GeneratedProtocolMessageType('CertificatePool', (_message.Message,), dict(
  DESCRIPTOR = _CERTIFICATEPOOL,
  __module__ = 'schema.asset.cert_pb2'
  # @@protoc_insertion_point(class_scope:asset.CertificatePool)
  ))
_sym_db.RegisterMessage(CertificatePool)


DESCRIPTOR._options = None
_CERTIFICATE.fields_by_name['certficate']._options = None
_CERTIFICATEPOOL.fields_by_name['include_named']._options = None
# @@protoc_insertion_point(module_scope)