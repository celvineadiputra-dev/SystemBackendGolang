package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"mvcGolang/app/Helpers"
	"mvcGolang/app/Http/Inputs"
	"mvcGolang/app/Http/Responses"
	"mvcGolang/app/Models/Campaigns"
	"mvcGolang/app/Models/Users"
)

type CampaignController interface {
	Index()
}

type campaignController struct {
	db *gorm.DB
}

func NewCampaignController(db *gorm.DB) *campaignController {
	return &campaignController{db}
}

func (cc *campaignController) Index(c *gin.Context) {
	var campaigns []Campaigns.Campaign

	err := cc.db.Preload(clause.Associations).Find(&campaigns).Error

	if err != nil {
		response := Helpers.ApiResponse(500, "Internal Server Error", nil)
		c.JSON(500, response)
		return
	}

	campaignResponse := Responses.CampaignFormat(campaigns)

	response := Helpers.ApiResponse(200, "List of campaigns", campaignResponse)
	c.JSON(200, response)
}

func (cc *campaignController) Show(c *gin.Context) {
	var campaign Campaigns.Campaign

	id := Helpers.Decrypt(c.Param("id"))
	err := cc.db.Preload(clause.Associations).Where("id", id).Find(&campaign).Error
	if err != nil {
		response := Helpers.ApiResponse(500, "Internal Server Error", nil)
		c.JSON(500, response)
		return
	}

	if campaign.ID == 0 {
		response := Helpers.ApiResponse(404, "Campaign Not Found", nil)
		c.JSON(404, response)
		return
	}

	campaignFormat := Responses.CampaignSingleFormat(campaign)
	response := Helpers.ApiResponse(200, "Camapign", campaignFormat)
	c.JSON(200, response)
}

func (cc *campaignController) Store(c *gin.Context) {
	var input Inputs.StoreCampaign

	err := c.ShouldBind(&input)
	if err != nil {
		errors := Helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := Helpers.ApiResponse(500, "Validation Error", errorMessage)
		c.JSON(500, response)
		return
	}

	campaign := Campaigns.Campaign{}
	campaign.Name = input.Name
	campaign.ShortDesc = input.ShortDesc
	campaign.Desc = input.Desc
	campaign.GoalAmount = input.GoalAmount
	campaign.CurrentAmount = 0
	campaign.Perks = input.Perks
	campaign.BackerCount = 0
	campaign.Slug = slug.Make(input.Name)
	campaign.UserId = c.MustGet("currentUser").(Users.User).ID

	errCreate := cc.db.Create(&campaign).Error

	if errCreate != nil {
		response := Helpers.ApiResponse(500, "Internal Server Error", nil)
		c.JSON(500, response)
		return
	}

	response := Helpers.ApiResponse(200, "Success add new campaign", nil)
	c.JSON(200, response)
}

func (cc *campaignController) Update(c *gin.Context) {
	var input Inputs.UpdateCampaign

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := Helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := Helpers.ApiResponse(500, "Validation Error", errorMessage)
		c.JSON(500, response)
		return
	}

	var campaign Campaigns.Campaign

	id := Helpers.Decrypt(c.Param("id"))
	err = cc.db.Where("id = ?", id).Find(&campaign).Error
	if err != nil {
		response := Helpers.ApiResponse(500, "Internal Server Error", nil)
		c.JSON(500, response)
		return
	}

	if campaign.ID == 0 {
		response := Helpers.ApiResponse(404, "Campaign Not Found", nil)
		c.JSON(404, response)
		return
	}

	if campaign.UserId != c.MustGet("currentUser").(Users.User).ID {
		response := Helpers.ApiResponse(402, "Not an owner of the campaign", nil)
		c.JSON(402, response)
		return
	}

	campaign.ShortDesc = input.ShortDesc
	campaign.Desc = input.Desc
	errUpdate := cc.db.Save(&campaign).Error

	if errUpdate != nil {
		response := Helpers.ApiResponse(402, "Update Failed", nil)
		c.JSON(402, response)
		return
	}

	campaignFormat := Responses.CampaignSingleFormat(campaign)
	response := Helpers.ApiResponse(200, "Update Successfully", campaignFormat)
	c.JSON(200, response)
}

func (cc *campaignController) Destroy(c *gin.Context) {
	id := Helpers.Decrypt(c.Param("id"))
	userId := c.MustGet("currentUser").(Users.User).ID

	var campaign Campaigns.Campaign
	err := cc.db.Where("id = ? and user_id = ?", id, userId).Find(&campaign).Error

	if err != nil {
		response := Helpers.ApiResponse(500, "Internal Server Error", nil)
		c.JSON(500, response)
		return
	}

	if campaign.ID == 0 {
		response := Helpers.ApiResponse(404, "Campaign Not Found", nil)
		c.JSON(404, response)
		return
	}

	errDelete := cc.db.Delete(&campaign).Error

	if errDelete != nil {
		response := Helpers.ApiResponse(500, "Internal Server Error", nil)
		c.JSON(500, response)
		return
	}

	response := Helpers.ApiResponse(200, "Success Delete Campaign", nil)
	c.JSON(200, response)
}

func (cc *campaignController) UploadImage(c *gin.Context) {
	var input Inputs.UploadCampaignImage

	err := c.ShouldBind(&input)
	if err != nil {
		errors := Helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := Helpers.ApiResponse(500, "Validation Error", errorMessage)
		c.JSON(500, response)
		return
	}

	currentUser := c.MustGet("currentUser").(Users.User)
	input.User = currentUser
	userId := currentUser.ID

	file, err := c.FormFile("file")
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}

		response := Helpers.ApiResponse(500, "Failed to upload campaign image", errorMessage)
		c.JSON(500, response)
		return
	}

	path := fmt.Sprintf("images/%d-%s", userId, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}

		response := Helpers.ApiResponse(500, "Failed to upload campaign image", errorMessage)
		c.JSON(500, response)
		return
	}

	input.CampaignId = Helpers.Decrypt(input.CampaignId)

	//cari campaign
	var campaign Campaigns.Campaign
	err = cc.db.Where("id = ?", input.CampaignId).Find(&campaign).Error
	if err != nil {
		response := Helpers.ApiResponse(500, "Internal Server Error", nil)
		c.JSON(500, response)
		return
	}

	if campaign.ID == 0 {
		response := Helpers.ApiResponse(404, "Campaign Not Found", nil)
		c.JSON(404, response)
		return
	}

	if campaign.UserId != userId {
		response := Helpers.ApiResponse(402, "Not an owner of the campaign", campaign.ID)
		c.JSON(402, response)
		return
	}

	if input.IsPrimary {
		err = cc.db.Where("campaign_id = ?", campaign.ID).Model(&Campaigns.CampaignImage{}).Update("is_primary", false).Error

		if err != nil {
			response := Helpers.ApiResponse(500, "Internal Server Error", nil)
			c.JSON(500, response)
			return
		}
	}

	campaignImage := Campaigns.CampaignImage{}
	campaignImage.CampaignId = campaign.ID
	campaignImage.FileName = path
	campaignImage.IsPrimary = input.IsPrimary
	err = cc.db.Create(&campaignImage).Error
	if err != nil {
		response := Helpers.ApiResponse(500, "Internal Server Error", nil)
		c.JSON(500, response)
		return
	}

	response := Helpers.ApiResponse(200, "Success add new campaign image", nil)
	c.JSON(200, response)
}
