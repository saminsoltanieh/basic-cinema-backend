package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user")
		if user == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		claims, ok := user.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot finish the process of token"})
		}
		role, ok := claims["role"].(string)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "invalid role",
			})
		}
		if role != "admin" && role != "superadmin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "admin access only",
			})
		}
		return c.Next()
	}
}
