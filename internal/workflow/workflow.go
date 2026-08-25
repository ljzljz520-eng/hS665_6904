package workflow

import (
	"fmt"
	"forkliftarchive/internal/domain"
	"forkliftarchive/internal/service"
	"time"
)

type Engine struct{ S *service.Service }

func New(s *service.Service) *Engine { return &Engine{S: s} }
func (e *Engine) CreateReviewArchive(id, code, title, loc string, cap int, actor string) (domain.Record, error) {
	r, err := e.S.Register(id, code, title, loc, cap)
	if err != nil {
		return r, err
	}
	if err = e.S.Approve(id, actor); err != nil {
		return r, err
	}
	if err = e.S.Archive(id, actor); err != nil {
		return r, err
	}
	return e.S.Get(id)
}
func (e *Engine) SearchUpdatePublish(id, title, actor string) (domain.Record, error) {
	if _, err := e.S.Search(domain.SearchFilter{Query: id}); err != nil {
		return domain.Record{}, err
	}
	if err := e.S.ChangeTitle(id, title, actor); err != nil {
		return domain.Record{}, err
	}
	if err := e.S.Approve(id, actor); err != nil {
		return domain.Record{}, err
	}
	return e.S.Get(id)
}
func StartWorkflow(id, record, owner string, now time.Time) domain.Workflow {
	return domain.Workflow{ID: id, RecordID: record, Owner: owner, Stage: "started", StartedAt: now}
}
func CompleteWorkflow(w domain.Workflow, stage string, now time.Time) (domain.Workflow, error) {
	if stage == "" {
		return w, fmt.Errorf("stage required")
	}
	w.Stage = stage
	w.CompletedAt = now
	return w, nil
}
