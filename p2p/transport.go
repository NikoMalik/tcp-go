package p2p

import "net"

type Peer interface {
	Send([]byte) (int, error)

	Close() error

	RemoteAddr() net.Addr
}

type Transport interface {
	Accept() (Peer, error)

	Close() error

	Addr() net.Addr
}
