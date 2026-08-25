package store

import (
	"forkliftarchive/internal/domain"
	"strings"
)

func (s *Store) Search(f domain.SearchFilter) ([]domain.Record, error) {
	rows, e := s.ListRecords()
	if e != nil {
		return nil, e
	}
	limit := domain.EnsureLimit(f.Limit)
	out := make([]domain.Record, 0, limit)
	for _, r := range rows {
		if f.Status != "" && r.Status != f.Status {
			continue
		}
		if f.Location != "" && !strings.EqualFold(r.Location, f.Location) {
			continue
		}
		if f.Query != "" && !strings.Contains(strings.ToLower(r.Title+" "+r.Code), strings.ToLower(f.Query)) {
			continue
		}
		out = append(out, r)
		if len(out) >= limit {
			break
		}
	}
	return out, nil
}
