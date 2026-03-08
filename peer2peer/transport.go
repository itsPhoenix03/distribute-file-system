package peer2peer

// Peer is an representation of a nodes
type Peer interface {
	Close() error
}

// Transport is an interface representing transport messages between nodes
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
