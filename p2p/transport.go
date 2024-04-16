package p2p

// Peer is a interface that represents the remote node
type Peer interface {
	Close() error
}

// Transport is anything that handles the communication
// between the nodes in the network This can be of the
// following types:
// 1. TCP
// 2. UDP
// 3. Websockets
// 4. ...
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
