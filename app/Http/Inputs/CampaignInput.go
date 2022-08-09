package Inputs

import "mvcGolang/app/Models/Users"

type StoreCampaign struct {
	Name       string `json:"name" binding:"required"`
	ShortDesc  string `json:"short_desc" binding:"required"`
	Desc       string `json:"desc" binding:"required"`
	GoalAmount int    `json:"goal_amount" binding:"required,min=1000"`
	Perks      string `json:"perks" binding:"required"`
}

type UpdateCampaign struct {
	ShortDesc string `json:"short_desc" binding:"required"`
	Desc      string `json:"desc" binding:"required"`
}

type UploadCampaignImage struct {
	CampaignId string `form:"campaign_id" binding:"required"`
	IsPrimary  bool   `form:"is_primary"`
	User       Users.User
}
