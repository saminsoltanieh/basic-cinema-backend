package routes

import (
	"cinema/controllers"
	"cinema/middleware"

	"github.com/gofiber/fiber/v2"
)

func MovieRoutes(app *fiber.App) {
	app.Get("/movies", controllers.GetMovies)
	app.Get("/movies/:id", controllers.GetMovieByID)
	admin := app.Group("/admin", middleware.Protected(), middleware.AdminOnly())
	admin.Post("/movies", controllers.CreateMovie)
	admin.Put("/movies/:id", controllers.UpdateMovie)
	admin.Delete("/movies/:id", controllers.DeleteMovie)
}
