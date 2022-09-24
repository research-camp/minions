package xerox

import (
	"github.com/amirhnajafiz/xerox/internal"
)

// Proxy type
type Proxy interface {
	Start()
	GetInfo() string
}

// NewProxyServer creates a new reverse proxy server on port 8080
func NewProxyServer(address string) Proxy {
	return internal.New(address)
}
