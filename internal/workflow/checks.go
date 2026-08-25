package workflow

import (
	"fmt"
	"forkliftarchive/internal/domain"
)

func CheckReady(r domain.Record) error {
	if !r.Valid() {
		return fmt.Errorf("record incomplete")
	}
	if r.Status == "archived" {
		return fmt.Errorf("record archived")
	}
	return nil
}
func NextStage(s string) string {
	switch s {
	case "started":
		return "review"
	case "review":
		return "confirmed"
	case "confirmed":
		return "archived"
	}
	return "started"
}
func Stages() []string { return []string{"started", "review", "confirmed", "archived"} }
