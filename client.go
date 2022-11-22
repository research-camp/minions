package xerox

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

// SSHClient
// using SSH client to run a shell command on a remote machine.
// Every SSH connection requires an ssh.ClientConfig object that defines configuration options such as authentication.
// Session is one of the parameters that acts as an entry point to the remote terminal.
type SSHClient struct {
	Server  *Endpoint
	Config  *ssh.ClientConfig
	Session *ssh.Session
}

// Connect
// opening a new connection to remote machine and creating a new session.
func (client *SSHClient) Connect() error {
	// opening connection to remove machine
	connection, err := ssh.Dial("tcp", client.Server.String(), client.Config)
	if err != nil {
		return fmt.Errorf("ssh dial failed:\n\t%v\n", err)
	}

	// creating a new session
	session, err := connection.NewSession()
	if err != nil {
		return fmt.Errorf("opening session failed:\n\t%v\n", err)
	}

	client.Session = session

	return nil
}
