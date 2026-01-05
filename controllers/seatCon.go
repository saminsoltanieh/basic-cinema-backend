package controllers

import (
	"cinema/config"
	"cinema/models"

	"github.com/gofiber/fiber/v2"
)

// create seat
func CreateSeat(c *fiber.Ctx) error {
	var seat models.Seat
	if err := c.BodyParser(&seat); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}
	if err := config.DB.Create(&seat).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "cannot create seat"})
	}
	return c.Status(201).JSON(seat)
}

// get seat
func GetSeat(c *fiber.Ctx) error {
	var seats []models.Seat
	if err := config.DB.Find(&seats).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch seats"})
	}
	return c.JSON(seats)
}

// get seat by id
func GetSeatByID(c *fiber.Ctx) error {
	hallID := c.Params("hall_id")
	var seats []models.Seat
	if err := config.DB.Where("hall_id=?", hallID).Find(&seats).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "seat not found"})
	}
	return c.JSON(seats)
}

// update seat
func UpdateSeat(c *fiber.Ctx) error {
	id := c.Params("id")
	var seat models.Seat
	if err := config.DB.First(&seat, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "seat not found"})
	}
	if err := c.BodyParser(&seat); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}
	config.DB.Save(&seat)
	return c.JSON(seat)
}

// delete seat
func DeleteSeat(c *fiber.Ctx) error {
	id := c.Params("id")
	var seat models.Seat
	if err := config.DB.First(&seat, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "cannot find seat"})
	}
	config.DB.Delete(&seat)
	return c.JSON(fiber.Map{"message": "seat deleted successfully"})
}
func ToggleSeatStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var seat models.Seat
	if err := config.DB.First(&seat, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "seat not found"})
	}
	seat.IsActive = !seat.IsActive
	config.DB.Save(&seat)
	return c.JSON(seat)
}

//return all seats for specific showtime

func GetSeatsByShowtime(c *fiber.Ctx) error {
	showtimeID := c.Params("id")
	var showtime models.Showtime
	if err := config.DB.First(&showtime, showtimeID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "showtime not found"})
	}
	var seats []models.Seat
	if err := config.DB.Where("hall_id=?", showtime.HallID).Find(&seats).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "seats not found"})
	}
	return c.JSON(seats)
}
