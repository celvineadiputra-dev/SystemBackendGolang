package Users

import (
	"gorm.io/gorm"
	"mvcGolang/app/Models/Roles"
	"time"
)

type User struct {
	ID         int            `json:"id,primary_key"`
	Name       string         `gorm:"size:50" json:"name"`
	Email      string         `gorm:"size:50" json:"email"`
	Occupation string         `gorm:"size:50" json:"occupation"`
	Avatar     string         `json:"avatar"`
	RoleId     int            `json:"role_id"`
	Role       Roles.Role     `grom:"foreignKey:RoleId"`
	Token      string         `json:"token"`
	Password   string         `json:"password"`
	CreatedAt  time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}
