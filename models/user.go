package models

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"unique;not null"`
	Email        string `gorm:"unique;not null"`
	Password     string `gorm:"not null"`
	FailedLogins int    `gorm:"not null;default:0"`
	LockedUntil  int64  `gorm:"not null;default:0"`
}
