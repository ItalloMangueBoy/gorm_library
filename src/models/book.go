package models

type Book struct {
	ID          int64    `json:"id"`
	Pages       int64    `json:"page"`
	Available   bool     `json:"available"`
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Description *string  `json:"description"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}
