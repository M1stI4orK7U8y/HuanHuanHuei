package dao

import (
	"errors"

	"github.com/boltdb/bolt"
)

// DeletePending delete pending
func (d *Dao) DeletePending(id string) error {
	return d.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(pendingBucket))
		if b == nil {
			return errors.New(pendingBucket + " Bucket not found")
		}

		return b.Delete([]byte(id)) // nil error
	})
}
