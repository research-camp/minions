package internal

import (
	"net/http"
	"strings"
)

// hop to hop headers
var hopToHopHeaders = []string{
	"Connection",
	"Keep-Alive",
	"Proxy-Authenticate",
	"Proxy-Authorization",
	"Te",
	"Trailers",
	"Transfer-Encoding",
	"Upgrade",
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func deleteHopHeaders(header http.Header) {
	for _, h := range hopToHopHeaders {
		header.Del(h)
	}
}

func appendHostToXProxy(header http.Header, host string) {
	if prior, ok := header["X-Forwarded-For"]; ok {
		host = strings.Join(prior, ", ") + ", " + host
	}

	header.Set("X-Forwarded-For", host)
}
