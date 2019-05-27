package pending

import (
	proto "github.com/golang/protobuf/proto"
)

// Serialize Serial Pending
func (p *Pending) Serialize() ([]byte, error) {
	return proto.Marshal(p)
}

// Deserialize Deserial Pending
func (p *Pending) Deserialize(input []byte) {
	proto.Unmarshal(input, p)
}

// NewPending get new pending with json str
func NewPending(input []byte) *Pending {
	ret := new(Pending)
	ret.Deserialize(input)

	return ret
}
