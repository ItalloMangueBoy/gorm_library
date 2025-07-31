package controllers

import (
	"encoding/json"
	"library/config/database"
	"library/src/models"
	"library/src/views"
	"net/http"
	"strconv"
	"strings"

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

// ListBooks handles fetching a list of books, optionally filtered by a search query
func ListBooks(w http.ResponseWriter, r *http.Request) {
	// Get the 'search' query parameter and trim whitespace
	search := strings.TrimSpace(r.URL.Query().Get("search"))
	var books []models.Book
	query := database.Conn

	// If a search term is provided, filter books by title or author (case-insensitive)
	if search != "" {
		searchPattern := "%" + strings.ToLower(search) + "%"
		query = query.Where("LOWER(title) LIKE ? OR LOWER(author) LIKE ?", searchPattern, searchPattern)
	}

	// Execute the query to find books
	if err := query.Find(&books).Error; err != nil {
		views.Message(w, http.StatusInternalServerError, "Failed to retrieve books")
		return
	}

	// If no books are found, return a not found message
	if len(books) == 0 {
		views.Message(w, http.StatusNotFound, "No books found")
		return
	}

	// Respond with the list of books
	views.JSON(w, http.StatusOK, books)
}

// UpdateBook handles updating an existing book by its ID
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	var updates models.Book

	// Parse the ID from the URL
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		views.Message(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	// Retrieve the book from the database
	if err := database.Conn.First(&book, id).Error; err != nil {
		views.Message(w, http.StatusNotFound, "ID not found")
		return
	}

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		views.Message(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Format and validate book data
	if err := updates.Format().Valid(); err != nil {
		views.ModelErrors(w, http.StatusUnprocessableEntity, err, "Invalid book data")
		return
	}

	// Update book in database
	if err := database.Conn.Model(&book).Updates(updates).Error; err != nil {
		views.Message(w, http.StatusInternalServerError, "Failed to update book")
		return
	}

	// Respond with updated book
	views.JSON(w, http.StatusOK, book)
}

// DeleteBook handles deleting a book by its ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Parse the ID from the URL
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		views.Message(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	// Attempt to delete the book directly
	response := database.Conn.Where("id = ?", id).Delete(&models.Book{})

	if response.Error != nil {
		views.Message(w, http.StatusInternalServerError, "Failed to delete book")
		return
	}

	// Check if any row was affected (book existed)
	if response.RowsAffected == 0 {
		views.Message(w, http.StatusNotFound, "ID not found")
		return
	}

	// Respond with a success message
	views.Message(w, http.StatusOK, "Book deleted successfully")
}
