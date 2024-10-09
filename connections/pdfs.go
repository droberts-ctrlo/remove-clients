package connections

import (
	"database/sql"
	"fmt"
	"os"
)

type PDFData struct {
	Id       int64
	FileName string
}

func GetPdfsByClientId(db *sql.DB, clientId int64) ([]PDFData, error) {
	var pdfs []PDFData

	rows, err := db.Query("SELECT ID, FileName FROM `PDFTable` WHERE ClientID = ?", clientId)
	if err != nil {
		return nil, fmt.Errorf("getPdfsByClient %d: %v", clientId, err)
	}

	defer func() {
		if err := rows.Close(); err != nil {
			fmt.Printf("Error closing rows: %v", err)
		}
	}()

	for rows.Next() {
		var pdf PDFData

		if err := rows.Scan(&pdf.Id, &pdf.FileName); err != nil {
			return nil, fmt.Errorf("getPdfsByClient %d: %v", clientId, err)
		}

		pdfs = append(pdfs, pdf)
	}

	return pdfs, nil
}

func DeletePDF(db *sql.DB, pdf PDFData) error {
	if _, err := db.Exec("DELETE FROM `PDFTable` WHERE PDFID = ?", pdf.Id); err != nil {
		return fmt.Errorf("DeletePDF %d: %v", pdf.Id, err)
	}
	location := "E:\\Sites\\Portal\\pdfstore\\" + pdf.FileName
	if err := os.Remove(location); err != nil {
		return fmt.Errorf("DeletePDF %d: %v", pdf.Id, err)
	}
	return nil
}
