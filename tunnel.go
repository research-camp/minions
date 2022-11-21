package xerox

import (
	"io"
	"log"
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
	// creating a listener based on local server
	listener, err := net.Listen("tcp", tunnel.Local.string())
	if err != nil {
		return err
	}
	defer func() {
		// closing local server
		if er := listener.Close(); er != nil {
			log.Printf("ssh tunnel failed to close:\n\t%v\n", er)
		}
	}()

	for {
		// accepting clients on local server
		conn, er := listener.Accept()
		if er != nil {
			return er
		}

		// forward client packets to intermediate server
		go tunnel.forward(conn)
	}
}

// Port forwarding is processed by establishing an SSH connection to the intermediate server.
// When we are connected to the intermediate server, we are able to access the target server.
// The data transfer between the client and the remote server is processed by io.Copy function in forward method.
func (tunnel *SSHTunnel) forward(localConn net.Conn) {
	// creating intermediate server connection
	serverConn, err := ssh.Dial("tcp", tunnel.Server.string(), tunnel.Config)
	if err != nil {
		log.Printf("intermediate server dial error:\n\t%v\n", err)

		return
	}

	// creating remove server connection
	remoteConn, err := serverConn.Dial("tcp", tunnel.Remote.string())
	if err != nil {
		log.Printf("remote server dial error:\n\t%v\n", err)

		return
	}

	// connecting two pipelines for transferring packets between two servers
	copyConn := func(writer, reader net.Conn) {
		// closing pipelines
		defer func() {
			if e := writer.Close(); e != nil {
				log.Printf("writer pipeline failed to close:\n\t%v\n", e)
			}
		}()
		defer func() {
			if e := reader.Close(); e != nil {
				log.Printf("reader pipeline failed to close:\n\t%v\n", e)
			}
		}()

		// using io.Copy to connect two pipelines
		if _, er := io.Copy(writer, reader); er != nil {
			log.Printf("io.Copy for creating pipeline error:\n\t%v\n", er)
		}
	}

	// starting pipelines
	go copyConn(localConn, remoteConn)
	go copyConn(remoteConn, localConn)
}
