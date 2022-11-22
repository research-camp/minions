package xerox

import (
	"fmt"
	"io"
	"strings"

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

// RunCommand
// executes commands.
func (client *SSHClient) RunCommand(cmd *SSHCommand) error {
	if err := client.prepareCommand(cmd); err != nil {
		return err
	}

	return client.session.Run(cmd.Path)
}

// To attach our os.Stdin, os.Stdout and os.Stderr to remote command we should open pipes between the
// local process and remote process.
// Fortunately, ssh.Session object provides that out of the box by
// invoking session.StdinPipe(), session.StdoutPipe() and session.StdoutErr() functions.
// Then we should asynchronously copy the end of the pipes to the right file descriptors by using
// go routines and os.Copy function.
func (client *SSHClient) prepareCommand(cmd *SSHCommand) error {
	for _, env := range cmd.Env {
		variable := strings.Split(env, "=")
		if len(variable) != 2 {
			continue
		}

		if err := client.session.Setenv(variable[0], variable[1]); err != nil {
			return fmt.Errorf("setting env variable failed:\n\t%v\n", err)
		}
	}

	if cmd.Stdin != nil {
		stdin, err := client.session.StdinPipe()
		if err != nil {
			return fmt.Errorf("failed to setup stdin for session:\n\t%v\n", err)
		}

		go io.Copy(stdin, cmd.Stdin)
	}

	if cmd.Stdout != nil {
		stdout, err := client.session.StdoutPipe()
		if err != nil {
			return fmt.Errorf("failed to setup stdout for session:\n\t%v\n", err)
		}

		go io.Copy(cmd.Stdout, stdout)
	}

	if cmd.Stderr != nil {
		stderr, err := client.session.StderrPipe()
		if err != nil {
			return fmt.Errorf("failed to setup stderr for session:\n\t%v\n", err)
		}

		go io.Copy(cmd.Stderr, stderr)
	}

	return nil
}
