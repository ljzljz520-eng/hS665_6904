package store

import (
	"forkliftarchive/internal/domain"
	"time"
)

func (s *Store) SaveMany(rows []domain.Record) error {
	for _, r := range rows {
		if e := s.SaveRecord(r); e != nil {
			return e
		}
	}
	return nil
}
func (s *Store) RecordsUpdatedSince(t time.Time) ([]domain.Record, error) {
	rows, e := s.ListRecords()
	if e != nil {
		return nil, e
	}
	out := []domain.Record{}
	for _, r := range rows {
		if r.UpdatedAt.After(t) {
			out = append(out, r)
		}
	}
	return out, nil
}
func (s *Store) Count(status string) (int, error) {
	rows, e := s.ListRecords()
	if e != nil {
		return 0, e
	}
	n := 0
	for _, r := range rows {
		if status == "" || r.Status == status {
			n++
		}
	}
	return n, nil
}
