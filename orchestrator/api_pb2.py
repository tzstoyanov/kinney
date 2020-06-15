# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: orchestrator/api.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='orchestrator/api.proto',
  package='kinney',
  syntax='proto3',
  serialized_options=b'Z*github.com/CamusEnergy/kinney/orchestrator',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x16orchestrator/api.proto\x12\x06kinney\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xc1\x01\n\x0e\x43hargerSession\x12\r\n\x05point\x18\x01 \x01(\t\x12\x0f\n\x07vehicle\x18\x02 \x01(\t\x12\r\n\x05watts\x18\x03 \x01(\x01\x12,\n\x08measured\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12)\n\x05start\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\'\n\x03\x65nd\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"[\n\x0e\x43hargerCommand\x12\r\n\x05point\x18\x01 \x01(\t\x12\r\n\x05limit\x18\x02 \x01(\x01\x12+\n\x08lifetime\x18\x03 \x01(\x0b\x32\x19.google.protobuf.Duration2O\n\x0cOrchestrator\x12?\n\x07\x43harger\x12\x16.kinney.ChargerSession\x1a\x16.kinney.ChargerCommand\"\x00(\x01\x30\x01\x42,Z*github.com/CamusEnergy/kinney/orchestratorb\x06proto3'
  ,
  dependencies=[google_dot_protobuf_dot_duration__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])




_CHARGERSESSION = _descriptor.Descriptor(
  name='ChargerSession',
  full_name='kinney.ChargerSession',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='point', full_name='kinney.ChargerSession.point', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='vehicle', full_name='kinney.ChargerSession.vehicle', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='watts', full_name='kinney.ChargerSession.watts', index=2,
      number=3, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='measured', full_name='kinney.ChargerSession.measured', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='start', full_name='kinney.ChargerSession.start', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='end', full_name='kinney.ChargerSession.end', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=100,
  serialized_end=293,
)


_CHARGERCOMMAND = _descriptor.Descriptor(
  name='ChargerCommand',
  full_name='kinney.ChargerCommand',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='point', full_name='kinney.ChargerCommand.point', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='limit', full_name='kinney.ChargerCommand.limit', index=1,
      number=2, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='lifetime', full_name='kinney.ChargerCommand.lifetime', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=295,
  serialized_end=386,
)

_CHARGERSESSION.fields_by_name['measured'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_CHARGERSESSION.fields_by_name['start'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_CHARGERSESSION.fields_by_name['end'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_CHARGERCOMMAND.fields_by_name['lifetime'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
DESCRIPTOR.message_types_by_name['ChargerSession'] = _CHARGERSESSION
DESCRIPTOR.message_types_by_name['ChargerCommand'] = _CHARGERCOMMAND
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ChargerSession = _reflection.GeneratedProtocolMessageType('ChargerSession', (_message.Message,), {
  'DESCRIPTOR' : _CHARGERSESSION,
  '__module__' : 'orchestrator.api_pb2'
  # @@protoc_insertion_point(class_scope:kinney.ChargerSession)
  })
_sym_db.RegisterMessage(ChargerSession)

ChargerCommand = _reflection.GeneratedProtocolMessageType('ChargerCommand', (_message.Message,), {
  'DESCRIPTOR' : _CHARGERCOMMAND,
  '__module__' : 'orchestrator.api_pb2'
  # @@protoc_insertion_point(class_scope:kinney.ChargerCommand)
  })
_sym_db.RegisterMessage(ChargerCommand)


DESCRIPTOR._options = None

_ORCHESTRATOR = _descriptor.ServiceDescriptor(
  name='Orchestrator',
  full_name='kinney.Orchestrator',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=388,
  serialized_end=467,
  methods=[
  _descriptor.MethodDescriptor(
    name='Charger',
    full_name='kinney.Orchestrator.Charger',
    index=0,
    containing_service=None,
    input_type=_CHARGERSESSION,
    output_type=_CHARGERCOMMAND,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_ORCHESTRATOR)

DESCRIPTOR.services_by_name['Orchestrator'] = _ORCHESTRATOR

# @@protoc_insertion_point(module_scope)