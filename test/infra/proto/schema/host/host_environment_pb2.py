# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema/host/host_environment.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from schema.asset import network_pb2 as schema_dot_asset_dot_network__pb2
from schema.common import file_reference_pb2 as schema_dot_common_dot_file__reference__pb2
from schema.common import validation_pb2 as schema_dot_common_dot_validation__pb2
from schema.gcp.compute import compute_api_pb2 as schema_dot_gcp_dot_compute_dot_compute__api__pb2
from schema.gcp.cloudkms import cloudkms_api_pb2 as schema_dot_gcp_dot_cloudkms_dot_cloudkms__api__pb2
from google.iam.admin.v1 import iam_pb2 as google_dot_iam_dot_admin_dot_v1_dot_iam__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='schema/host/host_environment.proto',
  package='host',
  syntax='proto3',
  serialized_options=_b('Z0chromium.googlesource.com/enterprise/cel/go/host'),
  serialized_pb=_b('\n\"schema/host/host_environment.proto\x12\x04host\x1a\x1aschema/asset/network.proto\x1a\"schema/common/file_reference.proto\x1a\x1eschema/common/validation.proto\x1a$schema/gcp/compute/compute-api.proto\x1a&schema/gcp/cloudkms/cloudkms-api.proto\x1a\x1dgoogle/iam/admin/v1/iam.proto\"`\n\x07Project\x12\x14\n\x04name\x18\x01 \x01(\tB\x06\x82\xb5\x18\x02\x08\x04\x12\x14\n\x04zone\x18\x02 \x01(\tB\x06\x82\xb5\x18\x02\x08\x02\x12)\n\x07project\x18\x64 \x01(\x0b\x32\x10.compute.ProjectB\x06\x82\xb5\x18\x02\x08\x05\"(\n\x0bLogSettings\x12\x19\n\tadmin_log\x18\x01 \x01(\tB\x06\x82\xb5\x18\x02\x08\x02\"M\n\x07Storage\x12\x16\n\x06\x62ucket\x18\x01 \x01(\tB\x06\x82\xb5\x18\x02\x08\x01\x12\x0e\n\x06prefix\x18\x02 \x01(\t\x12\x1a\n\ncreated_on\x18\n \x01(\tB\x06\x82\xb5\x18\x02\x08\x05\"\xb7\x01\n\x05Image\x12\x0c\n\x04name\x18\x01 \x01(\t\x12$\n\x06latest\x18\x02 \x01(\x0b\x32\x12.host.Image.FamilyH\x00\x12\x0f\n\x05\x66ixed\x18\x03 \x01(\tH\x00\x12\x0f\n\x07package\x18\x04 \x03(\t\x12\x13\n\x03url\x18\x05 \x01(\tB\x06\x82\xb5\x18\x02\x08\x05\x1a\x39\n\x06\x46\x61mily\x12\x17\n\x07project\x18\x01 \x01(\tB\x06\x82\xb5\x18\x02\x08\x04\x12\x16\n\x06\x66\x61mily\x18\x02 \x01(\tB\x06\x82\xb5\x18\x02\x08\x02\x42\x08\n\x06source\"[\n\x08NestedVM\x12\x15\n\x05image\x18\x01 \x01(\tB\x06\x82\xb5\x18\x02\x08\x01\x12\x11\n\tuser_name\x18\x02 \x01(\t\x12\x10\n\x08password\x18\x03 \x01(\t\x12\x13\n\x0bmachineType\x18\x04 \x01(\t\"\xa9\x01\n\x0bMachineType\x12\x0c\n\x04name\x18\x01 \x01(\t\x12:\n\x13instance_properties\x18\x05 \x01(\x0b\x32\x1b.compute.InstancePropertiesH\x00\x12#\n\x11instance_template\x18\x06 \x01(\tB\x06\x82\xb5\x18\x02\x08\x01H\x00\x12#\n\tnested_vm\x18\x07 \x01(\x0b\x32\x0e.host.NestedVMH\x00\x42\x06\n\x04\x62\x61se\"|\n\x0b\x41\x64\x64ressPool\x12\x0c\n\x04name\x18\x01 \x01(\t\x12%\n\rfixed_address\x18\x02 \x03(\x0b\x32\x0e.asset.Address\x12\x18\n\x10reserved_address\x18\x03 \x03(\t\x12\x1e\n\x16reserved_address_regex\x18\x04 \x03(\t\"\xe1\x01\n\x07Startup\x12\x32\n\x0bwin_startup\x18\x64 \x01(\x0b\x32\x15.common.FileReferenceB\x06\x82\xb5\x18\x02\x08\x05\x12\x34\n\rwin_agent_x64\x18\x65 \x01(\x0b\x32\x15.common.FileReferenceB\x06\x82\xb5\x18\x02\x08\x05\x12\x34\n\rlinux_startup\x18\x66 \x01(\x0b\x32\x15.common.FileReferenceB\x06\x82\xb5\x18\x02\x08\x05\x12\x36\n\x0flinux_agent_x64\x18g \x01(\x0b\x32\x15.common.FileReferenceB\x06\x82\xb5\x18\x02\x08\x05\"\xaf\x01\n\x0eRuntimeSupport\x12\x44\n\x0fservice_account\x18\x64 \x01(\x0b\x32#.google.iam.admin.v1.ServiceAccountB\x06\x82\xb5\x18\x02\x08\x05\x12/\n\ncrypto_key\x18\x65 \x01(\x0b\x32\x13.cloudkms.CryptoKeyB\x06\x82\xb5\x18\x02\x08\x05\x12&\n\x07startup\x18\x66 \x01(\x0b\x32\r.host.StartupB\x06\x82\xb5\x18\x02\x08\x01\"\xc9\x02\n\x0fHostEnvironment\x12&\n\x07project\x18\x01 \x01(\x0b\x32\r.host.ProjectB\x06\x82\xb5\x18\x02\x08\x01\x12&\n\x07storage\x18\x02 \x01(\x0b\x32\r.host.StorageB\x06\x82\xb5\x18\x02\x08\x01\x12/\n\x0clog_settings\x18\x03 \x01(\x0b\x32\x11.host.LogSettingsB\x06\x82\xb5\x18\x02\x08\x01\x12/\n\x0cmachine_type\x18\n \x03(\x0b\x32\x11.host.MachineTypeB\x06\x82\xb5\x18\x02\x08\x07\x12/\n\x0c\x61\x64\x64ress_pool\x18\x0b \x03(\x0b\x32\x11.host.AddressPoolB\x06\x82\xb5\x18\x02\x08\x07\x12\"\n\x05image\x18\x0c \x03(\x0b\x32\x0b.host.ImageB\x06\x82\xb5\x18\x02\x08\x07\x12/\n\tresources\x18\x64 \x01(\x0b\x32\x14.host.RuntimeSupportB\x06\x82\xb5\x18\x02\x08\x05\x42\x32Z0chromium.googlesource.com/enterprise/cel/go/hostb\x06proto3')
  ,
  dependencies=[schema_dot_asset_dot_network__pb2.DESCRIPTOR,schema_dot_common_dot_file__reference__pb2.DESCRIPTOR,schema_dot_common_dot_validation__pb2.DESCRIPTOR,schema_dot_gcp_dot_compute_dot_compute__api__pb2.DESCRIPTOR,schema_dot_gcp_dot_cloudkms_dot_cloudkms__api__pb2.DESCRIPTOR,google_dot_iam_dot_admin_dot_v1_dot_iam__pb2.DESCRIPTOR,])




