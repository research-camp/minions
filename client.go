package xerox

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

type SSHClient struct {
	Server  *Endpoint
	Config  *ssh.ClientConfig
	Session *ssh.Session
}

func (client *SSHClient) Connect() error {
	connection, err := ssh.Dial("tcp", client.Server.String(), client.Config)
	if err != nil {
		return fmt.Errorf("ssh dial failed:\n\t%v\n", err)
	}

	return nil
}
