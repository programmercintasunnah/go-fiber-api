package repositories

import (
	database "go-fiber-api/db"
	"go-fiber-api/models"
)

func FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func UpdateUser(user *models.User) error {
	return database.DB.Save(user).Error
}