_PROJECT = _descriptor.Descriptor(
  name='Project',
  full_name='host.Project',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='host.Project.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\004'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='zone', full_name='host.Project.zone', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\002'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='project', full_name='host.Project.project', index=2,
      number=100, type=11, cpp_type=10, label=1,
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
  ],
  serialized_start=249,
  serialized_end=345,
)


_LOGSETTINGS = _descriptor.Descriptor(
  name='LogSettings',
  full_name='host.LogSettings',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='admin_log', full_name='host.LogSettings.admin_log', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\002'), file=DESCRIPTOR),
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
  serialized_start=347,
  serialized_end=387,
)


_STORAGE = _descriptor.Descriptor(
  name='Storage',
  full_name='host.Storage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='bucket', full_name='host.Storage.bucket', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='prefix', full_name='host.Storage.prefix', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='created_on', full_name='host.Storage.created_on', index=2,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  ],
  serialized_start=389,
  serialized_end=466,
)


_IMAGE_FAMILY = _descriptor.Descriptor(
  name='Family',
  full_name='host.Image.Family',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='project', full_name='host.Image.Family.project', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\004'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='family', full_name='host.Image.Family.family', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\002'), file=DESCRIPTOR),
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
  serialized_start=585,
  serialized_end=642,
)

_IMAGE = _descriptor.Descriptor(
  name='Image',
  full_name='host.Image',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='host.Image.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='latest', full_name='host.Image.latest', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='fixed', full_name='host.Image.fixed', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='package', full_name='host.Image.package', index=3,
      number=4, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='url', full_name='host.Image.url', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\005'), file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_IMAGE_FAMILY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='source', full_name='host.Image.source',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=469,
  serialized_end=652,
)


