package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// OpenDB opens connection to database
func OpenDB() (*sql.DB, error) {
	return sql.Open("sqlite3", "database/ondetemroda.db")
}
