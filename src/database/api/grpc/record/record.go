package record

import (
	"github.com/golang/protobuf/proto"
)

// temp := new(protobuf.RecordReply)
// bytes, _ := json.Marshal(ret)
// json.Unmarshal(bytes, temp)

// Serialize Marshal Record
func (r *Record) Serialize() ([]byte, error) {
	return proto.Marshal(r)
}

// Deserialize Unmarshal Record
func (r *Record) Deserialize(input []byte) {
	proto.Unmarshal(input, r)
}

// NewRecord get new record with json str
func NewRecord(input []byte) *Record {
	ret := new(Record)
	ret.Deserialize(input)
	return ret
}
