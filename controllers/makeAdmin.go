package controllers

import (
	"cinema/config"
	"cinema/models"

	"github.com/gofiber/fiber/v2"
)

func MakeAdmin(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}
	if user.Role == "admin" || user.Role == "superadmin" {
		return c.Status(400).JSON(fiber.Map{"error": "user is already an admin or superadmin"})
	}
	user.Role = "admin"
	config.DB.Save(&user)
	return c.JSON(fiber.Map{"message": "user promoted to admin"})
}
