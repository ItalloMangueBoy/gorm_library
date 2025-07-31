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

	// Author routes
	router.HandleFunc("/authors", controllers.CreateAuthor).Methods("POST")
	router.HandleFunc("/authors", controllers.ListAuthors).Methods("GET")
	router.HandleFunc("/authors/{id}", controllers.GetAuthor).Methods("GET")
	router.HandleFunc("/authors/{id}", controllers.UpdateAuthor).Methods("PUT")
	router.HandleFunc("/authors/{id}", controllers.DeleteAuthor).Methods("DELETE")

	return router
}
