package xerox

import (
	"fmt"
)

// endpoint
// The tunneling protocol allows a network
// user to access or provide a network service that the
// underlying network does not support or provide directly.
// There are three type of server:
// - remote server
// - local server
// - target server
// each server can be represented by the following struct.
type endpoint struct {
	// server host address
	Host string
	// server port
	Port int
}

// returns the string of our endpoint.
func (endpoint *endpoint) string() string {
	return fmt.Sprintf("%s:%d", endpoint.Host, endpoint.Port)
}
