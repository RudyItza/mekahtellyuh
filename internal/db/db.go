package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func OpenDB() (*sql.DB, error) {
	return sql.Open("postgres", "postgres://user:pass@localhost/mekahtellyuh?sslmode=disable")
}
