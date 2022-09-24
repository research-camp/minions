package main

import (
	"os"

	"github.com/amirhnajafiz/xerox"
)

func main() {
	var (
		port = ":8080"
	)

	if value, ok := os.LookupEnv("SERVER_PORT"); ok {
		port = ":" + value
	}

	proxy := xerox.NewProxyServer(port)

	// starting the proxy server
	proxy.Start()
}
