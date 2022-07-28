package Router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mvcGolang/app/Http/Controllers"
	"mvcGolang/app/Repository"
)

func UserRouter(db *gorm.DB, api *gin.RouterGroup) {
	userRepository := Repository.NewRepository(db)
	userController := Controllers.NewUserController(userRepository)

	api.GET("/users", userController.GetAllUser)
	api.GET("/user/:id", userController.GetById)
}
