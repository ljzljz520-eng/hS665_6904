package report

import (
	"forkliftarchive/internal/domain"
	"testing"
)

func TestSummary(t *testing.T) {
	m := Summarize([]domain.Record{{Status: "pending"}, {Status: "pending"}})
	if m["pending"] != 2 {
		t.Fatal(m)
	}
}
