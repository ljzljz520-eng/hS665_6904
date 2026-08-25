package report

import (
	"forkliftarchive/internal/domain"
)

type Metrics struct {
	Total, Active, Archived int
	ByStatus                map[string]int
}

func BuildMetrics(rows []domain.Record) Metrics {
	m := Metrics{ByStatus: map[string]int{}, Total: len(rows)}
	for _, r := range rows {
		m.ByStatus[r.Status]++
		if r.Active() {
			m.Active++
		}
		if r.Status == "archived" {
			m.Archived++
		}
	}
	return m
}
func Coverage(rows []domain.Record) float64 {
	if len(rows) == 0 {
		return 0
	}
	n := 0
	for _, r := range rows {
		if r.Title != "" {
			n++
		}
	}
	return float64(n) / float64(len(rows))
}
