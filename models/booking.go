package models

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	UserID     uint     `json:"user_id"`
	User       User     `json:"user"`
	SeatID     uint     `json:"seat_id"`
	Seat       Seat     `json:"seat"`
	ShowtimeID uint     `json:"showtime_id"`
	Showtime   Showtime `json:"showtime"`
	Status     string   `json:"status"` //reserved/cancel
}
