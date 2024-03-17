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

func (th *Repository) FindTrashbin(trashBinID string, ctx *fiber.Ctx) (models.TrashBin, error) {
	trashBin := &models.TrashBin{}
	err := th.DB.Where("id = ?", trashBinID).First(trashBin).Error
	return *trashBin, err
}

func (th *Repository) ReadTrashBin(ctx *fiber.Ctx) error {
	trashBinID := ctx.Params("id")
	if err := utils.MatchUserTypeToUID(ctx, trashBinID); err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Request Get failed")
	}

	trashBin, err := th.FindTrashbin(trashBinID, ctx)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "trash bin not found")
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"level":    trashBin.TrashLevel,
		"location": trashBin.Location,
		"message":  "create trash bin successfully"})
}
func (th *Repository) UpdateTrashBin(ctx *fiber.Ctx) error {
	trashBinID := ctx.Params("id")
	if err := utils.MatchUserTypeToUID(ctx, trashBinID); err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Request Get failed")
	}

	trashBin, err := th.FindTrashbin(trashBinID, ctx)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "trashBin not found")
	}
	ctx.BodyParser(&trashBin)
	th.DB.Save(&trashBin)

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"info":    trashBin,
		"message": "Update Information Successfully"})
}
func (th *Repository) DeleteTrashBin(ctx *fiber.Ctx) error {
	trashBin := models.TrashBin{}
	id := ctx.Params("id")
	if id == "" {
		return utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "id cannot be empty")
	}
	err := th.DB.Where("id = ?", id).Delete(&trashBin)
	if err.Error != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "could not delete trashBin")
	}
	utils.HandleErrorResponse(ctx, http.StatusOK, "delete trashBin successfull")
	return nil
}
