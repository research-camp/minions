package proxy

import (
	"log"
	"net/http"
	"net/url"

	"github.com/amirhnajafiz/xerox/internal/config"
	"github.com/amirhnajafiz/xerox/internal/metric"
)

// New : creates a new reverse proxy server on port 8080
func New() {
	cfg := config.Load()
	mtc := metric.NewMetrics()

	metric.NewServer(cfg.Metric).Start()

	// forward client to the main server
	originServerURL, err := url.Parse("http://127.0.0.1:8081")
	if err != nil {
		log.Fatal("invalid origin server URL")
	}

	reverseProxy := http.HandlerFunc(HandleRequest(originServerURL))

	log.Fatal(http.ListenAndServe(":8080", reverseProxy))
}
