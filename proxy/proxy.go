package proxy

import (
	"net/http"
	"net/url"

	"github.com/amirhnajafiz/xerox/internal/metric"
	"go.uber.org/zap"
)

// New : creates a new reverse proxy server on port 8080
func New(logger *zap.Logger, cfg Config) {
	// forward client to the main server
	originServerURL, err := url.Parse(cfg.BaseURL)
	if err != nil {
		logger.Fatal("invalid origin server URL")
	}

	// reverse proxy server initialize
	reverseProxy := http.HandlerFunc(HandleRequest(originServerURL, logger, metric.NewMetrics()))

	logger.Fatal(http.ListenAndServe(cfg.Address, reverseProxy).Error())
}
