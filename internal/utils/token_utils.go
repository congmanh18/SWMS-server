package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"smart-waste-system/internal/app/models"

	jwt "github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

var SECRET_KEY string = os.Getenv("SECRET_KEY")

type SignedDetails struct {
	Uid        string
	FirstName  string
	MiddleName string
	LastName   string
	Phone      string
	Role       string
	jwt.StandardClaims
}

// Update Role sau
func GenerateAllTokens(phone string, uid string, first_name string, middle_name string, last_name string) (signedToken string, signedRefreshToken string) {
	claims := &SignedDetails{
		Uid:        uid,
		FirstName:  first_name,
		MiddleName: middle_name,
		LastName:   last_name,
		Phone:      phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken
}
func ValidateToken(signedToken string) (claims *SignedDetails, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},

		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("the token is invalid: %w", err)
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		return nil, fmt.Errorf("the token is invalid: %w", err)
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, fmt.Errorf("token is expired: %w", err)
	}

	return claims, nil
}

func UpdateAllTokens(db *gorm.DB, userId string, signedToken, signedRefreshToken string) error {
	var user models.User

	// Find the user by ID
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		return err
	}

	// Update the user's token and refresh token
	user.Token = &signedToken
	user.RefreshToken = &signedRefreshToken
	user.UpdatedAt = time.Now()

	// Save the updated user back to the database
	if err := db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}
