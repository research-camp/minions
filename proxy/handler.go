package proxy

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type ReqHandFunc func(w http.ResponseWriter, r *http.Request)

func HandleRequest(originServerURL *url.URL) ReqHandFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
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
	}
}
