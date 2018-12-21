package manager

import (
	"fmt"

	"github.com/onde-tem-roda-api/database"
	"github.com/onde-tem-roda-api/models"

	_ "github.com/mattn/go-sqlite3"
)

// SelectUserByID selects user from the database by its ID
func SelectUserByID(ID int) models.User {
	user := models.User{}
	db, err := database.OpenDB()
	checkErr(err)

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM users WHERE ID = %d", ID))
	checkErr(err)
	defer db.Close()

	// Verificar o uso de .Get ou Select
	for rows.Next() {
		rows.Scan(&user.ID, &user.Nome, &user.Email, &user.Telefone)
	}

	return user
}

// SelectUserByEmail selects user from the database by its e-mail
func SelectUserByEmail(email string) models.User {
	user := models.User{}
	db, err := database.OpenDB()
	checkErr(err)

	rows, err := db.Query(fmt.Sprintf(`SELECT * FROM users WHERE Email = "%s"`, email))
	checkErr(err)
	defer db.Close()

	// Verificar o uso de .Get ou Select
	for rows.Next() {
		rows.Scan(&user.ID, &user.Nome, &user.Email, &user.Telefone)
	}

	return user
}

// SelectAllUsers selects and returns all users from table users
func SelectAllUsers() []models.User {
	var users []models.User

	db, err := database.OpenDB()
	checkErr(err)

	rows, err := db.Query("SELECT * FROM users")
	checkErr(err)
	defer db.Close()

	// Verificar o uso de .Get ou Select
	for rows.Next() {
		var user models.User
		rows.Scan(&user.ID, &user.Nome, &user.Email, &user.Telefone)
		users = append(users, user)
	}

	return users
}

//InsertUser inserts a user in the database
func InsertUser(user models.User) {
	db, err := database.OpenDB()
	checkErr(err)

	statement, err := db.Prepare("INSERT INTO users (Nome, Email, Telefone) VALUES (?, ?, ?)")

	statement.Exec(user.Nome, user.Email, user.Telefone)
}

// DeleteUserByID deletes user from database by its ID
func DeleteUserByID(ID int) {
	db, err := database.OpenDB()
	checkErr(err)

	statement, err := db.Prepare("DELETE FROM users WHERE id = ?")
	checkErr(err)

	res, err := statement.Exec(ID)
	checkErr(err)

	affec, err := res.RowsAffected()
	checkErr(err)

	checkAffectedRows(affec)
}

// DeleteUserByEmail deletes user fromd database by its Email
func DeleteUserByEmail(email string) {
	db, err := database.OpenDB()
	checkErr(err)

	statement, err := db.Prepare(`DELETE FROM users WHERE Email = "?"`)
	checkErr(err)

	res, err := statement.Exec(email)
	checkErr(err)

	affec, err := res.RowsAffected()
	checkErr(err)

	checkAffectedRows(affec)
}

// UpdateUserEmail updates user email
func UpdateUserEmail(ID int, email string) {
	db, err := database.OpenDB()
	checkErr(err)

	statement, err := db.Prepare(`UPDATE users SET Email = "?" WHERE id = ?`)
	checkErr(err)

	res, err := statement.Exec(email, ID)
	checkErr(err)

	affec, err := res.RowsAffected()
	checkErr(err)

	checkAffectedRows(affec)
}

// UpdateUserTelephone updates user telephone
func UpdateUserTelephone(ID int, telephone string) {
	db, err := database.OpenDB()
	checkErr(err)

	statement, err := db.Prepare(`UPDATE users SET Telefone = "?" WHERE id = ?`)
	checkErr(err)

	res, err := statement.Exec(telephone, ID)
	checkErr(err)

	affec, err := res.RowsAffected()
	checkErr(err)

	checkAffectedRows(affec)
}

// UpdateUserName updates user name
func UpdateUserName(ID int, name string) {
	db, err := database.OpenDB()
	checkErr(err)

	statement, err := db.Prepare(`UPDATE users SET Nome = "?" WHERE id = ?`)
	checkErr(err)

	res, err := statement.Exec(name, ID)
	checkErr(err)

	affec, err := res.RowsAffected()
	checkErr(err)

	checkAffectedRows(affec)
}
