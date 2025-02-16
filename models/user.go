package models

type User struct {
	ID           uint   `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"-"` // Jangan kirim password dalam JSON response
	FailedLogins int    `json:"failed_logins"`
	LockedUntil  int64  `json:"locked_until"`
}
