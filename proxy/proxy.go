package proxy

import (
	"log"
	"net/http"
	"net/url"
)

// New : creates a new reverse proxy server on port 8080
func New() {
	// forward client to the main server
	originServerURL, err := url.Parse("http://127.0.0.1:8081")
	if err != nil {
		log.Fatal("invalid origin server URL")
	}

	reverseProxy := http.HandlerFunc(HandleRequest(originServerURL))

	log.Fatal(http.ListenAndServe(":8080", reverseProxy))
}