_NESTEDVM = _descriptor.Descriptor(
  name='NestedVM',
  full_name='host.NestedVM',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='image', full_name='host.NestedVM.image', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='user_name', full_name='host.NestedVM.user_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='password', full_name='host.NestedVM.password', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='machineType', full_name='host.NestedVM.machineType', index=3,
      number=4, type=9, cpp_type=9, label=1,
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
  serialized_start=654,
  serialized_end=745,
)


_MACHINETYPE = _descriptor.Descriptor(
  name='MachineType',
  full_name='host.MachineType',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='host.MachineType.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='instance_properties', full_name='host.MachineType.instance_properties', index=1,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='instance_template', full_name='host.MachineType.instance_template', index=2,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='nested_vm', full_name='host.MachineType.nested_vm', index=3,
      number=7, type=11, cpp_type=10, label=1,
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
    _descriptor.OneofDescriptor(
      name='base', full_name='host.MachineType.base',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=748,
  serialized_end=917,
)


_ADDRESSPOOL = _descriptor.Descriptor(
  name='AddressPool',
  full_name='host.AddressPool',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='host.AddressPool.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='fixed_address', full_name='host.AddressPool.fixed_address', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='reserved_address', full_name='host.AddressPool.reserved_address', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='reserved_address_regex', full_name='host.AddressPool.reserved_address_regex', index=3,
      number=4, type=9, cpp_type=9, label=3,
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
  serialized_start=919,
  serialized_end=1043,
)


_STARTUP = _descriptor.Descriptor(
  name='Startup',
  full_name='host.Startup',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='win_startup', full_name='host.Startup.win_startup', index=0,
      number=100, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\005'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='win_agent_x64', full_name='host.Startup.win_agent_x64', index=1,
      number=101, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\005'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='linux_startup', full_name='host.Startup.linux_startup', index=2,
      number=102, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\005'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='linux_agent_x64', full_name='host.Startup.linux_agent_x64', index=3,
      number=103, type=11, cpp_type=10, label=1,
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
  ],
  serialized_start=1046,
  serialized_end=1271,
)


_RUNTIMESUPPORT = _descriptor.Descriptor(
  name='RuntimeSupport',
  full_name='host.RuntimeSupport',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='service_account', full_name='host.RuntimeSupport.service_account', index=0,
      number=100, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\005'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='crypto_key', full_name='host.RuntimeSupport.crypto_key', index=1,
      number=101, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\005'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='startup', full_name='host.RuntimeSupport.startup', index=2,
      number=102, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=1274,
  serialized_end=1449,
)


_HOSTENVIRONMENT = _descriptor.Descriptor(
  name='HostEnvironment',
  full_name='host.HostEnvironment',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='project', full_name='host.HostEnvironment.project', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='storage', full_name='host.HostEnvironment.storage', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='log_settings', full_name='host.HostEnvironment.log_settings', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='machine_type', full_name='host.HostEnvironment.machine_type', index=3,
      number=10, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\007'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='address_pool', full_name='host.HostEnvironment.address_pool', index=4,
      number=11, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\007'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='image', full_name='host.HostEnvironment.image', index=5,
      number=12, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\202\265\030\002\010\007'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='resources', full_name='host.HostEnvironment.resources', index=6,
      number=100, type=11, cpp_type=10, label=1,
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
  ],
  serialized_start=1452,
  serialized_end=1781,
)

_PROJECT.fields_by_name['project'].message_type = schema_dot_gcp_dot_compute_dot_compute__api__pb2._PROJECT
_IMAGE_FAMILY.containing_type = _IMAGE
_IMAGE.fields_by_name['latest'].message_type = _IMAGE_FAMILY
_IMAGE.oneofs_by_name['source'].fields.append(
  _IMAGE.fields_by_name['latest'])
_IMAGE.fields_by_name['latest'].containing_oneof = _IMAGE.oneofs_by_name['source']
_IMAGE.oneofs_by_name['source'].fields.append(
  _IMAGE.fields_by_name['fixed'])
_IMAGE.fields_by_name['fixed'].containing_oneof = _IMAGE.oneofs_by_name['source']
_MACHINETYPE.fields_by_name['instance_properties'].message_type = schema_dot_gcp_dot_compute_dot_compute__api__pb2._INSTANCEPROPERTIES
_MACHINETYPE.fields_by_name['nested_vm'].message_type = _NESTEDVM
_MACHINETYPE.oneofs_by_name['base'].fields.append(
  _MACHINETYPE.fields_by_name['instance_properties'])
