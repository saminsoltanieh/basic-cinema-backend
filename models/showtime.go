package models

import (
	"time"

	"gorm.io/gorm"
)

type Showtime struct {
	gorm.Model
	MovieID   uint      `json:"movie_id"`
	Movie     Movie     `gorm:"foreignKey:MovieID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	HallID    uint      `json:"hall_id"`
	Hall      Hall      `gorm:"foreignKey:HallID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Price     int       `json:"price"`
	IsActive  bool      `json:"is_active"`
}
