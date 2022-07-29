package Router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mvcGolang/app/Http/Controllers"
	"mvcGolang/app/Http/Middleware"
)

func UserRouter(db *gorm.DB, api *gin.RouterGroup) {
	userController := Controllers.NewUserController(db)

	api.POST("/user", Middleware.LogMiddleware(), userController.RegisterUser)
}
