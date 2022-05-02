package proxy

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

// New : creates a new reverse proxy server on port 8080
func New() {
	// forward client to the main server
	originServerURL, err := url.Parse("http://127.0.0.1:8081")
	if err != nil {
		log.Fatal("invalid origin server URL")
	}

	reverseProxy := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Printf("[reverse proxy server] received request at: %s\n", time.Now())

		// set the parameters to forward our client to the main server
		req.Host = originServerURL.Host
		req.URL.Host = originServerURL.Host
		req.URL.Scheme = originServerURL.Scheme
		req.RequestURI = ""

		// send a request to the origin server
		originServerResponse, err := http.DefaultClient.Do(req)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)

			_, _ = fmt.Fprint(rw, err)

			return
		}

		// return response to the client
		rw.WriteHeader(http.StatusOK)
		_, _ = io.Copy(rw, originServerResponse.Body)
	})

	log.Fatal(http.ListenAndServe(":8080", reverseProxy))
}
