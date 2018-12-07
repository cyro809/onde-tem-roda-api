package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, _ := sql.Open("sqlite3", "./ondetemroda.db")

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS users (ID INTEGER PRIMARY KEY, Nome TEXT, Email TEXT, Telefone TEXT)")
	statement.Exec()
	statement, _ = database.Prepare(`
		CREATE TABLE IF NOT EXISTS event (
			ID INTEGER PRIMARY KEY,
			NomeEvento TEXT,
			Grupo TEXT,
			Responsavel TEXT,
			ResponsavelTel TEXT,
			ResponsavelEmail TEXT,
			Endereco TEXT,
			PlaceId TEXT,
			UserID INTEGER NOT NULL,
				FOREIGN KEY (UserID) REFERENCES users(ID)
		)
	`)
	statement.Exec()
}
