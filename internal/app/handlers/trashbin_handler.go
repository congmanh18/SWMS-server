package handlers

import (
	"fmt"
	"log"
	"net/http"
	"smart-waste-system/internal/app/models"
	"smart-waste-system/internal/utils"
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
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

// func (th *Repository) FindTrashbin(trashBinID string, ctx *fiber.Ctx) (models.TrashBin, error) {
// 	trashBin := &models.TrashBin{}
// 	err := th.DB.Where("id = ?", trashBinID).First(trashBin).Error
// 	return *trashBin, err
// }

// FindTrashbin tìm kiếm thông tin của trash bin dựa trên ID
func (th *Repository) FindTrashbin(trashBinID string) (*models.TrashBin, error) {
	trashBin := &models.TrashBin{}
	err := th.DB.Where("id = ?", trashBinID).First(trashBin).Error
	if err != nil {
		return nil, err
	}
	return trashBin, nil
}

func (th *Repository) ReadTrashBin(ctx *fiber.Ctx) error {
	trashBinID := ctx.Params("id")
	if err := utils.MatchUserTypeToUID(ctx, trashBinID); err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Request Get failed")
	}

	trashBin, err := th.FindTrashbin(trashBinID)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "trash bin not found")
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"level":    trashBin.TrashLevel,
		"location": trashBin.Location,
		"message":  "read trash bin successfully"})
}
func (th *Repository) UpdateTrashBin(ctx *fiber.Ctx) error {
	trashBinID := ctx.Params("id")
	if err := utils.MatchUserTypeToUID(ctx, trashBinID); err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Request Get failed")
	}

	trashBin, err := th.FindTrashbin(trashBinID)
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

// Định nghĩa cấu trúc cho các thông điệp WebSocket
type WebSocketMessage struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

func (th *Repository) WebSocketReadTrashBin(c *websocket.Conn) {
	// Lấy ID của trash bin từ thông điệp WebSocket
	var message WebSocketMessage
	for {
		if err := c.ReadJSON(&message); err != nil {
			log.Println("Error reading WebSocket message:", err)
			return
		}

		// Xác định hành động được yêu cầu từ thông điệp
		switch message.Action {
		case "read":
			trashBinID := fmt.Sprintf("%v", message.Data)
			trashBin, err := th.FindTrashbin(trashBinID)
			if err != nil {
				sendWebSocketError(c, "Trash bin not found")
				continue
			}

			if err := c.WriteJSON(trashBin); err != nil {
				log.Println("Error sending WebSocket message:", err)
				return
			}
		default:
			log.Println("Unknown action:", message.Action)
		}
	}
}

// Hàm xử lý WebSocket để cập nhật thông tin của trash bin
func (th *Repository) WebSocketUpdateTrashBin(c *websocket.Conn) {
	// Lặp để nhận và xử lý các thông điệp WebSocket
	var message WebSocketMessage
	for {
		if err := c.ReadJSON(&message); err != nil {
			log.Println("Error reading WebSocket message:", err)
			return
		}

		// Xác định hành động được yêu cầu từ thông điệp
		switch message.Action {
		case "update":
			// Lấy ID của trash bin và thông tin mới từ thông điệp WebSocket
			payload, ok := message.Data.(map[string]interface{})
			if !ok {
				log.Println("Invalid data format")
				continue
			}

			trashBinID, ok := payload["id"].(string)
			if !ok {
				log.Println("Invalid trash bin ID")
				continue
			}

			// Tìm trash bin từ ID và cập nhật thông tin
			trashBin, err := th.FindTrashbin(trashBinID)
			if err != nil {
				sendWebSocketError(c, "Trash bin not found")
				continue
			}

			// Cập nhật thông tin từ payload và lưu vào cơ sở dữ liệu
			if err := c.ReadJSON(&trashBin); err != nil {
				log.Println("Error decoding WebSocket message:", err)
				continue
			}

			if err := th.DB.Save(&trashBin).Error; err != nil {
				sendWebSocketError(c, "Error updating trash bin")
				continue
			}

			// Gửi thông điệp xác nhận cập nhật thành công qua WebSocket
			response := WebSocketMessage{
				Action: "update_success",
				Data:   "Trash bin updated successfully",
			}
			if err := c.WriteJSON(response); err != nil {
				log.Println("Error sending WebSocket message:", err)
				return
			}
		default:
			log.Println("Unknown action:", message.Action)
		}
	}
}

// Hàm gửi thông điệp lỗi qua WebSocket
func sendWebSocketError(c *websocket.Conn, message string) {
	errMessage := WebSocketMessage{
		Action: "error",
		Data:   message,
	}
	if err := c.WriteJSON(errMessage); err != nil {
		log.Println("Error sending WebSocket error message:", err)
	}
}
