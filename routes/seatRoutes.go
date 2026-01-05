package routes

import (
	"cinema/controllers"
	"cinema/middleware"

	"github.com/gofiber/fiber/v2"
)

func SeatRoutes(app *fiber.App) {
	app.Get("/seat", controllers.GetSeat)
	app.Get("/seat/:hall_id", controllers.GetSeatByID)
	app.Get("/showtime/:showtimes_id/seats", controllers.GetSeatsByShowtime)
	admin := app.Group("/admin", middleware.Protected(), middleware.AdminOnly())
	admin.Post("/seat", controllers.CreateSeat)
	admin.Put("/seat/:id", controllers.UpdateSeat)
	admin.Delete("/seat/:id", controllers.DeleteSeat)
	admin.Patch("/seat/:id/status", controllers.ToggleSeatStatus)
}
