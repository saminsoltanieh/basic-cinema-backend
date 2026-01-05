package routes

import (
	"cinema/controllers"
	"cinema/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Post("/user/register", controllers.RegisterUser)
	app.Post("/user/login", controllers.LoginUser)
	admin := app.Group("/admin", middleware.Protected(), middleware.AdminOnly())
	admin.Get("/user", controllers.GetAllUser)
	app.Get("/showtimes/:id/seats", middleware.Protected(), controllers.GetSeatByShowtime)
	app.Post("/book-seats", middleware.Protected(), controllers.ReserveSeats)
	app.Get("/my-bookings", middleware.Protected(), controllers.GetMyBooking)
	app.Put("/booking/:id/cancel", middleware.Protected(), controllers.CancelBooking)
}
