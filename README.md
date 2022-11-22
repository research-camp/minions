<p align="center">
    <img src="assets/logo.jpeg" width="400" alt="logo" />
</p>


<p align="center">
    SSH tunneling and SSH client implemented in Golang..
</p>

<p align="center">
    <img src="https://img.shields.io/badge/Go-1.18+-00ADD8?style=for-the-badge&logo=go" alt="go version" />
    <img src="https://img.shields.io/badge/Version-1.1.0-informational?style=for-the-badge&logo=github" alt="version" />
</p>

## What is Xerox?

**Xerox** assassin is a Golang library which helps you to make connection
to a remove machine, using SSH client or SSH tunnel.

## Install package

Import package by using the following command:

```shell
go get -v github.com/amirhnajafiz/xerox@latest
```

Check to see if the package is installed:

```go
import (
    _ "github.com/amirhnajafiz/xerox"
)
```

## Using Xerox to create SSH tunnel

```go
// creating the local, server, and remote servers
localEndpoint := &xerox.Endpoint{
    Host: "localhost",
    Port: 9000,
}
serverEndpoint := &xerox.Endpoint{
    Host: "example.com",
    Port: 22,
}
remoteEndpoint := &xerox.Endpoint{
    Host: "localhost",
    Port: 8080,
}
// creating ssh client config
sshConfig := &ssh.ClientConfig{
    User: "vcap",
    Auth: []ssh.AuthMethod{
        xerox.SSHAgent(os.Getenv("SSH_AUTH")),
    },
}
// creating xerox ssh tunnel
sshTunnel := &xerox.SSHTunnel{
    Local:  localEndpoint,
    Server: serverEndpoint,
    Remote: remoteEndpoint,
    Config: sshConfig,
}

log.Printf("ssh tunnel start...")

// starting ssh tunnel
if err := sshTunnel.Start(); err != nil {
    panic(err)
}
```

## Using Xerox to create SSH Client

```go
// creating ssh config
sshConfig := &ssh.ClientConfig{
    User: "root",
    Auth: []ssh.AuthMethod{
        xerox.PublicKeyFile("./.private-key"),
    },
}

// creating ssh client
client := &xerox.SSHClient{
    Config: sshConfig,
    Server: &xerox.Endpoint{
        Host: "129.0.4.22",
        Port: 80,
    },
    TerminalConfig: &xerox.SSHTerminal{
        Echo:             0,
        TtyOpInputSpeed:  14400,
        TtyOpOutputSpeed: 14400,
        Rows:             80,
        Columns:          40,
    },
}

// connecting to remove machine
if err := client.Connect(); err != nil {
    panic(err)
}
```

Executing commands:

```go
// generating our command
cmd := &xerox.SSHCommand{
    Path:   "ls -l $LC_DIR",
    Env:    []string{"LC_DIR=/usr"},
    Stdin:  os.Stdin,
    Stdout: os.Stdout,
    Stderr: os.Stderr,
}

// running command
if err := client.RunCommand(cmd); err != nil {
    panic(err)
}
```