package record

import "encoding/json"

// Serialize Marshal StatusTime
func (st *StatusTime) Serialize() ([]byte, error) {
	return json.Marshal(st)
}

// Deserialize Unmarshal StatusTime
func (st *StatusTime) Deserialize(input []byte) {
	json.Unmarshal(input, st)
}
