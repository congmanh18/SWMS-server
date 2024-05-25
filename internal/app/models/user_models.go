package models

import (
	"time"

	"gorm.io/gorm"
)

// username va email bo sung xac thuc dang nhap sau
// xem lai category nen dung string hay dung pointer string
// xem lai role nen dung gi
type User struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	FirstName    *string   `json:"first_name"`
	MiddleName   *string   `json:"middle_name"`
	LastName     *string   `json:"last_name"`
	Gender       *string   `json:"gender"`
	RoleName     *string   `json:"role_name"`
	DateOfBirth  *string   `json:"date_of_birth"`
	Nationality  *string   `json:"nationality"`
	POR          *string   `json:"por"`
	POO          *string   `json:"poo"`
	CIN          *string   `json:"cin"`
	Email        *string   `json:"email"`
	Username     *string   `json:"username"` // validate: user
	Phone        string    `json:"phone" validate:"required"`
	Password     *string   `json:"password"`
	Category     *string   `json:"category"`
	Token        *string   `json:"token"`
	RefreshToken *string   `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	// Role         string    `json:"role"`
}

func MigrateUser(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
