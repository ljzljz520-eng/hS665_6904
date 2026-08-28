package api

import (
	"forkliftarchive/internal/service"
	"forkliftarchive/internal/store"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/x")
	defer s.Close()
	r := httptest.NewRecorder()
	New(service.New(s, nil)).Handler().ServeHTTP(r, httptest.NewRequest("GET", "/health", nil))
	if r.Code != 200 {
		t.Fatal(r.Code)
	}
}
