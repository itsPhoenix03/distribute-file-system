package main

import (
	"log"

	"github.com/itsPhoenix03/distribute-file-system.git/peer2peer"
)

func OnPeer(peer *peer2peer.TCPPeer) error {
	peer.Close()
	return nil
}

func main() {
	options := peer2peer.TCPTransportOptions{
		ListenAddress: ":4000",
		ShakeHands:    peer2peer.NOPHandShake,
		Decoder:       peer2peer.DefaultDecoder{},
		OnPeer:        OnPeer,
	}

	te := peer2peer.NewTCPTransport(options)

	if err := te.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			msg := <-te.Consume()

			log.Printf("Received message: %+v\n", msg)
		}
	}()

	select {}
}
