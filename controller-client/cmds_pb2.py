# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: cmds.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\ncmds.proto\"\x91\x01\n\x0e\x43ommandRequest\x12\x0c\n\x04type\x18\x01 \x01(\x03\x12\x33\n\x14\x65\x63ho_command_request\x18\x02 \x01(\x0b\x32\x13.EchoCommandRequestH\x00\x12\x35\n\x15shell_command_request\x18\x03 \x01(\x0b\x32\x14.ShellCommandRequestH\x00\x42\x05\n\x03req\"4\n\x12\x45\x63hoCommandRequest\x12\x10\n\x08\x61gent_id\x18\x01 \x01(\x03\x12\x0c\n\x04text\x18\x02 \x01(\t\"#\n\x13\x45\x63hoCommandResponse\x12\x0c\n\x04text\x18\x01 \x01(\t\"4\n\x13ShellCommandRequest\x12\x10\n\x08\x61gent_id\x18\x01 \x01(\x03\x12\x0b\n\x03\x63md\x18\x02 \x01(\t\"1\n\x14ShellCommandResponse\x12\x0b\n\x03\x65rr\x18\x01 \x01(\t\x12\x0c\n\x04\x64\x61ta\x18\x02 \x01(\x0c\"A\n\x11UploadFileRequest\x12\x10\n\x08\x61gent_id\x18\x01 \x01(\x03\x12\x0c\n\x04path\x18\x02 \x01(\t\x12\x0c\n\x04\x64\x61ta\x18\x03 \x01(\x0c\"!\n\x12UploadFileResponse\x12\x0b\n\x03\x65rr\x18\x01 \x01(\t\"5\n\x13\x44ownloadFileRequest\x12\x10\n\x08\x61gent_id\x18\x01 \x01(\x03\x12\x0c\n\x04path\x18\x02 \x01(\t\"1\n\x14\x44ownloadFileResponse\x12\x0b\n\x03\x65rr\x18\x01 \x01(\t\x12\x0c\n\x04\x64\x61ta\x18\x02 \x01(\x0c\"<\n\tAgentInfo\x12\n\n\x02id\x18\x01 \x01(\x03\x12\r\n\x05\x61live\x18\x02 \x01(\x08\x12\x14\n\x0c\x63onnect_time\x18\x03 \x01(\t\"\x07\n\x05\x45mpty2\xac\x02\n\x0c\x41gentManager\x12=\n\x0eRunEchoCommand\x12\x13.EchoCommandRequest\x1a\x14.EchoCommandResponse\"\x00\x12@\n\x0fRunShellCommand\x12\x14.ShellCommandRequest\x1a\x15.ShellCommandResponse\"\x00\x12\x37\n\nUploadFile\x12\x12.UploadFileRequest\x1a\x13.UploadFileResponse\"\x00\x12=\n\x0c\x44ownloadFile\x12\x14.DownloadFileRequest\x1a\x15.DownloadFileResponse\"\x00\x12#\n\tGetAgents\x12\x06.Empty\x1a\n.AgentInfo\"\x00\x30\x01\x42\x08Z\x06./cmdsb\x06proto3')



_COMMANDREQUEST = DESCRIPTOR.message_types_by_name['CommandRequest']
_ECHOCOMMANDREQUEST = DESCRIPTOR.message_types_by_name['EchoCommandRequest']
_ECHOCOMMANDRESPONSE = DESCRIPTOR.message_types_by_name['EchoCommandResponse']
_SHELLCOMMANDREQUEST = DESCRIPTOR.message_types_by_name['ShellCommandRequest']
_SHELLCOMMANDRESPONSE = DESCRIPTOR.message_types_by_name['ShellCommandResponse']
_UPLOADFILEREQUEST = DESCRIPTOR.message_types_by_name['UploadFileRequest']
_UPLOADFILERESPONSE = DESCRIPTOR.message_types_by_name['UploadFileResponse']
_DOWNLOADFILEREQUEST = DESCRIPTOR.message_types_by_name['DownloadFileRequest']
_DOWNLOADFILERESPONSE = DESCRIPTOR.message_types_by_name['DownloadFileResponse']
_AGENTINFO = DESCRIPTOR.message_types_by_name['AgentInfo']
_EMPTY = DESCRIPTOR.message_types_by_name['Empty']
CommandRequest = _reflection.GeneratedProtocolMessageType('CommandRequest', (_message.Message,), {
  'DESCRIPTOR' : _COMMANDREQUEST,
  '__module__' : 'cmds_pb2'
  # @@protoc_insertion_point(class_scope:CommandRequest)
  })
_sym_db.RegisterMessage(CommandRequest)

EchoCommandRequest = _reflection.GeneratedProtocolMessageType('EchoCommandRequest', (_message.Message,), {
  'DESCRIPTOR' : _ECHOCOMMANDREQUEST,
  '__module__' : 'cmds_pb2'
  # @@protoc_insertion_point(class_scope:EchoCommandRequest)
  })
