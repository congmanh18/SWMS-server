package handlers

import "gorm.io/gorm"

// Repository đại diện cho một repository chung
type Repository struct {
	DB *gorm.DB
}

type Handler struct {
	Repository *Repository
}

// NewRepository tạo một thể hiện mới của Repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}
