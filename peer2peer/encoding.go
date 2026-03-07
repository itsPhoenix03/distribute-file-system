package peer2peer

import (
	"encoding/gob"
	"io"
)

type Decoder interface {
	Decode(io.Reader, *Message) error
}

// Global Decoder
type GlobalDecoder struct{}

func (g GlobalDecoder) Decode(r io.Reader, msg *Message) error {
	return gob.NewDecoder(r).Decode(msg)
}

// Default Decoder
type DefaultDecoder struct{}

func (d DefaultDecoder) Decode(r io.Reader, msg *Message) error {
	buf := make([]byte, 1024)

	n, err := r.Read(buf)

	if err != nil {
		// fmt.Printf("Error in reading the buffer: %+v\n", err)
		return err
	}

	msg.Payload = buf[:n]

	return nil
}
