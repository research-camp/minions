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
	Server         *Endpoint
	Config         *ssh.ClientConfig
	TerminalConfig *SSHTerminal
	session        *ssh.Session
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

	// building terminal modes
	modes := ssh.TerminalModes{
		ssh.ECHO:          client.TerminalConfig.Echo,
		ssh.TTY_OP_ISPEED: client.TerminalConfig.TtyOpInputSpeed,
		ssh.TTY_OP_OSPEED: client.TerminalConfig.TtyOpOutputSpeed,
	}

	// registering pseudo terminal
	if er := session.RequestPty("xterm", client.TerminalConfig.Rows, client.TerminalConfig.Columns, modes); er != nil {
		return fmt.Errorf("cannot request for pseudo terminal:\n\t%v\n", er)
	}

	client.session = session

	return nil
}