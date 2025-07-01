package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	// listenAddr := ":4000"
	// tr := NewTCPTransport(listenAddr)

	// assert.Equal(t, tr.listenAddress, listenAddr)

	// assert.Nil(t, tr.ListenAndAccept())

	opts := TCPTransportOpts{
		ListenAddr:    ":4000",
		HandshakeFunc: NOPHandshakeFunc,
		Decoder:       DefaultDecoder{},
	}

	tr := NewTCPTransport(opts)
	assert.Equal(t, tr.ListenAddr, opts.ListenAddr)
	assert.Nil(t, tr.ListenAndAccept())

}
