package connections

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

type ReportData struct {
	ID    int64
	Name  string
	Value string
}

func GetReportDataFromSubID(db *sql.DB, subId int64) ([]ReportData, error) {
	var reports []ReportData

	rows, err := db.Query("SELECT FormDataID, Name, Value FROM `ExtFormData` WHERE subId = ?", subId)
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
		if err := rows.Scan(&report.ID, &report.Name, &report.Value); err != nil {
			return nil, fmt.Errorf("getReportDataFromID %d: %v", subId, err)
		}

		reports = append(reports, report)
	}

	return reports, nil
}

func DeleteReportData(db *sql.DB, data ReportData) error {
	if _, err := db.Exec("DELETE FROM `ExtFormData` WHERE FormDataID = ?", data.ID); err != nil {
		return fmt.Errorf("deleteReportData: %v", err)
	}
	if strings.HasSuffix(data.Value, ".jpg") || strings.HasSuffix(data.Value, ".jpeg") {
		location := "E:\\Sites\\Portal\\ImageStore\\" + data.Value
		if err := os.Remove(location); err != nil {
			return fmt.Errorf("deleteReportData: %v", err)
		}
	}
	return nil
}
