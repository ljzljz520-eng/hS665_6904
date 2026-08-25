package workflow

import (
	"forkliftarchive/internal/service"
	"forkliftarchive/internal/store"
	"testing"
	"time"
)

func TestWorkflowCreateReviewArchive(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/x")
	defer s.Close()
	e := New(service.New(s, service.FixedClock{T: time.Unix(2, 0)}))
	r, err := e.CreateReviewArchive("1", "F", "T", "A", 2, "u")
	if err != nil || r.Status != "archived" {
		t.Fatalf("%+v %v", r, err)
	}
}
func TestWorkflowSearchUpdatePublish(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/x")
	defer s.Close()
	x := service.New(s, service.FixedClock{T: time.Unix(2, 0)})
	e := New(x)
	x.Register("1", "F", "T", "A", 2)
	x.Approve("1", "u")
	r, err := e.SearchUpdatePublish("1", "New", "u")
	if err != nil || r.Title != "New" {
		t.Fatal(err)
	}
}
