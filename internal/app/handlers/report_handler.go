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

func (th *Repository) CreateReport(ctx *fiber.Ctx) error {
	report := models.Report{}
	report.ID = uuid.New().String()
	err := ctx.BodyParser(&report)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusUnprocessableEntity, "Invalid request payload")
	}
	validationErr := validator.New().Struct(report)
	if validationErr != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Validation failed")
	}
	// slice transaction
	report.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	report.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	err = th.DB.Create(&report).Error
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Could not create trash bin")
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"info":    report,
		"message": "create trash bin successfully"})
}

func (th *Repository) FindReport(reportID string, ctx *fiber.Ctx) (models.Report, error) {
	report := &models.Report{}
	err := th.DB.Where("id = ?", reportID).First(report).Error
	return *report, err
}

func (th *Repository) ReadReport(ctx *fiber.Ctx) error {
	reportID := ctx.Params("id")
	if err := utils.MatchUserTypeToUID(ctx, reportID); err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Request Get failed")
	}

	report, err := th.FindReport(reportID, ctx)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "report not found")
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"description":      report.Description,
		"list transaction": report.TransactionIDs,
		"message":          "create report successfully"})
}
func (th *Repository) UpdateReport(ctx *fiber.Ctx) error {
	reportID := ctx.Params("id")
	if err := utils.MatchUserTypeToUID(ctx, reportID); err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Request Get failed")
	}

	report, err := th.FindReport(reportID, ctx)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "report not found")
	}
	ctx.BodyParser(&report)
	th.DB.Save(&report)

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"info":    report,
		"message": "Update Information Successfully"})
}
func (th *Repository) DeleteReport(ctx *fiber.Ctx) error {
	report := models.Report{}
	id := ctx.Params("id")
	if id == "" {
		return utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "id cannot be empty")
	}
	err := th.DB.Where("id = ?", id).Delete(&report)
	if err.Error != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "could not delete report")
	}
	utils.HandleErrorResponse(ctx, http.StatusOK, "delete report successfull")
	return nil
}
