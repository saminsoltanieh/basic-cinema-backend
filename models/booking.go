package models

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	UserID     uint     `json:"user_id" gorm:"not null"`
	User       User     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	SeatID     uint     `json:"seat_id" gorm:"not null"`
	Seat       Seat     `gorm:"foreignKey:SeatID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ShowtimeID uint     `json:"showtime_id" gorm:"not null"`
	Showtime   Showtime `gorm:"foreignKey:ShowtimeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Status     string   `json:"status" gorm:"check:status IN ('reserved','cancel')"` //reserved/cancel
}
