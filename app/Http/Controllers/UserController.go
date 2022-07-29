package Controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mvcGolang/app/Helpers"
	"mvcGolang/app/Http/Inputs"
	"mvcGolang/app/Models/Users"
	"time"
)

type UserController interface {
	RegisterUser()
}

type userController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *userController {
	return &userController{db}
}

func (uc *userController) RegisterUser(c *gin.Context) {
	var input Inputs.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := Helpers.ApiResponse(500, "Internal Server Error", nil)
		c.JSON(500, response)
		return
	}

	user := Users.User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	errCreate := uc.db.Create(&user).Error
	if errCreate != nil {
		response := Helpers.ApiResponse(500, "Internal Server Error", nil)
		c.JSON(500, response)
		return
	}

	response := Helpers.ApiResponse(200, "Successfully register new user", user)
	c.JSON(200, response)
}
