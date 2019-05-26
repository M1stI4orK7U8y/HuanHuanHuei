package record

import "encoding/json"

// Serialize Marshal TokenDetail
func (td *TokenDetail) Serialize() ([]byte, error) {
	return json.Marshal(td)
}

// Deserialize Unmarshal TokenDetail
func (td *TokenDetail) Deserialize(input []byte) {
	json.Unmarshal(input, td)
}
