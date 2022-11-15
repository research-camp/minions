package internal

import (
	"fmt"
	"log"

	"github.com/songgao/water"
)

const (
	// name of the tunnel
	tunnel = "utun7"
)

// CreateNewTunnel
// generates a new water tun interface.
func CreateNewTunnel() (*water.Interface, error) {
	// water configs
	config := water.Config{
		DeviceType: water.TUN,
	}
	config.Name = tunnel

	// creating a new tunnel
	inf, err := water.New(config)
	if err != nil {
		return nil, fmt.Errorf("error while creating a tun interface: %v\n", err)
	}

	log.Printf("tunnel created with name: %s\n", inf.Name())

	return inf, err
}