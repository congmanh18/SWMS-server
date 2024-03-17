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

func (th *Repository) CreateTransaction(ctx *fiber.Ctx) error {
	transaction := models.Transaction{}
	transaction.ID = uuid.New().String()
	err := ctx.BodyParser(&transaction)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusUnprocessableEntity, "Invalid request payload")
	}
	validationErr := validator.New().Struct(transaction)
	if validationErr != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Validation failed")
	}
	transaction.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	transaction.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	err = th.DB.Create(&transaction).Error
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Could not create transaction")
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"info":    transaction,
		"message": "create transaction successfully"})
}
func (th *Repository) ReadTransaction(ctx *fiber.Ctx) error {
	return nil
}
func (th *Repository) UpdateTransaction(ctx *fiber.Ctx) error {
	return nil
}
func (th *Repository) DeleteTransaction(ctx *fiber.Ctx) error {
	return nil
}
