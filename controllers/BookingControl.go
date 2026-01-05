package controllers

import (
	"cinema/config"
	"cinema/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// get all booking
func GetAllBooking(c *fiber.Ctx) error {
	var bookings []models.Booking
	if err := config.DB.Preload("User").Preload("Seat").Preload("Showtime").Find(&bookings).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch booking"})
	}
	return c.JSON(bookings)
}
func GetSeatByShowtime(c *fiber.Ctx) error {
	showtimeID := c.Params("id")
	var showtime models.Showtime
	if err := config.DB.First(&showtime, showtimeID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "showtime not found"})
	}
	var seats []models.Seat
	if err := config.DB.Where("hall_id=?", showtime.HallID).Find(&seats).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "seats not found"})
	}
	type SeatsWithStatus struct {
		ID       uint   `json:"id"`
		Row      string `json:"row"`
		Number   int    `json:"number"`
		IsBooked bool   `json:"is_booked"`
	}
	var seatsStatus []SeatsWithStatus
	for _, seat := range seats {
		var booking models.Booking
		err := config.DB.Where("seat_id = ? AND showtime_id = ? AND status = ?", seat.ID, showtimeID, "reserved").First(&booking).Error
		isBooked := err == nil
		seatsStatus = append(seatsStatus, SeatsWithStatus{
			ID:       seat.ID,
			Row:      seat.Row,
			Number:   seat.Number,
			IsBooked: isBooked,
		})
	}
	return c.JSON(seatsStatus)
}
func ReserveSeats(c *fiber.Ctx) error {
	userData := c.Locals("user").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	var data struct {
		ShowtimeID uint   `json:"showtime_id"`
		SeatIDs    []uint `json:"seat_ids"`
	}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse request"})
	}
	var showtime models.Showtime
	if err := config.DB.First(&showtime, data.ShowtimeID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "showtime not found"})
	}
	type ReservedResponse struct {
		SeatID     uint `json:"seat_id"`
		ShowtimeID uint `json:"showtime_id"`
	}
	var alreadyBooked []uint
	var reserved []ReservedResponse
	for _, seatID := range data.SeatIDs {
		var booking models.Booking
		err := config.DB.Where("seat_id=? AND showtime_id=? AND status=?", seatID, data.ShowtimeID, "reserved").First(&booking).Error
		if err == nil {
			alreadyBooked = append(alreadyBooked, seatID)
			continue
		}
		newBooking := models.Booking{
			UserID:     userID,
			SeatID:     seatID,
			ShowtimeID: data.ShowtimeID,
			Status:     "reserved",
		}
		if err := config.DB.Create(&newBooking).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "cannot create booking"})
		}
		reserved = append(reserved, ReservedResponse{
			SeatID:     seatID,
			ShowtimeID: data.ShowtimeID,
		})
	}
	if !showtime.IsActive {
		return c.Status(400).JSON(fiber.Map{
			"error": "this showtime is expired",
		})
	}
	return c.JSON(fiber.Map{
		"reserved":       reserved,
		"alreade_booked": alreadyBooked,
	})
}
func GetMyBooking(c *fiber.Ctx) error {
	userData := c.Locals("user").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	var bookings []models.Booking
	if err := config.DB.Preload("Seat").Preload("Showtime").Where("user_id=?", userID).Find(&bookings).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch bookings"})
	}
	return c.JSON(bookings)
}
func CancelBooking(c *fiber.Ctx) error {
	userData := c.Locals("user").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	bookingID := c.Params("id")
	var booking models.Booking
	if err := config.DB.First(&booking, bookingID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "booking not found"})
	}
	if booking.UserID != userID {
		return c.Status(403).JSON(fiber.Map{"error": "you can only cancel your own bookings"})
	}
	booking.Status = "cancel"
	if err := config.DB.Save(&booking).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "cannot cancel booking"})
	}
	return c.JSON(fiber.Map{"message": "booking canceled successfully"})
}
