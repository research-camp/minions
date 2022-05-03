package main

import (
	"github.com/amirhnajafiz/xerox/internal/config"
	"github.com/amirhnajafiz/xerox/proxy"
	"github.com/amirhnajafiz/xerox/server"
)

func main() {
	cfg := config.Load()

	// starting a proxy server
	go proxy.New(cfg.Proxy, cfg.Metric)

	// starting a new server
	server.New()
}
