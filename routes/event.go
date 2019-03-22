package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/onde-tem-roda-api/handler"
)

func DefineEventRoutes(router *httprouter.Router) {
	router.GET("/events", handler.GetAllEvents)
	router.GET("/events/:eventID", handler.GetEventByID)
}
