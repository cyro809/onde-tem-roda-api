package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/onde-tem-roda-api/manager"
)

// GetAllEvents returns all events on the database
func GetAllEvents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	events := manager.SelectAllEvents()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(events); err != nil {
		panic(err)
	}
}

// GetEventByID returns a event by it's ID
func GetEventByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	eventID, err := strconv.Atoi(ps.ByName("eventID"))
	if err != nil {
		panic(err)
	}
	event := manager.SelectEventByID(eventID)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(event); err != nil {
		panic(err)
	}
}
