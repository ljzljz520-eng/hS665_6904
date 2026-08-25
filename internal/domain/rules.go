package domain

import "fmt"

func ValidateTransition(from, to string) error {
	allowed := map[string][]string{"pending": {"approved", "rejected"}, "approved": {"changed", "archived"}, "changed": {"approved", "archived"}, "rejected": {"pending"}}
	for _, x := range allowed[from] {
		if x == to {
			return nil
		}
	}
	return fmt.Errorf("invalid transition %s -> %s", from, to)
}
func NormalizeStatus(s string) string {
	if s == "" {
		return "pending"
	}
	return s
}
func NormalizeTitle(s string) string {
	if len(s) > 80 {
		return s[:80]
	}
	return s
}
func EnsureLimit(n int) int {
	if n <= 0 {
		return 50
	}
	if n > 500 {
		return 500
	}
	return n
}
