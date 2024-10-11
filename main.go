package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"blog.davetheitguy/remove-clients/connections"

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

	pdfs, err := connections.GetPdfs(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, pdf := range pdfs {
		if err := connections.DeletePDF(db, pdf); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Deleted PDF: %v\n", pdf)
	}
}
