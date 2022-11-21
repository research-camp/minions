package xerox

import (
	"net"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

// SSHAgent
// manages to create the ssh auth method with sshAuth parameter.
func SSHAgent(sshAuth string) ssh.AuthMethod {
	if sshAgent, err := net.Dial("unix", sshAuth); err == nil {
		return ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers)
	}

	return nil
}
