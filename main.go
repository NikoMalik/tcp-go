package main

import (
	"runtime"
	"tcp-go/p2p"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	tr := p2p.NewTcpTransport(":3000")

	go tr.ListenAndAccept()

	select {}
}
