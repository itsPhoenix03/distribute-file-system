package main

import (
	"log"

	"github.com/itsPhoenix03/distribute-file-system.git/peer2peer"
)

func main() {
	options := peer2peer.TCPTransportOptions{
		ListenAddress: ":4000",
		ShakeHands:    peer2peer.NOPHandShake,
		Decoder:       peer2peer.DefaultDecoder{},
	}

	te := peer2peer.NewTCPTransport(options)

	if err := te.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
