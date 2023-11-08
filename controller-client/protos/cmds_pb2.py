# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/cmds.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x11protos/cmds.proto\"\x9d\x03\n\x0e\x43ommandRequest\x12\x33\n\x14\x65\x63ho_command_request\x18\x01 \x01(\x0b\x32\x13.EchoCommandRequestH\x00\x12\x35\n\x15shell_command_request\x18\x02 \x01(\x0b\x32\x14.ShellCommandRequestH\x00\x12\x35\n\x15\x64ownload_file_request\x18\x03 \x01(\x0b\x32\x14.DownloadFileRequestH\x00\x12\x31\n\x13upload_file_request\x18\x04 \x01(\x0b\x32\x12.UploadFileRequestH\x00\x12\x30\n\x12screenshot_request\x18\x05 \x01(\x0b\x32\x12.ScreenshotRequestH\x00\x12>\n\x1astart_socks_server_request\x18\x06 \x01(\x0b\x32\x18.StartSocksServerRequestH\x00\x12<\n\x19stop_socks_server_request\x18\x07 \x01(\x0b\x32\x17.StopSocksServerRequestH\x00\x42\x05\n\x03req\"4\n\x12\x45\x63hoCommandRequest\x12\x10\n\x08\x61gent_id\x18\x01 \x01(\x03\x12\x0c\n\x04\x64\x61ta\x18\x02 \x01(\t\"#\n\x13\x45\x63hoCommandResponse\x12\x0c\n\x04\x64\x61ta\x18\x01 \x01(\t\"4\n\x13ShellCommandRequest\x12\x10\n\x08\x61gent_id\x18\x01 \x01(\x03\x12\x0b\n\x03\x63md\x18\x02 \x01(\t\"1\n\x14ShellCommandResponse\x12\x0b\n\x03\x65rr\x18\x01 \x01(\t\x12\x0c\n\x04\x64\x61ta\x18\x02 \x01(\x0c\"5\n\x13\x44ownloadFileRequest\x12\x10\n\x08\x61gent_id\x18\x01 \x01(\x03\x12\x0c\n\x04path\x18\x02 \x01(\t\"1\n\x14\x44ownloadFileResponse\x12\x0b\n\x03\x65rr\x18\x01 \x01(\t\x12\x0c\n\x04\x64\x61ta\x18\x02 \x01(\x0c\"A\n\x11UploadFileRequest\x12\x10\n\x08\x61gent_id\x18\x01 \x01(\x03\x12\x0c\n\x04path\x18\x02 \x01(\t\x12\x0c\n\x04\x64\x61ta\x18\x03 \x01(\x0c\"!\n\x12UploadFileResponse\x12\x0b\n\x03\x65rr\x18\x01 \x01(\t\"%\n\x11ScreenshotRequest\x12\x10\n\x08\x61gent_id\x18\x01 \x01(\x03\"/\n\x12ScreenshotResponse\x12\x0b\n\x03\x65rr\x18\x01 \x01(\t\x12\x0c\n\x04\x64\x61ta\x18\x02 \x01(\x0c\"&\n\x10GetAgentsRequest\x12\x12\n\nalive_only\x18\x01 \x01(\x08\"+\n\x17StartSocksServerRequest\x12\x10\n\x08\x61gent_id\x18\x01 \x01(\x03\"(\n\x18StartSocksServerResponse\x12\x0c\n\x04\x61\x64\x64r\x18\x01 \x01(\t\"*\n\x16StopSocksServerRequest\x12\x10\n\x08\x61gent_id\x18\x01 \x01(\x03\"\x19\n\x17StopSocksServerResponse\"<\n\tAgentInfo\x12\n\n\x02id\x18\x01 \x01(\x03\x12\r\n\x05\x61live\x18\x02 \x01(\x08\x12\x14\n\x0c\x63onnect_time\x18\x03 \x01(\t2\x83\x04\n\x0c\x41gentManager\x12=\n\x0eRunEchoCommand\x12\x13.EchoCommandRequest\x1a\x14.EchoCommandResponse\"\x00\x12@\n\x0fRunShellCommand\x12\x14.ShellCommandRequest\x1a\x15.ShellCommandResponse\"\x00\x12=\n\x0c\x44ownloadFile\x12\x14.DownloadFileRequest\x1a\x15.DownloadFileResponse\"\x00\x12\x37\n\nUploadFile\x12\x12.UploadFileRequest\x1a\x13.UploadFileResponse\"\x00\x12\x37\n\nScreenshot\x12\x12.ScreenshotRequest\x1a\x13.ScreenshotResponse\"\x00\x12.\n\tGetAgents\x12\x11.GetAgentsRequest\x1a\n.AgentInfo\"\x00\x30\x01\x12I\n\x10StartSocksServer\x12\x18.StartSocksServerRequest\x1a\x19.StartSocksServerResponse\"\x00\x12\x46\n\x0fStopSocksServer\x12\x17.StopSocksServerRequest\x1a\x18.StopSocksServerResponse\"\x00\x42\x08Z\x06./cmdsb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'protos.cmds_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\006./cmds'
  _globals['_COMMANDREQUEST']._serialized_start=22
  _globals['_COMMANDREQUEST']._serialized_end=435
  _globals['_ECHOCOMMANDREQUEST']._serialized_start=437
  _globals['_ECHOCOMMANDREQUEST']._serialized_end=489
  _globals['_ECHOCOMMANDRESPONSE']._serialized_start=491
  _globals['_ECHOCOMMANDRESPONSE']._serialized_end=526
  _globals['_SHELLCOMMANDREQUEST']._serialized_start=528
  _globals['_SHELLCOMMANDREQUEST']._serialized_end=580
  _globals['_SHELLCOMMANDRESPONSE']._serialized_start=582
  _globals['_SHELLCOMMANDRESPONSE']._serialized_end=631
  _globals['_DOWNLOADFILEREQUEST']._serialized_start=633
  _globals['_DOWNLOADFILEREQUEST']._serialized_end=686
  _globals['_DOWNLOADFILERESPONSE']._serialized_start=688
  _globals['_DOWNLOADFILERESPONSE']._serialized_end=737
  _globals['_UPLOADFILEREQUEST']._serialized_start=739
  _globals['_UPLOADFILEREQUEST']._serialized_end=804
  _globals['_UPLOADFILERESPONSE']._serialized_start=806
  _globals['_UPLOADFILERESPONSE']._serialized_end=839
  _globals['_SCREENSHOTREQUEST']._serialized_start=841
  _globals['_SCREENSHOTREQUEST']._serialized_end=878
  _globals['_SCREENSHOTRESPONSE']._serialized_start=880
  _globals['_SCREENSHOTRESPONSE']._serialized_end=927
  _globals['_GETAGENTSREQUEST']._serialized_start=929
  _globals['_GETAGENTSREQUEST']._serialized_end=967
  _globals['_STARTSOCKSSERVERREQUEST']._serialized_start=969
  _globals['_STARTSOCKSSERVERREQUEST']._serialized_end=1012
  _globals['_STARTSOCKSSERVERRESPONSE']._serialized_start=1014
  _globals['_STARTSOCKSSERVERRESPONSE']._serialized_end=1054
  _globals['_STOPSOCKSSERVERREQUEST']._serialized_start=1056
  _globals['_STOPSOCKSSERVERREQUEST']._serialized_end=1098
  _globals['_STOPSOCKSSERVERRESPONSE']._serialized_start=1100
  _globals['_STOPSOCKSSERVERRESPONSE']._serialized_end=1125
  _globals['_AGENTINFO']._serialized_start=1127
  _globals['_AGENTINFO']._serialized_end=1187
  _globals['_AGENTMANAGER']._serialized_start=1190
  _globals['_AGENTMANAGER']._serialized_end=1705
# @@protoc_insertion_point(module_scope)