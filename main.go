package main

import (
	"time"

	"github.com/amirhnajafiz/xerox/internal"

	"github.com/songgao/water"
)

func main() {
	// testing our tunnel building
	tun, err := internal.CreateNewTunnel()
	if err != nil {
		panic(err)
	}

	// closing tunnel after the application is closed
	defer func(tun *water.Interface) {
		if er := tun.Close(); er != nil {
			panic(er)
		}
	}(tun)

	// busy waiting
	time.Sleep(10 * time.Minute)
}
