package main

import (
	"fmt"

	"github.com/onde-tem-roda-api/manager"
)

//olhar pacote flag

// criar um pacote cmd com um main para instrumentar os command lines
func main() {
	events := manager.SelectAllEvents()
	// user := models.User{
	// 	Nome:     "Brutus",
	// 	Email:    "brutus@brutus.com",
	// 	Telefone: "11111111",
	// }
	// manager.InsertUser(user)
	// user := manager.SelectUserByEmail("cyro809@gmail.com")
	// users := manager.SelectAllUsers()
	fmt.Println(events)
}
