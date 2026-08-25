package flow052

import (
	"forkliftarchive/internal/domain"
	"forkliftarchive/internal/service"
	"forkliftarchive/internal/store"
	"testing"
	"time"
)

func Test665BusinessRegression(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/x")
	defer s.Close()
	h := New(service.New(s, service.FixedClock{T: time.Unix(1, 0)}))
	if e := h.ImportBatch([]domain.ImportRow{{Code: "A", Title: "旧标题", Location: "L", Capacity: 1}, {Code: "B", Title: "新标题", Location: "L", Capacity: 1}}); e != nil {
		t.Fatal(e)
	}
	r, e := h.Detail("batch-B")
	if e != nil {
		t.Fatal(e)
	}
	if r.Title != "新标题" {
		t.Fatalf("unexpected title %q", r.Title)
	}
}
