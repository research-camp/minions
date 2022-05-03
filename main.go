package main

import (
	"github.com/amirhnajafiz/xerox/internal/config"
	"github.com/amirhnajafiz/xerox/internal/metric"
	"github.com/amirhnajafiz/xerox/proxy"
	"github.com/amirhnajafiz/xerox/server"
)

func main() {
	// loading configs
	cfg := config.Load()

	if cfg.Metric.Enable {
		metric.NewServer(cfg.Metric).Start()
	}

	// starting a proxy server
	go proxy.New(cfg.Proxy)

	// starting a new server
	server.New()
}
