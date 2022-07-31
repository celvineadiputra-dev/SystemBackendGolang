package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mvcGolang/app/Http/Middleware"
	"mvcGolang/routes/Router"
)

func Api(db *gorm.DB) *gin.Engine {
	route := gin.Default()
	api := route.Group("api/v1", Middleware.LogMiddleware())

	Router.UserRouter(db, api)
	Router.CampaignRouter(db, api)

	return route
}
