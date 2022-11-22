package bolt

/*
	存储示例
	hash0:block0
	hash1:block1
	hash2:block2
	...
	Last:LastHash
*/

import (
	"sync"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

var (
	once   = &sync.Once{}
	boltdb *bolt.DB
)

const (
	_db    = "blockchain.db"
	_table = "blockchain.table"
)

type Dao struct {
	Bolt *bolt.DB
}

func New() *Dao {
	once.Do(func() {
		var err error
		boltdb, err = bolt.Open(_db, 0600, nil)
		if err != nil {
			panic(err)
		}
	})
	return &Dao{
		Bolt: boltdb,
	}
}

func (d *Dao) SetBlock(key []byte, value []byte) error {
	if err := boltdb.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(_table))
		if err != nil {
			return err
		}
		if err := bucket.Put(key, value); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (d *Dao) GetBlock(key []byte) ([]byte, error) {
	var value []byte
	if err := boltdb.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(_table))
		if bucket == nil {
			return errors.Errorf("failed to find table %s", _table)
		}
		value = bucket.Get(key)
		if len(value) == 0 {
			return errors.Errorf("no value found %s", string(key))
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return value, nil
}

func (d *Dao) Name() string {
	return "bolt"
}

func (d *Dao) Close() {
	d.Bolt.Close()
}

func (d *Dao) Reset() {
	boltdb.Update(func(tx *bolt.Tx) error {
		if err := tx.DeleteBucket([]byte(_table)); err != nil {
			return err
		}
		return nil
	})
}
