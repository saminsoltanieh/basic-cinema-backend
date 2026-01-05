package main

import (
	"cinema/config"
	"cinema/controllers"
	"cinema/routes"
	"cinema/superadmin"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDB()
	superadmin.CreateSuperAdmin()
	app := fiber.New()
	routes.AdminRoutes(app)
	routes.UserRoutes(app)
	routes.MovieRoutes(app)
	routes.ShowtimeRoutes(app)
	routes.SeatRoutes(app)
	app.Listen(":3001")
	go func() {
		for {
			controllers.Deactiver()
			time.Sleep(1 * time.Minute)
		}
	}()
}
