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

func (th *Repository) CreateTrashBin(ctx *fiber.Ctx) error {
	trashBin := models.TrashBin{}
	trashBin.ID = uuid.New().String()
	err := ctx.BodyParser(&trashBin)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusUnprocessableEntity, "Invalid request payload")
	}
	validationErr := validator.New().Struct(trashBin)
	if validationErr != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Validation failed")
	}

	//trash level?
	trashBin.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	trashBin.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	err = th.DB.Create(&trashBin).Error
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Could not create trash bin")
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"info":    trashBin,
		"message": "create trash bin successfully"})
}

func (th *Repository) ReadTrashBin(ctx *fiber.Ctx) error {
	trashBinID := ctx.Params("id")
	if err := utils.MatchUserTypeToUID(ctx, trashBinID); err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Request Get failed")
	}
	trashBin := models.TrashBin{}
	err := th.DB.Where("id = ?", trashBinID).First(trashBin).Error
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "trash bin not found")
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"level":   trashBin.TrashLevel,
		"info":    trashBin.Location,
		"message": "create trash bin successfully"})
}
func (th *Repository) UpdateTrashBin(ctx *fiber.Ctx) error {
	return nil
}
func (th *Repository) DeleteTrashBin(ctx *fiber.Ctx) error {
	return nil
}
