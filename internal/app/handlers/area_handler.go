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

func (th *Repository) CreateArea(ctx *fiber.Ctx) error {
	area := models.Area{}
	area.ID = uuid.New().String()
	err := ctx.BodyParser(&area)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusUnprocessableEntity, "Invalid request payload")
	}
	validationErr := validator.New().Struct(area)
	if validationErr != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Validation failed")
	}
	// slice transaction
	area.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	area.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	err = th.DB.Create(&area).Error
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Could not create area")
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"info":    area,
		"message": "create trash bin successfully"})
}

func (th *Repository) FindArea(areaID string, ctx *fiber.Ctx) (models.Area, error) {
	area := &models.Area{}
	err := th.DB.Where("id = ?", areaID).First(area).Error
	return *area, err
}

func (th *Repository) ReadArea(ctx *fiber.Ctx) error {
	areaID := ctx.Params("id")
	area, err := th.FindArea(areaID, ctx)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "area not found")
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"name":         area.Name,
		"area address": area.Address,
		"message":      "Successfully fetched area"})
}

func (th *Repository) ReadListArea(ctx *fiber.Ctx) error {
	var list []models.Area
	if err := th.DB.Find(&list).Error; err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "Failed to fetch list of area")
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"info":    list,
		"message": "Successfully fetched list of area"})
}

func (th *Repository) UpdateArea(ctx *fiber.Ctx) error {
	areaID := ctx.Params("id")
	if err := utils.MatchUserTypeToUID(ctx, areaID); err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Request Get failed")
	}

	area, err := th.FindArea(areaID, ctx)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "area not found")
	}
	ctx.BodyParser(&area)
	th.DB.Save(&area)

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"info":    area,
		"message": "Update Information Successfully"})
}
func (th *Repository) DeleteArea(ctx *fiber.Ctx) error {
	area := models.Area{}
	id := ctx.Params("id")
	if id == "" {
		return utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "id cannot be empty")
	}
	err := th.DB.Where("id = ?", id).Delete(&area)
	if err.Error != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "could not delete area")
	}
	utils.HandleErrorResponse(ctx, http.StatusOK, "delete area successfull")
	return nil
}
