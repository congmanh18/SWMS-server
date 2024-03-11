package repository

import (
	"net/http"
	"smart-waste-system/internal/app/models"
	"smart-waste-system/internal/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{Repository: Repository{DB: db}}
}

func (ur *UserRepository) CreateAccount(ctx *fiber.Ctx, user models.User) error {
	validationErr := validator.New().Struct(user)
	if validationErr != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Validation failed")
	}

	token, refreshToken := utils.GenerateAllTokens(*user.Phone, user.ID, *user.FirstName, *user.MiddleName, *user.LastName)
	user.Token = &token
	user.RefreshToken = &refreshToken

	password := utils.HashPassword(*user.Password)
	user.Password = &password

	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	// err = NewUserRepository(ur.DB.Create(&user))
	err := ur.DB.Create(&user).Error
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "Failed to create user")
	}
	return nil
}

// Signin by phone number
func (ur *UserRepository) SigninByPhoneNumber(ctx *fiber.Ctx, user models.User) (*models.User, error) {
	foundUser := models.User{}

	if err := ur.DB.Where("phone = ?", user.Phone).First(&foundUser).Error; err != nil {
		return nil, utils.HandleErrorResponse(ctx, http.StatusUnprocessableEntity, err.Error())
	}

	passwordIsValid, msg := utils.VerifyPassword(*user.Password, *foundUser.Password)
	if !passwordIsValid {
		return nil, utils.HandleErrorResponse(ctx, http.StatusInternalServerError, msg)
	}

	if foundUser.Phone == nil {
		return nil, utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "user not found")
	}

	token, refreshToken := utils.GenerateAllTokens(*foundUser.Phone, foundUser.ID, *foundUser.FirstName, *foundUser.MiddleName, *foundUser.LastName)
	utils.UpdateAllTokens(ur.DB, foundUser.ID, token, refreshToken)

	return &foundUser, nil
}

func (ur *UserRepository) GetUserByID(ctx *fiber.Ctx, userID string) (*models.User, error) {
	user := models.User{}
	if err := utils.MatchUserTypeToUID(ctx, userID); err != nil {
		return nil, utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Request Get failed")
	}

	err := ur.DB.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "user not found")
	}
	return &user, nil
}

func (ur *UserRepository) EditInfo(ctx *fiber.Ctx, userID string, updatedInfo models.User) (*models.User, error) {
	user, err := ur.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// update fields like FirstName, LastName, etc.
	user.FirstName = updatedInfo.FirstName
	user.MiddleName = updatedInfo.MiddleName
	user.LastName = updatedInfo.LastName
	user.Password = updatedInfo.Password

	// Save the updated user information to the database
	if err := ur.DB.Save(&user).Error; err != nil {
		return nil, utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "Failed to update user")
	}

	return user, nil
}
