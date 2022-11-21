package main

import (
	"log"
	"os"

	"github.com/amirhnajafiz/xerox"

	"golang.org/x/crypto/ssh"
)

// example of using xerox tunneling module
// to create 3 servers and execute the port forwarding operation.
func main() {
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
}
