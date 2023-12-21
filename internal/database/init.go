package database

import (
	"database/sql"
	"log"

	"git.sula.io/solevis/poultracker/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// InitDatabase connect to the database
// returning the DB connection
func Init() (*sql.DB, error) {
	var err error

	db, err = sql.Open("sqlite3", config.DatabasePath())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to database")

	return db, nil
}

// GetDB get the DB connection
func GetDB() *sql.DB {
	return db
}
