package p2p

const (
	IncomingMessage = 0x1
	IncomingStream  = 0x2
)

// RPC holds data is being sent over
type RPC struct {
	From    string
	Payload []byte
	Stream  bool
}
