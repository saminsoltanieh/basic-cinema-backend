// This file handles database connection and initialization using gorm
package config

import (
	"cinema/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is a global database connection
var DB *gorm.DB

// connect to db
func ConnectDB() {
	//does env exist
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	//database connection setting
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "")
	password := getEnv("DB_PASSWORD", "")
	dbname := getEnv("DB_NAME", "")
	//data source name
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	//open a connection
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//error
	if err != nil {
		panic("failed to connect to the database!" + err.Error())
	}
	DB = database
	fmt.Println("Database connection successful!")
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Hall{})
	DB.AutoMigrate(&models.Seat{})
	DB.AutoMigrate(&models.Movie{})
	DB.AutoMigrate(&models.Showtime{})
	DB.AutoMigrate(&models.Booking{})
}
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
