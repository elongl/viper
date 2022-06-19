from dataclasses import dataclass
from typing import Union

import grpc

import cmds_pb2
import cmds_pb2_grpc
import errors


@dataclass
class Agent:
    id: int
    _stub: cmds_pb2_grpc.AgentManagerStub

    def echo(self, data: str) -> str:
        req = cmds_pb2.EchoCommandRequest(agent_id=self.id, data=data)
        resp = self._stub.RunEchoCommand(req)
        return resp.data

    def shell(self, cmd: str, decode: bool = True) -> bytes:
        req = cmds_pb2.ShellCommandRequest(agent_id=self.id, cmd=cmd)
        resp = self._stub.RunShellCommand(req)
        if resp.err:
            raise errors.ShellCommandError(cmd, resp.err)
        return resp.data.decode() if decode else resp.data

    def download_file(self, remote_path: str, local_path: str = None) -> Union[bytes, None]:
        req = cmds_pb2.DownloadFileRequest(agent_id=self.id, path=remote_path)
        resp = self._stub.DownloadFile(req)
        if resp.err:
            raise errors.DownloadFileError(remote_path, resp.err)
        if local_path:
            with open(local_path, 'wb') as f:
                f.write(resp.data)
        else:
            return resp.data

    def upload_file(self, local_path: str, remote_path: str) -> None:
        with open(local_path, 'rb') as f:
            data = f.read()
        req = cmds_pb2.UploadFileRequest(
            agent_id=self.id, path=remote_path, data=data)
        resp = self._stub.UploadFile(req)
        if resp.err:
            raise errors.UploadFileError(remote_path, resp.err)

    def __repr__(self) -> str:
        return f'Agent(id={self.id})'


@dataclass
class ControllerClient:
    addr: str
    _stub: cmds_pb2_grpc.AgentManagerStub = None

    def connect(self) -> None:
        channel = grpc.insecure_channel(self.addr)
        self._stub = cmds_pb2_grpc.AgentManagerStub(channel)

    def get_agent(self, agent_id: int) -> Agent:
        return Agent(agent_id, self._stub)

    def get_agents(self) -> [cmds_pb2.AgentInfo]:
        return self._stub.GetAgents(cmds_pb2.Empty())


if __name__ == '__main__':
    cnc = ControllerClient('localhost:8159')
    cnc.connect()
