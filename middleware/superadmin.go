package middleware

import (
	"cinema/config"
	"cinema/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func IsSuperAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userData := c.Locals("user")
		if userData == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
		}

		claims, ok := userData.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid user data"})
		}

		idFloat, ok := claims["id"].(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid user id in token"})
		}
		userID := uint(idFloat)

		var user models.User
		if err := config.DB.First(&user, userID).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "user not found"})
		}

		if !user.IsSuperAdmin {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "only superadmin can perform this action"})
		}

		return c.Next()
	}
}
