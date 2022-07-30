package Campaigns

import (
	"gorm.io/gorm"
	"mvcGolang/app/Models/Users"
	"time"
)

type Campaign struct {
	ID             int             `json:"id"`
	UserId         int             `json:"user_id"`
	User           Users.User      `gorm:"foreignKey:UserId"`
	ShortDesc      string          `gorm:"size:100" json:"short_desc"`
	Desc           string          `json:"desc"`
	GoalAmount     int             `json:"goal_amount"`
	CurrentAmount  int             `json:"current_amount"`
	Perks          string          `json:"perks"`
	BackerCount    int             `json:"backer_count"`
	Slug           string          `json:"slug"`
	CampaignImages []CampaignImage `gorm:"foreignKey:CampaignId"`
	CreatedAt      time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt      gorm.DeletedAt  `json:"deleted_at"`
}
