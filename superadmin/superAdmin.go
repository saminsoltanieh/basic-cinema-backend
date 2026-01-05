package superadmin

import (
	"cinema/config"
	"cinema/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func CreateSuperAdmin() {
	var count int64
	config.DB.Model(&models.User{}).Where("role=?", "superadmin").Count(&count)
	if count == 0 {
		password, err := bcrypt.GenerateFromPassword([]byte("supersecret123"), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("Error generating password", err)
		}
		superAdmin := models.User{
			Name:         "Super Admin",
			Email:        "superadmin@test.com",
			Password:     string(password),
			Role:         "superadmin",
			IsSuperAdmin: true,
		}
		if err := config.DB.Create(&superAdmin).Error; err != nil {
			fmt.Println("Error creating super admin:", err)
			return
		}
		println("super admin created")
	}
}
