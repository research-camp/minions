package internal

import (
	"fmt"
	"log"

	"github.com/songgao/water"
)

// createNewTunnel
// generates a new water tun interface.
func createNewTunnel() (*water.Interface, error) {
	// creating a new tunnel
	inf, err := water.New(water.Config{
		DeviceType: water.TUN,
	})
	if err != nil {
		return nil, fmt.Errorf("error while creating a tun interface: %v\n", err)
	}

	log.Printf("tunnel created with name: %s\n", inf.Name())

	// prepare network configurations for interface
	if er := prepareNetworkConfigs(inf.Name()); er != nil {
		return nil, fmt.Errorf("failed to set interface network configurations: %v\n", er)
	}

	return inf, err
}
