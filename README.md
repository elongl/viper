# Viper

Remote control software using gRPC and Go.

## Components

- **Agent**: Runs on the endpoint.
- **Controller**: The server that the _agents_ connect to.
- **Agent Manager**: The server that runs alongside the _controller_ to manage the _agents_.
- **Controller Client**: Talks to the _agent manager_.

## Features

- **Cross-platform Support**: Runs on all modern operating systems (Windows, Linux, macOS, etc).
- **Authentication & Encryption**: Connections are encrypted and verified using client certificates.
- **Persistence**: Keeps connection to the controller even if rebooted or disconnected momentarily.
- **Shell**: Execute shell commands.
- **File I/O**: Download and upload files.
- **Screenshots**: Capture screenshots.
- **Network Proxy**: Connect into the agent's network using [SOCKS](https://en.wikipedia.org/wiki/SOCKS).

## Build

1. Update the `config.json` according to your needs.
2. Run `make` to build the executables in `bin`.

## Usage

1. Run `controller` to accept agent connections.
2. Run the controller client to manage the agents: `ipython -i controller-client/main.py`.

```py
cnc.get_agents()
> [
    id: 0
    alive: true
    connect_time: "2022-06-19 15:22:13.17828 +0300 IDT m=+61.594837928"
]

agent = cnc.get_agent(agent_id=0)
agent.shell('whoami')
> 'root'
```
