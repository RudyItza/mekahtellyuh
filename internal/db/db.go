package db

import (
	"database/sql"
	"log" // Importing log to use for error logging

	_ "github.com/lib/pq"
)

var db *sql.DB

// Initialize Database Connection
func InitDB() (*sql.DB, error) {
	connStr := "postgres://rudy:tella@localhost/folklore?sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err) // Logging the error if db connection fails
		return nil, err
	}

	// Example of using the db variable for a query
	err = db.Ping() // This checks if the connection is valid
	if err != nil {
		log.Fatal("Failed to ping the database:", err) // Logging the error if ping fails
		return nil, err
	}
	return db, nil
}
