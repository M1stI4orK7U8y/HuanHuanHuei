package types

// IDBType interface of DB types
type IDBType interface {
	Marshal() ([]byte, error)
	Unmarshal(input []byte)
}
