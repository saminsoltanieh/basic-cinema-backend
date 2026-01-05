// This file handles database connection and initialization using gorm
package config

import (
	"cinema/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is a global database connection
var DB *gorm.DB

// connect to db
func ConnectDB() {
	//database connection setting
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "samin1383"
	dbname := "cinema_db"
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
