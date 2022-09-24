package main

import "github.com/amirhnajafiz/xerox"

func main() {
	// creating a proxy server on port 8080 and
	proxy := xerox.NewProxyServer("8080")

	// starting the proxy server
	proxy.Start()
}
