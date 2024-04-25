package models

import (
	"time"

	"gorm.io/gorm"
)

type TrashBin struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	TrashLevel *float32  `json:"trash_level"` // chua biet
	Address    *string   `json:"address"`
	Location   *string   `json:"location_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func MigrateTrashBins(db *gorm.DB) error {
	return db.AutoMigrate(&TrashBin{})
}
