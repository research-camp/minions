package xerox

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/amirhnajafiz/xerox/internal"
)

// Proxy type
type Proxy interface {
	Start()

	GetInfo() string
}

// Server type
type Server struct {
	Address string
	Target  *url.URL

	Handler internal.ReqHandFunc
}

// NewProxyServer creates a new reverse proxy server on port 8080
func NewProxyServer(target string, address string) Server {
	// forward client to the main server
	originServerURL, err := url.Parse(target)
	if err != nil {
		log.Fatal(internal.ErrInvalidUrl)
	}

	// reverse proxy server initialize
	return Server{
		Address: address,
		Target:  originServerURL,
		Handler: internal.HandleRequest(originServerURL),
	}
}

// GetInfo returns the proxy server information
func (s Server) GetInfo() string {
	return fmt.Sprintf("proxy_server: %s, targeting: %s\n", s.Address, s.Target)
}

// Start the proxy server
func (s Server) Start() {
	go func() {
		log.Fatal(http.ListenAndServe(s.Address, http.HandlerFunc(s.Handler)))
	}()
}
