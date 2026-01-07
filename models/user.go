package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"unique"`
	Password     string `json:"-" gorm:"column:password"`
	Role         string `json:"role"`
	IsActive     bool   `json:"is_active"`
	IsSuperAdmin bool   `json:"is_superadmin"`
}
