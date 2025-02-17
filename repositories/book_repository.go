package repositories

import (
	database "go-fiber-api/db"
	"go-fiber-api/models"
	"go-fiber-api/utils"
)

func GetBooks() ([]models.Book, error) {
	var books []models.Book
	err := database.DB.Find(&books).Error
	return books, err
}

func GetBookByID(id string) (*models.Book, error) {
	var book models.Book
	err := database.DB.First(&book, id).Error
	return &book, err
}

func CreateBook(book *models.Book) error {
	return database.DB.Create(book).Error
}

func UpdateBook(id string, book *models.Book) error {
	return database.DB.Model(&models.Book{}).Where("id = ?", id).Updates(book).Error
}

func DeleteBook(id string) error {
	return database.DB.Delete(&models.Book{}, id).Error
}

func SearchBooks(params utils.SearchParams) ([]models.Book, int64, error) {
	var books []models.Book
	var total int64

	query := database.DB.Model(&models.Book{})

	// Filter berdasarkan search value (judul atau penulis)
	if params.SearchValue != "" {
		query = query.Where("title ILIKE ? OR author ILIKE ?", "%"+params.SearchValue+"%", "%"+params.SearchValue+"%")
	}

	// Filter tambahan jika ada
	for key, value := range params.Filters {
		query = query.Where(key+" = ?", value)
	}

	// Sorting
	if params.OrderBy.Field != "" && (params.OrderBy.Type == "asc" || params.OrderBy.Type == "desc") {
		query = query.Order(params.OrderBy.Field + " " + params.OrderBy.Type)
	}

	// Populate (preload)
	for _, relation := range params.Populate {
		query = query.Preload(relation)
	}

	// Hitung total data
	query.Count(&total)

	// Pagination
	offset := (params.Page - 1) * params.PageSize
	query = query.Offset(offset).Limit(params.PageSize)

	// Eksekusi query
	if err := query.Find(&books).Error; err != nil {
		return nil, 0, err
	}

	return books, total, nil
}
