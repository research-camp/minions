package xerox

import "golang.org/x/crypto/ssh"

type SSHClient struct {
	Server  *Endpoint
	Config  *ssh.Config
	Session *ssh.Session
}
