package peer2peer

import (
	"errors"
	"fmt"
	"net"
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
	OnPeer        func(*TCPPeer) error
}

type TCPTransport struct {
	TCPTransportOptions
	listener net.Listener
	rpcChan  chan RPC
}

// Close is used to close the connection of a peer from the network
func (p *TCPPeer) Close() error {
	return p.conn.Close()
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
		rpcChan:             make(chan RPC),
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

func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcChan
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

func (t *TCPTransport) handleConnection(conn net.Conn) {
	var err error

	defer func() {
		fmt.Printf("dropping peer connection: %s\n", err)
		conn.Close()
	}()

	peer := NewTCPPeer(conn, true)

	if err := t.ShakeHands(peer); err != nil {
		return
	}

	if t.OnPeer != nil {
		if err = t.OnPeer(peer); err != nil {
			return
		}
	}

	rpcMsg := RPC{}
	// Read loop
	for {
		err = t.Decoder.Decode(conn, &rpcMsg)

		// fmt.Printf("err: ", err)
		// panic(err)
		if errors.Is(err, net.ErrClosed) {
			return
		}

		if err != nil {
			fmt.Printf("Error decoding message: %s\n", err)
			continue
		}

		rpcMsg.From = conn.RemoteAddr()
		t.rpcChan <- rpcMsg
	}

}
