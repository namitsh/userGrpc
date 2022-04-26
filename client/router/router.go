package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// it would contain gorilla mux router setup

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", GetUserByIdHandler).Methods(http.MethodGet)
	r.HandleFunc("/users", GetAllUsersHandler).Methods(http.MethodGet)
	return r
}
