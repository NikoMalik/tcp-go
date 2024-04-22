package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TcpPeer struct {
	conn    net.Conn
	onbound bool
}

type TctpTransport struct {
	listenAddr string
	listener   net.Listener

	peers map[net.Addr]*TcpPeer
	mu    sync.RWMutex
}

func NewTcpPeer(conn net.Conn, onbound bool) *TcpPeer {
	return &TcpPeer{conn: conn, onbound: onbound}
}

func NewTcpTransport(listenAddr string) *TctpTransport {
	return &TctpTransport{
		listenAddr: listenAddr,
		peers:      make(map[net.Addr]*TcpPeer),
	}
}

func (t *TctpTransport) ListenAndAccept() error {
	ln, err := net.Listen("tcp", t.listenAddr)
	if err != nil {
		return err
	}

	t.listener = ln

	go t.acceptLoop()
	return nil
}

func (t *TctpTransport) acceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			return
		}

		t.handleConnection(conn)
	}
}

func (t *TctpTransport) handleConnection(conn net.Conn) {
	t.mu.Lock()
	defer t.mu.Unlock()

	peer := NewTcpPeer(conn, true)
	t.peers[conn.RemoteAddr()] = peer

	fmt.Printf("new connection: %v\n", peer)
}
