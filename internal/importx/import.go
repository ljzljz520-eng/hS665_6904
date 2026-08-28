package importx

import (
	"fmt"
	"forkliftarchive/internal/domain"
	"forkliftarchive/internal/service"
)

func ImportRows(s *service.Service, rows []domain.ImportRow, actor string) domain.ImportReport {
	rep := domain.ImportReport{}
	for i, row := range rows {
		if row.Code == "" || row.Title == "" || row.Location == "" || row.Capacity <= 0 {
			rep.Rejected++
			rep.Errors = append(rep.Errors, fmt.Sprintf("row %d invalid", i+1))
			continue
		}
		id := fmt.Sprintf("import-%d-%s", i+1, row.Code)
		if _, e := s.Register(id, row.Code, row.Title, row.Location, row.Capacity); e != nil {
			rep.Rejected++
			rep.Errors = append(rep.Errors, e.Error())
			continue
		}
		rep.Accepted++
	}
	return rep
}
