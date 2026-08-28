package domain

import (
	"testing"
	"time"
)

func TestRecordValidation(t *testing.T) {
	if NewRecord("", "c", "t", "l", 1, time.Time{}).Valid() {
		t.Fatal("expected invalid")
	}
}
