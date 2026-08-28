package service

import (
	"forkliftarchive/internal/store"
	"testing"
	"time"
)

func TestRegisterApprove(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/x")
	defer s.Close()
	x := New(s, FixedClock{T: time.Unix(1, 0)})
	if _, e := x.Register("1", "F", "Title", "A", 1); e != nil {
		t.Fatal(e)
	}
	if e := x.Approve("1", "a"); e != nil {
		t.Fatal(e)
	}
}
