// CRUD for hall
package controllers

import (
	"cinema/config"
	"cinema/models"

	"github.com/gofiber/fiber/v2"
)

// create hall
func CreateHall(c *fiber.Ctx) error {
	var hall models.Hall
	if err := c.BodyParser(&hall); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}
	if err := config.DB.Create(&hall).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "cannot creat hall"})
	}
	return c.Status(201).JSON(hall)
}

// get all halls
func GetHall(c *fiber.Ctx) error {
	var halls []models.Hall
	result := config.DB.Find(&halls)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch hall"})
	}
	return c.JSON(halls)
}

// get a hall by id
func GetHallById(c *fiber.Ctx) error {
	id := c.Params("id")
	var hall models.Hall
	if err := config.DB.First(&hall, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "hall not found",
		})
	}
	return c.JSON(hall)
}

// update a hall
func UpdateHall(c *fiber.Ctx) error {
	id := c.Params("id")
	var hall models.Hall
	if err := config.DB.First(&hall, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "hall not found"})
	}
	if err := c.BodyParser(&hall); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannpt parse JSON"})
	}
	config.DB.Save(&hall)
	return c.JSON(hall)
}

// delete hall
func DeleteHall(c *fiber.Ctx) error {
	id := c.Params("id")
	var hall models.Hall
	if err := config.DB.First(&hall, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "hall not found",
		})
	}
	config.DB.Delete(&hall)
	return c.JSON(fiber.Map{
		"message": "hall deleted successfully",
	})
}
