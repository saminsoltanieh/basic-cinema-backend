package controllers

import (
	"cinema/config"
	"cinema/models"

	"github.com/gofiber/fiber/v2"
)

func CreateMovie(c *fiber.Ctx) error {
	var movie models.Movie
	if err := c.BodyParser(&movie); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}
	if movie.Title == "" {
		return c.Status(400).JSON(fiber.Map{"error": "title is required"})
	}
	if err := config.DB.Create(&movie).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "cannot create movie",
		})
	}
	return c.Status(201).JSON(movie)
}
func GetMovies(c *fiber.Ctx) error {
	var movies []models.Movie
	if err := config.DB.Find(&movies).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Movie not found"})
	}
	return c.JSON(movies)
}
func GetMovieByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var movie models.Movie
	if err := config.DB.First(&movie, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "movie not found"})
	}
	return c.JSON(movie)
}
func UpdateMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	var movie models.Movie
	if err := config.DB.First(&movie, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "cannot find movie"})
	}
	if err := c.BodyParser(&movie); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}
	config.DB.Save(&movie)
	return c.JSON(movie)
}
func DeleteMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	var movie models.Movie
	if err := config.DB.First(&movie, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "cannot find movie"})
	}
	config.DB.Delete(&movie)
	return c.JSON(fiber.Map{"message": "movie deleted successfully"})
}
