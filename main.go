package main

import (
	"log"

	"github.com/itsPhoenix03/distribute-file-system.git/peer2peer"
)

func main() {
	te := peer2peer.NewTCPTransport(":4000")

	if err := te.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
