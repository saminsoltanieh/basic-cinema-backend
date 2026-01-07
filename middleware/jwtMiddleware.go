package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// check if a request has a valid jwt token or not
func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing or invalid token"})
		}
		tokenString := authHeader[len("Bearer "):]
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "setting server"})
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
		}
		c.Locals("user", token.Claims)
		return c.Next() //permission to continue
	}
}
