package report

import (
	"encoding/json"
	"forkliftarchive/internal/domain"
	"sort"
)

func Summarize(rows []domain.Record) map[string]int {
	m := map[string]int{}
	for _, r := range rows {
		m[r.Status]++
	}
	return m
}
func Sorted(rows []domain.Record) []domain.Record {
	out := append([]domain.Record(nil), rows...)
	sort.Slice(out, func(i, j int) bool { return out[i].Code < out[j].Code })
	return out
}
func JSON(rows []domain.Record) ([]byte, error) { return json.Marshal(Sorted(rows)) }
