package service

import (
	"forkliftarchive/internal/domain"
)

func (s *Service) Search(f domain.SearchFilter) ([]domain.Record, error) { return s.Store.Search(f) }
func (s *Service) Attach(id, name, typ string, data []byte) error {
	if _, e := s.Get(id); e != nil {
		return e
	}
	return s.Store.SaveAttachment(domain.Attachment{ID: id + "/" + name, RecordID: id, Name: name, MediaType: typ, Data: data})
}
