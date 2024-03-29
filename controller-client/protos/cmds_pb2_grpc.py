# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from protos import cmds_pb2 as protos_dot_cmds__pb2


class AgentManagerStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.RunEchoCommand = channel.unary_unary(
                '/AgentManager/RunEchoCommand',
                request_serializer=protos_dot_cmds__pb2.EchoCommandRequest.SerializeToString,
                response_deserializer=protos_dot_cmds__pb2.EchoCommandResponse.FromString,
                )
        self.RunShellCommand = channel.unary_unary(
                '/AgentManager/RunShellCommand',
                request_serializer=protos_dot_cmds__pb2.ShellCommandRequest.SerializeToString,
                response_deserializer=protos_dot_cmds__pb2.ShellCommandResponse.FromString,
                )
        self.DownloadFile = channel.unary_unary(
                '/AgentManager/DownloadFile',
                request_serializer=protos_dot_cmds__pb2.DownloadFileRequest.SerializeToString,
                response_deserializer=protos_dot_cmds__pb2.DownloadFileResponse.FromString,
                )
        self.UploadFile = channel.unary_unary(
                '/AgentManager/UploadFile',
                request_serializer=protos_dot_cmds__pb2.UploadFileRequest.SerializeToString,
                response_deserializer=protos_dot_cmds__pb2.UploadFileResponse.FromString,
                )
        self.Screenshot = channel.unary_unary(
                '/AgentManager/Screenshot',
                request_serializer=protos_dot_cmds__pb2.ScreenshotRequest.SerializeToString,
                response_deserializer=protos_dot_cmds__pb2.ScreenshotResponse.FromString,
                )
        self.GetAgents = channel.unary_stream(
                '/AgentManager/GetAgents',
                request_serializer=protos_dot_cmds__pb2.GetAgentsRequest.SerializeToString,
                response_deserializer=protos_dot_cmds__pb2.AgentInfo.FromString,
                )
        self.StartSocksServer = channel.unary_unary(
                '/AgentManager/StartSocksServer',
                request_serializer=protos_dot_cmds__pb2.StartSocksServerRequest.SerializeToString,
                response_deserializer=protos_dot_cmds__pb2.StartSocksServerResponse.FromString,
                )
        self.StopSocksServer = channel.unary_unary(
                '/AgentManager/StopSocksServer',
                request_serializer=protos_dot_cmds__pb2.StopSocksServerRequest.SerializeToString,
                response_deserializer=protos_dot_cmds__pb2.StopSocksServerResponse.FromString,
                )


class AgentManagerServicer(object):
    """Missing associated documentation comment in .proto file."""

    def RunEchoCommand(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RunShellCommand(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DownloadFile(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UploadFile(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Screenshot(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetAgents(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def StartSocksServer(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def StopSocksServer(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AgentManagerServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'RunEchoCommand': grpc.unary_unary_rpc_method_handler(
                    servicer.RunEchoCommand,
                    request_deserializer=protos_dot_cmds__pb2.EchoCommandRequest.FromString,
                    response_serializer=protos_dot_cmds__pb2.EchoCommandResponse.SerializeToString,
            ),
            'RunShellCommand': grpc.unary_unary_rpc_method_handler(
                    servicer.RunShellCommand,
                    request_deserializer=protos_dot_cmds__pb2.ShellCommandRequest.FromString,
                    response_serializer=protos_dot_cmds__pb2.ShellCommandResponse.SerializeToString,
            ),
            'DownloadFile': grpc.unary_unary_rpc_method_handler(
                    servicer.DownloadFile,
                    request_deserializer=protos_dot_cmds__pb2.DownloadFileRequest.FromString,
                    response_serializer=protos_dot_cmds__pb2.DownloadFileResponse.SerializeToString,
            ),
            'UploadFile': grpc.unary_unary_rpc_method_handler(
                    servicer.UploadFile,
                    request_deserializer=protos_dot_cmds__pb2.UploadFileRequest.FromString,
                    response_serializer=protos_dot_cmds__pb2.UploadFileResponse.SerializeToString,
            ),
            'Screenshot': grpc.unary_unary_rpc_method_handler(
                    servicer.Screenshot,
                    request_deserializer=protos_dot_cmds__pb2.ScreenshotRequest.FromString,
                    response_serializer=protos_dot_cmds__pb2.ScreenshotResponse.SerializeToString,
            ),
            'GetAgents': grpc.unary_stream_rpc_method_handler(
                    servicer.GetAgents,
                    request_deserializer=protos_dot_cmds__pb2.GetAgentsRequest.FromString,
                    response_serializer=protos_dot_cmds__pb2.AgentInfo.SerializeToString,
            ),
            'StartSocksServer': grpc.unary_unary_rpc_method_handler(
                    servicer.StartSocksServer,
                    request_deserializer=protos_dot_cmds__pb2.StartSocksServerRequest.FromString,
                    response_serializer=protos_dot_cmds__pb2.StartSocksServerResponse.SerializeToString,
            ),
            'StopSocksServer': grpc.unary_unary_rpc_method_handler(
                    servicer.StopSocksServer,
                    request_deserializer=protos_dot_cmds__pb2.StopSocksServerRequest.FromString,
                    response_serializer=protos_dot_cmds__pb2.StopSocksServerResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'AgentManager', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class AgentManager(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def RunEchoCommand(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/AgentManager/RunEchoCommand',
            protos_dot_cmds__pb2.EchoCommandRequest.SerializeToString,
            protos_dot_cmds__pb2.EchoCommandResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def RunShellCommand(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/AgentManager/RunShellCommand',
            protos_dot_cmds__pb2.ShellCommandRequest.SerializeToString,
            protos_dot_cmds__pb2.ShellCommandResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DownloadFile(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/AgentManager/DownloadFile',
            protos_dot_cmds__pb2.DownloadFileRequest.SerializeToString,
            protos_dot_cmds__pb2.DownloadFileResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UploadFile(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/AgentManager/UploadFile',
            protos_dot_cmds__pb2.UploadFileRequest.SerializeToString,
            protos_dot_cmds__pb2.UploadFileResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Screenshot(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/AgentManager/Screenshot',
            protos_dot_cmds__pb2.ScreenshotRequest.SerializeToString,
            protos_dot_cmds__pb2.ScreenshotResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetAgents(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/AgentManager/GetAgents',
            protos_dot_cmds__pb2.GetAgentsRequest.SerializeToString,
            protos_dot_cmds__pb2.AgentInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def StartSocksServer(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/AgentManager/StartSocksServer',
            protos_dot_cmds__pb2.StartSocksServerRequest.SerializeToString,
            protos_dot_cmds__pb2.StartSocksServerResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def StopSocksServer(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/AgentManager/StopSocksServer',
            protos_dot_cmds__pb2.StopSocksServerRequest.SerializeToString,
            protos_dot_cmds__pb2.StopSocksServerResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
