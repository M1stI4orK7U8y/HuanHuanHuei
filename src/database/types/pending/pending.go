package pending

import "encoding/json"

// Marshal Marshal Pending
func (p *Pending) Marshal() ([]byte, error) {
	return json.Marshal(p)
}

// Unmarshal Unmarshal Pending
func (p *Pending) Unmarshal(input []byte) {
	json.Unmarshal(input, p)
}

// NewPending get new pending with json str
func NewPending(input []byte) *Pending {
	ret := new(Pending)
	ret.Unmarshal(input)

	return ret
}
