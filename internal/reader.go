package internal

import (
	"bufio"
	"fmt"
	"log"

	"github.com/songgao/water"
)

// ReadFromInterface
// reads packet from the tunnel that
// we just created.
func ReadFromInterface(inf *water.Interface) error {
	// creating a new reader
	r := bufio.NewReader(inf)

	// creating a 1500 bytes buffer
	packet := make([]byte, 1500)

	// start reading from interface
	for {
		n, err := r.Read(packet)
		if err != nil {
			return fmt.Errorf("failed to read packets: %s\n", err)
		}

		log.Printf("got packet: %v\n", packet[:n])
	}
}
