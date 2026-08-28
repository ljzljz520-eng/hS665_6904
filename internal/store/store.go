package store

import (
	"encoding/json"
	"fmt"
	"forkliftarchive/internal/domain"
	"go.etcd.io/bbolt"
	"path/filepath"
)

var recordsBucket = []byte("records")
var auditBucket = []byte("audit")
var workflowBucket = []byte("workflow")
var attachmentBucket = []byte("attachment")

type Store struct{ db *bbolt.DB }

func Open(path string) (*Store, error) {
	db, err := bbolt.Open(filepath.Clean(path), 0600, nil)
	if err != nil {
		return nil, err
	}
	s := &Store{db: db}
	err = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range [][]byte{recordsBucket, auditBucket, workflowBucket, attachmentBucket} {
			if _, e := tx.CreateBucketIfNotExists(b); e != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		db.Close()
	}
	return s, err
}
func (s *Store) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	return s.db.Close()
}
func put[T any](tx *bbolt.Tx, b []byte, key string, v T) error {
	data, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return tx.Bucket(b).Put([]byte(key), data)
}
func get[T any](tx *bbolt.Tx, b []byte, key string, out *T) error {
	v := tx.Bucket(b).Get([]byte(key))
	if v == nil {
		return fmt.Errorf("not found: %s", key)
	}
	return json.Unmarshal(v, out)
}
func (s *Store) SaveRecord(r domain.Record) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, recordsBucket, r.ID, r) })
}
func (s *Store) GetRecord(id string) (domain.Record, error) {
	var r domain.Record
	e := s.db.View(func(tx *bbolt.Tx) error { return get(tx, recordsBucket, id, &r) })
	return r, e
}
func (s *Store) ListRecords() ([]domain.Record, error) {
	out := []domain.Record{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(recordsBucket).ForEach(func(_, v []byte) error {
			var r domain.Record
			if e := json.Unmarshal(v, &r); e != nil {
				return e
			}
			out = append(out, r)
			return nil
		})
	})
	return out, e
}
func (s *Store) SaveAudit(a domain.AuditEvent) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, auditBucket, a.ID, a) })
}
func (s *Store) SaveWorkflow(w domain.Workflow) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, workflowBucket, w.ID, w) })
}
func (s *Store) SaveAttachment(a domain.Attachment) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, attachmentBucket, a.ID, a) })
}
