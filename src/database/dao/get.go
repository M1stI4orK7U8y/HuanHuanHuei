package dao

import (
	"errors"

	"github.com/boltdb/bolt"
)

// GetRecord : get member password from owner ID
func (d *Dao) GetRecord(id string) ([]byte, error) {

	ret := []byte{}

	err := d.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(recordBucket))
		if b == nil {
			return errors.New(recordBucket + " Bucket not found")
		}

		retbyte := b.Get([]byte(id))

		if retbyte == nil {
			return errors.New(id + " not found in record bucket")
		}
		ret = retbyte
		return nil
	})

	return ret, err
}

// GetPendings get all pending items
func (d *Dao) GetPendings() (map[string][]byte, error) {

	ret := map[string][]byte{}
	err := d.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(pendingBucket))
		if b == nil {
			return errors.New(pendingBucket + " Bucket not found")
		}

		// 獲得全部 k v pair
		cur := b.Cursor()
		for k, v := cur.First(); k != nil; k, v = cur.Next() {
			ret[string(k)] = v
		}
		return nil // nil error
	})

	return ret, err
}
