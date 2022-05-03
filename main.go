package main

import (
	"github.com/amirhnajafiz/xerox/internal/config"
	"github.com/amirhnajafiz/xerox/internal/metric"
	"github.com/amirhnajafiz/xerox/proxy"
)

func main() {
	// loading configs
	cfg := config.Load()

	if cfg.Metric.Enable {
		metric.NewServer(cfg.Metric).Start()
	}

	// starting a proxy server
	proxy.New(cfg.Proxy)
}
