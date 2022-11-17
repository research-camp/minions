package xerox

import "github.com/amirhnajafiz/xerox/internal"

type VPN interface {
	Run() error
}

func NewVPN() VPN {
	return internal.NewVPN()
}
