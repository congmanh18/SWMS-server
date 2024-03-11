package middleware

import (
	"net/http"
	"smart-waste-system/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func Authenticate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		clientToken := c.Get("token")
		if clientToken == "" {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "No Authorization header provided"})
		}

		claims, err := utils.ValidateToken(clientToken)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}

		c.Locals("phone", claims.Phone)
		c.Locals("first_name", claims.FirstName)
		c.Locals("middle_name", claims.MiddleName)
		c.Locals("last_name", claims.LastName)
		return c.Next()
	}
}
