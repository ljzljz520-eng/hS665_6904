package importx

import (
	"forkliftarchive/internal/domain"
	"forkliftarchive/internal/service"
	"forkliftarchive/internal/store"
	"testing"
	"time"
)

func TestWorkflowImportReport(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/x")
	defer s.Close()
	r := ImportRows(service.New(s, service.FixedClock{T: time.Unix(1, 0)}), []domain.ImportRow{{Code: "A", Title: "T", Location: "L", Capacity: 1}, {Code: "", Title: "X"}}, "u")
	if r.Accepted != 1 || r.Rejected != 1 {
		t.Fatal(r)
	}
}
