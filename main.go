package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/onde-tem-roda-api/handler"
)

//olhar pacote flag

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// criar um pacote cmd com um main para instrumentar os command lines
func main() {
	router := httprouter.New()
	router.GET("/", Index)

	router.GET("/users", handler.GetAllUsers)
	router.GET("/users/:userID", handler.GetUserByID)

	router.GET("/events", handler.GetAllEvents)
	router.GET("/events/:eventID", handler.GetEventByID)

	log.Fatal(http.ListenAndServe(":3000", router))
}
