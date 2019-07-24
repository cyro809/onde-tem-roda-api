package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/onde-tem-roda-api/database"

	"github.com/julienschmidt/httprouter"
	"github.com/onde-tem-roda-api/routes"
)

//olhar pacote flag

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// criar um pacote cmd com um main para instrumentar os command lines
func main() {
	database.Create()
	database.AddEventDateColumn()

	router := httprouter.New()
	router.GET("/", Index)

	routes.DefineUserRoutes(router)
	routes.DefineEventRoutes(router)

	log.Fatal(http.ListenAndServe(":3000", router))
}