_MACHINETYPE.fields_by_name['instance_properties'].containing_oneof = _MACHINETYPE.oneofs_by_name['base']
_MACHINETYPE.oneofs_by_name['base'].fields.append(
  _MACHINETYPE.fields_by_name['instance_template'])
_MACHINETYPE.fields_by_name['instance_template'].containing_oneof = _MACHINETYPE.oneofs_by_name['base']
_MACHINETYPE.oneofs_by_name['base'].fields.append(
  _MACHINETYPE.fields_by_name['nested_vm'])
_MACHINETYPE.fields_by_name['nested_vm'].containing_oneof = _MACHINETYPE.oneofs_by_name['base']
_ADDRESSPOOL.fields_by_name['fixed_address'].message_type = schema_dot_asset_dot_network__pb2._ADDRESS
_STARTUP.fields_by_name['win_startup'].message_type = schema_dot_common_dot_file__reference__pb2._FILEREFERENCE
_STARTUP.fields_by_name['win_agent_x64'].message_type = schema_dot_common_dot_file__reference__pb2._FILEREFERENCE
_STARTUP.fields_by_name['linux_startup'].message_type = schema_dot_common_dot_file__reference__pb2._FILEREFERENCE
_STARTUP.fields_by_name['linux_agent_x64'].message_type = schema_dot_common_dot_file__reference__pb2._FILEREFERENCE
_RUNTIMESUPPORT.fields_by_name['service_account'].message_type = google_dot_iam_dot_admin_dot_v1_dot_iam__pb2._SERVICEACCOUNT
_RUNTIMESUPPORT.fields_by_name['crypto_key'].message_type = schema_dot_gcp_dot_cloudkms_dot_cloudkms__api__pb2._CRYPTOKEY
_RUNTIMESUPPORT.fields_by_name['startup'].message_type = _STARTUP
_HOSTENVIRONMENT.fields_by_name['project'].message_type = _PROJECT
_HOSTENVIRONMENT.fields_by_name['storage'].message_type = _STORAGE
_HOSTENVIRONMENT.fields_by_name['log_settings'].message_type = _LOGSETTINGS
_HOSTENVIRONMENT.fields_by_name['machine_type'].message_type = _MACHINETYPE
_HOSTENVIRONMENT.fields_by_name['address_pool'].message_type = _ADDRESSPOOL
_HOSTENVIRONMENT.fields_by_name['image'].message_type = _IMAGE
_HOSTENVIRONMENT.fields_by_name['resources'].message_type = _RUNTIMESUPPORT
DESCRIPTOR.message_types_by_name['Project'] = _PROJECT
DESCRIPTOR.message_types_by_name['LogSettings'] = _LOGSETTINGS
DESCRIPTOR.message_types_by_name['Storage'] = _STORAGE
DESCRIPTOR.message_types_by_name['Image'] = _IMAGE
DESCRIPTOR.message_types_by_name['NestedVM'] = _NESTEDVM
DESCRIPTOR.message_types_by_name['MachineType'] = _MACHINETYPE
DESCRIPTOR.message_types_by_name['AddressPool'] = _ADDRESSPOOL
DESCRIPTOR.message_types_by_name['Startup'] = _STARTUP
DESCRIPTOR.message_types_by_name['RuntimeSupport'] = _RUNTIMESUPPORT
DESCRIPTOR.message_types_by_name['HostEnvironment'] = _HOSTENVIRONMENT
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Project = _reflection.GeneratedProtocolMessageType('Project', (_message.Message,), dict(
  DESCRIPTOR = _PROJECT,
  __module__ = 'schema.host.host_environment_pb2'
  # @@protoc_insertion_point(class_scope:host.Project)
  ))
_sym_db.RegisterMessage(Project)

LogSettings = _reflection.GeneratedProtocolMessageType('LogSettings', (_message.Message,), dict(
  DESCRIPTOR = _LOGSETTINGS,
  __module__ = 'schema.host.host_environment_pb2'
  # @@protoc_insertion_point(class_scope:host.LogSettings)
  ))
_sym_db.RegisterMessage(LogSettings)

Storage = _reflection.GeneratedProtocolMessageType('Storage', (_message.Message,), dict(
  DESCRIPTOR = _STORAGE,
  __module__ = 'schema.host.host_environment_pb2'
  # @@protoc_insertion_point(class_scope:host.Storage)
  ))
_sym_db.RegisterMessage(Storage)

