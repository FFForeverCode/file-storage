package p2p

import "net"

// Peer net connect 端点
type Peer interface {
	net.Conn
	Send([]byte) error
	CloseStream()
}

// Transport 数据传输接口定义
type Transport interface {
	Addr() string
	Dial(string) error
	ListenAndAccept() error
	Consume() <-chan RPC
	Close() error
}
