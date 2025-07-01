package p2p

// Peer models the remote node.
type Peer interface {
	Close() error
}

// Transport handles the communication modes
// TCP, UDP, gRPC etc.
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
