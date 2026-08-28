package store

import (
	"forkliftarchive/internal/domain"
	"testing"
	"time"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := t.TempDir() + "/x.db"
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	r := domain.NewRecord("1", "F-1", "Reach", "A", 100, time.Unix(1, 0))
	if e = s.SaveRecord(r); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	got, e := s.GetRecord("1")
	if e != nil || got.Title != "Reach" {
		t.Fatalf("%+v %v", got, e)
	}
}
