package database

import (
	"gorm.io/gorm"
	"mvcGolang/app/Models/Roles"
)

func RoleSeed(db *gorm.DB) {
	var role = []Roles.Role{
		{Role: "User"},
		{Role: "Admin"},
	}

	db.Create(&role)
}
