package database

import (
	"database/sql"
	"fmt"
)

func insert(data) {

}

func selectUserById(id int, database *sql.DB) {
	openDataBase(database)
	statement, _ := database.Query(fmt.Sprintf("SELECT * FROM users WHERE id = %s", id))

}

func openDataBase(database *sql.DB) {
	database, _ := sql.Open("sqlite3", "./ondetemroda.db")
}
