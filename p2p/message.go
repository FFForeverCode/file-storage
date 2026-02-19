package p2p

// RPC holds data is being sent over
type RPC struct {
	From    string
	Payload []byte
	Stream  bool
}
