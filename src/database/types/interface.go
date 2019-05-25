package types

// IRecordType interface of record types
type IRecordType interface {
	Marshal() ([]byte, error)
	Unmarshal(input []byte)
	GetID() string
}
