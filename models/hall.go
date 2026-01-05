package models

import (
	"gorm.io/gorm"
)

type Hall struct {
	gorm.Model
	Name  string `json:"name"`
	Seats []Seat `gorm:"foreignKey:HallID"`
}
