package peer2peer

import "net"

// RPC represents the data which is being sent by the nodes to each other in the network
type RPC struct {
	Payload []byte
	From    net.Addr
}
