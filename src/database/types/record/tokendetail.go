package record

import "encoding/json"

// Marshal Marshal TokenDetail
func (td *TokenDetail) Marshal() ([]byte, error) {
	return json.Marshal(td)
}

// Unmarshal Unmarshal TokenDetail
func (td *TokenDetail) Unmarshal(input []byte) {
	json.Unmarshal(input, td)
}
