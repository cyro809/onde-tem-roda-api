package database

import (
	"database/sql"
)

// OpenDB opens connection to database
func OpenDB() (*sql.DB, error) {
	return sql.Open("sqlite3", "./ondetemroda.db")
}
