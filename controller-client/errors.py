class ControllerError(Exception):
    """Base class for all controller errors"""


class CommandError(ControllerError):
    """Base class for all command errors"""


class ShellCommandError(CommandError):
    def __init__(self, cmd: str, err: str, data: bytes):
        self.cmd = cmd
        self.err = err
        self.data = data
        super().__init__(
            f'Failed to execute shell command "{cmd}": {err} : {data}')


class ScreenshotError(CommandError):
    def __init__(self, err: str):
        self.err = err
        super().__init__(f'Failed to take screenshot: {err}')


class DownloadFileError(CommandError):
    def __init__(self, remote_path: str, err: str):
        self.remote_path = remote_path
        self.err = err
        super().__init__(f'Failed to download file "{remote_path}": {err}')


class UploadFileError(CommandError):
    def __init__(self, remote_path: str, err: str):
        self.remote_path = remote_path
        self.err = err
        super().__init__(f'Failed to upload file "{remote_path}": {err}')
