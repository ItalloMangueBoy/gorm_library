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

// CreateAuthor handles the creation of a new author
func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		views.Message(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate author data
	if err := author.Format().Valid(); err != nil {
		views.ModelErrors(w, http.StatusUnprocessableEntity, err, "Invalid author data")
		return
	}

	// Save author to database
	if err := database.Conn.Create(&author).Error; err != nil {
		views.Message(w, http.StatusInternalServerError, "Failed to create author")
		return
	}

	// Respond with created author
	views.JSON(w, http.StatusCreated, author)
}

// GetAuthor handles fetching an author by its ID
func GetAuthor(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		views.Message(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	var author models.Author
	if err := database.Conn.First(&author, id).Error; err != nil {
		views.Message(w, http.StatusNotFound, "Author not found")
		return
	}

	views.JSON(w, http.StatusOK, author)
}

// ListAuthors handles fetching a list of authors
func ListAuthors(w http.ResponseWriter, r *http.Request) {
	// "Get the 'search' query parameter and trim whitespace
	search := strings.TrimSpace(r.URL.Query().Get("search"))
	var authors []models.Author
	query := database.Conn

	// If a search term is provided, filter authors by name (case-insensitive)
	if search != "" {
		searchPattern := "%" + strings.ToLower(search) + "%"
		query = query.Where("LOWER(name) LIKE ?", searchPattern)
	}

	// Execute the query to find authors
	if err := query.Find(&authors).Error; err != nil {
		views.Message(w, http.StatusInternalServerError, "Failed to retrieve authors")
		return
	}

	// If no authors are found, return a not found message
	if len(authors) == 0 {
		views.Message(w, http.StatusNotFound, "No authors found")
		return
	}

	// Respond with the list of authors
	views.JSON(w, http.StatusOK, authors)
}

// UpdateAuthor handles updating an existing author by its ID
func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	var updates models.Author

	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		views.Message(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	if err := database.Conn.First(&author, id).Error; err != nil {
		views.Message(w, http.StatusNotFound, "ID not found")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		views.Message(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := updates.Format().Valid(); err != nil {
		views.ModelErrors(w, http.StatusUnprocessableEntity, err, "Invalid author data")
		return
	}

	if err := database.Conn.Model(&author).Updates(updates).Error; err != nil {
		views.Message(w, http.StatusInternalServerError, "Failed to update author")
		return
	}

	views.JSON(w, http.StatusOK, author)
}

// DeleteAuthor handles deleting an author by its ID
func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		views.Message(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	response := database.Conn.Where("id = ?", id).Delete(&models.Author{})
	if response.Error != nil {
		views.Message(w, http.StatusInternalServerError, "Failed to delete author")
		return
	}

	if response.RowsAffected == 0 {
		views.Message(w, http.StatusNotFound, "ID not found")
		return
	}

	views.Message(w, http.StatusOK, "Author deleted successfully")
}
