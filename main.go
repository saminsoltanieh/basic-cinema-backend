package main

import (
	"cinema/config"
	"cinema/controllers"
	"cinema/routes"
	"cinema/superadmin"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("JWT_SECRET =", os.Getenv("JWT_SECRET"))
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
