package p2p

import (
	"file-storage/logger"
	"net"
	"sync"
)

type TCPPeer struct {
	net.Conn

	outbound bool

	wg *sync.WaitGroup
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		Conn:     conn,
		outbound: outbound,
		wg:       &sync.WaitGroup{},
	}
}
func (p *TCPPeer) CloseStream() {
	p.wg.Done()
}
func (p *TCPPeer) Send(b []byte) error {
	_, err := p.Conn.Write(b)
	if err != nil {
		logger.Logger.Error("tcp_transport send message error")
	}
	return err
}
