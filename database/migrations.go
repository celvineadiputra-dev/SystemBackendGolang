package database

import (
	"mvcGolang/app/Models/Campaigns"
	"mvcGolang/app/Models/Roles"
	"mvcGolang/app/Models/Transactions"
	"mvcGolang/app/Models/Users"
	"mvcGolang/config"
)

func Migrate() {
	config.Database().AutoMigrate(Users.User{}, Roles.Role{}, Campaigns.Campaign{}, Campaigns.CampaignImage{}, Transactions.Transaction{})

	config.Database().Model(&Users.User{})
	config.Database().Model(&Roles.Role{})
	config.Database().Model(&Campaigns.Campaign{})
	config.Database().Model(&Campaigns.CampaignImage{})
	config.Database().Model(&Transactions.Transaction{})
}
