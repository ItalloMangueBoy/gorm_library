package routes

import (
	"library/src/controllers"

	"github.com/gorilla/mux"
)

// SetupRoutes initializes the routes for the application
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Define books routes
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")

	return router
}
