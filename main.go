package main

import (
	"fmt"

	"github.com/onde-tem-roda-api/models"

	"github.com/onde-tem-roda-api/manager"
)

//olhar pacote flag

// criar um pacote cmd com um main para instrumentar os command lines
func main() {

	event := models.Event{
		NomeEvento:       "Roda Mister Capoeira",
		Grupo:            "Mister Capoeira Grupo",
		Responsavel:      "Mister M",
		ResponsavelTel:   "999999999",
		ResponsavelEmail: "mail@mail.com",
		Endereco:         "R. Foo Bar, 123",
		PlaceID:          "0o9i8u7y",
		UserID:           1,
	}
	manager.InsertEvent(event)
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
