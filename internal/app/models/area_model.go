package models

import (
	"time"

	"gorm.io/gorm"
)

type Area struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      *string   `json:"name"`
	Address   *string   `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MigrateArea(db *gorm.DB) error {
	return db.AutoMigrate(&TrashBin{})
}
