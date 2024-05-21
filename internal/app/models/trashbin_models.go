package models

import (
	"time"

	"gorm.io/gorm"
)

type TrashBin struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Weight    *float32  `json:"weight"`
	Level     *float32  `json:"level"`
	Address   *string   `json:"address"`
	Latitude  *float64  `json:"latitude"`
	Longitude *float64  `json:"longitude"`
	AreaID    *string   `json:"area_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	//Distance - Time
}

func MigrateTrashBins(db *gorm.DB) error {
	return db.AutoMigrate(&TrashBin{})
}
