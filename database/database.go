package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase(dsn string) {
	var err error

	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("DATABASE CONNECTION FAILED => ", err)
	}

	log.Println("Database connected")
}
