package handlers

import (
	"fmt"
	"net/http"
	"smart-waste-system/internal/app/models"
	"smart-waste-system/internal/app/repository"
	"smart-waste-system/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler struct {
	UserRepository *repository.UserRepository
}

func NewUserHandler(userRepository *repository.UserRepository) *UserHandler {
	return &UserHandler{UserRepository: userRepository}
}

func (uh *UserHandler) Register(ctx *fiber.Ctx) error {
	user := models.User{}
	user.ID = uuid.New().String()

	// Đọc dữ liệu từ request body và chuyển đổi thành đối tượng User
	err := ctx.BodyParser(&user)
	fmt.Println(user)
	fmt.Println(&user)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusUnprocessableEntity, "Invalid request payload")
	}

	// Gọi hàm đăng ký user từ Repository
	if err := uh.UserRepository.CreateAccount(ctx, user); err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "Failed to register user")
	}

	// Trả về thành công nếu không có lỗi
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "register successfully"})
}

func (uh *UserHandler) Login(ctx *fiber.Ctx) error {
	user := models.User{}

	err := ctx.BodyParser(&user)
	// fmt.Println(err)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusUnprocessableEntity, "Invalid request payload")
	}

	foundUser, err := uh.UserRepository.SigninByPhoneNumber(ctx, user)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Request Login failed")
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message":    "login successfully",
		"found user": foundUser})
}

func (uh *UserHandler) ViewInfo(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")

	user, err := uh.UserRepository.GetUserByID(ctx, userID)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "Failed to get user by ID")
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "information",
		"data":    user})
}

func (uh *UserHandler) UpdateInfo(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Update Information Successfully"})
}

func (uh *UserHandler) DeleteUser(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Delete User Successfully"})
}
