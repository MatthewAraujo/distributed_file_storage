package p2p

import "net"

// RPC represents any or arbitrary data
// that is being sent over each transport
// between two nodes in the networks
type RPC struct {
	From    net.Addr
	Payload []byte
}
