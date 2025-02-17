package repositories

import (
	database "go-fiber-api/db"
	"go-fiber-api/models"
)

var BookRepository *BaseRepository[models.Book]

func InitRepositories() {
	if database.DB == nil {
		panic("Database connection is nil")
	}
	BookRepository = NewBaseRepository[models.Book](database.DB)
}
