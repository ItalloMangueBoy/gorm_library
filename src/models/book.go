package models

import (
	"library/src/helpers"
	"time"
)

// Book represents a book in the library
type Book struct {
	ID          uint64    `json:"id"`
	Pages       uint64    `json:"pages"        gorm:"not null"               validate:"required"`
	Avaliable   *bool     `json:"avaliable"    gorm:"not null;default:false" validate:"required"`
	Title       string    `json:"title"        gorm:"not null"               validate:"required"`
	Author      string    `json:"author"       gorm:"not null"               validate:"required"`
	Description *string   `json:"description"                                validate:"required"`
	CreatedAt   time.Time `json:"created_at"   gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at"   gorm:"not null"`
}

// Fornmat formats the Book struct to a more readable format
func (b *Book) Format() *Book {
	b.Title = helpers.FormatString(b.Title)
	b.Author = helpers.FormatString(b.Author)
	b.Description = helpers.FormatOptionalString(b.Description)

	return b
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
