package Campaigns

import (
	"gorm.io/gorm"
	"time"
)

type CampaignImage struct {
	ID         int            `json:"id"`
	CampaignId int            `json:"campaign_id"`
	Campaign   Campaign       `gorm:"foreignKey:CampaignId"`
	FileName   int            `json:"file_name"`
	IsPrimary  int            `json:"is_primary"`
	CreatedAt  time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}
