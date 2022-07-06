package internal

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

// New : creates a new reverse proxy server on port 8080
func New(cfg Config) {
	// forward client to the main server
	if bURL, ok := os.LookupEnv("BASE_URL"); ok {
		cfg.BaseURL = bURL
	}

	originServerURL, err := url.Parse(cfg.BaseURL)
	if err != nil {
		log.Fatal("invalid origin server URL")
	}

	// reverse proxy server initialize
	reverseProxy := http.HandlerFunc(HandleRequest(originServerURL))

	log.Fatal(http.ListenAndServe(cfg.Address, reverseProxy).Error())
}
