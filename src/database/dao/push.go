package dao

import (
	"github.com/boltdb/bolt"
	types "gitlab.com/packtumi9722/huanhuanhuei/src/database/api/grpc"
)

// UpdateRecord add / update a record
func (d *Dao) UpdateRecord(input types.IRecordType) error {
	return d.update(input, recordBucket)
}

// UpdatePending add pending item
func (d *Dao) UpdatePending(input types.IRecordType) error {
	return d.update(input, pendingBucket)
}

func (d *Dao) update(input types.IRecordType, bname string) error {
	return d.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bname))
		if b == nil {
			newBucket, err := tx.CreateBucket([]byte(bname))
			if err != nil {
				return err
			}
			b = newBucket
		}
		// write to database
		save, _ := input.Serialize()
		return b.Put([]byte(input.GetId()), save)
	})
}
