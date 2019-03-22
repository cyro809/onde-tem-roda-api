package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/onde-tem-roda-api/handler"
)

// DefineUserRoutes defines all user routes from the API
func DefineUserRoutes(router *httprouter.Router) {
	router.GET("/users", handler.GetAllUsers)
	router.GET("/users/:userID", handler.GetUserByID)
}
