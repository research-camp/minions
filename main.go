package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/amirhnajafiz/xerox/internal"

	"github.com/songgao/water"
)

func main() {
	// creating two channels for application termination
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// notifying signal
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// waiting for done signal
	go func() {
		<-sigs
		done <- true
	}()

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

	// waiting for interrupt signal
	<-done

	log.Println("Exiting...")
}
