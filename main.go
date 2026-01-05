package main

import (
	"cinema/config"
	"cinema/controllers"
	"cinema/routes"
	"cinema/superadmin"
	"fmt"
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
	go func() {
		for {
			controllers.Deactiver()
			time.Sleep(1 * time.Minute)
		}
	}()
	if err := app.Listen(":3001"); err != nil {
		panic(fmt.Errorf("failed to start server: %v", err))
	}

}
