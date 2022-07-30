package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mvcGolang/routes/Router"
)

func Api(db *gorm.DB) *gin.Engine {
	route := gin.Default()
	api := route.Group("api/v1")

	Router.UserRouter(db, api)

	return route
}
