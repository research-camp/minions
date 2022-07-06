package internal

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
)

// ReqHandFunc for request handling
type ReqHandFunc func(w http.ResponseWriter, r *http.Request)

// HandleRequest for proxy request handling
func HandleRequest(originServerURL *url.URL) ReqHandFunc {
	// handle request method will return a proxy handler by forwarding our client
	return func(rw http.ResponseWriter, req *http.Request) {
		// set the parameters to forward our client to the main server
		req.Host = originServerURL.Host
		req.URL.Host = originServerURL.Host
		req.URL.Scheme = originServerURL.Scheme
		req.RequestURI = ""

		// supporting only http and https
		if req.URL.Scheme != "http" && req.URL.Scheme != "https" {
			msg := ErrUnsupportedProtocol + req.URL.Scheme

			http.Error(rw, msg, http.StatusBadRequest)

			return
		}

		// deleting the hop to hop headers
		deleteHopHeaders(req.Header)

		// appending host to x forward header in proxy server
		if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
			appendHostToXProxy(req.Header, clientIP)
		}

		// send a request to the origin server
		originServerResponse, err := http.DefaultClient.Do(req)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)

			_, _ = fmt.Fprint(rw, err)

			return
		}

		// deleting the hop to hop headers
		deleteHopHeaders(originServerResponse.Header)
		// adding the response headers from origin server
		copyHeader(rw.Header(), originServerResponse.Header)

		// return response to the client
		rw.WriteHeader(http.StatusOK)
		_, _ = io.Copy(rw, originServerResponse.Body)
	}
}
