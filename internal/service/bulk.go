package service

import (
	"fmt"
	"forkliftarchive/internal/domain"
)

func (s *Service) ValidateRows(rows []domain.ImportRow) []string {
	errs := []string{}
	for i, r := range rows {
		if r.Code == "" {
			errs = append(errs, fmt.Sprintf("%d code", i))
		}
		if r.Capacity <= 0 {
			errs = append(errs, fmt.Sprintf("%d capacity", i))
		}
	}
	return errs
}
func (s *Service) BulkApprove(ids []string, actor string) []error {
	errs := []error{}
	for _, id := range ids {
		if e := s.Approve(id, actor); e != nil {
			errs = append(errs, e)
		}
	}
	return errs
}
func (s *Service) Snapshot() ([]domain.Record, error) { return s.Store.ListRecords() }
