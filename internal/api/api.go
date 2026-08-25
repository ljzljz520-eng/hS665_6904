package api

import (
	"encoding/json"
	"forkliftarchive/internal/domain"
	"forkliftarchive/internal/service"
	"net/http"
)

type Server struct{ S *service.Service }

func New(s *service.Service) *Server { return &Server{S: s} }
func (x *Server) Handler() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK); w.Write([]byte("ok")) })
	m.HandleFunc("/records", x.records)
	return m
}
func (x *Server) records(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		q := r.URL.Query()
		rows, e := x.S.Search(domain.SearchFilter{Query: q.Get("q"), Status: q.Get("status"), Location: q.Get("location"), Limit: 50})
		if e != nil {
			http.Error(w, e.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(rows)
		return
	}
	http.Error(w, "method not allowed", 405)
}
