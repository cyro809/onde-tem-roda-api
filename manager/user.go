package manager

import (
	"fmt"

	"github.com/onde-tem-roda-api/database"
	"github.com/onde-tem-roda-api/models"
)

// SelectUserByID selects user from the database by its ID
func SelectUserByID(id int) models.User {
	fmt.Println("selecting")
	connection, err := database.OpenDB()
	if err != nil {
		return models.User{}
	}
	defer connection.Close()

	fmt.Println("opened")
	// Verificar o uso de .Get ou Select
	rows, _ := connection.Query("SELECT * FROM users WHERE id = 1")
	fmt.Println(rows)
	var user models.User
	for rows.Next() {
		rows.Scan(&user)
	}

	return user
}
