package Responses

import (
	"gorm.io/gorm"
	"mvcGolang/app/Helpers"
	"mvcGolang/app/Models/Campaigns"
	"strconv"
	"time"
)

type CampaignResponse struct {
	ID            string              `json:"id"`
	Name          string              `json:"name"`
	ShortDesc     string              `json:"short_desc"`
	Desc          string              `json:"desc"`
	GoalAmount    int                 `json:"goal_amount"`
	CurrentAmount int                 `json:"current_amount"`
	Perks         string              `json:"perks"`
	BackerCount   int                 `json:"backer_count"`
	Slug          string              `json:"slug"`
	CreatedAt     time.Time           `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time           `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt     gorm.DeletedAt      `json:"deleted_at"`
	User          UserResponseRegular `json:"user"`
}

func CampaignSingleFormat(campaign Campaigns.Campaign) CampaignResponse {
	format := CampaignResponse{
		ID:            Helpers.Encrypt(strconv.Itoa(campaign.ID)),
		Name:          campaign.Name,
		ShortDesc:     campaign.ShortDesc,
		Desc:          campaign.Desc,
		GoalAmount:    campaign.GoalAmount,
		CurrentAmount: campaign.CurrentAmount,
		Perks:         campaign.Perks,
		BackerCount:   campaign.BackerCount,
		Slug:          campaign.Slug,
		CreatedAt:     campaign.CreatedAt,
		UpdatedAt:     campaign.UpdatedAt,
		DeletedAt:     campaign.DeletedAt,
		User:          UserFormatRegular(campaign.User),
	}

	return format
}

func CampaignFormat(campaigns []Campaigns.Campaign) []CampaignResponse {
	var listCampaigns []CampaignResponse

	for _, item := range campaigns {
		listCampaigns = append(listCampaigns, CampaignSingleFormat(item))
	}

	return listCampaigns
}
