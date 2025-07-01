package main

import (
	"fmt"
	"log"

	"github.com/ishaan29/Distributed_File_System/p2p"
)

func OnPeer(peer p2p.Peer) error {
	peer.Close()
	fmt.Println("logic with the peer outside of the transport")
	return nil
}

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":4000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("message received: %+v\n", msg)
		}
	}()
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
	//fmt.Println("Hello!, welcome to the world of go.")
}
