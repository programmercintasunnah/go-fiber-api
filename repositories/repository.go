package repositories

import (
	database "go-fiber-api/db"
	"go-fiber-api/models"
	"log"
)

var BookRepository *BaseRepository[models.Book]
var IktikafRepository *BaseRepository[models.Iktikaf]

// Pastikan fungsi init() tidak langsung dipanggil sebelum database diinisialisasi
func Init() {
	if database.DB == nil {
		log.Fatal("Database is not initialized. Initialization failed.")
	}

	// Initialize repositories
	BookRepository = NewBaseRepository[models.Book](database.DB)
	IktikafRepository = NewBaseRepository[models.Iktikaf](database.DB)
}
