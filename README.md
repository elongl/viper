# Viper

Remote control software using gRPC and Go.

## Components

- Agent - Runs on the endpoint.
- Controller - The server that the _agents_ connect to.
- Agent Manager - The server that runs alongside the _controller_ to manage the _agents_.
- Controller Client - Talks to the _agent manager_.

## Features

- Cross-platform Support - Viper currently runs on all modern operating systems (Windows, Linux, macOS, etc).
- Authentication & Encryption - Using TLS, agent connections are encrypted and verified using client certificates.
- Persistence - The agent keeps connection to the controller even if rebooted or disconnected momentarily.
- Shell - Execute shell commands.
- File I/O - Download and upload files.
- Screenshots - Take a screenshots.
- SOCKS - Connect into the agents' network.

## Usage

1. Update the `config.json` according to your needs.
2. Run the controller servers: `go run controller/main.go`.
3. Build and run the agent: `make build_agent` or `go run agent/main.go`.
4. Control the agents: `ipython -i controller-client/main.py`.

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
