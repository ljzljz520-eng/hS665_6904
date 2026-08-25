package flow052

import (
	"forkliftarchive/internal/domain"
	"forkliftarchive/internal/service"
)

type Handler struct {
	S *service.Service
}

func New(s *service.Service) *Handler { return &Handler{S: s} }
func (h *Handler) ImportBatch(rows []domain.ImportRow) error {
	for _, r := range rows {
		id := "batch-" + r.Code
		if _, e := h.S.Register(id, r.Code, r.Title, r.Location, r.Capacity); e != nil {
			return e
		}
	}
	return nil
}
func (h *Handler) Detail(id string) (domain.Record, error) {
	return h.S.Get(id)
}
