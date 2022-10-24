package internal

import (
	"log"
	"net/http"
	"time"
)

// Server type
type Server struct {
	address   string
	createdAt time.Time

	handler ReqHandFunc
}

// New creates a new reverse proxy server on port 8080
func New(address string) Server {
	// reverse proxy server initialize
	return Server{
		address:   ":" + address,
		createdAt: time.Now(),
		handler:   HandleRequest(),
	}
}

// Start the proxy server
func (s Server) Start() {
	log.Fatal(http.ListenAndServe(s.address, http.HandlerFunc(s.handler)))
}
