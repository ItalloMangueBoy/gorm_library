package routes

import (
	"library/src/controllers"

	"github.com/gorilla/mux"
)

// SetupRoutes initializes the application's routes.
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Book routes
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books", controllers.ListBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")

	return router
}
