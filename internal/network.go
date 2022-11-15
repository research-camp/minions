package internal

import (
	"fmt"
	"github.com/vishvananda/netlink"
)

const (
	// tunnel ip
	tunnelIP = "192.168.9.10/24"
)

// PrepareNetwork
// will set the network configs for our tunnel.
func prepareNetworkConfigs(tunnel string) error {
	// getting the tunnel link by name
	link, err := netlink.LinkByName(tunnel)
	if err != nil {
		return fmt.Errorf("getting tunnel failed: %v", err)
	}

	// parsing the tunnel ip address
	addr, err := netlink.ParseAddr(tunnelIP)
	if err != nil {
		return fmt.Errorf("parsing the tunnel ip failed: %v", err)
	}

	// setting the mtu for tunnel
	if er := netlink.LinkSetMTU(link, 1300); er != nil {
		return fmt.Errorf("setting tunnel mtu failed: %v", er)
	}

	// adding ip address for tunnel
	if er := netlink.AddrAdd(link, addr); er != nil {
		return fmt.Errorf("adding ip address for tunnel failed: %v", er)
	}

	// setting the tunnel link in user pc network
	if er := netlink.LinkSetUp(link); er != nil {
		return fmt.Errorf("setting the tunnel link failed: %v", er)
	}

	return nil
}
