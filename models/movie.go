package models

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Duration    int       `json:"duration"`
	ReleaseDate time.Time `json:"release_date"`
}
