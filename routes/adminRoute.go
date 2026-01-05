package routes

import (
	"cinema/controllers"
	"cinema/middleware"

	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(app *fiber.App) {
	admin := app.Group("/admin", middleware.Protected(), middleware.AdminOnly())
	//Halls route for admin
	admin.Post("/halls", controllers.CreateHall)
	admin.Get("/halls", controllers.GetHall)
	admin.Get("/halls/:id", controllers.GetHallById)
	admin.Put("/halls/:id", controllers.UpdateHall)
	admin.Delete("/halls/:id", controllers.DeleteHall)
	app.Put("/make-admin/:id", middleware.Protected(), middleware.IsSuperAdmin(), controllers.MakeAdmin)
}
