package internal

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// Server type
type Server struct {
	Address string
	Target  *url.URL

	Handler ReqHandFunc
}

// New creates a new reverse proxy server on port 8080
func New(target string, address string) Server {
	// forward client to the main server
	originServerURL, err := url.Parse(target)
	if err != nil {
		log.Fatal(errInvalidUrl)
	}

	// reverse proxy server initialize
	return Server{
		Address: address,
		Target:  originServerURL,
		Handler: HandleRequest(originServerURL),
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
