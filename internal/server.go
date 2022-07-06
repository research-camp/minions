package internal

import (
	"log"
	"net/http"
	"net/url"
)

// New : creates a new reverse proxy server on port 8080
func New(target string, address string) {
	// forward client to the main server
	originServerURL, err := url.Parse(target)
	if err != nil {
		log.Fatal("invalid origin server URL")
	}

	// reverse proxy server initialize
	reverseProxy := http.HandlerFunc(HandleRequest(originServerURL))

	// creating a new server
	log.Fatal(http.ListenAndServe(address, reverseProxy).Error())
}
