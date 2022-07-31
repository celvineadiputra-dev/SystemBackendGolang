package Router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mvcGolang/app/Http/Controllers"
)

func UserRouter(db *gorm.DB, api *gin.RouterGroup) {
	userController := Controllers.NewUserController(db)

	api.POST("/register", userController.RegisterUser)
	api.POST("/login", userController.Login)
}
