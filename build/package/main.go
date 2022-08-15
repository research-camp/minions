package main

import (
	"fmt"
	"os"

	"github.com/amirhnajafiz/xerox"
)

func main() {
	var (
		target string
		port   = "8080"
	)

	if value, ok := os.LookupEnv("SERVER_PORT"); ok {
		port = value
	}

	if value, ok := os.LookupEnv("SERVER_TARGET"); ok {
		target = value
	} else {
		panic(fmt.Errorf("not target is set"))
	}

	proxy := xerox.NewProxyServer(target, port)

	// starting the proxy server
	proxy.Start()
}
