package service

import (
	"forkliftarchive/internal/domain"
	"forkliftarchive/internal/store"
	"testing"
)

func TestSearch(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/x")
	defer s.Close()
	x := New(s, nil)
	x.Register("1", "A", "Alpha", "L", 1)
	rows, e := x.Search(domain.SearchFilter{Query: "alp"})
	if e != nil || len(rows) != 1 {
		t.Fatal(e, len(rows))
	}
}
