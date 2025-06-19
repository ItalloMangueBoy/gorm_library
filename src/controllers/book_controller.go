package controllers

import (
	"encoding/json"
	"library/config/database"
	"library/src/models"
	"library/src/views"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateBook handles the creation of a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		views.Message(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate book data
	if err := book.Format().Valid(); err != nil {
		views.ModelErrors(w, http.StatusUnprocessableEntity, err, "Invalid book data")
		return
	}

	// Save book to database
	if err := database.Conn.Create(&book).Error; err != nil {
		views.Message(w, http.StatusInternalServerError, "Failed to create book")
		return
	}

	// Respond with created book
	views.JSON(w, http.StatusCreated, book)
}

// GetBook handles fetching a book by its ID
func GetBook(w http.ResponseWriter, r *http.Request) {
	// Parse the ID from the URL
	idStr := mux.Vars(r)["id"]

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		views.Message(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	// Retrieve the book from the database
	var book models.Book

	if err := database.Conn.First(&book, id).Error; err != nil {
		views.Message(w, http.StatusNotFound, "Book not found")
		return
	}

	// Respond with the found book
	views.JSON(w, http.StatusOK, book)
}
