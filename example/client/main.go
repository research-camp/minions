package main

import (
	"os"

	"github.com/amirhnajafiz/xerox"

	"golang.org/x/crypto/ssh"
)

// example of creating an ssh client and running commands to our remote machine.
func main() {
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
}
