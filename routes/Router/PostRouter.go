package Router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mvcGolang/app/Http/Controllers"
)

func PostRouter(db *gorm.DB, api *gin.RouterGroup) {
	postController := Controllers.NewPostController(db)

	api.GET("/posts", postController.GetAllPosts)
}
