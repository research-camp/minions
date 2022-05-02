package proxy

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// New : creates a new reverse proxy server on port 8080
func New() {
	reverseProxy := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Printf("[reverse proxy server] received request at: %s\n", time.Now())
	})

	log.Fatal(http.ListenAndServe(":8080", reverseProxy))
}
