package dao

import (
	"sync"

	"github.com/boltdb/bolt"
)

const (
	// dbname
	dbname string = "RECORD.db"

	// table / bucket name
	recordBucket  string = "RECORD"
	pendingBucket string = "PENDING"
)

var instance *Dao
var once sync.Once

// Dao db struct
type Dao struct {
	db *bolt.DB
}

// Instance get database instance
func Instance() *Dao {
	once.Do(func() {
		instance = &Dao{}
		instance.db, _ = bolt.Open(dbname, 0600, nil)
		defer instance.Close()
	})
	return instance
}

// Close : close db
func (d *Dao) Close() error {
	return d.db.Close()
}
