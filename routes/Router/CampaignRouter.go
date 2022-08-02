package Router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mvcGolang/app/Http/Controllers"
	"mvcGolang/app/Http/Middleware"
)

func CampaignRouter(db *gorm.DB, api *gin.RouterGroup) {
	campaignController := Controllers.NewCampaignController(db)

	api.GET("/campaigns", Middleware.AuthMiddleware(db), campaignController.Index)
	api.GET("/campaign/:id", Middleware.AuthMiddleware(db), campaignController.Show)
	api.POST("/campaign", Middleware.AuthMiddleware(db), campaignController.Store)
	api.PUT("/campaign/:id", Middleware.AuthMiddleware(db), campaignController.Update)
	api.DELETE("/campaign/:id", Middleware.AuthMiddleware(db), campaignController.Destroy)
}
