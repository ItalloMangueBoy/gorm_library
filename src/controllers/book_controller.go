package controllers

import (
	"encoding/json"
	"library/config/database"
	"library/src/models"
	"library/src/views"
	"net/http"
)

// CreateBook handles the creation of a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	// Recive data
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		views.Message(w, http.StatusUnprocessableEntity, "Invalid request payload")
		return
	}

	// Format and validate data
	if err := book.Format().Valid(); err != nil {
		views.ModelErrors(w, http.StatusUnprocessableEntity, err, "Invalid book data")
		return
	}

	// Save data
	if err := database.Conn.Create(&book).Error; err != nil {
		views.Message(w, http.StatusInternalServerError, "Failed to create book")
		return
	}

	// Return response
	views.JSON(w, http.StatusCreated, book)
}
