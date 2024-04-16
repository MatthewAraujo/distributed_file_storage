package p2p

import "net"

// message represents any or arbitrary data
// that is being sent over each transport
// between two nodes in the networks
type Message struct {
	From    net.Addr
	Payload []byte
}
