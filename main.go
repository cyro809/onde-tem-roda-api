package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//olhar pacote flag

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// criar um pacote cmd com um main para instrumentar os command lines
func main() {
	router := httprouter.New()
	router.GET("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
	// event := models.Event{
	// 	NomeEvento:       "Roda Mister Capoeira",
	// 	Grupo:            "Mister Capoeira Grupo",
	// 	Responsavel:      "Mister M",
	// 	ResponsavelTel:   "999999999",
	// 	ResponsavelEmail: "mail@mail.com",
	// 	Endereco:         "R. Foo Bar, 123",
	// 	PlaceID:          "0o9i8u7y",
	// 	UserID:           1,
	// }
	// manager.UpdateEventByField(2, "NomeEvento", "Mr Caopeira Roda")
	// manager.DeleteEventByID(2)
	// manager.InsertEvent(event)
	// events := manager.SelectAllEvents()
	// event := manager.SelectEventByID(1)
	// user := models.User{
	// 	Nome:     "Brutus",
	// 	Email:    "brutus@brutus.com",
	// 	Telefone: "11111111",
	// }

	// manager.UpdateUserByField(1, "Telefone", "777777777")
	// manager.InsertUser(user)
	// user := manager.SelectUserByEmail("cyro809@gmail.com")

	// users := manager.SelectAllUsers()

	// fmt.Println(user)
	// fmt.Println(events)
	// fmt.Println(event)
}
