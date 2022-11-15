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

	defer func(tun *water.Interface) {
		err := tun.Close()
		if err != nil {
			panic(err)
		}
	}(tun)

	time.Sleep(10 * time.Minute)
}