Image = _reflection.GeneratedProtocolMessageType('Image', (_message.Message,), dict(

  Family = _reflection.GeneratedProtocolMessageType('Family', (_message.Message,), dict(
    DESCRIPTOR = _IMAGE_FAMILY,
    __module__ = 'schema.host.host_environment_pb2'
    # @@protoc_insertion_point(class_scope:host.Image.Family)
    ))
  ,
  DESCRIPTOR = _IMAGE,
  __module__ = 'schema.host.host_environment_pb2'
  # @@protoc_insertion_point(class_scope:host.Image)
  ))
_sym_db.RegisterMessage(Image)
_sym_db.RegisterMessage(Image.Family)

NestedVM = _reflection.GeneratedProtocolMessageType('NestedVM', (_message.Message,), dict(
  DESCRIPTOR = _NESTEDVM,
  __module__ = 'schema.host.host_environment_pb2'
  # @@protoc_insertion_point(class_scope:host.NestedVM)
  ))
_sym_db.RegisterMessage(NestedVM)

MachineType = _reflection.GeneratedProtocolMessageType('MachineType', (_message.Message,), dict(
  DESCRIPTOR = _MACHINETYPE,
  __module__ = 'schema.host.host_environment_pb2'
  # @@protoc_insertion_point(class_scope:host.MachineType)
  ))
_sym_db.RegisterMessage(MachineType)

AddressPool = _reflection.GeneratedProtocolMessageType('AddressPool', (_message.Message,), dict(
  DESCRIPTOR = _ADDRESSPOOL,
  __module__ = 'schema.host.host_environment_pb2'
  # @@protoc_insertion_point(class_scope:host.AddressPool)
  ))
_sym_db.RegisterMessage(AddressPool)

Startup = _reflection.GeneratedProtocolMessageType('Startup', (_message.Message,), dict(
  DESCRIPTOR = _STARTUP,
  __module__ = 'schema.host.host_environment_pb2'
  # @@protoc_insertion_point(class_scope:host.Startup)
  ))
_sym_db.RegisterMessage(Startup)

RuntimeSupport = _reflection.GeneratedProtocolMessageType('RuntimeSupport', (_message.Message,), dict(
  DESCRIPTOR = _RUNTIMESUPPORT,
  __module__ = 'schema.host.host_environment_pb2'
  # @@protoc_insertion_point(class_scope:host.RuntimeSupport)
  ))
_sym_db.RegisterMessage(RuntimeSupport)

HostEnvironment = _reflection.GeneratedProtocolMessageType('HostEnvironment', (_message.Message,), dict(
  DESCRIPTOR = _HOSTENVIRONMENT,
  __module__ = 'schema.host.host_environment_pb2'
  # @@protoc_insertion_point(class_scope:host.HostEnvironment)
  ))
_sym_db.RegisterMessage(HostEnvironment)


DESCRIPTOR._options = None
_PROJECT.fields_by_name['name']._options = None
_PROJECT.fields_by_name['zone']._options = None
_PROJECT.fields_by_name['project']._options = None
_LOGSETTINGS.fields_by_name['admin_log']._options = None
_STORAGE.fields_by_name['bucket']._options = None
_STORAGE.fields_by_name['created_on']._options = None
_IMAGE_FAMILY.fields_by_name['project']._options = None
_IMAGE_FAMILY.fields_by_name['family']._options = None
_IMAGE.fields_by_name['url']._options = None
_NESTEDVM.fields_by_name['image']._options = None
_MACHINETYPE.fields_by_name['instance_template']._options = None
_STARTUP.fields_by_name['win_startup']._options = None
_STARTUP.fields_by_name['win_agent_x64']._options = None
_STARTUP.fields_by_name['linux_startup']._options = None
_STARTUP.fields_by_name['linux_agent_x64']._options = None
_RUNTIMESUPPORT.fields_by_name['service_account']._options = None
_RUNTIMESUPPORT.fields_by_name['crypto_key']._options = None
_RUNTIMESUPPORT.fields_by_name['startup']._options = None
_HOSTENVIRONMENT.fields_by_name['project']._options = None
_HOSTENVIRONMENT.fields_by_name['storage']._options = None
_HOSTENVIRONMENT.fields_by_name['log_settings']._options = None
_HOSTENVIRONMENT.fields_by_name['machine_type']._options = None
_HOSTENVIRONMENT.fields_by_name['address_pool']._options = None
_HOSTENVIRONMENT.fields_by_name['image']._options = None
_HOSTENVIRONMENT.fields_by_name['resources']._options = None
# @@protoc_insertion_point(module_scope)
