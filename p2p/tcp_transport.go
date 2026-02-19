package p2p

import (
	"errors"
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

type TCPTransportOpts struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
	OnPeer        func(Peer) error
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener
	rrpc     chan RPC
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rrpc:             make(chan RPC, 1024),
	}
}

func (t *TCPTransport) Addr() string {
	return t.ListenAddr
}

func (t *TCPTransport) Consume() <-chan RPC {
	return t.rrpc
}

func (t *TCPTransport) Close() error {
	return t.listener.Close()
}

func (t *TCPTransport) Dial(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		logger.Logger.Error("tcp dial failed")
		return err
	}
	go t.handleConn(conn, true)

	return nil

}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return nil
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) handleConn(conn net.Conn, outbound bool) {

	defer func() {
		logger.Logger.Info("dropping peer connection")
		conn.Close()
	}()
	peer := NewTCPPeer(conn, outbound)

	if err := t.HandshakeFunc(peer); err != nil {
		logger.Logger.Error("tcp handShakeFunc failed")
		return
	}

	if t.OnPeer != nil {
		if err := t.OnPeer(peer); err != nil {
			return
		}
	}
	for {
		rpc := RPC{}
		err := t.Decoder.Decode(conn, &rpc)
		if err != nil {
			return
		}

		rpc.From = conn.RemoteAddr().String()

		if rpc.Stream {
			peer.wg.Add(1)
			logger.Logger.Info(rpc.From + "incoming stream,waiting...\n")
			peer.wg.Wait()
		}
		t.rrpc <- rpc
	}
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if errors.Is(err, net.ErrClosed) {
			return
		}
		if err != nil {
			logger.Logger.Error("startAcceptLoop error")
		}
		go t.handleConn(conn, false)
	}
}
