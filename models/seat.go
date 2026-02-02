package models

import "gorm.io/gorm"

type Seat struct {
	gorm.Model
	HallID   uint   `json:"hall_id"`
	Hall     Hall   `gorm:"foreignKey:HallID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Row      string `json:"row"` //abcd
	Number   int    `json:"number"`
	IsActive bool   `json:"is_active"`
}
