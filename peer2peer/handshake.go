package peer2peer

import "errors"

// Handshake Error Messages
var (
	InvalidHandshakeError = errors.New("Invalid Handshake")
)

// HandshakeFunc is a function to handle the handshake
type HandshakeFunc func(any) error

// Default Handshake for peers which does nothing
func NOPHandShake(any) error { return nil }
