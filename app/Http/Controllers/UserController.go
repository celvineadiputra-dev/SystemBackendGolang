package Controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mvcGolang/app/Helpers"
	"mvcGolang/app/Http/Inputs"
	"mvcGolang/app/Models/Users"
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
		errors := Helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := Helpers.ApiResponse(500, "Validation Error", errorMessage)
		c.JSON(500, response)
		return
	}

	user := Users.User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Password = Helpers.HashPassword(input.Password)
	user.RoleId = 1

	errCreate := uc.db.Create(&user).Error
	if errCreate != nil {
		response := Helpers.ApiResponse(500, "Internal Server Error", nil)
		c.JSON(500, response)
		return
	}

	response := Helpers.ApiResponse(200, "Successfully register new user", user)
	c.JSON(200, response)
}

func (uc *userController) Login(c *gin.Context) {
	var input Inputs.LoginUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		response := Helpers.ApiResponse(500, "Internal Server Error", nil)
		c.JSON(500, response)
	}

	var user Users.User
	err = uc.db.Where("email = ?", input.Email).Find(&user).Error

	if err != nil {
		response := Helpers.ApiResponse(500, "Internal Server error", nil)
		c.JSON(500, response)
		return
	}

	if user.ID == 0 {
		response := Helpers.ApiResponse(404, "User Not Found", nil)
		c.JSON(404, response)
		return
	}

	isCorrect := Helpers.CheckPasswordHash(input.Password, user.Password)

	if !isCorrect {
		response := Helpers.ApiResponse(401, "Password and email doesn't correct", nil)
		c.JSON(401, response)
		return
	}

	response := Helpers.ApiResponse(200, "Success", user)
	c.JSON(200, response)
}
