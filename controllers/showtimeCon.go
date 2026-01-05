package controllers

import (
	"cinema/config"
	"cinema/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateShowtime(c *fiber.Ctx) error {
	var showtime models.Showtime
	if err := c.BodyParser(&showtime); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse json"})
	}
	showtime.IsActive = true
	//check the time
	hallID := showtime.HallID
	newStart := showtime.StartTime
	newEnd := showtime.EndTime
	var conflict models.Showtime
	err := config.DB.Where("hall_id=? AND start_time<? AND end_time>?", hallID, newEnd, newStart).First(&conflict).Error
	if err == nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "hall is already booked in this time",
		})
	}
	if err := config.DB.Create(&showtime).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "cannot create showtime"})
	}
	return c.Status(201).JSON(showtime)
}
func GetShowtime(c *fiber.Ctx) error {
	var showtimes []models.Showtime
	if err := config.DB.Preload("Movie").Preload("Hall").Where("is_active=?", true).Find(&showtimes).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch showtimes",
		})
	}
	return c.JSON(showtimes)
}
func GetShowtimeById(c *fiber.Ctx) error {
	id := c.Params("id")
	var showtime models.Showtime
	if err := config.DB.Preload("Movie").Preload("Hall").Where("id=? AND is_active=?", id, true).First(&showtime, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "showtime not found",
		})
	}
	return c.JSON(showtime)
}
func UpdateShowtime(c *fiber.Ctx) error {
	id := c.Params("id")
	var showtime models.Showtime
	if err := config.DB.First(&showtime, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "showtime not found!"})
	}
	if err := c.BodyParser(&showtime); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "cannot parse showtime",
		})
	}
	config.DB.Save(&showtime)
	return c.JSON(showtime)
}
func DeleteShowtime(c *fiber.Ctx) error {
	id := c.Params("id")
	var showtime models.Showtime
	if err := config.DB.First(&showtime, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "cannot find show time"})
	}
	config.DB.Delete(&showtime)
	return c.JSON(fiber.Map{
		"message": "showtime deleted successfully",
	})
}
func Deactiver() {
	now := time.Now()
	config.DB.
		Model(&models.Showtime{}).
		Where("end_time<? AND is_active =?", now, true).
		Update("is_active", false)
}
