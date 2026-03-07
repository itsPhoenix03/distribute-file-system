package peer2peer

import "net"

// Message represents the data which is being sent by the nodes to each other in the network
type Message struct {
	Payload []byte
	From    net.Addr
}
