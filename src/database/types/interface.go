package types

// IRecordType interface of record types
type IRecordType interface {
	Serialize() ([]byte, error)
	Deserialize(input []byte)
	GetId() string
}
