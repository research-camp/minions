package xerox

import (
	"io/ioutil"
	"net"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

// SSHAgent
// is a program that runs during user session in *nix system.
// It stores the private keys in an encrypted form.
// Because typing the passphrase can be tedious, many users would prefer to use it to store their private keys.
// You can obtain all stored keys via sshAuth variable which stores the SSH agent unix socket.
// We should access the keys by calling net.Dial and then instance an agent client used
// by ssh.PublicKeysCallback factory auth method.
func SSHAgent(sshAuth string) ssh.AuthMethod {
	if sshAgent, err := net.Dial("unix", sshAuth); err == nil {
		return ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers)
	}

	return nil
}

// PublicKeyFile
// If you want to authenticate by using SSH certificate you need to create a public key file.
// You can parse your private key file by using ssh.ParsePrivateKey function.
// This is required by ssh.PublicKeys auth method function that creates ssh.AuthMethod instance from private key.
func PublicKeyFile(file string) ssh.AuthMethod {
	// reading the public key file
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	// parsing private key file
	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}

	// returning public keys
	return ssh.PublicKeys(key)
}
