package models

import (
	"library/src/helpers"
)

type Book struct {
	ID          int64   `json:"id"`
	Pages       int64   `json:"page" validate:"required"`
	Available   bool    `json:"available" validate:"required"`
	Title       string  `json:"title" validate:"required"`
	Author      string  `json:"author" validate:"required"`
	Description *string `json:"description" validate:"required"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// Valid detect errors in the Book struct fields
func (b *Book) Valid() helpers.ErrorMessages {
	customMessages := helpers.ErrorMessages{
		"Book.Pages.required":       "Pages is required",
		"Book.Available.required":   "Available is required",
		"Book.Title.required":       "Title is required",
		"Book.Author.required":      "Author is required",
		"Book.Description.required": "Description is required",
	}

	return helpers.ValidateStruct(b, customMessages)
}
