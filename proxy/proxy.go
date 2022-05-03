package proxy

import (
	"log"
	"net/http"
	"net/url"

	"github.com/amirhnajafiz/xerox/internal/metric"
)

// New : creates a new reverse proxy server on port 8080
func New(cfg Config, mCfg metric.Config) {
	if mCfg.Enable {
		metric.NewServer(mCfg).Start()
	}

	// forward client to the main server
	originServerURL, err := url.Parse(cfg.BaseURL)
	if err != nil {
		log.Fatal("invalid origin server URL")
	}

	reverseProxy := http.HandlerFunc(HandleRequest(originServerURL, metric.NewMetrics()))

	log.Fatal(http.ListenAndServe(cfg.Address, reverseProxy))
}
