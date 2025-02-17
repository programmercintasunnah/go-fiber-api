package repositories

import (
	"go-fiber-api/utils"

	"gorm.io/gorm"
)

// BaseRepository adalah struct untuk repository generik
type BaseRepository[T any] struct {
	DB *gorm.DB
}

// NewBaseRepository membuat instance baru dari BaseRepository
func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{DB: db}
}

// GetAll mengambil semua data
func (r *BaseRepository[T]) GetAll() ([]T, error) {
	var items []T
	err := r.DB.Find(&items).Error
	return items, err
}

// GetByID mengambil satu data berdasarkan ID
func (r *BaseRepository[T]) GetByID(id string) (T, error) {
	var item T
	err := r.DB.First(&item, id).Error
	return item, err
}

// Create menambahkan data baru
func (r *BaseRepository[T]) Create(item *T) error {
	return r.DB.Create(item).Error
}

// Update mengupdate data berdasarkan ID
func (r *BaseRepository[T]) Update(id string, item *T) error {
	var entity T
	return r.DB.Model(&entity).Where("id = ?", id).Updates(item).Error
}

func (r *BaseRepository[T]) Delete(id string) error {
	var entity T
	return r.DB.Delete(&entity, id).Error
}

func (r *BaseRepository[T]) Search(params utils.SearchParams) ([]T, int64, error) {
	var items []T
	var total int64

	var entity T
	query := r.DB.Model(&entity)

	// Filter berdasarkan search value
	if params.SearchValue != "" {
		query = query.Where("title ILIKE ? OR author ILIKE ?", "%"+params.SearchValue+"%", "%"+params.SearchValue+"%")
	}

	// Filter tambahan
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
	if err := query.Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}
