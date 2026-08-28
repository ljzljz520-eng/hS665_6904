package api

import (
	"encoding/json"
	"forkliftarchive/internal/domain"
	"net/http"
)

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
func parseRecord(r *http.Request) (domain.Record, error) {
	var x domain.Record
	e := json.NewDecoder(r.Body).Decode(&x)
	return x, e
}
func methodAllowed(r *http.Request, methods ...string) bool {
	for _, m := range methods {
		if r.Method == m {
			return true
		}
	}
	return false
}
