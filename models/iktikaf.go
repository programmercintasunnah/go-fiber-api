package models

type Iktikaf struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Notes     string `json:"notes"`
}
