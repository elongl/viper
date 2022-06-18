class ControllerError(Exception):
    """Base class for all controller errors"""


class CommandError(ControllerError):
    """Base class for all command errors"""


class ShellCommandError(CommandError):
    def __init__(self, cmd: str, err_msg: str):
        self.cmd = cmd
        self.err_msg = err_msg
        super().__init__(f'Failed to execute shell command "{cmd}": {err_msg}')
