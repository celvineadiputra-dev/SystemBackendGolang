package Controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mvcGolang/app/Helpers"
	"mvcGolang/app/Http/Responses"
	"mvcGolang/app/Models/Posts"
)

type PostController interface {
	GetAllPosts()
}

type postController struct {
	db *gorm.DB
}

func NewPostController(db *gorm.DB) *postController {
	return &postController{db}
}

func (ps *postController) GetAllPosts(c *gin.Context) {
	var posts []Posts.Post

	err := ps.db.Find(&posts).Error
	if err != nil {
		response := Helpers.ApiResponse(500, err.Error(), nil)
		c.JSON(500, response)
		return
	}

	response := Helpers.ApiResponse(200, "List of Posts", Responses.FormatPosts(posts))
	c.JSON(200, response)
}
