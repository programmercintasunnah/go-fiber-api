package models

type Iktikaf struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	Email     string `gorm:"unique;not null" json:"email"`
	Phone     string `gorm:"not null" json:"phone"`
	StartDate string `gorm:"not null" json:"start_date"`
	EndDate   string `gorm:"not null" json:"end_date"`
	Notes     string `json:"notes"`
}
