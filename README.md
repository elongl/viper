# Viper

Remote control software using gRPC and Go.

## Components

- Agent - Runs on the endpoint.
- Controller - The server that the _agents_ connect to.
- Agent Manager - The server that runs alongside the _controller_ to manage the _agents_.
- Controller Client - Talks to the _agent manager_.
