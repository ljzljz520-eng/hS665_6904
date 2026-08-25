package importx

import (
	"encoding/csv"
	"forkliftarchive/internal/domain"
	"io"
	"strconv"
)

func ParseCSV(r io.Reader) ([]domain.ImportRow, error) {
	cr := csv.NewReader(r)
	out := []domain.ImportRow{}
	for {
		row, e := cr.Read()
		if e == io.EOF {
			break
		}
		if e != nil {
			return nil, e
		}
		if len(row) < 4 {
			continue
		}
		cap, e := strconv.Atoi(row[3])
		if e != nil {
			cap = 0
		}
		out = append(out, domain.ImportRow{Code: row[0], Title: row[1], Location: row[2], Capacity: cap})
	}
	return out, nil
}
func Header() []string { return []string{"code", "title", "location", "capacity"} }
