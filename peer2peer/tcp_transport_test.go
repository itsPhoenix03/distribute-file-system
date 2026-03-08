package peer2peer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	opts := TCPTransportOptions{
		ListenAddress: ":4000",
		ShakeHands:    NOPHandShake,
		Decoder:       DefaultDecoder{},
	}
	tr := NewTCPTransport(opts)

	assert.Equal(t, ":4000", tr.ListenAddress)

	// Server
	assert.Nil(t, tr.ListenAndAccept())
}
