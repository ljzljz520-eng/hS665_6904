package service

import (
	"fmt"
	"forkliftarchive/internal/domain"
	"forkliftarchive/internal/store"
	"time"
)

type Clock interface{ Now() time.Time }
type FixedClock struct{ T time.Time }

func (c FixedClock) Now() time.Time { return c.T }

type Service struct {
	Store *store.Store
	Clock Clock
}

func New(st *store.Store, c Clock) *Service {
	if c == nil {
		c = FixedClock{T: time.Unix(0, 0)}
	}
	return &Service{Store: st, Clock: c}
}
func (s *Service) Register(id, code, title, loc string, cap int) (domain.Record, error) {
	r := domain.NewRecord(id, code, domain.NormalizeTitle(title), loc, cap, s.Clock.Now())
	if !r.Valid() {
		return r, fmt.Errorf("invalid record")
	}
	return r, s.Store.SaveRecord(r)
}
func (s *Service) Get(id string) (domain.Record, error) { return s.Store.GetRecord(id) }
func (s *Service) Approve(id, actor string) error       { return s.transition(id, "approved", actor) }
func (s *Service) Reject(id, actor string) error        { return s.transition(id, "rejected", actor) }
func (s *Service) Archive(id, actor string) error       { return s.transition(id, "archived", actor) }
func (s *Service) ChangeTitle(id, title, actor string) error {
	r, e := s.Get(id)
	if e != nil {
		return e
	}
	if e = domain.ValidateTransition(r.Status, "changed"); e != nil {
		return e
	}
	r.Title = domain.NormalizeTitle(title)
	r.Status = "changed"
	r.Version++
	r.UpdatedAt = s.Clock.Now()
	if e = s.Store.SaveRecord(r); e != nil {
		return e
	}
	return s.Store.SaveAudit(domain.AuditEvent{ID: fmt.Sprintf("%s-%d", id, r.Version), RecordID: id, Actor: actor, Action: "change", Detail: r.Title, At: s.Clock.Now()})
}
func (s *Service) transition(id, to, actor string) error {
	r, e := s.Get(id)
	if e != nil {
		return e
	}
	if e = domain.ValidateTransition(r.Status, to); e != nil {
		return e
	}
	r.Status = to
	r.Version++
	r.UpdatedAt = s.Clock.Now()
	if e = s.Store.SaveRecord(r); e != nil {
		return e
	}
	return s.Store.SaveAudit(domain.AuditEvent{ID: fmt.Sprintf("%s-%d", id, r.Version), RecordID: id, Actor: actor, Action: to, At: s.Clock.Now()})
}
