package Campaigns

import (
	"gorm.io/gorm"
	"time"
)

type CampaignImage struct {
	ID         int            `json:"id"`
	CampaignId int            `json:"campaign_id"`
	Campaign   Campaign       `gorm:"foreignKey:CampaignId"`
	FileName   string         `json:"file_name"`
	IsPrimary  bool           `json:"is_primary"`
	CreatedAt  time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}
