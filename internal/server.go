package internal

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

// Server type
type Server struct {
	address   string
	target    *url.URL
	createdAt time.Time

	handler ReqHandFunc
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
		address:   address,
		target:    originServerURL,
		createdAt: time.Now(),
		handler:   HandleRequest(originServerURL),
	}
}

// GetInfo returns the proxy server information
func (s Server) GetInfo() string {
	return fmt.Sprintf("[%v] proxy_server: %s, targeting: %s\n", s.createdAt, s.address, s.target)
}

// Start the proxy server
func (s Server) Start() {
	log.Fatal(http.ListenAndServe(s.address, http.HandlerFunc(s.handler)))
}
