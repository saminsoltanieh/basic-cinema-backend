package models

import (
	"gorm.io/gorm"
)

type Hall struct {
	gorm.Model
	Name  string `json:"name" gorm:"unique"`
	Seats []Seat `gorm:"foreignKey:HallID"`
}
