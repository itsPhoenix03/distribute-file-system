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

type TCPTransportOptions struct {
	ListenAddress string
	ShakeHands    HandshakeFunc
	Decoder       Decoder
}

type TCPTransport struct {
	TCPTransportOptions
	listener net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPPeer(conn net.Conn, isOutbound bool) *TCPPeer {
	return &TCPPeer{
		conn:       conn,
		isOutbound: isOutbound,
	}
}

func NewTCPTransport(opts TCPTransportOptions) *TCPTransport {
	return &TCPTransport{
		TCPTransportOptions: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddress)

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

		fmt.Printf("New incoming connection from %+v\n", conn)

		go t.handleConnection(conn)
	}
}

// TEMP
type Temp struct{}

func (t *TCPTransport) handleConnection(conn net.Conn) {
	NewTCPPeer(conn, true)

	if err := t.ShakeHands(conn); err != nil {
		conn.Close()
		fmt.Printf("Error during handshake: %s\n", err)
		return
	}

	msg := &Message{}
	// Read loop
	for {
		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Printf("Error decoding message: %s\n", err)
			continue
		}

		msg.From = conn.RemoteAddr()
		fmt.Printf("Received message: %+v\n", msg)
	}

}
