package handlers

import (
	"net/http"
	"smart-waste-system/internal/app/models"
	"smart-waste-system/internal/utils"
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (uh *Repository) Register(ctx *fiber.Ctx) error {
	user := models.User{}
	user.ID = uuid.New().String()

	err := ctx.BodyParser(&user)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusUnprocessableEntity, "Invalid request payload")
	}
	validationErr := validator.New().Struct(user)
	if validationErr != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Validation failed")
	}

	token, _ := utils.GenerateAllTokens(user.Phone, user.ID, *user.FirstName, *user.MiddleName, *user.LastName)
	user.Token = &token

	password := utils.HashPassword(*user.Password)
	user.Password = &password

	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	err = uh.DB.Create(&user).Error
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Could not create user")
	}

	// Trả về thành công nếu không có lỗi
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"info":    user,
		"message": "register successfully"})
}

func (uh *Repository) Login(ctx *fiber.Ctx) error {
	user := models.User{}
	foundUser := models.User{}

	err := ctx.BodyParser(&user)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusUnprocessableEntity, "Invalid request payload")
	}
	if err := uh.DB.Where("id = ?", user.ID).First(&foundUser).Error; err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusUnprocessableEntity, err.Error())
	}

	passwordIsValid, msg := utils.VerifyPassword(*user.Password, *foundUser.Password)
	if !passwordIsValid {
		return utils.HandleErrorResponse(ctx, http.StatusInternalServerError, msg)
	}

	token, refreshToken := utils.GenerateAllTokens(foundUser.Phone, foundUser.ID, *foundUser.FirstName, *foundUser.MiddleName, *foundUser.LastName)
	utils.UpdateAllTokens(uh.DB, foundUser.ID, token, refreshToken)

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message":    "login successfully",
		"found user": foundUser})
}

func (uh *Repository) ViewInfo(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")
	user := &models.User{}
	if err := utils.MatchUserTypeToUID(ctx, userID); err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Request Get failed")
	}

	err := uh.DB.Where("id = ?", userID).First(user).Error
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "user not found")
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "information",
		"data":    user})
}

func (uh *Repository) UpdateInfo(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Update Information Successfully"})
}

func (uh *Repository) DeleteUser(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Delete User Successfully"})
}
