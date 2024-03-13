package handlers

import (
	"fmt"
	"net/http"
	"smart-waste-system/internal/app/models"
	"smart-waste-system/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (uh *Repository) Muoba(context *fiber.Ctx) error {
	user := models.User{}
	user.ID = uuid.New().String()

	err := context.BodyParser(&user)
	fmt.Println(user)
	if err != nil {
		return utils.HandleErrorResponse(context, http.StatusUnprocessableEntity, "Request failed")
	}
	if err := uh.CreateAccount(context, user); err != nil {
		return utils.HandleErrorResponse(context, http.StatusInternalServerError, "Failed to register user")
	}

	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user has been added"})
}

func (uh *Repository) Register(ctx *fiber.Ctx) error {
	user := models.User{}
	user.ID = uuid.New().String()

	err := ctx.BodyParser(&user)
	fmt.Println(user)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusUnprocessableEntity, "Invalid request payload")
	}

	// Gọi hàm đăng ký user từ Repository
	if err := uh.CreateAccount(ctx, user); err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "Failed to register user")
	}

	// Trả về thành công nếu không có lỗi
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "register successfully"})
}

func (uh *Repository) Login(ctx *fiber.Ctx) error {
	user := models.User{}

	err := ctx.BodyParser(&user)
	// fmt.Println(err)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusUnprocessableEntity, "Invalid request payload")
	}

	foundUser, err := uh.SigninByPhoneNumber(ctx, user)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Request Login failed")
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message":    "login successfully",
		"found user": foundUser})
}

func (uh *Repository) ViewInfo(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")

	user, err := uh.GetUserByID(ctx, userID)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "Failed to get user by ID")
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
