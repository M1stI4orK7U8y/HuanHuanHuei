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
