package Controllers

import (
	"github.com/gin-gonic/gin"
	"mvcGolang/app/Http/Inputs"
	"mvcGolang/app/Models/Users"
	"mvcGolang/app/Repository"
	"time"
)

type UserController interface {
	GetAllUser()
	GetById(c *gin.Context)
	RegisterUser(c *gin.Context)
}

type userController struct {
	userRepository Repository.UserRepository
}

func NewUserController(userRepository Repository.UserRepository) *userController {
	return &userController{userRepository}
}

func (uc *userController) GetAllUser(c *gin.Context) {
	users, err := uc.userRepository.GetAll()

	if err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, users)
}

func (uc *userController) GetById(c *gin.Context) {
	user, err := uc.userRepository.GetById(c)

	if err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, user)
}

func (uc *userController) RegisterUser(c *gin.Context) {
	var input Inputs.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(500, err.Error())
	}

	user := Users.User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	c.JSON(200, user)

	newUser, err := uc.userRepository.RegisterUser(user)

	if err != nil {
		c.JSON(200, err.Error())
	}

	c.JSON(200, newUser)
}
