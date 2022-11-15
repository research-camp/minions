package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/amirhnajafiz/xerox/internal"
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
	defer func() {
		if er := tun.Close(); er != nil {
			panic(er)
		}
	}()

	// reading packets from tunnel interface
	go func() {
		if er := internal.ReadFromInterface(tun); er != nil {
			panic(er)
		}
	}()

	// waiting for interrupt signal
	<-done

	log.Println("Exiting...")
}
