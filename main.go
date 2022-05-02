package main

import (
	"amirhnajafiz/xerox/internal/proxy"
	"amirhnajafiz/xerox/internal/server"
)

func main() {
	// starting a proxy server
	go proxy.New()

	// starting a new server
	server.New()
}
