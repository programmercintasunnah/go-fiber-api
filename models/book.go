package models

type Book struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"not null" json:"title"`
	Author      string `gorm:"not null" json:"author"`
	Publisher   string `gorm:"not null" json:"publisher"`
	Year        int    `gorm:"not null" json:"year"`
	Description string `json:"description"`
}
