package main

import (
	"fmt"

	"github.com/onde-tem-roda-api/models"

	"github.com/onde-tem-roda-api/manager"
)

//olhar pacote flag

// criar um pacote cmd com um main para instrumentar os command lines
func main() {
	// user := manager.SelectUserByID(1)
	user := models.User{
		Nome:     "Brutus",
		Email:    "brutus@brutus.com",
		Telefone: "11111111",
	}
	manager.InsertUser(user)
	users := manager.SelectAllUsers()
	fmt.Println(users)
}
