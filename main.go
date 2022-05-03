package main

import (
	"amirhnajafiz/xerox/proxy"
	"amirhnajafiz/xerox/server"
)

func main() {
	// starting a proxy server
	go proxy.New()

	// starting a new server
	server.New()
}
