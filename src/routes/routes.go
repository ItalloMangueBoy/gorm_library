package routes

import (
	"library/src/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.HandleHome).Methods("GET")

	return router
}
