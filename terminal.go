package xerox

type SSHTerminal struct {
	Echo             uint32
	TtyOpInputSpeed  uint32
	TtyOpOutputSpeed uint32
	Columns          int
	Rows             int
}
