package database

import (
	"database/sql"
	"log"
)

func ConnectionDb() (*sql.DB, error) {
	connStr := "user=postgres dbname=db_inventory sslmode=disable password=admin host=localhost"
	db, err := sql.Open("postgres", connStr)
	

	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}

	return db, err
}