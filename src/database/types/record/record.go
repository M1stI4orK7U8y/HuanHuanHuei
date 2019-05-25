package record

import "encoding/json"

// temp := new(protobuf.RecordReply)
// bytes, _ := json.Marshal(ret)
// json.Unmarshal(bytes, temp)

// Marshal Marshal Record
func (r *Record) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Unmarshal Unmarshal Record
func (r *Record) Unmarshal(input []byte) {
	json.Unmarshal(input, r)
}

// NewRecord get new record with json str
func NewRecord(input []byte) *Record {
	ret := new(Record)
	ret.Unmarshal(input)
	return ret
}