_sym_db.RegisterMessage(EchoCommandRequest)

EchoCommandResponse = _reflection.GeneratedProtocolMessageType('EchoCommandResponse', (_message.Message,), {
  'DESCRIPTOR' : _ECHOCOMMANDRESPONSE,
  '__module__' : 'cmds_pb2'
  # @@protoc_insertion_point(class_scope:EchoCommandResponse)
  })
_sym_db.RegisterMessage(EchoCommandResponse)

ShellCommandRequest = _reflection.GeneratedProtocolMessageType('ShellCommandRequest', (_message.Message,), {
  'DESCRIPTOR' : _SHELLCOMMANDREQUEST,
  '__module__' : 'cmds_pb2'
  # @@protoc_insertion_point(class_scope:ShellCommandRequest)
  })
_sym_db.RegisterMessage(ShellCommandRequest)

ShellCommandResponse = _reflection.GeneratedProtocolMessageType('ShellCommandResponse', (_message.Message,), {
  'DESCRIPTOR' : _SHELLCOMMANDRESPONSE,
  '__module__' : 'cmds_pb2'
  # @@protoc_insertion_point(class_scope:ShellCommandResponse)
  })
_sym_db.RegisterMessage(ShellCommandResponse)

UploadFileRequest = _reflection.GeneratedProtocolMessageType('UploadFileRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPLOADFILEREQUEST,
  '__module__' : 'cmds_pb2'
  # @@protoc_insertion_point(class_scope:UploadFileRequest)
  })
_sym_db.RegisterMessage(UploadFileRequest)

UploadFileResponse = _reflection.GeneratedProtocolMessageType('UploadFileResponse', (_message.Message,), {
  'DESCRIPTOR' : _UPLOADFILERESPONSE,
  '__module__' : 'cmds_pb2'
  # @@protoc_insertion_point(class_scope:UploadFileResponse)
  })
_sym_db.RegisterMessage(UploadFileResponse)

DownloadFileRequest = _reflection.GeneratedProtocolMessageType('DownloadFileRequest', (_message.Message,), {
  'DESCRIPTOR' : _DOWNLOADFILEREQUEST,
  '__module__' : 'cmds_pb2'
  # @@protoc_insertion_point(class_scope:DownloadFileRequest)
  })
_sym_db.RegisterMessage(DownloadFileRequest)

DownloadFileResponse = _reflection.GeneratedProtocolMessageType('DownloadFileResponse', (_message.Message,), {
  'DESCRIPTOR' : _DOWNLOADFILERESPONSE,
  '__module__' : 'cmds_pb2'
  # @@protoc_insertion_point(class_scope:DownloadFileResponse)
  })
_sym_db.RegisterMessage(DownloadFileResponse)

AgentInfo = _reflection.GeneratedProtocolMessageType('AgentInfo', (_message.Message,), {
  'DESCRIPTOR' : _AGENTINFO,
  '__module__' : 'cmds_pb2'
  # @@protoc_insertion_point(class_scope:AgentInfo)
  })
_sym_db.RegisterMessage(AgentInfo)

Empty = _reflection.GeneratedProtocolMessageType('Empty', (_message.Message,), {
  'DESCRIPTOR' : _EMPTY,
  '__module__' : 'cmds_pb2'
  # @@protoc_insertion_point(class_scope:Empty)
  })
_sym_db.RegisterMessage(Empty)

_AGENTMANAGER = DESCRIPTOR.services_by_name['AgentManager']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\006./cmds'
  _COMMANDREQUEST._serialized_start=15
  _COMMANDREQUEST._serialized_end=160
  _ECHOCOMMANDREQUEST._serialized_start=162
  _ECHOCOMMANDREQUEST._serialized_end=214
  _ECHOCOMMANDRESPONSE._serialized_start=216
  _ECHOCOMMANDRESPONSE._serialized_end=251
  _SHELLCOMMANDREQUEST._serialized_start=253
  _SHELLCOMMANDREQUEST._serialized_end=305
  _SHELLCOMMANDRESPONSE._serialized_start=307
  _SHELLCOMMANDRESPONSE._serialized_end=356
  _UPLOADFILEREQUEST._serialized_start=358
  _UPLOADFILEREQUEST._serialized_end=423
  _UPLOADFILERESPONSE._serialized_start=425
  _UPLOADFILERESPONSE._serialized_end=458
  _DOWNLOADFILEREQUEST._serialized_start=460
  _DOWNLOADFILEREQUEST._serialized_end=513
  _DOWNLOADFILERESPONSE._serialized_start=515
  _DOWNLOADFILERESPONSE._serialized_end=564
  _AGENTINFO._serialized_start=566
  _AGENTINFO._serialized_end=626
  _EMPTY._serialized_start=628
  _EMPTY._serialized_end=635
  _AGENTMANAGER._serialized_start=638
  _AGENTMANAGER._serialized_end=938
# @@protoc_insertion_point(module_scope)
