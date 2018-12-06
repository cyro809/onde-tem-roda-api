package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, _ := sql.Open("sqlite3", "./ondetemroda.db")

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, nome TEXT, email TEXT, telefone TEXT)")
	statement.Exec()
	statement, _ = database.Prepare(`
		CREATE TABLE IF NOT EXISTS event (
			id INTEGER PRIMARY KEY,
			nome_evento TEXT,
			grupo TEXT,
			responsavel TEXT,
			responsavel_tel TEXT,
			responsavel_email TEXT,
			endereco TEXT,
			place_id TEXT,
			user_id INTEGER NOT NULL,
				FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`)
	statement.Exec()
}
