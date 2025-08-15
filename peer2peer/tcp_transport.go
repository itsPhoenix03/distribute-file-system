package peer2peer

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer is the representation of a node over TCP
type TCPPeer struct {
	conn net.Conn

	// If a request and connect => isOutbound = true
	// If a accepting and connect => isOutbound = false
	isOutbound bool
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPPeer(conn net.Conn, isOutbound bool) *TCPPeer {
	return &TCPPeer{
		conn:       conn,
		isOutbound: isOutbound,
	}
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)

	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()

		if err != nil {
			fmt.Printf("Error accepting connection: %s\n", err)
		}

		go t.handleConnection(conn)
	}
}

func (t *TCPTransport) handleConnection(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	fmt.Printf("New incoming connection from %+v\n", peer)
}
