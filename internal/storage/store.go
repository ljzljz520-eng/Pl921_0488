package storage

import (
	"encoding/json"
	"errors"
	"go.etcd.io/bbolt"
)

var buckets = [][]byte{[]byte("customers"), []byte("visits"), []byte("reviews"), []byte("archives")}

type Store struct{ db *bbolt.DB }

func Open(path string) (*Store, error) {
	db, e := bbolt.Open(path, 0600, nil)
	if e != nil {
		return nil, e
	}
	s := &Store{db: db}
	e = s.init()
	if e != nil {
		db.Close()
	}
	return s, e
}
func (s *Store) init() error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, e := tx.CreateBucketIfNotExists(b); e != nil {
				return e
			}
		}
		return nil
	})
}
func (s *Store) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	return s.db.Close()
}
func encode(v any) ([]byte, error) { return json.Marshal(v) }
func decode(data []byte, v any) error {
	if len(data) == 0 {
		return errors.New("not found")
	}
	return json.Unmarshal(data, v)
}
func (s *Store) Put(bucket, key string, v any) error {
	d, e := encode(v)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(bucket)).Put([]byte(key), d) })
}
func (s *Store) Get(bucket, key string, v any) error {
	return s.db.View(func(tx *bbolt.Tx) error { return decode(tx.Bucket([]byte(bucket)).Get([]byte(key)), v) })
}
func (s *Store) List(bucket string) ([][]byte, error) {
	var out [][]byte
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(bucket)).ForEach(func(k, v []byte) error { out = append(out, append([]byte(nil), v...)); return nil })
	})
	return out, e
}
