package test

import (
	"log"
	"testing"

	"github.com/songgao/water"
)

func TestTunnel(_ *testing.T) {
	ifce, err := water.New(water.Config{
		DeviceType: water.TUN,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Interface Name: %s\n", ifce.Name())

	packet := make([]byte, 2000)
	for {
		n, err := ifce.Read(packet)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Packet Received: % x\n", packet[:n])
	}
}
