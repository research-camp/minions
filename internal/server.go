package internal

import (
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

// NewProxyServer creates a new reverse proxy server on port 8080
func NewProxyServer(target string, address string) Server {
	// forward client to the main server
	originServerURL, err := url.Parse(target)
	if err != nil {
		log.Fatal("invalid origin server URL")
	}

	// reverse proxy server initialize
	return Server{
		Address: address,
		Target:  originServerURL,
		Handler: HandleRequest(originServerURL),
	}
}

// Start the proxy server
func (s Server) Start() {
	go func() {
		log.Fatal(http.ListenAndServe(s.Address, http.HandlerFunc(s.Handler)))
	}()
}
