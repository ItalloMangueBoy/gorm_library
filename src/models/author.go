package models

import (
	"library/src/helpers"
	"time"
)

// Author represents a book writer in the library
type Author struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"       gorm:"not null"               validate:"required"`
	Age       uint64    `json:"age"        gorm:"not null"               validate:"required"`
	Biography *string   `json:"biography"                                validate:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}

// Format formats Author fields for better presentation
func (a *Author) Format() *Author {
	a.Name = helpers.FormatString(a.Name)
	a.Biography = helpers.FormatOptionalString(a.Biography)
	return a
}

// Valid returns custom validation error messages for Author
func (a *Author) Valid() helpers.ErrorMessages {
	customMessages := helpers.ErrorMessages{
		"Author.Name.required":      "Name is required",
		"Author.Age.required":       "Age is required",
		"Author.Biography.required": "Biography is required",
	}
	return helpers.ValidateStruct(a, customMessages)
}
