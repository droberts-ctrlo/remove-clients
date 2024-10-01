package main

import (
	"blog.davetheitguy/remove-clients/connections"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_HOST"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Successfully connected to database!")

	clients, err := connections.ClientsByName(db, "Connells")
	if err != nil {
		log.Fatal(err)
	}

	for _, client := range clients {
		props, err := connections.PropsByClientID(db, client.Id)
		if err != nil {
			log.Fatal(err)
		}
		for _, prop := range props {
			jobs, err := connections.GetJobsFromPropertyID(db, prop.PropID)
			if err != nil {
				log.Fatal(err)
			}
			for _, job := range jobs {
				reports, err := connections.GetReportDataFromSubID(db, job.SubID)
				if err != nil {
					log.Fatal(err)
				}

				for _, report := range reports {
					fmt.Printf("%s: %s\n", report.Name, report.Value)
				}
			}
		}

		pdfs, err := connections.GetPdfsByClientId(db, client.Id)
		if err != nil {
			log.Fatal(err)
		}

		for _, pdf := range pdfs {
			fmt.Printf("File: %s\n", pdf.FileName)
		}
	}
}
