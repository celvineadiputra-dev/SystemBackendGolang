package Transactions

import (
	"gorm.io/gorm"
	"mvcGolang/app/Models/Campaigns"
	"mvcGolang/app/Models/Users"
	"time"
)

type Transaction struct {
	ID         int                `json:"id"`
	CampaignId int                `json:"campaigns_id"`
	Campaign   Campaigns.Campaign `gorm:"foreignKey:CampaignId"`
	UserId     int                `json:"user_id"`
	User       Users.User         `gorm:"foreignKey:UserId"`
	Amount     int                `json:"amount"`
	Status     int                `json:"status"`
	Code       int                `json:"code"`
	CreatedAt  time.Time          `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time          `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  gorm.DeletedAt     `json:"deleted_at"`
}
