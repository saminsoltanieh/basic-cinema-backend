package models

import (
	"time"

	"gorm.io/gorm"
)

type Showtime struct {
	gorm.Model
	MovieID   uint      `json:"movie_id"`
	Movie     Movie     `json:"movie"`
	HallID    uint      `json:"hall_id"`
	Hall      Hall      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Price     int       `json:"price"`
	IsActive  bool      `json:"is_active"`
}
