package internal

import (
	"bufio"
	"fmt"
	"log"

	"github.com/songgao/water"
	"golang.org/x/net/ipv4"
)

// readFromInterface
// reads packet from the tunnel that
// we just created.
func readFromInterface(inf *water.Interface) error {
	// creating a new reader
	r := bufio.NewReader(inf)

	// creating a 1500 bytes buffer
	packet := make([]byte, 1500)

	// start reading from interface
	for {
		// reading packet
		n, err := r.Read(packet)
		if err != nil {
			return fmt.Errorf("failed to read packets: %s\n", err)
		}

		// parsing header
		hdr, err := ipv4.ParseHeader(packet[:n])
		if err != nil {
			log.Printf("error while parsing ip header: %v", err)

			continue
		}

		log.Printf("got packet: %+v\n", hdr)
	}
}
