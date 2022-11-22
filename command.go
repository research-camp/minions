package xerox

import "io"

// SSHCommand
// is used for sending and receiving commands
// to our remote machine.
type SSHCommand struct {
	Path   string
	Env    []string
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}
