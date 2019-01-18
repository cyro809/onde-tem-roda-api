package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/onde-tem-roda-api/manager"
)

// GetAllUsers returns all users on the database
func GetAllUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	users := manager.SelectAllUsers()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
}

func GetUserByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID, err := strconv.Atoi(ps.ByName("userID"))
	if err != nil {
		panic(err)
	}
	user := manager.SelectUserByID(userID)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}
