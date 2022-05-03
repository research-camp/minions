package main

import (
	"github.com/amirhnajafiz/xerox/proxy"
	"github.com/amirhnajafiz/xerox/server"
)

func main() {
	// starting a proxy server
	go proxy.New()

	// starting a new server
	server.New()
}
