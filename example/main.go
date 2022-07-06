package main

import "github.com/amirhnajafiz/xerox"

func main() {
	// creating a proxy server on port 8080 and
	// bind to localhost:8081
	proxy := xerox.NewProxyServer("localhost:8081", "8080")

	// starting the proxy server
	proxy.Start()
}
