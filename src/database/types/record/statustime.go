package record

import "encoding/json"

// Marshal Marshal StatusTime
func (st *StatusTime) Marshal() ([]byte, error) {
	return json.Marshal(st)
}

// Unmarshal Unmarshal StatusTime
func (st *StatusTime) Unmarshal(input []byte) {
	json.Unmarshal(input, st)
}
