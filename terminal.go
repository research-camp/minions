package xerox

// SSHTerminal
// to run the command on the remote machine, we should create a pseudo terminal on the remote machine.
// A pseudo-terminal (or “pty”) is a pair of virtual character devices that provide a
// bidirectional communication channel.
type SSHTerminal struct {
	Echo             uint32
	TtyOpInputSpeed  uint32
	TtyOpOutputSpeed uint32
	Columns          int
	Rows             int
}
