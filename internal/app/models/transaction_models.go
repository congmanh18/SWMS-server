package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID        string `json:"id" gorm:"primaryKey"`
	StaffName string `json:"staff_name"`
	Address   string `json:"address"`
	Weight    string `json:"weight"`
	///
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MigrateTransaction(db *gorm.DB) error {
	return db.AutoMigrate(&Transaction{})

}
