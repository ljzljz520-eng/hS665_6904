package domain

import "time"

type Record struct {
	ID, Code, Title, Location, Status string
	Capacity                          int
	UpdatedAt                         time.Time
	Version                           int
}
type AuditEvent struct {
	ID, RecordID, Actor, Action, Detail string
	At                                  time.Time
}
type Workflow struct {
	ID, RecordID, Stage, Owner string
	StartedAt, CompletedAt     time.Time
}
type Attachment struct {
	ID, RecordID, Name, MediaType string
	Data                          []byte
}
type SearchFilter struct {
	Query, Status, Location string
	Limit                   int
}
type ImportRow struct {
	Code, Title, Location string
	Capacity              int
}
type ImportReport struct {
	Accepted, Rejected int
	Errors             []string
}

func (r Record) Valid() bool {
	return r.ID != "" && r.Code != "" && r.Title != "" && r.Location != "" && r.Capacity > 0
}
func (r Record) Active() bool { return r.Status == "pending" || r.Status == "approved" }
func NewRecord(id, code, title, location string, capacity int, now time.Time) Record {
	return Record{ID: id, Code: code, Title: title, Location: location, Capacity: capacity, Status: "pending", UpdatedAt: now, Version: 1}
}
