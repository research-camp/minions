package xerox

import (
	"fmt"
	"io"
	"net"

	"golang.org/x/crypto/ssh"
)

// SSHTunnel
// The client is connecting to local endpoint.
// Then the server endpoint mediates between local endpoint and remote endpoint.
// The algorithm is encapsulated in SSHTunnel struct.
type SSHTunnel struct {
	Local  *endpoint
	Server *endpoint
	Remote *endpoint

	Config *ssh.ClientConfig
}

func (tunnel *SSHTunnel) Start() error {
	listener, err := net.Listen("tcp", tunnel.Local.string())
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go tunnel.forward(conn)
	}
}

// Port forwarding is processed by establishing an SSH connection to the intermediate server.
// When we are connected to the intermediate server, we are able to access the target server.
// The data transfer between the client and the remote server is processed by io.Copy function in forward method.
func (tunnel *SSHTunnel) forward(localConn net.Conn) {
	serverConn, err := ssh.Dial("tcp", tunnel.Server.string(), tunnel.Config)
	if err != nil {
		fmt.Printf("Server dial error: %s\n", err)
		return
	}

	remoteConn, err := serverConn.Dial("tcp", tunnel.Remote.string())
	if err != nil {
		fmt.Printf("Remote dial error: %s\n", err)
		return
	}

	copyConn := func(writer, reader net.Conn) {
		defer writer.Close()
		defer reader.Close()

		_, err := io.Copy(writer, reader)
		if err != nil {
			fmt.Printf("io.Copy error: %s", err)
		}
	}

	go copyConn(localConn, remoteConn)
	go copyConn(remoteConn, localConn)
}
