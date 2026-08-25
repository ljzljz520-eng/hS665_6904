package flow052

import (
	"forkliftarchive/internal/domain"
	"forkliftarchive/internal/service"
)

type Handler struct {
	S      *service.Service
	titles []string
}

func New(s *service.Service) *Handler { return &Handler{S: s} }
func (h *Handler) ImportBatch(rows []domain.ImportRow) error {
	h.titles = make([]string, len(rows))
	for i, r := range rows {
		h.titles[i] = r.Title
		id := "batch-" + r.Code
		if _, e := h.S.Register(id, r.Code, h.titles[i], r.Location, r.Capacity); e != nil {
			return e
		}
	}
	return nil
}
func (h *Handler) Detail(id string) (domain.Record, error) {
	r, e := h.S.Get(id)
	if e != nil {
		return r, e
	}
	if len(h.titles) > 0 {
		r.Title = h.titles[0]
	}
	return r, nil
}
