# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/location.building.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from proto.location import street_pb2 as proto_dot_location_dot_street__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/location.building.proto',
  package='infrastructure',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x1dproto/location.building.proto\x12\x0einfrastructure\x1a\x1bproto/location.street.proto\"I\n\x08\x42uilding\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0e\n\x06number\x18\x02 \x01(\x05\x12\x1f\n\x06street\x18\x03 \x01(\x0b\x32\x0f.council.Streetb\x06proto3')
  ,
  dependencies=[proto_dot_location_dot_street__pb2.DESCRIPTOR,])




_BUILDING = _descriptor.Descriptor(
  name='Building',
  full_name='infrastructure.Building',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='infrastructure.Building.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='number', full_name='infrastructure.Building.number', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='street', full_name='infrastructure.Building.street', index=2,
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
  serialized_start=78,
  serialized_end=151,
)

_BUILDING.fields_by_name['street'].message_type = proto_dot_location_dot_street__pb2._STREET
DESCRIPTOR.message_types_by_name['Building'] = _BUILDING
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Building = _reflection.GeneratedProtocolMessageType('Building', (_message.Message,), {
  'DESCRIPTOR' : _BUILDING,
  '__module__' : 'proto.location.building_pb2'
  # @@protoc_insertion_point(class_scope:infrastructure.Building)
  })
_sym_db.RegisterMessage(Building)


# @@protoc_insertion_point(module_scope)
