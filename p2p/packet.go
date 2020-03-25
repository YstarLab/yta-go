package p2p

import (
	"github.com/ystar-foundation/yta-go"
)

type Envelope struct {
	Sender   *Peer
	Receiver *Peer
	Packet   *yta.Packet `json:"envelope"`
}

func NewEnvelope(sender *Peer, receiver *Peer, packet *yta.Packet) *Envelope {
	return &Envelope{
		Sender:   sender,
		Receiver: receiver,
		Packet:   packet,
	}
}
