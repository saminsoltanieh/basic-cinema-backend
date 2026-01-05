package controllers

import (
	"cinema/config"
	"cinema/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// register new user
func RegisterUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}
	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "cannot hash password",
		})
	}
	user.Password = string(hashedPassword)
	user.Role = "user"

	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "cannot creat user",
		})
	}
	return c.Status(201).JSON(user)
}

// verify jwt token
var jwtSecret = []byte("mysecretkey")

// login
func LoginUser(c *fiber.Ctx) error {
	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "cannot parse JSON"})
	}
	var user models.User
	if err := config.DB.Where("email = ?", data.Email).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "invalid credentials",
		})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "invalid credentials",
		})
	}
	//making jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "cannot generate token"})
	}
	return c.JSON(fiber.Map{"token": tokenString})
}
func GetAllUser(c *fiber.Ctx) error {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch users",
		})
	}
	return c.JSON(users)
}
