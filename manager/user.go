package manager

import (
	"fmt"

	"github.com/onde-tem-roda-api/database"
	"github.com/onde-tem-roda-api/models"

	_ "github.com/mattn/go-sqlite3"
)

// SelectUserByID selects user from the database by its ID
func SelectUserByID(ID int) models.User {
	fmt.Println("selecting")
	user := models.User{}
	db, err := database.OpenDB()
	if err != nil {
		fmt.Println(err)
	}

	rows, err2 := db.Query(fmt.Sprintf("SELECT * FROM users WHERE ID = %d", ID))
	if err2 != nil {
		fmt.Println(err2)
	}
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
	if err != nil {
		fmt.Println(err)
	}

	rows, err2 := db.Query("SELECT * FROM users")
	if err2 != nil {
		fmt.Println(err2)
	}
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
	if err != nil {
		fmt.Println(err)
	}

	statement, err := db.Prepare("INSERT INTO users (Nome, Email, Telefone) VALUES (?, ?, ?)")

	statement.Exec(user.Nome, user.Email, user.Telefone)
}
