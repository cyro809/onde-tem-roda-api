package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/onde-tem-roda-api/handler"
)

//olhar pacote flag

type Handle func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func WrapperHandler(handle Handle) httprouter.Handle {
	wrapper := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		err := handle(w, r, ps)
		if err != nil {
			panic(err)
		}
	}
	return wrapper
}

// criar um pacote cmd com um main para instrumentar os command lines
func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/users", handler.GetAllUsers)
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
