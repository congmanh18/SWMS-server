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

func (th *Repository) FindTransaction(transactionID string, ctx *fiber.Ctx) (models.Transaction, error) {
	transaction := &models.Transaction{}
	err := th.DB.Where("id = ?", transactionID).First(transaction).Error
	return *transaction, err
}

func (th *Repository) ReadTransaction(ctx *fiber.Ctx) error {
	transactionID := ctx.Params("id")
	if err := utils.MatchUserTypeToUID(ctx, transactionID); err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Request Get failed")
	}

	transaction, err := th.FindTransaction(transactionID, ctx)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "transaction not found")
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"info":    transaction,
		"message": "create transaction successfully"})
}
func (th *Repository) UpdateTransaction(ctx *fiber.Ctx) error {
	transactionID := ctx.Params("id")
	if err := utils.MatchUserTypeToUID(ctx, transactionID); err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Request Get failed")
	}

	transaction, err := th.FindTransaction(transactionID, ctx)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "transaction not found")
	}
	ctx.BodyParser(&transaction)
	th.DB.Save(&transaction)

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"info":    transaction,
		"message": "Update Information Successfully"})
}
func (th *Repository) DeleteTransaction(ctx *fiber.Ctx) error {
	transaction := models.Transaction{}
	id := ctx.Params("id")
	if id == "" {
		return utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "id cannot be empty")
	}
	err := th.DB.Where("id = ?", id).Delete(&transaction)
	if err.Error != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "could not delete transaction")
	}
	utils.HandleErrorResponse(ctx, http.StatusOK, "delete transaction successfull")
	return nil
}
