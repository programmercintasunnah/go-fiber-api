package models

type Book struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
	Year        int    `json:"year"`
	Description string `json:"description"`
}
