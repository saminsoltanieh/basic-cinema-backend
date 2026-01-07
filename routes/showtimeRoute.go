package routes

import (
	"cinema/controllers"
	"cinema/middleware"

	"github.com/gofiber/fiber/v2"
)

func ShowtimeRoutes(app *fiber.App) {
	app.Get("/showtime", controllers.GetShowtimes)
	app.Get("/showtime/:id", controllers.GetShowtimeById)
	admin := app.Group("/admin", middleware.Protected(), middleware.AdminOnly())
	admin.Post("/showtime", controllers.CreateShowtime)
	admin.Put("/showtime/:id", controllers.UpdateShowtime)
	admin.Delete("/showtime/:id", controllers.DeleteShowtime)
}
