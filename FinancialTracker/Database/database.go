package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func GetDatabase() *sql.DB {

	connStr := "user=postgres password=kraken1288 dbname=financial_tracker_apps sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	return db
}
