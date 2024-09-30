package connections

import (
	"database/sql"
	"fmt"
	"log"
)

type ReportData struct {
	Name  string
	Value string
}

func GetReportDataFromSubID(db *sql.DB, subId int64) ([]ReportData, error) {
	var reports []ReportData

	rows, err := db.Query("SELECT Name, Value FROM `ExtFormData` WHERE subId = ?", subId)
	if err != nil {
		return nil, fmt.Errorf("getReportDataFromID %d: %v", subId, err)
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Printf("GetReportDataFromID %d: %v", subId, err)
		}
	}()

	for rows.Next() {
		var report ReportData
		if err := rows.Scan(&report.Name, &report.Value); err != nil {
			return nil, fmt.Errorf("getReportDataFromID %d: %v", subId, err)
		}

		reports = append(reports, report)
	}

	return reports, nil
}
